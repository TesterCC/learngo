## 内建变量类型

#### bool, string

#### (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr  
- 加u是无符号整数，int长度不设置则根据操作系统32位还是64位来决定

#### byte, rune    
- rune是go语言的字符类型(类似C语言中的char)，因为utf-8编码很多字符是3字节，所以用int32 4字节代表rune; byte 8位 

#### float32，float64, complex64, complex128
- float浮点数(属于实数)，complex复数
- complex64的实部和虚部是float32,complex128的实部和虚部是float64

Go运行结果：
```go
(0+1.2246467991473515e-16i)
```

Python运行结果：
```python
In [35]: import cmath
In [36]: cmath.exp(1j*cmath.pi) + 1
Out[36]: 1.2246467991473532e-16j
```

## 强制类型转换
#### Go的类型转换是强制的