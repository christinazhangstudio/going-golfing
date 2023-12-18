## md5me

```
./md5me routines=(number of reader worker threads for md5 hash'ing files in the current dir)
```
routines default = 3

see Dockerfile...

```
docker build . -t md5me
```

```
docker run -it md5me
```

...which opens /bin/bash as `CMD`


#### wonks
if get: `$GOPATH/go.mod exists but should not`,

should set `WORKDIR`
