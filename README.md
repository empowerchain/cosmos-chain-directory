# Cosmos chain directory

This is a simple http server that exposes a simple API to the Cosmos Chain Registry (github.com/cosmos/chain-registry).

TODO: Where it is hosted (an Akash)

## API

TODO

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