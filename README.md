# Senkyou
Senkyou provides an Ethereum RPC gateway over message broker systems such as Kafka.

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
      --broker-type string   message broker type (nats, kafka) (default "nats")
  -h, --help                 help for senkyou
      --http-enabled         start http server for administration
      --http-port int        http port (default 8080)
      --kafka-url string     kafka bootstrap server (default "127.0.0.1:9092")
      --nats-url string      nats server url (default "nats://127.0.0.1:4222")
      --rpc-url string       ethereum rpc url (default "127.0.0.1:8545")
```

## Examples
```sh
senkyou --http-enabled --http-port=9000 \
--rpc-url=127.0.0.1:8545 \
--nats-url=nats://127.0.0.1:4222 --broker-type=nats
```
