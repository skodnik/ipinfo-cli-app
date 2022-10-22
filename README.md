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

## Integration test

```shell
make test
```

## Build

```shell
make build
```

## Build and install with the go tool

```shell
make install
```

This command builds the `ipinfo` command, producing an executable binary. It then installs that binary as
`$HOME/go/bin/ipinfo` (or, under Windows, `%USERPROFILE%\go\bin\hello.exe`).

You can use the go env command to portably set the default value for an environment variable for future go commands:

```shell
go env -w GOBIN=/somewhere/else/bin
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
