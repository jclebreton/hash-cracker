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
$ echo d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1 > hashes.txt 
$ echo d2rsph111lxo3twk39e169d94697bc5fc3e9da8bd17b0c23677a7583 >> hashes.txt
$ go run main.go hashes.txt crackstation-human-only.txt

INFO[0000] 4 workers                                    
Dictionary 63941069 / 63941069 [===================================] 100.00% 27s
    Hashes 2 / 2 [=================================================] 100.00% 27s
   Cracked 2 / 2 [=====================================================] 100.00%
   
$ cat output.txt
d2rsph111lxo3twka829f192f7fd38700cacdc5c645596ce3e9d09b1:qwerty1234
d2rsph111lxo3twk39e169d94697bc5fc3e9da8bd17b0c23677a7583:12345xxx
```
