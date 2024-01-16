# base64 shortcuts

Some shortcuts for base64 encoding and decoding.

## Overview

* en/decode string
* en/decode kubernetes secret file

## Index

* [b64](#b64)
* [b64d](#b64d)
* [b64k8s](#b64k8s)
* [b64dk8s](#b64dk8s)

### b64

base64 encode with pipe

#### Example

```bash
b64 your-string
```

#### Arguments

* **$1** (string): A value to encode

#### Output on stdout

* encoded string

### b64d

base64 decode with pipe

#### Example

```bash
b64d your-string
```

#### Arguments

* **$1** (string): A value to dencode

#### Output on stdout

* decoded string

### b64k8s

encode the kubernetes secret data

#### Example

```bash
b64k input-file output-file
b64k8s input-file output-file
```

### b64dk8s

dencode the kubernetes secret data

#### Example

```bash
b64dk input-file output-file
b64dk8s input-file output-file
```

