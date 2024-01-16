# tempfile

Create temp notebook(folder) and file

## Overview

quick create file/folder

## Index

* [tmpnb](#tmpnb)
* [tmpf](#tmpf)

### tmpnb

create a temp folder to store your works

#### Example

```bash
tmpnb
```

#### Output on stdout

* You'll cd to the new temp folder

### tmpf

collect the stdin to a temp file

#### Example

```bash
echo "test: a" | tmpf
echo "test: a" | tmpf yaml
```

#### Output on stdout

* temp file path

