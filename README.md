# IPinfo cli app

The unofficial [ipinfo.io](https://ipinfo.io) cli app for IP address information.

Get information on your actual external IP.

```shell
go run ./cmd/ipinfo/ipinfo.go
```

Get information on 8.8.8.8

```shell
go run ./cmd/ipinfo/ipinfo.go --ip 8.8.8.8
```

If you have a token, you can specify it.

```shell
go run ./cmd/ipinfo/ipinfo.go --ip 8.8.8.8 --token xxxxxxxxxxxx
```

## Build

```shell
go build -o ./build/bin/ipinfo ./cmd/ipinfo/ipinfo.go
```

## Build and install with the go tool

```shell
go install ./cmd/ipinfo/ipinfo.go
```

## Usage

```shell
ipinfo
./build/bin/ipinfo --ip 8.8.8.8
./build/bin/ipinfo --ip 8.8.8.8 --token xxxxxxxxxxxx
```

## Output

```shell
./build/bin/ipinfo --ip 8.8.8.8 --token xxxxxxxxxxxx

8.8.8.8 - AS15169 Google LLC
US, California, Mountain View
```
