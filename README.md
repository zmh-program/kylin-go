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
- 语法简单，易学易用
- 语法高效，编译速度快
- 内存占用小
- 跨平台编译


## 语言示例

```kylin
fn main() {
    print("Hello, World!")
}

main()
```

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
- `break` 跳出循环
- `continue` 跳过本次循环
- `import` 导入
- `use` 国际化导入语言包

## 内置函数
- `print` 打印
- `input` 输入
- `range` 范围
- `len` 长度
- `sum` 求和
- `max` 最大值
- `min` 最小值
- `abs` 绝对值
- `all` 全部为真
- `any` 任意为真
- `join` 连接字符串
- `type` 类型
- `time` 时间 (ms)
- `sleep` 阻塞等待 (ms)
- `timenano` CPU 时间 (ns)
- `exit` 退出程序

## 语言设计
#### 1. 词法分析
- Lexer 编译 AST 语法树
- 词法分析器使用递归下降进行语法分析

#### 2. 语言解释
- 动态解释器

#### 3. 语言编译
- 编译器使用 Golang SSA 进行编译
