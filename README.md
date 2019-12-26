# freefish 是freefishgo的一个项目构建工具（https://github.com/freefishgo/freeFishGo）

## Installation

To install `freefish` use the `go get` command:

```bash
go get github.com/freefishgo/freefish
```

Then you can add `freefish` binary to PATH environment variable in your `~/.bashrc` or `~/.bash_profile` file:

```bash
export PATH=$PATH:<your_main_gopath>/bin
```

> If you already have `freefish` installed, updating `freefish` is simple:

```bash
go get -u github.com/freefishgo/freefish
```

## Basic commands

freefish provides a variety of commands which can be helpful at various stages of development. The top level commands include:

```
    -h          获得帮助文档
    new         在当前文件夹下 构建一个freefishgo mvc项目
    new -gopath 在gopath目录下的src构建一个freefishgo mvc项目

`
