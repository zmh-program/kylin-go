<div align="center">

# ⚡  kylin-go
Kylin Go,  一款基于 Golang 的跨平台简洁高效轻量的编程语言

![Last Commit](https://img.shields.io/github/last-commit/zmh-program/kylin-go)
![Status](https://img.shields.io/github/actions/workflow/status/zmh-program/kylin-go/build.yaml?branch=main)
![Dependence](https://img.shields.io/badge/dependencies-0-blue)
![License](https://img.shields.io/github/license/zmh-program/kylin-go)

[» Kylin Jvm »](https://github.com/Linwin-Cloud/Kylin-Language)

</div>

## 语言特性
- ✨ 语法简单，易学易用
- ⚡ 语法高效，编译速度快
- 📦 内存占用小
- 🎃 跨平台编译
- 🎈 国际化支持 （英，中文）

## [发行版](https://github.com/zmh-program/kylin-go/releases)
- [x] Windows
- [x] Linux Ubuntu
- [x] MacOS

## 语言示例

```kylin
fn main() {
    print("Hello, World!")
}

main()
```

## 国际化
```kylin
use 'chinese'

遍历 变量 在 范围(2) {
    输出(变量, "hi")
     尝试 {
        输出(变量位置, "hi")
    } 捕获 {
        输出("报错：", error)
    }
}

```

> ```shell
> 0 hi
> 报错： ReferenceError(type="ReferenceError", message="Variable 变量位置 not defined", line=7, column=15)
> 1 hi
> 报错： ReferenceError(type="ReferenceError", message="Variable 变量位置 not defined", line=7, column=15)
> 2 hi
> 报错： ReferenceError(type="ReferenceError", message="Variable 变量位置 not defined", line=7, column=15)
> ```

## 语言规范
#### 1. 赋值计算
```kylin
a = 1024
a += 2
a -= 4
a *= 8
a **= 12
a /= 16

b = a + 1
```

#### 2. 函数
```kylin
fn add(a, b) {
    return a + b
}

fn main() {
    print("Hello, World!")
    return add(1, 2)
}
```

#### 3. 条件判断
```kylin
if val > 10 {
    print("val > 10")
} elif val < 10 {
    print("val < 10")
} else {
    print("val = 10")
}

if condition {
    print("condition")
}

if n + 1 != 0 && n {
    print("n & (n + 1) != 0")
}
```

#### 4. 循环
```kylin
for i in range(0, 10) {
    print(i)
}
```
```kylin
n = 0
while n < 10 {
    print(n)
    n += 1
}
```

#### 5. 数组
```kylin
arr = [1, 2, 3, 4, 5]

for i in arr {
  ...
}

for i in ["hello", "world"] {
  print(i)
}
```

#### 6. 异常截获
```
for i in range(1,10,2) {
    try {
        print(id)
    } catch {
        print("a", error, "b")
    }
}
```
> ```shell
> $ kylin test.ky
> a ReferenceError(message="Variable id not defined", type="ReferenceError",
>  line=3, column=22) b
> a ReferenceError(message="Variable id not defined", type="ReferenceError",
>  line=3, column=22) b
> a ReferenceError(message="Variable id not defined", type="ReferenceError",
>  line=3, column=22) b
> a ReferenceError(message="Variable id not defined", type="ReferenceError",
>  line=3, column=22) b
> a ReferenceError(message="Variable id not defined", type="ReferenceError",
>  line=3, column=22) b
> a ReferenceError(message="Variable id not defined", type="ReferenceError",
>  line=3, column=22) b
> ```

## 运行
```shell
kylin main.ky
```
> .ky 后缀可省略
> ```shell
> kylin server
> ```


## 关键字
- `fn` 函数
- `if` 条件判断
- `else` 条件判断
- `for` 循环
- `in` 循环
- `while` 循环
- `return` 返回
- `true` 真
- `false` 假
- `null` 空
- `try` 异常截获
- `catch` 异常截获
- `break` 跳出循环
- `continue` 跳过本次循环
- `import` 导入
- `use` 国际化导入语言包

## 内置函数
- `print` 打印
- `input` 输入
- `str` 字符串
- `int` 整数
- `float` 浮点数
- `bool` 布尔值
- `array` 数组
- `range` 范围
- `len` 长度
- `sum` 求和
- `max` 最大值
- `min` 最小值
- `abs` 绝对值
- `all` 全部为真
- `any` 任意为真
- `join` 连接字符串
- `split` 分割字符串
- `type` 类型
- `time` 时间 (ms)
- `sleep` 阻塞等待 (ms)
- `timenano` CPU 时间 (ns)
- `read` 读取文件
- `write` 写入文件
- `shell` 执行 shell 命令
- `exit` 退出程序

## 语言设计
#### 1. 词法分析
- Lexer 编译 AST 语法树
- 词法分析器使用递归下降进行语法分析

#### 2. 语言解释
- 动态解释器

#### 3. 语言编译
- 编译器使用 Golang SSA 进行编译


## 基准测试

#### 性能测试
```kylin
n = 0
val = 2345
t = time()
while n < 9999999 {
    n += 1
    val **= 1289
    val = n * 999
}

print(time() - t, "ms")
```
1. C (GCC) `891.9ms`
2. Golang `1012.5ms`
3. NodeJS `1170.0ms`
4. **Kylin Go** `1751.2ms`
5. Python `4681.8ms`

#### 内存占用 （整体）
```kylin
n = 0
while n < 1000000 {
    n += 1
    print('hello world')
}
```
1. C (GCC) `0.9 MB`
2. Golang `3.2 MB`
3. **Kylin Go** `8 MB`
4. Node `10 MB`
5. Python `14MB`
6. Kylin JVM `84MB`

#### 内存溢出测试
1. C ❌
2. Golang ❌
3. **Kylin Go** ❌
4. NodeJS ❌
5. Python ❌
6. Kylin JVM ✔

#### 风格对比
C 语言
```c
#include <stdio.h>

int main() {
    int n = 0;
    
    while (n < 1000000) {
        n++;
        printf("hello world\n");
    }
    
    return 0;
}
```
Golang 
```go
package main

import "fmt"

func main() {
	n := 0
	for n < 1000000 {
		n++
		fmt.Println("hello world")
	}
}
```
**Kylin Go**
```kylin
n = 0
while n < 1000000 {
    n += 1
    print('hello world')
}
```
**Kylin JVM**
```kylin
var n = 0
while n < 1000000 {
    n = <n + 1>
    out('hello world')
}
```
NodeJS
```js
let n = 0;
while (n < 1000000) {
    n++;
    console.log('hello world');
}
```
Python
```python
n = 0
while n < 1000000:
    n += 1
    print('hello world')
```
