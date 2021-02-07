# http counter

HttpCounter is a web server that on each request returns a number of requests for a specific time window.

## How to build

In project directory execute

```sh
go build
```

## How to use

To start the server with the default values (window=60 sec, port=8080) execute

```sh
./httpcounter
```

To specify a different params set environment variables `TS_COUNT_WINDOW_SEC` and `TS_PORT`, for example

```sh
TS_PORT=8181 TS_COUNT_WINDOW_SEC=10 ./httpcounter
```

Now it's possible to query the localhost like this

```sh
curl localhost:8181
```
