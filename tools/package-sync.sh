#!/usr/bin/env bash
git clone https://github.com/golang/crypto.git vendor/golang.org/x/crypto/
git clone https://github.com/golang/net.git vendor/golang.org/x/net/
git clone https://github.com/golang/sys.git vendor/golang.org/x/sys/
go get -u github.com/kardianos/govendor
govendor sync