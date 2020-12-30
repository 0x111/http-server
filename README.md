# nhttp

Attention: This is a work in progress.

A simple zero-dependency webserver. I present you [nhttp](https://nhttp.org)
You can use this, to serve static files also wasm binaries with the correct mime type or use it for testing as a simple web server.
There is a lack of support for serving the correct mime type for WASM binaries. (at the time of writing of this README, for example nodejs http-server does not serve wasm binaries correctly)

This app is here to help you with that, you can simply install it by downloading the binary from releases for any platform or compiling from source!
The package makes it easy to have a webserver locally, testing your WASM binary for example instead of resorting to various hacks. I will try to implement this in par with the [http-server](https://www.npmjs.com/package/http-server) package to provide full compatibility.

This is not intended for production use, only for development and testing.

## Install
To install it you simply need to download the binary.

Note: Replace darwin, with your preferred platform, like linux for example. Windows binary also available.

### macOS
```shell
curl https://github.com/0x111/nhttp/releases/download/v0.1/nhttp-darwin-amd64 -L -s -o /usr/local/bin/nhttp
chmod +x /usr/local/bin/http-server
http-server -version # prints http-server v0.1
```

### Linux
```shell
curl https://github.com/0x111/nhttp/releases/download/v0.1/nhttp-linux-amd64 -L -s -o /usr/local/bin/nhttp
chmod +x /usr/local/bin/http-server
http-server -version # prints http-server v0.1
```


## Usage
The app right now accepts two parameters, `-path` and `-host`.

If you leave out the parameters, the defaults are listening on `localhost:8080` and serving the current directory `.`.
```shell
http-server
# Listening on localhost:8080 and serving path .
```

In any other case, you can specify a static folder and for example the port 9090 still listening on localhost.
```shell
http-server -host=localhost:9090 -path=./static
# Listening on localhost:9090 and serving path ./static
```

This is strictly opinionated and it will likely stay as-is for the foreseable future.
