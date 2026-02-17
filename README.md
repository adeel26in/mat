gat

Linux cat command but multicoloured 🐱🌈

Usage: Check gat help

Installation instructions:

Download the appropriate binaries for your system and then do this (Open the terminal): cd Downloads/ && mv gat_ gat && mv gat /usr/bin # Change with the binary name

Compiling from source:

Make sure you have Go installed: [go.dev](https://go.dev/dl/)

git clone https://github.com/adeel26in/gat.git

cd gat/src

CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -buildid="

sudo mv gat /usr/bin

gat --help
