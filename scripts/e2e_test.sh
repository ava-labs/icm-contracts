#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

ICM_CONTRACTS_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd ../ && pwd
)

function printHelp() {
    echo "Usage: ./scripts/e2e_test.sh [--components component]"
    echo ""
    printUsage
}

function printUsage() {
    cat << EOF
Arguments:
    --components component1,component2            Comma separated list of test suites to run. Valid components are:
                                                  $(echo $valid_components | tr ' ' '\n' | sort | tr '\n' ' ')
                                                  (default: all)
    --network-dir path                            Path to the network directory. 
                                                  If the directory does not exist or is empty, it will be used as the root network directory for a new network.
                                                  If the directory exists and is non-empty, the network will be reused.
                                                  If not set, a new network will be created at the default root network directory.
    --epoch-duration duration                     Set to override the default test epoch duration.
Options:
    --help                                        Print this help message
EOF
}

valid_components=$(ls -d $ICM_CONTRACTS_PATH/tests/suites/*/ | xargs -n 1 basename)
components=
reuse_network_dir=
root_dir=
network_dir=
reuse_network=false
epoch_duration=
while [ $# -gt 0 ]; do
    case "$1" in
        --components)  
            if [[ $2 != --* ]]; then
                components=$2
            else 
                echo "Invalid components $2" && printHelp && exit 1
            fi 
            shift;;
        --network-dir)
            if [[ $2 != --* ]]; then
                reuse_network_dir=$2
            else 
                echo "Invalid network directory $2" && printHelp && exit 1
            fi 
            shift;;
        --epoch-duration)
            if [[ $2 != --* ]]; then
                epoch_duration=$2
            else 
                echo "Invalid epoch duration $2" && printHelp && exit 1
            fi 
            shift;;
        --help) 
            printHelp && exit 0 ;;
        *) 
            echo "Invalid option: $1" && printHelp && exit 1;;
    esac
    shift
done

# Run all suites if no component is provided
if [ -z "$components" ]; then
    components=$valid_components
fi

# Exit if invalid component is provided
for component in $(echo $components | tr ',' ' '); do
    if [[ $valid_components != *$component* ]]; then
        echo "Invalid component $component" && exit 1
    fi
done

if [ -n "$reuse_network_dir" ]; then
    if [ -d "$reuse_network_dir" ] && [ "$(ls -A "$reuse_network_dir")" ]; then
        network_dir=$reuse_network_dir
        reuse_network=true
        echo "Reuse network directory: $network_dir"
    else
        echo "Network directory $reuse_network_dir does not exist or is empty. Creating a new network at root $reuse_network_dir."
        mkdir -p "$reuse_network_dir"
        root_dir=$reuse_network_dir
    fi
fi

if [ -n "$epoch_duration" ]; then
    export GRANITE_EPOCH_DURATION=$epoch_duration
    echo "GRANITE_EPOCH_DURATION: $GRANITE_EPOCH_DURATION"
fi

source "$ICM_CONTRACTS_PATH"/scripts/constants.sh
source "$ICM_CONTRACTS_PATH"/scripts/versions.sh

BASEDIR=${BASEDIR:-"$HOME/.icm-contracts-e2e-deps"}

cwd=$(pwd)
# Install the avalanchego and subnet-evm binaries
rm -rf $BASEDIR/avalanchego
BASEDIR=$BASEDIR AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego "${ICM_CONTRACTS_PATH}/scripts/install_avalanchego_release.sh"
BASEDIR=$BASEDIR "${ICM_CONTRACTS_PATH}/scripts/install_subnetevm_release.sh"

cp ${BASEDIR}/subnet-evm/subnet-evm ${BASEDIR}/avalanchego/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy
echo "Copied ${BASEDIR}/subnet-evm/subnet-evm binary to ${BASEDIR}/avalanchego/plugins/"

export AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego
export AVALANCHEGO_PATH=$AVALANCHEGO_BUILD_PATH/avalanchego
export AVAGO_PLUGIN_DIR=$AVALANCHEGO_BUILD_PATH/plugins

ICM_SERVICES_BUILD_PATH=$BASEDIR/icm-services

cd $ICM_CONTRACTS_PATH
# Install signature-aggregator binary
BASEDIR=$BASEDIR ICM_SERVICES_BUILD_PATH=$ICM_SERVICES_BUILD_PATH "${ICM_CONTRACTS_PATH}/scripts/install_sig_agg_release.sh"
echo "Installed signature-aggregator from icm-services release ${ICM_SERVICES_VERSION}"

cd $ICM_CONTRACTS_PATH
if command -v forge &> /dev/null; then
  forge build --skip test
else
  echo "Forge command not found, attempting to use from $HOME"
  $HOME/.foundry/bin/forge build
fi

cd "$ICM_CONTRACTS_PATH"

for component in $(echo $components | tr ',' ' '); do
    echo "Building e2e tests for $component"
    go run github.com/onsi/ginkgo/v2/ginkgo build ./tests/suites/$component

    echo "Running e2e tests for $component"

    RUN_E2E=true SIG_AGG_PATH=$ICM_SERVICES_BUILD_PATH/signature-aggregator ./tests/suites/$component/$component.test \
    --root-network-dir=${root_dir} \
    --reuse-network=${reuse_network} \
    --network-dir=${network_dir} \
    --ginkgo.vv \
    --ginkgo.label-filter=${GINKGO_LABEL_FILTER:-""} \
    --ginkgo.focus=${GINKGO_FOCUS:-""} \
    --ginkgo.trace

    echo "$component e2e tests passed"
    echo ""
done
exit 0
