#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

ICM_CONTRACTS_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd ../ && pwd
)

function printHelp() {
    echo "Usage: ./scripts/e2e_test.sh [--component component]"
    echo ""
    printUsage
}

function printUsage() {
    cat << EOF
Arguments:
    --components component1,component2            Comma separated list of test suites to run. Valid components are:
                                                  $(echo $valid_components | tr ' ' '\n' | sort | tr '\n' ' ')
                                                  (default: all)
Options:
    --help                                        Print this help message
EOF
}

valid_components=$(ls -d $ICM_CONTRACTS_PATH/tests/suites/*/ | xargs -n 1 basename)
components=

while [ $# -gt 0 ]; do
    case "$1" in
        --components)  
            if [[ $2 != --* ]]; then
                components=$2
            else 
                echo "Invalid components $2" && printHelp && exit 1
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

source "$ICM_CONTRACTS_PATH"/scripts/constants.sh
source "$ICM_CONTRACTS_PATH"/scripts/versions.sh

BASEDIR=${BASEDIR:-"$HOME/.teleporter-deps"}

cwd=$(pwd)
# Install the avalanchego and subnet-evm binaries
rm -rf $BASEDIR/avalanchego
BASEDIR=$BASEDIR AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego "${ICM_CONTRACTS_PATH}/scripts/install_avalanchego_release.sh"
BASEDIR=$BASEDIR "${ICM_CONTRACTS_PATH}/scripts/install_subnetevm_release.sh"

cp ${BASEDIR}/subnet-evm/subnet-evm ${BASEDIR}/avalanchego/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy
echo "Copied ${BASEDIR}/subnet-evm/subnet-evm binary to ${BASEDIR}/avalanchego/plugins/"

export AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego

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
# Build ginkgo
# to install the ginkgo binary (required for test build and run)
go install -v github.com/onsi/ginkgo/v2/ginkgo@${GINKGO_VERSION}

for component in $(echo $components | tr ',' ' '); do
    echo "Building e2e tests for $component"
    ginkgo build ./tests/suites/$component

    echo "Running e2e tests for $component"

    RUN_E2E=true SIG_AGG_PATH=$ICM_SERVICES_BUILD_PATH/signature-aggregator ./tests/suites/$component/$component.test \
    --ginkgo.vv \
    --ginkgo.label-filter=${GINKGO_LABEL_FILTER:-""} \
    --ginkgo.focus=${GINKGO_FOCUS:-""} \
    --ginkgo.trace

    echo "$component e2e tests passed"
    echo ""
done
exit 0
