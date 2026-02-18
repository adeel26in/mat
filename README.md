mat

Linux cat command but multicoloured like a mat 🐱🌈

Usage: Check mat help

Installation instructions:

Download the appropriate binaries for your system and then do this (Open the terminal): cd Downloads/ && mv mat_ mat && mv mat /usr/bin # Change with the binary name

Compiling from source:

Make sure you have Go installed: [go.dev](https://go.dev/dl/)

git clone https://github.com/adeel26in/mat.git

cd mat/src

CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -buildid="

sudo mv mat /usr/bin

mat --help
