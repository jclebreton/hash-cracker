#!/usr/bin/env bash

# Version
source ./scripts/version.sh
deb_version=$(getVersionFromGitTag)

# Build binary
./scripts/build.sh

# Build Debian package
docker build -t fpm -f fpm.Dockerfile .
docker run -ti -v $PWD:/packaging fpm fpm --verbose -s dir -t deb -n hash-cracker -v $deb_version \
  --description "hash-cracker is a tool to crack cryptographic hash function" \
  hash-cracker=/usr/bin/hash-cracker

# Create shasums files
sha256sum *.deb > SHA256SUMS
md5sum *.deb > MD5SUMS
