# http-server

Attention: This is a work in progress.

A simple zero-dependency webserver to serve static files also wasm binaries with the correct mime type.
There is a lack of support for serving the correct mime type for WASM binaries.

This package is here to help you with that, you can simply install it by downloading the binary from releases for any platform or compiling from source!
The package makes it easy to test your binary locally instead of resorting to various hacks. I will try to implement this in par with the [http-server](https://www.npmjs.com/package/http-server) package to provide full compatibility.

This is not intended for production use, only for development and testing.

## Install
To install it you simply need to download the binary.

Note: Replace darwin, with your preferred platform, like linux for example. Windows binary also available.
```shell
curl https://github.com/0x111/serve-wasm/releases/download/v0.1/serve-wasm-darwin-amd64 -L -s -o /usr/local/bin/serve-wasm
chmod +x /usr/local/bin/serve-wasm
serve-wasm -version # prints serve-wasm v0.1
```

## Usage
The app right now accepts two parameters, `-path` and `-host`.

If you leave out the parameters, the defaults are listening on `localhost:8080` and serving the current directory `.`.
```shell
serve-wasm 
# Listening on localhost:8080 and serving path .
```

In any other case, you can specify a static folder and for example the port 9090 still listening on localhost.
```shell
serve-wasm -host=localhost:9090 -path=./static
# Listening on localhost:9090 and serving path ./static
```

This is strictly opinionated and it will likely stay as-is for the foreseable future.
