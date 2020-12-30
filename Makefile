binary_name = nhttp

.PHONY: all
all: linux_amd64 darwin_amd64 windows_amd64 checksums

.PHONY: linux_amd64
linux_amd64:
	GOOS=linux GOARCH=amd64 go build -v -a -gcflags=-trimpath=$$PWD -asmflags=-trimpath=$$PWD -o build/$(binary_name)-linux-amd64

.PHONY: darwin_amd64
darwin_amd64:
	GOOS=darwin GOARCH=amd64 go build -v -a -gcflags=-trimpath=$$PWD -asmflags=-trimpath=$$PWD -o build/$(binary_name)-darwin-amd64

.PHONY: windows_amd64
windows_amd64:
	GOOS=windows GOARCH=amd64 go build -v -a -gcflags=-trimpath=$$PWD -asmflags=-trimpath=$$PWD -o build/$(binary_name)-windows-amd64.exe

.PHONY: checksums
checksums:
	shasum -a 256 build/* > build/checksum.txt
