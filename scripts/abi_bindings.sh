#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

ICM_CONTRACTS_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source $ICM_CONTRACTS_PATH/scripts/constants.sh
source $ICM_CONTRACTS_PATH/scripts/versions.sh

echo "Avalanche EVM Version: $AVALANCHE_EVM_VERSION"
echo "Avalanche Solidity Version: $AVALANCHE_SOLIDITY_VERSION"

AVALANCHE_ICM_PATH=${ICM_CONTRACTS_PATH}/avalanche

export ARCH=$(uname -m)
[ $ARCH = x86_64 ] && ARCH=amd64
echo "ARCH set to $ARCH"

DEFAULT_AVALANCHE_CONTRACT_LIST="TeleporterMessenger TeleporterRegistry ExampleERC20 ExampleRewardCalculator TestMessenger ValidatorSetSig NativeTokenStakingManager ERC20TokenStakingManager
TokenHome TokenRemote ERC20TokenHome ERC20TokenHomeUpgradeable ERC20TokenRemote ERC20TokenRemoteUpgradeable NativeTokenHome NativeTokenHomeUpgradeable NativeTokenRemote NativeTokenRemoteUpgradeable
WrappedNativeToken MockERC20SendAndCallReceiver MockNativeSendAndCallReceiver ExampleERC20Decimals IStakingManager ACP99Manager ValidatorManager PoAManager NativeWarp"
PROXY_LIST="TransparentUpgradeableProxy ProxyAdmin"
ACCESS_LIST="OwnableUpgradeable"

SUBNET_EVM_LIST="INativeMinter"

EXTERNAL_LIBS="ValidatorMessages"

AVALANCHE_CONTRACT_LIST=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -ac | --avalanche-contracts) AVALANCHE_CONTRACT_LIST=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/abi_bindings.sh [OPTIONS]"
    echo "Build contracts and generate Go bindings"
    echo ""
    echo "Options:"
    echo "  -ac, --avalanche-contracts contract1 contract2    Generate Go bindings for the contract. If empty, generate Go bindings for a default list of Avalanche contracts"
    echo "  -h,  --help                                       Print this help message"
    exit 0
fi

if ! command -v forge &> /dev/null; then
    echo "forge not found. You can install by calling $ICM_CONTRACTS_PATH/scripts/install_foundry.sh" && exit 1
fi

if ! command -v svm &> /dev/null; then
    echo "svm not found, installing..."
    # install Rust toolchain
    curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | bash -s -- -y
    ~/.cargo/bin/cargo install svm-rs
fi

# Install abigen
# Temporarily hardcode the version to v0.7.8 since cmd/abigen is removed in the following versions.
# TODO: Remove the hardcoded version and switch to using libevm abigen once the rest of the code is updated to use it. 
echo "Building subnet-evm abigen"
go install github.com/ava-labs/subnet-evm/cmd/abigen@v0.7.8

# Solc does not recursively expand remappings, so we must construct them manually
avalanche_remappings=$(cat $AVALANCHE_ICM_PATH/remappings.txt)

# Helper function to expand remappings for a given base path and variable name
expand_remappings() {
    local base_path="$1"
    local remappings_var="$2"
    local remappings_value

    # Recursively search for all remappings.txt files in the lib directory.
    # For each file, prepend the remapping with the relative path to the file.
    while read -r filepath; do
        relative_path="${filepath#$base_path/}"
        dir_path=$(dirname "$relative_path")
        echo $dir_path

        # Use sed to transform each line with the relative path if it matches the @token=remapping pattern,
        # so that each remapping is of the form @token=lib/path/to/remapping
        transformed_lines=$(sed -n "s|^\(@[^=]*=\)\(.*\)|\1$dir_path/\2|p" "$filepath")
        remappings_value+=" $transformed_lines "
    done < <(find "$base_path/lib" -type f -name "remappings.txt" )

    # Use indirect reference to set the variable by name
    eval "$remappings_var=\"\${$remappings_var}\$remappings_value\""
}

# Expand remappings for Avalanche and Ethereum
expand_remappings "$AVALANCHE_ICM_PATH" "avalanche_remappings"

function convertToLower() {
    if [ "$ARCH" = 'arm64' ]; then
        echo $1 | perl -ne 'print lc'
    else
        echo $1 | sed -e 's/\(.*\)/\L\1/'
    fi
}

# Removes a matching string from a comma-separated list
remove_matching_string() {
    input_list="$1"
    match="$2"
    # Split the input list by commas
    IFS=',' read -ra elements <<< "$input_list"
    
    # Initialize an empty result array
    result=()

    # Iterate over each element
    for element in "${elements[@]}"; do
        # Check if the part after the colon matches the given string
        if [[ "${element#*:}" != "$match" ]]; then
        # If it doesn't match, add the element to the result array
        result+=("$element")
        fi
    done

    # Join the result array with commas and print
    (IFS=','; echo "${result[*]}")
}

function generate_bindings() {
    local project_base_path="$1"
    local evm_version="$2"
    local solc_version="$3"
    local remappings="$4"
    local additional_flags="$5"
    shift 5
    local contract_names=("$@")

    # Use svm to switch to the required Solidity version
    echo "Switching to Solidity version $solc_version..."
    svm install $solc_version --non-interactive
    svm use $solc_version

    echo "Project: $project_base_path"
    echo "EVM Version: $evm_version"
    echo "Solidity Version: $solc_version"
    echo "Additional flags: $additional_flags"

    for contract_name in "${contract_names[@]}"
    do
        path=$(find . -name $contract_name.sol)
        dir=$(dirname $path)
        dir="${dir#./}"

        echo "Building $contract_name..."
        mkdir -p $project_base_path/out/$contract_name.sol
        
        combined_json=$project_base_path/out/$contract_name.sol/combined-output.json

        cwd=$(pwd)
        cd $project_base_path
        solc --optimize --evm-version $evm_version $additional_flags --combined-json abi,bin,metadata,ast,devdoc,userdoc --pretty-json $cwd/$dir/$contract_name.sol $remappings > $combined_json
        cd $cwd

        # construct the exclude list
        contracts=$(jq -r '.contracts | keys | join(",")' $combined_json)

        # Filter out the contract we are generating bindings for
        filtered_contracts=$(remove_matching_string $contracts $contract_name)
        
        gen_path=$ICM_CONTRACTS_PATH/abi-bindings/go/$dir/$contract_name
        mkdir -p $gen_path
        echo "Generating Go bindings for $contract_name..."
        
        if [ -z "$filtered_contracts" ]; then
            echo "No external libraries found"
            $GOPATH/bin/abigen --pkg $(convertToLower $contract_name) \
                            --combined-json $combined_json \
                            --type $contract_name \
                            --out $gen_path/$contract_name.go
        else
            # Filter out external libraries
            for lib in $EXTERNAL_LIBS; do
                filtered_contracts=$(remove_matching_string $filtered_contracts $lib)
            done

            $GOPATH/bin/abigen --pkg $(convertToLower $contract_name) \
                            --combined-json $combined_json \
                            --type $contract_name \
                            --out $gen_path/$contract_name.go \
                            --exc $filtered_contracts
        fi
        
        echo "Done generating Go bindings for $contract_name."
    done
}

contract_names=($CONTRACT_LIST)

# If AVALANCHE_CONTRACT_LIST is empty, use DEFAULT_AVALANCHE_CONTRACT_LIST
if [[ -z "${AVALANCHE_CONTRACT_LIST}" ]]; then
    AVALANCHE_CONTRACT_LIST=($DEFAULT_AVALANCHE_CONTRACT_LIST)
fi


contract_names=($AVALANCHE_CONTRACT_LIST)
cd $AVALANCHE_ICM_PATH/contracts
generate_bindings "$AVALANCHE_ICM_PATH" "$AVALANCHE_EVM_VERSION" "$AVALANCHE_SOLIDITY_VERSION" "$avalanche_remappings" "" "${contract_names[@]}"

contract_names=($PROXY_LIST)
cd $AVALANCHE_ICM_PATH/lib/openzeppelin-contracts-upgradeable/lib/openzeppelin-contracts/contracts/proxy/transparent
generate_bindings "$AVALANCHE_ICM_PATH" "$AVALANCHE_EVM_VERSION" "$AVALANCHE_SOLIDITY_VERSION" "$avalanche_remappings" "" "${contract_names[@]}"

contract_names=($ACCESS_LIST)
cd $AVALANCHE_ICM_PATH/lib/openzeppelin-contracts-upgradeable/contracts/access
generate_bindings "$AVALANCHE_ICM_PATH" "$AVALANCHE_EVM_VERSION" "$AVALANCHE_SOLIDITY_VERSION" "$avalanche_remappings" "" "${contract_names[@]}"

exit 0
