# GoMud

This is an early version of GoMud, having only been in development a couple of months.

It has been refactored on the fly, which is why some aspects might seem less than ideal.

The network layer still needs to be cleaned up, since it start with very different assumptions than where things ended up, but it works fine as it is for now.

Screenshots for some of the features can be found [here](https://imgur.com/a/90y6OGS).

Colorization is handled through extensive use of my [github.com/Volte6/ansitags](https://github.com/Volte6/ansitags) library.

Can be run locally as a standard go program or via docker container. The default port is `33333`.

There is some stubbed out folders/files/code bits for a web service and web client, but nothing substantial or even moderately functional yet. Later this should use websockets to connect, and be able to server game-aware pages up.

Certain admin in-game commands can be destructive. For example, the `build` command is notoriously finicky if you don't understand what you are doing. Although there is some documentation, it doesn't mean stuff won't get missed, plus it's possible to accidentally mess up typing something and then can be tricky to fix if you don't first understand the underlying mechanisms. Now that there is a user prompt system working this can probably be improved considerably in the near future, and building or modifying a room can be a series of prompts.

The network layer will eventually be overhauled and possibly include support for the `alternative screen buffer` mode at some point.

# Quick Start

A youtube playlist to getting started has been set up here:

[![Getting Started Videos](https://i.ytimg.com/vi/OOZqX01aHt8/hqdefault.jpg)](https://www.youtube.com/watch?v=OOZqX01aHt8&list=PL20JEmG_bxBuaOE9oFziAhAmx1pyXhQ1p)

You can compile and run it locally with:
> `go run .`

Or you can just build the binary if you prefer:
> `go build -o GoMudServer`

> `./GoMudServer`

Or if you have docker installed:
> `docker-compose -f provisioning/docker-compose.yml up --build --remove-orphans server`

From there you should see some logging, and once ready, connect to `localhost` on port `33333` with a telnet client and use the default admin login:

**Username:** _admin_

**Password:** _password_

## Makefile usage

There are a number of make targets that might be useful for building/running the MUD.

You can type `make help` to see a couple make targets worth knowing about.

_________________

### **Go Specific Makefile targets**

_These require Go to be installed locally_

Run go vendor/tidy/verify:
> `make mod`

Run in a container (port 33333):
> `make run`

Connect to running container via a container client:
> `make client`

Build the `go-mud-server` binary:
> `make build`

_________________

### **To Restart Docker Daemon**

_From powershell w/ admin priv:_

> `restart-service *docker*`
_________________

### **Dockerfile specific Makefile targets**

_These require Docker to be installed locally_

Build/Run in Docker container:

_Will run on port `33333` in the container and publicly exposes port `33333` ( per [provisioning/dockerdocker-compose.yml](dockerdocker-compose.yml) ):_

>  `make run`



Get exposed port of running container:

>  `make port`


_________________

## Connecting to server

_Connection can be made with any terminal program (telnet, nc, etc)_
>  `telnet localhost 33333`

_________________
_NOTE:_ Windows default telnet client is no longer compatible with typical ANSI Escape codes.
_________________


### **Connect using custom client** 

_This will build a lightweight linux container and use it to telnet to the server. This is useful especially on windows where ANSI color escape sequences are borked._
> `make client`

_Or using docker (localhost connection)_
>  `docker run --rm -u root -it busybox:latest telnet host.docker.internal 33333`

_Or using docker (external connection)_
>  `docker run --rm -u root -it busybox:latest telnet {hostname/ip} 33333`
>
[Some Notes](notes.md)

