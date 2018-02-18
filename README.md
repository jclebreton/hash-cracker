# hash-cracker [![Build Status](https://travis-ci.org/jclebreton/hash-cracker.svg?branch=master)](https://travis-ci.org/jclebreton/hash-cracker)

*hash-cracker*  is a tool to crack *cryptographic hash function* using *dictionaries*
and *hashers* interfaces


### Available dictionaries

- text file
- text file with randomization

### Available hashers

- $salt.sha1($salt.$pass)

### Build & Run

```
$ dep ensure
$ go install
$ hash-cracker <hash-path> <dictionary-path> --random
```

#### Example using dictionary example

```
$ hash-cracker examples/hashes.txt examples/dico-passwords.txt --random

INFO[0000] 8 logical cpu                                
INFO[0000] random dictionary mode enable                
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

$ cat error.txt output.txt
d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1:qwerty1234
d2rsph111lxo3twk39e169d94697bc5fc3e9da8bd17b0c23677a7583:12345xxx
```
