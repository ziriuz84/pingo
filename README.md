# pingo

Pingo is a simple ping tool written in Golang. Born and developed as my first experiment in Golang.

## Requirements

The only requirements is [go-ping](https://pkg.go.dev/github.com/go-ping/ping#section-readme) package. You can install it via
```go get -u github.com/go-ping/ping```

## Before usage

Before usage you must set some settings as written in go-ping documentation

### Linux

This library attempts to send an "unprivileged" ping via UDP. On Linux, this must be enabled with the following sysctl command:

```sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"```

If you do not wish to do this, you can call pinger.SetPrivileged(true) in your code and then use setcap on your binary to allow it to bind to raw sockets (or just run it as root):

```setcap cap_net_raw=+ep /path/to/your/compiled/binary```

Maybe you must build it and run as root via sudo or similar

### Windows

You must use pinger.SetPrivileged(true), otherwise you will receive the following error:

```socket: The requested protocol has not been configured into the system, or no implementation for it exists.```

Despite the method name, this should work without the need to elevate privileges and has been tested on Windows 10. Please note that accessing packet TTL values is not supported due to limitations in the Go x/net/ipv4 and x/net/ipv6 packages.

## Usage

Apart previous configuration, simply run it and follow instruction in video
