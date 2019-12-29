# freefish 是freefishgo的一个项目快速构建工具（https://github.com/freefishgo/freeFishGo ）


## 详细文档请查看 http://freefishgo.com

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
C:\Users\huzhouyu>freefish -h
freefishgo version: 1.00
Usage: freefish h look help

Options:
   -c
        在freefish生成的项目中操作视图 具体命令有:
        freefish -c [controllerName] ：在controllers文件夹下生成 controllerName+"Controller" 控制器
   -h
        freeFishGo 帮助信息
   -v
        在freefish生成的项目中操作视图 具体命令有:
        freefish -v check ：检查Mvc视图文件是否存在，打印缺视图的控制器和视图
        freefish -v create：遍历Mvc控制器文件，创建缺失的视图
   new
        创建一个新的mvc项目 具体有:
        freefish new [ProjectName]                :在当前目录下创建mvc项目
        freefish new [ProjectName] -path [dirPath]:在当前目录dirPath创建mvc项目
        freefish new -gopath [ProjectName]        :在GOPATH下创建一个新的mvc项目

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
>在当前目录下构建一个叫mysqlwebhelp的项目
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
>在gopath目录下构建一个叫mysqlwebhelp的项目
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
>在当前项目创建一个名叫create的控制器
```bash
C:\Users\huzhouyu\go\src\freefishgodoc>freefish -c create
2019/12/28 19:39:57 freeFish:->Controller:createController创建成功,文件地址为:controllers\createController.go
```
>检查是否有使用了的视图 没有创建
```bash
C:\Users\huzhouyu\go\src\freefishgodoc>freefish -v check
2019/12/28 19:45:21 freeFish:->路径:controllers\createController.go Controller:createController Action:Index 缺失视图:create\Index.fish 行号:16
```
>创建使用了没有创建的视图
```bash
C:\Users\huzhouyu\go\src\freefishgodoc>freefish -v create
2019/12/28 19:47:19 freeFish:->路径:controllers\createController.go Controller:createController Action:Index 成功创建视图:create\Index.fish 行号:16
```

