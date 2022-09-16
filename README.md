# IPinfo cli app
The unofficial [ipinfo.io](https://ipinfo.io) cli app for IP address information.

## Usage
Get information on your actual external IP.
```shell
go run ipinfo.go
```

Get information on 8.8.8.8
```shell
go run ipinfo.go --ip 8.8.8.8
```

## Build
```shell
go build -o ./build/ipinfo ipinfo.go
```

## Usage
```shell
./build/ipinfo --ip 8.8.8.8
```

## Output
```shell
./build/ipinfo --ip 8.8.8.8

8.8.8.8 - AS15169 Google LLC
US, California, Mountain View
```