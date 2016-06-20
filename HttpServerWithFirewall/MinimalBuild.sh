#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o HttpServerWithFirewall .
docker build -t theminimalserver -f MinimalGoDockerfile .