# Senkyou
Senkyou provides an Ethereum RPC gateway over message broker systems such as Kafka.

Meaning: Senkyou is a Japanese word for boat's bridge ( 桟橋 )

* [Install](#install)
* [Examples](#examples)

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u https://github.com/abdelhamidbakhta/senkyou
```

## Usage

```
Usage:
  senkyou [flags]

Flags:
      --broker-type string           message broker type (nats, kafka) (default "nats")
  -h, --help                         help for senkyou
      --http-enabled                 start http server for administration
      --http-port int                http port (default 8080)
      --kafka-url string             kafka bootstrap server (default "127.0.0.1:9092")
      --logging logLevel             log level (DEBUG, INFO, WARN, ERROR) (default DEBUG)
      --nats-url string              nats server url (default "nats://127.0.0.1:4222")
      --rpc-url string               ethereum rpc url (default "http://127.0.0.1:8545")
      --topic-errors string          topic to use for error handling (default "errors")
      --topic-rpc-requests string    topic to use for receiving incoming RPC requests (default "rpc.request")
      --topic-rpc-responses string   topic to use for pushing RPC responses (default "rpc.response")
```

## Examples

Run NATS server.
```sh 
docker-compose -f build/package/nats/docker-compose.yml up -d
```

Run Senkyou.
```sh
senkyou \
--topic-rpc-requests=ethereum.rpc.requests \
--topic-rpc-responses=ethereum.rpc.responses \ 
--topic-errors=senkyou.errors \
--http-enabled --http-port=9000 \
--rpc-url=127.0.0.1:8545 \
--nats-url=nats://127.0.0.1:4222 --broker-type=nats
```
