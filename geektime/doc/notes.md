#### The master has failed more times than beginner has tried.

## 编写测试程序
#### 1.源码文件以_test结尾：xxx_test.go
#### 2.测试方法名以Test开头: func TestXXX(t *teststing.T) {...}

#### 备注：大写的方法名表示包外可以访问

#### Go更推荐统一的赋值方式，Go 2.0可能会取消一些赋值方式


## 变量赋值
与其他主要编程语言的差异
- 赋值可以进行自动类型推断
- 在一个赋值语句中可以对多个变量进行同时赋值

## 类型转化
与其他主要编程语言的差异
- Go语言不允许隐式类型转换
- 别名和原有类型也不能进行隐式类型转换
- 不支持指针运算（能使用指针）
- string是值类型，其默认的初始化值为空字符串，而不是nil

## 关于打包问题要注意平台兼容性
https://blog.csdn.net/jinjianghai/article/details/85293380
