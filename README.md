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
  -h, --help              help for senkyou
      --http-enabled      start http server for administration
      --http-port int     kafka bootstrap server (default is 127.0.0.1:9092) (default 8080)
      --kafkaUrl string   kafka bootstrap server (default is 127.0.0.1:9092) (default "127.0.0.1:9092")
```

## Examples
```sh
senkyou --http-enabled --http-port=9000
```
