#!/usr/bin/env bash

#gometalinter.v2 --install
gometalinter.v2 --exclude vendor --deadline=60s -D aligncheck -D gotype -D dupl -D gocyclo -D vetshadow -D ineffassign -D gas -D staticcheck --min-occurrences=7 ./...
