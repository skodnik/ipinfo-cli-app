![integration tests](https://github.com/skodnik/ipinfo-cli-app/actions/workflows/ci.yaml/badge.svg)

# IPinfo cli app

The unofficial [ipinfo.io](https://ipinfo.io) cli app for IP address information.

## Usage

### Simple way

#### Download for Windows

- [ipinfo-windows-amd64.exe](https://github.com/skodnik/ipinfo-cli-app/raw/main/build/bin/ipinfo-windows-amd64.exe)
- [ipinfo-windows-386.exe](https://github.com/skodnik/ipinfo-cli-app/raw/main/build/bin/ipinfo-windows-386.exe)

#### Download for Linux

- [ipinfo-linux-amd64](https://github.com/skodnik/ipinfo-cli-app/raw/main/build/bin/ipinfo-linux-amd64)

#### Download for MacOS

- [ipinfo-darwin-arm64](https://github.com/skodnik/ipinfo-cli-app/raw/main/build/bin/ipinfo-darwin-arm64)
- [ipinfo-darwin-amd64](https://github.com/skodnik/ipinfo-cli-app/raw/main/build/bin/ipinfo-darwin-amd64)

### The hard way - install

```shell
git clone git@github.com:skodnik/ipinfo-cli-app.git
cd ipinfo-cli-app
make install
```

### Help

```shell
ipinfo --help

NAME:
   ipinfo - get ip information

USAGE:
   ipinfo [global options] command [command options] [arguments...]

VERSION:
   v1.0.7

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --alt, -a                use an alternative host ip.zxq.co (default: false)
   --help, -h               show help (default: false)
   --ip value, -i value     ip to search
   --json, -j               result to json (default: false)
   --pretty, -p             prettier json (default: false)
   --sly, -s                rich info without token (only main host) (default: false)
   --token value, -t value  access token
   --version, -v            print the version (default: false)
```

### Usage and output

Short-hand.

```shell
ipinfo --ip 8.8.8.8 --token xxxxxxxxxxxx

8.8.8.8 - AS15169 Google LLC
US, California, Mountain View
```

Full information.

```shell
ipinfo --ip 8.8.8.8 --json --pretty
  
{
    "ip": "8.8.8.8",
    "hostname": "dns.google",
    "anycast": true,
    "city": "Mountain View",
    "region": "California",
    "country": "US",
    "loc": "37.4056,-122.0775",
    "org": "AS15169 Google LLC",
    "postal": "94043",
    "timezone": "America/Los_Angeles",
    "readme": "https://ipinfo.io/missingauth"
}
```

Rich full information with `--sly` flag (**experimental option**).

```shell
ipinfo --ip 8.8.8.8 --sly --json -pretty

{
    "ip": "8.8.8.8",
    "hostname": "dns.google",
    "anycast": true,
    "city": "Mountain View",
    "region": "California",
    "country": "US",
    "loc": "37.4056,-122.0775",
    "org": "AS15169 Google LLC",
    "postal": "94043",
    "timezone": "America/Los_Angeles",
    "asn": {
        "asn": "AS15169",
        "name": "Google LLC",
        "domain": "google.com",
        "route": "8.8.8.0/24",
        "type": "hosting"
    },
    "company": {
        "name": "Google LLC",
        "domain": "google.com",
        "type": "hosting"
    },
    "privacy": {
        "vpn": false,
        "proxy": false,
        "tor": false,
        "relay": false,
        "hosting": true,
        "service": ""
    },
    "abuse": {
        "address": "US, CA, Mountain View, 1600 Amphitheatre Parkway, 94043",
        "country": "US",
        "email": "network-abuse@google.com",
        "name": "Abuse",
        "network": "8.8.8.0/24",
        "phone": "+1-650-253-0000"
    },
    "domains": {
        "ip": "8.8.8.8",
        "total": 13685,
        "domains": [
            "aonode.com",
            "000180.top",
            "chaoxz.com",
            "dns.google",
            "server-panel.net"
        ]
    }
}
```

## Developers section

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
