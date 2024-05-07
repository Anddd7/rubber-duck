# DoW - Grpc TLS

use `kustomization/overlays/tls` to deploy grpcbin with tls enabled.

## GRPCS

First, if you enabled `GRPCS` for grpcbin, nginx will take care of the tls connection between the backend and nginx.  

> client --<http/grpc>--> nginx --<https/grpcs>--> backend

So you can use a non-tls client to connect to a tls server.

```sh
grpcbin unary --message hello --host=grpcbin.example.com --port=8080

# ---------------------------------------------------
Request Headers
- :authority=grpcbin.example.com
- content-type=application/grpc
- user-agent=grpc-go/1.63.2
- x-forwarded-for=127.0.0.1
- x-forwarded-host=grpcbin.example.com
- x-forwarded-port=80
- x-forwarded-proto=http
- x-forwarded-scheme=http
- x-real-ip=127.0.0.1
- x-request-id=a8408bb28c5699d47b5ff8e3f7e280c1
- x-scheme=http
```

### With Ingress TLS

Second, you still have a chance to enable tls between client and nginx - enable tls in ingress.

> client --<https/grpcs>--> nginx --<https/grpcs>--> backend

But they are not the same tls connection, nginx will proxy the request to the backend with a new tls connection.

- backend is using the built-in cert of grpcbin.
- nginx is using a self-signed cert with the `CN=grpcbin.example.com`.

```sh
grpcbin unary --message hello --host=grpcbin.example.com --port=8443 --tls-cert=certs/tls.crt

# ---------------------------------------------------
Request Headers
- :authority=grpcbin.example.com
- content-type=application/grpc
- user-agent=grpc-go/1.63.2
- x-forwarded-for=127.0.0.1
- x-forwarded-host=grpcbin.example.com
- x-forwarded-port=443
- x-forwarded-proto=https
- x-forwarded-scheme=https
- x-real-ip=127.0.0.1
- x-request-id=3e2d44cdc65d891fada3f720260186c3
- x-scheme=https
```
