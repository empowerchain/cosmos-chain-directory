VERSION 0.6
FROM golang:1.16-alpine
WORKDIR /app

project-files:
    COPY go.mod go.sum .
    COPY *.go .
    COPY chain-registry ./chain-registry

build:
    FROM +project-files
    RUN go build -o dist/server
    SAVE ARTIFACT dist/server dist/server AS LOCAL dist/server

docker:
    ARG COSMOS_CHAIN_DIRECTORY_VERSION
    ARG CHAIN_REGISTRY_VERSION
    FROM alpine
    ENV COSMOS_CHAIN_DIRECTORY_VERSION $COSMOS_CHAIN_DIRECTORY_VERSION
    ENV CHAIN_REGISTRY_VERSION $CHAIN_REGISTRY_VERSION
    WORKDIR /app
    COPY +build/dist/server .
    CMD [ "/app/server" ]
    SAVE IMAGE --push empowergjermund/cosmos-chain-directory:$COSMOS_CHAIN_DIRECTORY_VERSION empowergjermund/cosmos-chain-directory:latest

deploy-to-akash:
    ARG COSMOS_CHAIN_DIRECTORY_VERSION
    FROM ubuntu:20.04
    WORKDIR /akash
    COPY deploy.yml .
    RUN sed -i 's/CHANGE_ME/'$COSMOS_CHAIN_DIRECTORY_VERSION'/' deploy.yml
    ENV AKASH_HOME=/akash/.akash
    ENV AKASH_KEY_NAME=AKASH_EARTHLY
    RUN apt-get update -yq \
            && apt-get install --no-install-recommends -yq \
            curl wget jq ca-certificates
    RUN --secret GITHUB_TOKEN wget $(curl -H "Authorization: token $GITHUB_TOKEN" -s https://api.github.com/repos/ovrclk/akash/releases/latest | jq -r ".assets[] | select(.name | test(\"linux_amd64.deb\")) | .browser_download_url") -O akash.deb
    RUN dpkg -i akash.deb
    RUN --secret AKASH_WALLET_KEY echo "$AKASH_WALLET_KEY" > key.key
    RUN --secret AKASH_WALLET_KEY_PASSWORD echo "$AKASH_WALLET_KEY_PASSWORD" | akash keys import $AKASH_KEY_NAME key.key --keyring-backend test
    RUN --secret AKASH_DEPLOY_CERTIFICATE echo "$AKASH_DEPLOY_CERTIFICATE" > $AKASH_HOME/$(akash keys show $AKASH_KEY_NAME --keyring-backend test --output=json | jq -r ".address").pem
    RUN AKASH_VERSION="$(curl -s "https://raw.githubusercontent.com/ovrclk/net/master/mainnet/version.txt")" \
        && AKASH_CHAIN_ID="$(curl -s "https://raw.githubusercontent.com/ovrclk/net/master/mainnet/chain-id.txt")" \
        && AKASH_NODE=https://rpc.akash.forbole.com:443 \
        && echo $AKASH_VERSION $AKASH_NODE $AKASH_CHAIN_ID \
        && akash tx deployment update deploy.yml -y --dseq 5342523 --from $AKASH_KEY_NAME --keyring-backend test --chain-id $AKASH_CHAIN_ID --node $AKASH_NODE --gas-prices="0.025uakt" --gas="auto" --gas-adjustment=1.5 \
        && akash provider send-manifest deploy.yml --node $AKASH_NODE --dseq 5342523 --provider akash1u5cdg7k3gl43mukca4aeultuz8x2j68mgwn28e --home $AKASH_HOME --from $AKASH_KEY_NAME --keyring-backend test
