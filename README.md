# Cosmos chain directory

This is a simple http server that exposes a simple API to the Cosmos Chain Registry (github.com/cosmos/chain-registry).

## API

The API is hosted on https://cosmos-chain.directory (With Akash, see deployment workflow and Earthfile if interested)

### Endpoints
- GET https://cosmos-chain.directory/chains/ (list of all registered chains in the directory)
- GET https://cosmos-chain.directory/chains/{chainName}/ (specific chain info)
- GET https://cosmos-chain.directory/version (version of the api + version of the registry (in the form of git hash))

### Example outputs:

https://cosmos-chain.directory/chains
```json
{
    "chains": ["agoric","akash","arkh", "etc"]
} 
```

https://cosmos-chain.directory/chains/cosmoshub
```json
{
    "$schema": "../chain.schema.json",
    "chain_name": "cosmoshub",
    "chain_id": "cosmoshub-4",
    "pretty_name": "Cosmos Hub",
    "status": "live",
    "network_type": "mainnet",
    "bech32_prefix": "cosmos",
    "genesis": {
        "genesis_url": "https://github.com/cosmos/mainnet/raw/master/genesis.cosmoshub-4.json.gz"
    },
    "daemon_name": "gaiad",
    "node_home": "$HOME/.gaia",
    "key_algos": [
        "secp256k1"
    ],
    "slip44": 118,
    "fees": {
        "fee_tokens": [
            {
                "denom": "uatom",
                "fixed_min_gas_price": 0
            }
        ]
    },
    "codebase": {
        "git_repo": "https://github.com/cosmos/gaia",
        "recommended_version": "v6.0.4",
        "compatible_versions": [
            "v6.0.0",
            "v6.0.4"
        ],
        "binaries": {
            "linux/amd64": "https://github.com/cosmos/gaia/releases/download/v6.0.4/gaiad-v6.0.4-linux-amd64",
            "linux/arm64": "https://github.com/cosmos/gaia/releases/download/v6.0.4/gaiad-v6.0.4-linux-arm64",
            "darwin/amd64": "https://github.com/cosmos/gaia/releases/download/v6.0.4/gaiad-v6.0.4-darwin-amd64",
            "windows/amd64": "https://github.com/cosmos/gaia/releases/download/v6.0.4/gaiad-v6.0.4-windows-amd64.exe"
        }
    },
    "peers": {
        "seeds": [
            {
                "id": "bf8328b66dceb4987e5cd94430af66045e59899f",
                "address": "public-seed.cosmos.vitwit.com:26656",
                "provider": "vitwit"
            }
        ],
        "persistent_peers": [
            {
                "id": "ee27245d88c632a556cf72cc7f3587380c09b469",
                "address": "45.79.249.253:26656"
            }
        ]
    },
    "apis": {
        "rpc": [
            {
                "address": "https://rpc-cosmoshub.blockapsis.com",
                "provider": "chainapsis"
            }
        ],
        "rest": [
            {
                "address": "https://lcd-cosmoshub.blockapsis.com",
                "provider": "chainapsis"
            }
        ],
        "grpc": [
            {
                "address": "cosmoshub.strange.love:9090",
                "provider": "strangelove"
            }
        ]
    },
    "explorers": [
        {
            "kind": "mintscan",
            "url": "https://www.mintscan.io/cosmos",
            "tx_page": "https://www.mintscan.io/cosmos/txs/${txHash}"
        }
    ]
}
```

https://cosmos-chain.directory/version
```json 
{
    "cosmosChainDirectoryVersion": "da4eb1e",
    "chainRegistryVersion": "8ffcb2d"
}
```

## Build

To build you first need to get the chain-registry itself (which is from github.com/cosmos/chain-registry):
```shell
$ git submodule update --init --recursive
```

You can of course build with Go directly, but a more reproducible build is provided with the Earthfile using Earthly (think Makefile + Docker).

Download and install Earthly here: https://earthly.dev/get-earthly

Then you can simply run:
```shell
$ earthly +build
```

The Earthfile also has target for a Dockerfile