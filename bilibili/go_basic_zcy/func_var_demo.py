def foo(n: int) -> None:
    print(f"n={n}, n的地址 {id(n)}")
    n += 1   # 因为python中int类型是常量，所以n+=1需要copy一份出来
    print(f"n={n}, n的地址 {id(n)}")  # 所以这里的打印id不一样了

if __name__ == '__main__':
    i = 8
    print(f"i={i}, i的地址 {id(i)}")
    foo(i)
    print(f"i={i}, i的地址 {id(i)}")

"""
i=8, i的地址 140714461186712
n=8, n的地址 140714461186712
n=9, n的地址 140714461186744
i=8, i的地址 140714461186712
"""