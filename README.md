# redis
golang redis client / go语言实现的redis客户端

[![Build Status](https://travis-ci.org/chyroc/redis.svg?branch=master)](https://travis-ci.org/chyroc/redis)
[![codecov](https://codecov.io/gh/chyroc/redis/branch/master/graph/badge.svg)](https://codecov.io/gh/chyroc/redis)

[godoc文档](https://godoc.org/github.com/chyroc/redis)

## TODO
* [x] connect
* [x] key
* [x] string
* [x] hash
* [x] list
* [x] set
* [x] sorted set
* [x] hyper log log
* [x] geo
* [x] pub / sub
* [x] transaction
* [ ] script
* [ ] server

## test

### run redis
```
docker run -p 6379:6379 -d redis # instance 1
docker run -p 7777:6379 -d redis # instance 2
```

### run test
```
./test.sh
```
