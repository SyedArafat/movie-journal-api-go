#!/bin/sh

if [ "$HOST_OS" = "mac" ]; then
    exec /go/bin/linux_amd64/air -c .air.toml
else
    exec air -c .air.toml
fi