authorization
=============

A small tool to manage authorized_keys written in go

## Usage examples

``` bash
# add all public keys of user tranquility that are on github to ~/.ssh/authorized_keys
$ authorization add tranquility

# remove all keys of tranquility from ~/.ssh/authorized_keys
$ authorization remove tranquility
```
