# ICM Contracts

For help getting started with building ICM contracts, refer to [the avalanche-starter-kit repository](https://github.com/ava-labs/avalanche-starter-kit).

- [Setup](#setup)
  - [Initialize the repository](#initialize-the-repository)
  - [Dependencies](#dependencies)
- [Structure](#structure)
- [E2E tests](#e2e-tests)
  - [Run specific E2E tests](#run-specific-e2e-tests)
- [ABI Bindings](#abi-bindings)
- [Docs](#docs)
- [Resources](#resources)

## Setup

### Initialize the repository

- Get all submodules: `git submodule update --init --recursive`

### Dependencies

- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo) for running the end-to-end tests.
- [Foundry](https://book.getfoundry.sh/) Use `./scripts/install_foundry.sh` to install Foundry for building contracts.

## Structure

- `contracts/`
  - [`governance/`](./contracts/governance/README.md) includes contracts related to L1 governance.
  - [`ictt/`](./contracts/ictt/README.md) Interchain Token Transfer contracts. Facilitates the transfer of tokens among L1s.
  - [`teleporter/`](./contracts/teleporter/README.md) includes `TeleporterMessenger`, which serves as the interface for most contracts to use ICM.
    - [`registry/`](./contracts/teleporter/registry/README.md) includes a registry contract for managing different versions of `TeleporterMessenger`.
  - [`validator-manager/`](./contracts/validator-manager/README.md) includes contracts for managing the validator set of an L1.
- `abi-bindings/` includes Go ABI bindings for the contracts in `contracts/`.
- [`audits/`](./audits/README.md) includes all audits conducted on contracts in this repository.
- `tests/` includes integration tests for the contracts in `contracts/`, written using the [Ginkgo](https://onsi.github.io/ginkgo/) testing framework.
- `utils/` includes Go utility functions for interacting with the contracts in `contracts/`. Included are Golang scripts to derive the expected EVM contract address deployed from a given EOA at a specific nonce, and also construct a transaction to deploy provided byte code to the same address on any EVM chain using [Nick's method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c#).
- `scripts/` includes bash scripts for interacting with TeleporterMessenger in various environments, as well as utility scripts.
  - `abi_bindings.sh` generates ABI bindings for the contracts in `contracts/` and outputs them to `abi-bindings/`.
  - `lint.sh` performs Solidity and Golang linting.

## E2E tests

In addition to the docker setup, end-to-end integration tests written using Ginkgo are provided in the `tests/` directory. E2E tests are run as part of CI, but can also be run locally. Any new features or cross-chain example applications checked into the repository should be accompanied by an end-to-end tests. See the [Contribution Guide](./CONTRIBUTING.md) for additional details.

To run the E2E tests locally, you'll need to install Gingko following the instructions [here](https://onsi.github.io/ginkgo/#installing-ginkgo).

Then run the following command from the root of the repository:

```bash
./scripts/e2e_test.sh
```

### Run specific E2E tests

To run a specific E2E test, specify the environment variable `GINKGO_FOCUS`, which will then look for test descriptions that match the provided input. For example, to run the `Calculate Teleporter message IDs` test:

```bash
GINKGO_FOCUS="Calculate Teleporter message IDs" ./scripts/e2e_test.sh
```

A substring of the full test description can be used as well:

```bash
GINKGO_FOCUS="Calculate Teleporter" ./scripts/e2e_test.sh
```

The E2E test script also supports a `--components` flag, making it easy to run all the test cases for a particular project. For example, to run all E2E tests for the `tests/flows/ictt/` folder:

```bash
./scripts/e2e_test.sh --components "ictt"
```

## ABI Bindings

The E2E tests written in Golang interface with the solidity contracts by use of generated ABI bindings. To regenerate Golang ABI bindings for the Solidity smart contracts, run:

```bash
./scripts/abi_bindings.sh
```

The auto-generated bindings should be written under the `abi-bindings/` directory.

## Docs

- [ICM Protocol Overview](./contracts/teleporter/README.md)
- [Teleporter Registry and Upgrades](./contracts/teleporter/registry/README.md)
- [Contract Deployment](./utils/contract-deployment/README.md)
- [Teleporter CLI](./cmd/teleporter-cli/README.md)

## Resources

- List of blockchain signing cryptography algorithms [here](http://ethanfast.com/top-crypto.html).
- Background on stateful precompiles [here](https://medium.com/avalancheavax/customizing-the-evm-with-stateful-precompiles-f44a34f39efd).
- Background on BLS signature aggregation [here](https://crypto.stanford.edu/~dabo/pubs/papers/BLSmultisig.html).
