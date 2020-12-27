# serve-wasm
A simple webserver to serve static files with wasm binary file

Try your wasm binary locally. There is a lack of support for serving the correct mime type for WASM binaries.

This package is here to help you with that, you can simply install it by downloading the binary from releases for any platform or compiling from source!
The package makes it easy to test your binary locally instead of resorting to various hacks. A drop in replacement, no installation needed!

I was using `http-server` node package in the past, but this does not serve wasm files with the `application/wasm` mime type.

The builtin golang types already do this for us, so this is a really minimalistic sample that I use for personal development.

This is not something groundbreaking, but makes things easier, I've seen many tutorials on how to modify python's builtin webserver to pass proper type and etc.

This is a download and use solution, nothing else needed really.

## Install
To install it you simply need to download the binary.

Note: Replace darwin, with your preferred platform, like linux for example. Windows binary also available.
```shell
curl https://github.com/0x111/serve-wasm/releases/download/v0.1/serve-wasm-darwin-amd64 -O /usr/local/bin/serve-wasm
chmod +x /usr/local/bin/serve-wasm
serve-wasm -version # prints serve-wasm v0.1
```

## Usage
The app accepts two parameters, `-path` and `-host`.

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

Now you can test your wasm binary also locally. This app is opinionated, not intended for production use and it is simply a tool for personal use but maybe someone can have a use for it!