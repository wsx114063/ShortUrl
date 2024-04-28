#!/bin/bash
apt-get update
apt-get install wget sudo

GO_VERSION="1.22.2"
GO_DOWNLOAD_URL="https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz"

wget -q "${GO_DOWNLOAD_URL}" -O /tmp/go.tar.gz
sudo tar -C /usr/local -xzf /tmp/go.tar.gz

go version

go build -o gobuild