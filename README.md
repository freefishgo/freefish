# freefish 是freefishgo的一个项目快速构建工具（https://github.com/freefishgo/freeFishGo）


# 详细文档请查看 http://freefishgo.com

## Installation

To install `freefish` use the `go get` command:

```bash
go get github.com/freefishgo/freefishgo
go get github.com/freefishgo/freefish
```

Then you can add `freefish` binary to PATH environment variable in your `~/.bashrc` or `~/.bash_profile` file:

```bash
export PATH=$PATH:<your_main_gopath>/bin
```

> If you already have `freefish` installed, updating `freefish` is simple:

```bash
go get -u github.com/freefishgo/freefishgo
go get -u github.com/freefishgo/freefish
```

## Basic commands

freefish provides a variety of commands which can be helpful at various stages of development. The top level commands include:

```bash
    -h          获得帮助文档
    new         在当前文件夹下 构建一个freefishgo mvc项目
    new -gopath 在gopath目录下的src构建一个freefishgo mvc项目

```
```bash
PS D:\> freefish -h
freefishgo version: 1.00
Usage: freefish h look help

Options:
   -gopath
        在GOPATH下创建一个新的mvc项目 如:freefish new -gopath [ProjectName]
   -h
        freeFishGo 帮助信息
   new
        创建一个新的mvc项目 如:freefish new [ProjectName]
PS D:\>
```

```bash

PS D:\> freefish new  mysqlwebhelp
2019/12/26 11:13:43 生成MVC项目:mysqlwebhelp 中.....
2019/12/26 11:13:43 C:\Users\JackShan\go\src\github.com\freefishgo\freefish\template
创建文件：D:/mysqlwebhelp/conf/app.conf 成功
创建目录:D:/mysqlwebhelp/conf/
创建文件：D:/mysqlwebhelp/conf/config.go 成功
创建文件：D:/mysqlwebhelp/controllers/homeController.go 成功
创建目录:D:/mysqlwebhelp/controllers/
创建文件：D:/mysqlwebhelp/controllers/staticController.go 成功
创建文件：D:/mysqlwebhelp/fishgo/init.go 成功
创建目录:D:/mysqlwebhelp/fishgo/
创建文件：D:/mysqlwebhelp/go.mod 成功
创建文件：D:/mysqlwebhelp/main.go 成功
创建文件：D:/mysqlwebhelp/middlewares/youMiddleware.go 成功
创建目录:D:/mysqlwebhelp/middlewares/
创建文件：D:/mysqlwebhelp/routers/router.go 成功
创建目录:D:/mysqlwebhelp/routers/
创建文件：D:/mysqlwebhelp/static/img/fish.ico 成功
创建目录:D:/mysqlwebhelp/static/
创建目录:D:/mysqlwebhelp/static/img/
创建文件：D:/mysqlwebhelp/static/img/fish.jpg 成功
创建文件：D:/mysqlwebhelp/views/Home/Index.fish 成功
创建目录:D:/mysqlwebhelp/views/
创建目录:D:/mysqlwebhelp/views/Home/
2019/12/26 11:13:43 MVC项目:mysqlwebhelp 生成成功.....请查看目录:D:\mysqlwebhelp
PS D:\>
```
```bash
PS C:\Users\JackShan\go\src\github.com\freefishgo> freefish new -gopath myfreefish
2019/12/26 13:43:44 生成MVC项目:myfreefish 中.....
2019/12/26 13:43:44 C:\Users\JackShan\go\src\github.com\freefishgo\freefish\template
创建文件：C:/Users/JackShan/go/src/myfreefish/conf/app.conf 成功
创建目录:C:/Users/JackShan/go/src/myfreefish/conf/
创建文件：C:/Users/JackShan/go/src/myfreefish/conf/config.go 成功
创建文件：C:/Users/JackShan/go/src/myfreefish/controllers/homeController.go 成功
创建目录:C:/Users/JackShan/go/src/myfreefish/controllers/
创建文件：C:/Users/JackShan/go/src/myfreefish/controllers/staticController.go 成功
创建文件：C:/Users/JackShan/go/src/myfreefish/fishgo/init.go 成功
创建目录:C:/Users/JackShan/go/src/myfreefish/fishgo/
创建文件：C:/Users/JackShan/go/src/myfreefish/go.mod 成功
创建文件：C:/Users/JackShan/go/src/myfreefish/main.go 成功
创建文件：C:/Users/JackShan/go/src/myfreefish/middlewares/youMiddleware.go 成功
创建目录:C:/Users/JackShan/go/src/myfreefish/middlewares/
创建文件：C:/Users/JackShan/go/src/myfreefish/routers/router.go 成功
创建目录:C:/Users/JackShan/go/src/myfreefish/routers/
创建文件：C:/Users/JackShan/go/src/myfreefish/static/img/fish.ico 成功
创建目录:C:/Users/JackShan/go/src/myfreefish/static/
创建目录:C:/Users/JackShan/go/src/myfreefish/static/img/
创建文件：C:/Users/JackShan/go/src/myfreefish/static/img/fish.jpg 成功
创建文件：C:/Users/JackShan/go/src/myfreefish/views/Home/Index.fish 成功
创建目录:C:/Users/JackShan/go/src/myfreefish/views/
创建目录:C:/Users/JackShan/go/src/myfreefish/views/Home/
2019/12/26 13:43:44 MVC项目:myfreefish 生成成功.....请查看目录:C:\Users\JackShan\go\src\myfreefish
PS C:\Users\JackShan\go\src\github.com\freefishgo>
```
