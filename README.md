# IPinfo cli app

The unofficial [ipinfo.io](https://ipinfo.io) cli app for IP address information.

Get information on your actual external IP.

```shell
go run ipinfo.go
```

Get information on 8.8.8.8

```shell
go run ipinfo.go --ip 8.8.8.8
```

If you have a token, you can specify it.

```shell
go run ipinfo.go --ip 8.8.8.8 --token xxxxxxxxxxxx
```

## Build

```shell
go build -o ./build/ipinfo ipinfo.go
```

## Usage

```shell
./build/ipinfo
./build/ipinfo --ip 8.8.8.8
./build/ipinfo --ip 8.8.8.8 --token xxxxxxxxxxxx
```

## Output

```shell
./build/ipinfo --ip 8.8.8.8 --token xxxxxxxxxxxx

8.8.8.8 - AS15169 Google LLC
US, California, Mountain View
```
