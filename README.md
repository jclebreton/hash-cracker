# hash-cracker [![Build Status](https://travis-ci.org/jclebreton/hash-cracker.svg?branch=master)](https://travis-ci.org/jclebreton/hash-cracker) [![codecov](https://codecov.io/gh/jclebreton/hash-cracker/branch/master/graph/badge.svg)](https://codecov.io/gh/jclebreton/hash-cracker)

*hash-cracker*  is a tool to crack *cryptographic hash function* using *Providers* and *Comparators* interfaces

### Available providers

- text file

### Available comparators

- LBC hash implementation

### Build & Run

```
# go run main.go <dictionary-path> <lbc-hash>
```

#### Example using CrackStation dictionary

```
$ wget https://crackstation.net/files/crackstation-human-only.txt.gz
$ gunzip -d crackstation-human-only.txt.gz
$ go run main.go crackstation-human-only.txt d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1

INFO[0000] cracking hash: d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1 
 510.10 MiB / 683.25 MiB [===============================================>--------------]  74.66% 19s
INFO[0056] password found                                plain=qwerty1234
```
