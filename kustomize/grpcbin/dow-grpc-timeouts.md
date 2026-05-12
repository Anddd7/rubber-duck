# DoW - Grpc Timeouts

To test how to set grpc timeouts and proxy timeouts

- <https://nginx.org/en/docs/http/ngx_http_grpc_module.html>
- <https://nginx.org/en/docs/http/ngx_http_proxy_module.html>

all default timeout is 60s, adjust to 10s for testing

**test cases**

- unary
  - 0s / 5s / 15s
- server streaming
  - 0s / 5s / 15s
  - 0s / 5s / 15s * 3
- client streaming
  - 0s / 5s / 15s
  - 0s / 5s / 15s * 3
- bidi streaming
  - 0s / 5s / 15s
  - 0s / 5s / 15s * 3

## scenario: same timeouts - 10s

```sh
grpcbin {call} --message grpc-nginx-test --host grpcbin.example.com --port 8080 \
  --headers caller=grpcbin-cli \
  --response-headers=responder=grpcbin-behind-nginx \
  --delay {delay}
```

unary / client-streaming / client-streaming * 3

- :white_check_mark: 0s pass
- :white_check_mark: 5s pass
- :rotating_light: 15s unexpected HTTP status code received from server: 504 (Gateway Timeout)

server-streaming / server-streaming * 3

- :white_check_mark: 0s pass
- :white_check_mark: 5s pass
- :rotating_light: 15s stream terminated by RST_STREAM with error code: INTERNAL_ERROR"

bidi-streaming / bidi-streaming * 3

- :white_check_mark: 0s pass
- :white_check_mark: 5s pass
- :rotating_light: 15s
  - first stream successed,
  - following stream terminated by RST_STREAM with error code: INTERNAL_ERROR"

## scenario: smaller grpc timeouts - grpc 10s, proxy 20s

same as above

## scenario: bigger grpc timeouts - grpc 20s, proxy 10s

:white_check_mark: all pass (0s / 5s / 15s)

> the grpc connections will be impacted by grpc timeouts only,

:rotating_light: same as above when set delay to 20s

> if you want longer timeouts for grpc, you have to set grpc_send_timeout/grpc_read_timeout

## Summary

default r/w timouts for proxy and grpc are 60s. so the grpc connection will be cut off if exceed 60s.

if you set longer proxy timeouts, it won't help for grpc connections.

so, we need server-snippet or new annotations to set the grpc timeouts.

### Addtional Test Case

- set proxy-read-timeout and proxy-send-timeout to 3600
- grpc connection is still cut off after 60s
