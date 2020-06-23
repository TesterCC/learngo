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


## 变量定义要点：

#### 1.变量类型写在变量名之后
#### 2.编译器可推测变量类型
#### 3.没有char，只有rune (32 bits)
#### 4.原生支持复数类型

## if语句

#### if的条件里可以赋值
#### if的条件里赋值的变量变量作用域就在这个if语句里

## switch语句

#### switch后可以没有表达式
#### switch会自动break，除非使用fallthrough

## for语句

#### for的条件里不需要括号
#### for的条件里可以省略初始条件、结束条件、递增表达式

## 基本语法回顾

#### for,if 后面的条件没有括号
#### if条件里也可以定义变量
#### 没有while
#### switch不需要break，也可以直接switch多个条件

## 函数

#### 函数可返回多个值
#### 函数返回多个值时可以起名字
#### 仅用于非常简单的函数
#### 对于调用者而言没有区别
#### 多个返回值的使用场景通常是返回一个error
#### 函数可以作为参数
#### 返回值的类型写在最后面
#### 没有默认参数，可选参数

## 指针
（go语言的指针比C语言的指针好理解）
#### go语言的指针不能运算
#### 参数传递
```markdown
C、C++既可以值传递也可以引用传递
Java、Python绝大部分是引用传递
Go语言只有值传递一种方式
```

## 数组

#### 数组遍历
```
求和，可通过_省略变量
不仅range，任何地方都可通过_省略边量
如果只要i，可写成 for i:= range numbers
```

#### 为什么要用range
```markdown
意义明确、美观
C++没有类似能力
Java/Python：只能for each value，不能同时获取i,v
```

#### 数组是值类型
- `[10]int` 和 `[20]int` 是不同类型
- 调用`func f(arr [10]int)` 会**拷贝**数组
- 在go语言中一般不直接使用数组，更多是用slice(切片)

## 切片 Slice
- Slice本身没有数据，是对底层array的一个view
- Slice可以向后扩展，不可以向前扩展
- s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)

### 向Slice添加元素
- 添加元素时如果超越cap最大值，系统会重新分配更大的底层数组 （原来的数组如果没用，则会被垃圾回收）
- 由于值传递的关系，必须接收append的返回值。
```go
s = append(s, val)
```

## Map

### Map的操作

- 创建: `make(map[string]int)`
- 获取元素：`m[key]`
- 当key不存在时，获得value类型的初始值或者零值
- 用`value, ok := m[key]`来判断是否存在key

### Map的遍历

- 使用range遍历key，或者遍历key,value对
- 不保证遍历顺序，如需顺序，需手动对key排序
- 使用len获得元素个数

### Map的key

- map使用哈希表，必须可以比较相等
- 除了slice,map,function的内建类型都可以作为key
- Struct类型不包含上述字段，也可作为key （编译时会检查）

例题：
[无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/?utm_source=LCUS&utm_medium=ip_redirect_q_uns&utm_campaign=trans)
[英文版](https://leetcode.com/problems/longest-substring-without-repeating-characters)
同：[剑指Offer 48 寻找最长不含有重复字符的子串](https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/)

解题思路：

对于每一个字母x
- `lastOccurred[x]`不存在，或者`< start` --> 无需操作
- `lastOccurred[x] >= start ` --> 更新start
- 更新`lastOccurred[x]`，更新maxLength


REF：
[在线学习](https://www.bilibili.com/video/BV18Q4y1M7NV)