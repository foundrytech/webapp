#!/bin/bash

set -e 

sudo dnf install wget

wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz

rm go1.21.6.linux-amd64.tar.gz

{
  echo "export PATH=$PATH:/usr/local/go/bin"
  echo "export GOPATH=$HOME/go"
  echo "export PATH=$PATH:$GOPATH/bin"
} >> ~/.bashrc

# shellcheck disable=SC1090
source ~/.bashrc

go version