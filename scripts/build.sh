#!/usr/bin/env zsh

mkdir $HOME/.mgh/
cp ./configs/config.yaml $HOME/.mgh/config.yaml
go build -o /usr/local/bin/mgh -ldflags "-X 'main.GitTag=$(git describe --tags)' -X 'main.Timestamp=$(date -u)'"