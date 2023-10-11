# Near Validator Watcher

[![License](https://img.shields.io/badge/license-MIT-blue)](https://opensource.org/licenses/MIT)

**Near Validator Watcher** is a Prometheus exporter to help you monitor missed blocks & chunks of your validator.

## ✨ Usage

Example on mainnet using the default RPC endpoint.

### Via compiled binary

Compiled binary can be found on the [Releases page](https://github.com/kilnfi/near-validator-watcher/releases).

```bash
near-validator-watcher \
  --node https://rpc.mainnet.near.org \
  --validator kiln-1.poolv1.near
```

### Via Docker

Latest Docker image can be found on the [Packages page](https://github.com/kilnfi/near-validator-watcher/pkgs/container/near-validator-watcher).

```bash
docker run --rm ghcr.io/kilnfi/near-validator-watcher:latest \
  --node https://rpc.mainnet.near.org \
  --validator kiln-1.poolv1.near
```

### Available options

```
near-validator-watcher --help

NAME:
   near-validator-watcher - NEAR validators monitoring tool

USAGE:
   near-validator-watcher [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --http-addr value                        http server address (default: ":8080")
   --log-level value                        log level (debug, info, warn, error) (default: "info")
   --namespace value                        prefix for Prometheus metrics (default: "near_validator_watcher")
   --no-color                               disable colored output (default: false)
   --node value                             rpc node endpoint to connect to (default: "https://rpc.mainnet.near.org")
   --refresh-rate value                     how often to call the rpc endpoint (default: 10s)
   --validator value [ --validator value ]  validator pool id to track
   --help, -h                               show help
   --version, -v                            print the version
```


## ❇️ Endpoints

- `/metrics` exposed Prometheus metrics (see next section)
- `/ready` responds OK when at the node is synced
- `/live` responds OK as soon as server is up & running correctly


## 📊 Prometheus metrics

All metrics are by default prefixed by `near_validator_watcher` but this can be changed through options.

Metrics (without prefix)    | Description
----------------------------|-------------------------------------------------------------------------
`block_number`              | The number of most recent block
`chain_id`                  | Near chain id
`current_proposals_stake`   | Current proposals
`epoch_length`              | Near epoch length as specified in the protocol
`epoch_start_height`        | Near epoch start height
`next_validator_stake`      | The next validators
`prev_epoch_kickout`        | Near previous epoch kicked out validators
`protocol_version`          | Current protocol version deployed to the blockchain
`seat_price`                | Validator seat price
`sync_state`                | Sync state
`validator_blocks_expected` | Current amount of validator expected blocks
`validator_blocks_produced` | Current amount of validator produced blocks
`validator_chunks_expected` | Current amount of validator expected chunks
`validator_chunks_produced` | Current amount of validator produced chunks
`validator_slashed`         | Validators slashed
`validator_stake`           | Current amount of validator stake
`version_build`             | The Near node version build


## 📃 License

[MIT License](LICENSE).
