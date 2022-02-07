# nhttp

Attention: This is a work in progress.

A simple zero-configuration and zero-dependency webserver. I present you [nhttp](https://nhttp.org)
You do not need to install any packages except of downloading one binary and using it right away!
nhttp uses the internal net/http for serving the files. Other than that, we rely on 3rd party implementations for compressions.
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
curl https://github.com/0x111/nhttp/releases/download/v0.1.1/nhttp-darwin-amd64 -L -s -o /usr/local/bin/nhttp
chmod +x /usr/local/bin/nhttp
nhttp -version # prints nhttp v0.1.1
```

### Linux
```shell
curl https://github.com/0x111/nhttp/releases/download/v0.1.1/nhttp-linux-amd64 -L -s -o /usr/local/bin/nhttp
chmod +x /usr/local/bin/nhttp
nhttp -version # prints nhttp v0.1.1
```

### Windows
Download the latest version of the exe from the [Releases](https://github.com/0x111/nhttp/releases) tab.

## Usage
The app right now accepts the parameters below.

`-gzip` if set, it will serve content with a gzip encoding, turned off by default

`-host` if set, you can set the listen host, default is `localhost`

`-port` you can set the port to listen on, default is `8080`

`-spa` if set, the app will host index.html whenever no file is found. You can set your own path of the static directory and the index.html file within it.

`-staticPath` a path to the static folder, default is the current directory i.e. empty

`-indexPath` a path to the index file, within the set `-staticPath`, default is `index.html`

`-version` prints version information

If you leave out the parameters, the defaults are listening on `localhost:8080` and serving the current directory `.`.
```shell
nhttp
# Listening on localhost:8080 and serving path .
```

In any other case, you can specify a static folder and for example the port 9090 still listening on localhost.
```shell
nhttp -host localhost -port 9090 -path=./static
# Listening on localhost:9090 and serving path ./static
```

This is strictly opinionated for now and it will likely stay as-is for the foreseable future.
