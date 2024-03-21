# Kubectl Plugins

Some simple plugins to extend kubectl functionality.

- tail: cosolidate logs from multiple pods
- ...

## Usage

```shell
k tail --tail 10        # then select pod from list
```

## Completion

There are 2 kinds of completion

- zsh completion
  - if you want completion when you execute the command directly in terminal, you need terminal completion for zsh or bash
  - you can create fabulous completion with zsh completion functions, like compdef, _arguments,_describe, etc.
- kubectl completion
  - if you want to enable completion when you execute it as kubectl plugin, you need to follow kubectl's specification
  - create a script `kubectl_complete-<plugin>` and add into path
