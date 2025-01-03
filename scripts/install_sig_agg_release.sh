#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

# Load the versions
ICM_CONTRACTS_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)
source "$ICM_CONTRACTS_PATH"/scripts/versions.sh
source "$ICM_CONTRACTS_PATH"/scripts/constants.sh

############################
# download icm-services
# https://github.com/ava-labs/icm-services/releases
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
BASEDIR=${BASEDIR:-"/tmp/icm-services-release"}
ICM_SERVICES_BUILD_PATH=${ICM_SERVICES_BUILD_PATH-${BASEDIR}/icm-services}

mkdir -p ${BASEDIR}

ICM_SERVICES_DOWNLOAD_URL=https://github.com/ava-labs/icm-services/releases/download/${ICM_SERVICES_VERSION}/icm-services_${ICM_SERVICES_VERSION#v}_linux_${GOARCH}.tar.gz
ICM_SERVICES_DOWNLOAD_PATH=${BASEDIR}/icm-services-linux-${GOARCH}-${ICM_SERVICES_VERSION}.tar.gz

if [[ ${GOOS} == "darwin" ]]; then
  ICM_SERVICES_DOWNLOAD_URL=https://github.com/ava-labs/icm-services/releases/download/${ICM_SERVICES_VERSION}/icm-services_${ICM_SERVICES_VERSION#v}_darwin_${GOARCH}.tar.gz
  ICM_SERVICES_DOWNLOAD_PATH=${BASEDIR}/icm-services-darwin-${GOARCH}-${ICM_SERVICES_VERSION}.tar.gz
fi

BUILD_DIR=${ICM_SERVICES_BUILD_PATH}-${ICM_SERVICES_VERSION}

extract_archive() {
  mkdir -p ${BUILD_DIR}

  if [[ ${ICM_SERVICES_DOWNLOAD_PATH} == *.tar.gz ]]; then
    tar xzvf ${ICM_SERVICES_DOWNLOAD_PATH} --directory ${BUILD_DIR}
  elif [[ ${ICM_SERVICES_DOWNLOAD_PATH} == *.zip ]]; then
    unzip ${ICM_SERVICES_DOWNLOAD_PATH} -d ${BUILD_DIR}
    mv ${BUILD_DIR}/build/* ${BUILD_DIR}
    rm -rf ${BUILD_DIR}/build/
  fi
}

# first check if we already have the archive
if [[ -f ${ICM_SERVICES_DOWNLOAD_PATH} ]]; then
  # if the download path already exists, extract and exit
  echo "found icm-services ${ICM_SERVICES_VERSION} at ${ICM_SERVICES_DOWNLOAD_PATH}"

  extract_archive
else
  # try to download the archive if it exists
  if curl -s --head --request GET ${ICM_SERVICES_DOWNLOAD_URL} | grep "302" > /dev/null; then
    echo "${ICM_SERVICES_DOWNLOAD_URL} found"
    echo "downloading to ${ICM_SERVICES_DOWNLOAD_PATH}"
    curl -L ${ICM_SERVICES_DOWNLOAD_URL} -o ${ICM_SERVICES_DOWNLOAD_PATH}

    extract_archive
  else
    # else the version is a git commit (or it's invalid)
    GIT_CLONE_URL=https://github.com/ava-labs/icm-services.git
    GIT_CLONE_PATH=${BASEDIR}/icm-services-repo/

    # check to see if the repo already exists, if not clone it
    if [[ ! -d ${GIT_CLONE_PATH} ]]; then
      echo "cloning ${GIT_CLONE_URL} to ${GIT_CLONE_PATH}"
      git clone --no-checkout ${GIT_CLONE_URL} ${GIT_CLONE_PATH}
    fi

    # check to see if the commitish exists in the repo
    WORKDIR=$(pwd)

    cd ${GIT_CLONE_PATH}

    git fetch

    echo "checking out ${ICM_SERVICES_VERSION}"

    # Try to checkout the branch. If it fails, try the commit.
    if ! git checkout "origin/${ICM_SERVICES_VERSION}" > /dev/null 2>&1; then
      if ! git checkout "${ICM_SERVICES_VERSION}" > /dev/null 2>&1; then
        # If the version is in the format of tag-commit, try to extract the commit and checkout.
        ICM_SERVICES_VERSION=$(extract_commit "${ICM_SERVICES_VERSION}")
        if ! git checkout "${ICM_SERVICES_VERSION}" > /dev/null 2>&1; then
          echo
          echo "'${ICM_SERVICES_VERSION}' is not a valid release tag, commit hash, or branch name"
          exit 1
        fi
      fi
    fi

    # initialize the submodules
    git submodule update --init --recursive

    COMMIT=$(git rev-parse HEAD)

    # use the commit hash instead of the branch name or tag
    BUILD_DIR=${ICM_SERVICES_BUILD_PATH}-${COMMIT}

    # if the build-directory doesn't exist, build icm-services
    if [[ ! -d ${BUILD_DIR} ]]; then
      echo "building icm-services ${COMMIT} to ${BUILD_DIR}"
      ./scripts/build_signature_aggregator.sh
      mkdir -p ${BUILD_DIR}
      cp ./build/signature-aggregator ${BUILD_DIR}/signature-aggregator
    fi

    cd $WORKDIR
  fi
fi

SIGNATURE_AGGREGATOR_PATH=${ICM_SERVICES_BUILD_PATH}/signature-aggregator
mkdir -p ${ICM_SERVICES_BUILD_PATH}

cp ${BUILD_DIR}/signature-aggregator ${SIGNATURE_AGGREGATOR_PATH}


echo "Installed signature-aggregator from icm-services release ${ICM_SERVICES_VERSION}"
echo "signature-aggregator Path: ${SIGNATURE_AGGREGATOR_PATH}"

