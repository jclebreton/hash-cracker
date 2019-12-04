# hash-cracker [![Build Status](https://travis-ci.org/jclebreton/hash-cracker.svg?branch=master)](https://travis-ci.org/jclebreton/hash-cracker) [![Go Report Card](https://goreportcard.com/badge/github.com/jclebreton/hash-cracker)](https://goreportcard.com/report/github.com/jclebreton/hash-cracker) [![GoDoc](https://godoc.org/github.com/jclebreton/hash-cracker?status.svg)](https://godoc.org/github.com/jclebreton/hash-cracker)

*hash-cracker*  is a tool to crack *cryptographic hash function* using *dictionaries*
and *hashers* interfaces


### Available dictionaries

- text file (one by line)
- text file (with passwords generation)

### Available hashers

- $salt.sha1($salt.$pass)

### Build & Run

```
$ go install
$ hash-cracker <hash-path> <dictionary-path> --generate
```

#### Example using dictionary example

```
$ hash-cracker examples/hashes.txt examples/dico-passwords.txt --generate

INFO[0000] 8 logical CPUs                                
INFO[0000] passwords dictionary generation enable                
Dictionary 1300 / 1300 [===============================================================] 100.00% 0s
INFO[0000] 8 workers                                    
worker 1 10 / 163 [=====>-----------------------------------------------------------------]   6.13%
worker 2 163 / 163 [======================================================================] 100.00%
worker 3 163 / 163 [======================================================================] 100.00%
worker 4 106 / 163 [=====================================>--------------------------------]  65.03%
worker 5 134 / 162 [=====================================================>----------------]  82.72%
worker 6 162 / 162 [======================================================================] 100.00%
worker 7 162 / 162 [======================================================================] 100.00%
worker 8 162 / 162 [======================================================================] 100.00%
Hashes 2 / 2 [=========================================================================] 100.00% 1s
Cracked 2 / 2 [===========================================================================] 100.00%   
INFO[0000] finish in 1.884776ms

$ cat error.txt output.txt
d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1    qwerty1234
d2rsph111lxo3twk39e169d94697bc5fc3e9da8bd17b0c23677a7583    12345xxx
```
