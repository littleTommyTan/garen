#!/usr/bin/env bash
git clone https://github.com/golang/crypto.git vendor/golang.org/x/crypto/
git clone https://github.com/golang/net.git vendor/golang.org/x/net/
git clone https://github.com/golang/sys.git vendor/golang.org/x/sys/
git clone https://github.com/golang/time.git vendor/golang.org/x/time/
echo "go get -u github.com/kardianos/govendor"
go get -u github.com/kardianos/govendor
echo "govendor sync"
govendor sync