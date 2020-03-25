## Black Hat Go -- Go Programming for Hackers and Pentesters

![Black Hat Go](https://cdn.jsdelivr.net/gh/TesterCC/pic_bed/img/20200325234257.png)

### 关键知识笔记

P6-7  
```
go run执行后不会生成独立的二进制文件，要用下面的命令生成可执行的二进制文件：
go build xxx.go
./xxx

默认go build包含debug信息和符号表，减少生成二进制文件大小约30%的方法
go build -ldflags "-w -s" xxx.go

交叉编译：指定os和架构
go build编译时可以通过指定GOOS和GOARCH参数来编译指定系统和架构的可执行文件

e.g.:

GOOS="linux" GOARCH="amd64" go build c1_demo.go

PS:Go的交叉编译相对其他现代编程语言要简单很多。

file检测：
c1_demo: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped


用file命令查看可执行文件的类型
e.g.:
file c1_demo
c1_demo: Mach-O 64-bit executable x86_64

```

P8-9  

go doc 命令使用

go doc的文档内容源自源代码中的注释
```
-> go doc fmt.Println

warning: pattern "all" matched no module dependencies
package fmt // import "fmt"

func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.
```


go get 命令使用

go get 命定用于安装go的第三方包

e.g.:
```
go get github.com/stacktitan/ldapauth
```

go fmt 自动格式化源代码，现在很多IDE自带格式自动整理了，所以用得不多
```
go fmt ./c1_demo.go 
```