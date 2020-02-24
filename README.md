# Goal
To allow you to connect to tcp socket and have that redirected to another tcp socket by way of a unix domain socket.

```
 +-----+         +-------------------+         +-----+
 | tcp | <---->  | unix domainsocket | <---->  | tcp |
 +-----+         +-------------------+         +-----+
```

## Components

The proxy has two components

1. serve
    1. listen to a tcp port and proxying the connection to a unix domain socket
1. forward
    1. listen to a unix domain socket and forwarding to the desired tcp socket

## Usage

Lets say you have an existing app on port 80 and you want to expose it on port 9999

launch the forwarder

```txt
.\domain-socket-proxy.exe forward --address :80 --path .\unix.socket
```

launch the server

```txt
.\domain-socket-proxy.exe serve --address localhost:9999 --path .\unix.socket
```

The above example uses the file path `.\unix.socket` for the unix domain socket

## Container
You can use a container for the serve as well

```txt
docker run -v c:\tmp:c:\tmp --rm --name socket knicknic/domain-socket-proxy:v1.0.0 --path c:\tmp\unix.socket --address :9999
```
