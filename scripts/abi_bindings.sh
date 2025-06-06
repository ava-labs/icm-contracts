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

export ARCH=$(uname -m)
[ $ARCH = x86_64 ] && ARCH=amd64
echo "ARCH set to $ARCH"

DEFAULT_CONTRACT_LIST="TeleporterMessenger TeleporterRegistry ExampleERC20 ExampleRewardCalculator TestMessenger ValidatorSetSig NativeTokenStakingManager ERC20TokenStakingManager
TokenHome TokenRemote ERC20TokenHome ERC20TokenHomeUpgradeable ERC20TokenRemote ERC20TokenRemoteUpgradeable NativeTokenHome NativeTokenHomeUpgradeable NativeTokenRemote NativeTokenRemoteUpgradeable
WrappedNativeToken MockERC20SendAndCallReceiver MockNativeSendAndCallReceiver ExampleERC20Decimals IStakingManager ACP99Manager ValidatorManager PoAManager"
PROXY_LIST="TransparentUpgradeableProxy ProxyAdmin"
ACCESS_LIST="OwnableUpgradeable"

SUBNET_EVM_LIST="INativeMinter"

EXTERNAL_LIBS="ValidatorMessages"

CONTRACT_LIST=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -c | --contract) CONTRACT_LIST=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/abi_bindings.sh [OPTIONS]"
    echo "Build contracts and generate Go bindings"
    echo ""
    echo "Options:"
    echo "  -c, --contract <contract_name>          Generate Go bindings for the contract. If empty, generate Go bindings for a default list of contracts"
    echo "  -c, --contract "contract1 contract2"    Generate Go bindings for multiple contracts"
    echo "  -h, --help                              Print this help message"
    exit 0
fi

if ! command -v forge &> /dev/null; then
    echo "forge not found. You can install by calling $ICM_CONTRACTS_PATH/scripts/install_foundry.sh" && exit 1
fi

if ! command -v solc &> /dev/null; then
    echo "solc not found. See https://docs.soliditylang.org/en/latest/installing-solidity.html for installation instructions" && exit 1
fi
    
# Get the version from solc output
solc_version_output=$(solc --version 2>&1)

# Extract the semver version from the output
extracted_version=$(solc --version 2>&1 | awk '/Version:/ {print $2}' | awk -F'+' '{print $1}')

# Check if the extracted version matches the expected version
if ! [[ "$extracted_version" == "$SOLIDITY_VERSION" ]]; then
    echo "Expected solc version $SOLIDITY_VERSION, but found $extracted_version. Please install the correct version." && exit 1
fi

# Install abigen
echo "Building subnet-evm abigen"
go install github.com/ava-labs/subnet-evm/cmd/abigen@${SUBNET_EVM_VERSION}

# Solc does not recursively expand remappings, so we must construct them manually
remappings=$(cat $ICM_CONTRACTS_PATH/remappings.txt)

# Recursively search for all remappings.txt files in the lib directory.
# For each file, prepend the remapping with the relative path to the file.
while read -r filepath; do
    relative_path="${filepath#$ICM_CONTRACTS_PATH/}"
    dir_path=$(dirname "$relative_path")
    echo $dir_path
  
    # Use sed to transform each line with the relative path if it matches the @token=remapping pattern,
    # so that each remapping is of the form @token=lib/path/to/remapping
    transformed_lines=$(sed -n "s|^\(@[^=]*=\)\(.*\)|\1$dir_path/\2|p" "$filepath")
    remappings+=" $transformed_lines "
done < <(find "$ICM_CONTRACTS_PATH/lib" -type f -name "remappings.txt" )

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
    local contract_names=("$@")
    for contract_name in "${contract_names[@]}"
    do
        path=$(find . -name $contract_name.sol)
        dir=$(dirname $path)
        dir="${dir#./}"

        echo "Building $contract_name..."
        mkdir -p $ICM_CONTRACTS_PATH/out/$contract_name.sol
        
        combined_json=$ICM_CONTRACTS_PATH/out/$contract_name.sol/combined-output.json

        cwd=$(pwd)
        cd $ICM_CONTRACTS_PATH
        solc --optimize --evm-version $EVM_VERSION --combined-json abi,bin,metadata,ast,devdoc,userdoc --pretty-json $cwd/$dir/$contract_name.sol $remappings > $combined_json
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

# If CONTRACT_LIST is empty, use DEFAULT_CONTRACT_LIST
if [[ -z "${CONTRACT_LIST}" ]]; then
    contract_names=($DEFAULT_CONTRACT_LIST)
fi

cd $ICM_CONTRACTS_PATH/contracts
generate_bindings "${contract_names[@]}"

contract_names=($PROXY_LIST)
cd $ICM_CONTRACTS_PATH/lib/openzeppelin-contracts-upgradeable/lib/openzeppelin-contracts/contracts/proxy/transparent
generate_bindings "${contract_names[@]}"

contract_names=($SUBNET_EVM_LIST)
cd $ICM_CONTRACTS_PATH/lib/subnet-evm/contracts/contracts/interfaces
generate_bindings "${contract_names[@]}"

contract_names=($ACCESS_LIST)
cd $ICM_CONTRACTS_PATH/lib/openzeppelin-contracts-upgradeable/contracts/access
generate_bindings "${contract_names[@]}"

exit 0
