# base64 shortcuts

Some shortcuts for base64 encoding and decoding.

## Usage

- `b64 xxx` -> `echo xxx | base64`, reverse the pipe for base64 encode
- `b64d xxx` -> `echo xxx | base64 -d`, reverse the pipe for base64 decode
- `b64k`,`b64k8s`, encode the kubernetes secret data
- `b64dk`,`b64dk8s`, decode the kubernetes secret data
