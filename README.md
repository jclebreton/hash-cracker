# hash-cracker [![Build Status](https://travis-ci.org/jclebreton/hash-cracker.svg?branch=master)](https://travis-ci.org/jclebreton/hash-cracker)

*hash-cracker*  is a tool to crack *cryptographic hash function* using *dictionaries* and *hashers* interfaces

### Available dictionaries

- text file

### Available hashers

- $salt.sha1($salt.$pass)

### Build & Run

```
$ dep ensure
$ go run main.go <hash-path> <dictionary-path>
```

#### Example using CrackStation dictionary

```
$ wget https://crackstation.net/files/crackstation-human-only.txt.gz
$ gunzip -d crackstation-human-only.txt.gz
$ go run main.go hashes.txt crackstation-human-only.txt

INFO[0000] 4 workers                                    
Dictionary 26 / 26 [=====================================================================] 100.00% 1s
    Hashes 11 / 11 [=====================================================================] 100.00% 1s
   Cracked 1 / 11 [===============>------------------------------------------------------]   9.09%
   
$ cat output.txt
d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1:qwerty1234
```
