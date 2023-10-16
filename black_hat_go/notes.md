## Black Hat Go -- Go Programming for Hackers and Pentesters

![Black Hat Go](https://cdn.jsdelivr.net/gh/TesterCC/pic_bed/img/20200325234257.png)

- [Black Hat Go](https://nostarch.com/blackhatgo)

## Code

- [Black Hat Go - Code](https://github.com/blackhat-go/bhg/)

## After install go
```shell
安装后的设置：
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## Install libpcap

- Ubuntu/Kali：`apt-get install libpcap-dev`
- CentOS7：`sudo yum install libpcap-devel`
- MacOS: `brew xxx`
- Windows: install winpcap (https://www.winpcap.org)

## Wiki

### PCAP-FILTER MAN PAGE

- [packet filter syntax](https://www.tcpdump.org/manpages/pcap-filter.7.html)

### Golang Debug

- [Golang修养之路](https://www.yuque.com/aceld/golang/ga6pb1)

```
test cmd: time go run 2-4_tcp-scanner-too-fast.go

time是Unix/Linux内置多命令，使用时一般不用传过多参数，直接跟上需要调试多程序即可。
- real：从程序开始到结束，实际度过的时间；
- user：程序在用户态度过的时间；
- sys：程序在内核态度过的时间。

一般情况下 real >= user + sys，因为系统还有其它进程(切换其他进程中间对于本进程会有空白期)。

```

### 关键知识笔记

为何Go语言中并没有形成一种统一的用于处理和记录错误的最佳方法？
- Go的内置错误类型没有隐式包含堆栈跟踪以帮助查明错误的上下文或位置。

scanner target: `scanme.nmap.org`
scanner target: `testphp.vulnweb.com`

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

golint 代码规范检查，作用类似于pylint

安装：
```
go get -u github.com/golang/lint/golint
```

使用：
```
golint [file/dir]

e.g.

➜  (master) ✗ golint ./black_hat_go
black_hat_go/c1_demo.go:31:9: var Id should be ID
```

[Golint代码规范检测](https://blog.csdn.net/chenguolinblog/article/details/90665161)

go vet 检查go代码语法
[Go Vet 常见warning总结](https://www.jianshu.com/p/19a44cbc69fb)