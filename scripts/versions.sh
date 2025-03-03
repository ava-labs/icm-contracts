#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e
set -o pipefail

ICM_CONTRACTS_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

# Pass in the full name of the dependency.
# Parses go.mod for a matching entry and extracts the version number.
function getDepVersion() {
    grep -m1 "^\s*$1" $ICM_CONTRACTS_PATH/go.mod | cut -d ' ' -f2
}

function extract_commit() {
  local version=$1

  # Regex for a commit hash (assumed to be a 12+ character hex string)
  commit_hash_regex="-([0-9a-f]{12,})$"

  if [[ "$version" =~ $commit_hash_regex ]]; then
      # Extract the substring after the last '-'
      version=${BASH_REMATCH[1]}
  fi
  echo "$version"
}

# ICM_SERVICES_VERSION is needed for the E2E tests but is not a direct dependency since that would create a circular dependency.
# ICM_SERVICES_VERSION=${ICM_SERVICES_VERSION:-'signature-aggregator-v1.0.0-rc.0'}
ICM_SERVICES_VERSION=${ICM_SERVICES_VERSION:-'1ff2e4f1313e5d0a4961cc6dd680b27d9331fa1f'}

# Don't export them as they're used in the context of other calls
AVALANCHEGO_VERSION=${AVALANCHEGO_VERSION:-$(extract_commit "$(getDepVersion github.com/ava-labs/avalanchego)")}
GINKGO_VERSION=${GINKGO_VERSION:-$(extract_commit "$(getDepVersion github.com/onsi/ginkgo/v2)")}
SUBNET_EVM_VERSION=${SUBNET_EVM_VERSION:-$(extract_commit "$(getDepVersion github.com/ava-labs/subnet-evm)")}


# Set golangci-lint version
GOLANGCI_LINT_VERSION=${GOLANGCI_LINT_VERSION:-'v1.64.5'}

# Extract the Solidity version from foundry.toml
SOLIDITY_VERSION=$(awk -F"'" '/^solc_version/ {print $2}' $ICM_CONTRACTS_PATH/foundry.toml)
EVM_VERSION=$(awk -F"'" '/^evm_version/ {print $2}' $ICM_CONTRACTS_PATH/foundry.toml)
