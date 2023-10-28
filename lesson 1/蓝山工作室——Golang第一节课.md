# 蓝山工作室——Golang第一节课

## 前言

首先，我要热烈欢迎你们加入我们的Go语言课程。Go语言是一门强大且灵活的编程语言，它具备高性能、简单易学的特点，已经在许多领域取得了广泛的应用。本课程将帮助你掌握Go语言的核心概念、语法和最佳实践，无论你是初学者还是有一些编程经验的学员。

Go语言是一门开源编程语言，由Google开发并维护。它的设计目标是简单性、效率和可读性，使得开发者可以更轻松地构建高性能的应用程序。Go语言在云计算、网络编程、大数据处理等领域表现出色，是许多知名公司和大厂的首选语言之一。

这节课的主要内容是认识：变量，常量，基本数据类型，运算符，流程控制，函数基础，fmt包

## 基础语法

### var

#### 变量

##### 变量类型
变量（Variable）的功能是存储数据。不同的变量保存的数据类型可能会不一样。经过半个多世纪的发展，编程语言已经基本形成了一套固定的类型，常见变量的数据类型有：整型、浮点型、布尔型等。

Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用。

1. **整数类型（Integer Types）**：
   - `int`：根据你的操作系统架构，可以是32位或64位的整数。
   - `int8`、`int16`、`int32`、`int64`：有符号整数类型，分别表示8位、16位、32位和64位整数。
   - `uint`：无符号整数类型，根据操作系统架构，可以是32位或64位。
   - `uint8`、`uint16`、`uint32`、`uint64`：无符号整数类型，分别表示8位、16位、32位和64位整数。
2. **浮点数类型（Floating-Point Types）**：
   - `float32`：单精度浮点数。
   - `float64`：双精度浮点数。
3. **复数类型（Complex Types）**：
   - `complex64`：包含32位实部和32位虚部的复数。
   - `complex128`：包含64位实部和64位虚部的复数。
4. **布尔类型（Boolean Type）**：
   - `bool`：表示真（true）或假（false）。
5. **字符串类型（String Type）**：
   - `string`：用于存储文本数据的字符序列。
6. **字节类型（Byte Type）**：
   - `byte`：`uint8` 的别名，通常用于表示ASCII字符。
7. **符文类型（Rune Type）**：
   - `rune`：`int32` 的别名，通常用于表示**Unicode字符**。
   - Unicode（统一码、万国码、单一码）是一种字符编码标准，它用于表示世界上几乎所有的书写系统中的字符，**包括各种文字、标点符号和特殊符号**。Unicode的目标是提供一个统一的、跨语言的字符编码系统，以消除不同字符编码之间的混乱和兼容性问题。

##### 变量声明
Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后**必须使用**。

##### 变量的初始化
Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如： 整型和浮点型变量的默认值为0。 字符串变量的默认值为空字符串。 布尔型变量默认为false。 切片、函数、指针变量的默认为nil。

```go
var 变量名 类型 = 表达式

var a = "initial" // 类型推导，不指定类型自动判断

var b, c int = 1, 2 // 一次初始化多个变量

var d = true // 布尔型变量

var e float64 // 普通声明未赋值

f := float32(e) // 类型转换初始化赋值

g := a + "apple"
fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
fmt.Println(g)                // initialapple
```

#### 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。

```go
const s string = "constant"
const h = 500000000
const i = 3e20 / h
fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
```

### for

for 是设置条件并让代码进行循环的关键字

#### 三段式

```go
for A; B; C {
    // 这里是循环体，每次循环需要执行的代码在这里面
}
```

`A` 位置是单次表达式，循环开始时会执行一次这里，一般用于初始化变量。

`B` 位置是条件表达式，即循环条件，只要满足循环条件就会执行循环体。

`C` 位置是末尾循环体，每次执行完一遍循环体之后会执行一次 `C` 中的表达式。

执行末尾循环体后将再次进行条件判断，若条件还成立，则继续重复上述循环，当条件不成立时则跳出当下for循环。

并且，你可以选择性的留空，就是让A，B，C**任何位置**为空。例如：

```go
for ;i<5;{
  // 循环体
}
```

#### 一段式

一段式那就是只写条件

```go
for A{
  // 循环体
}
```

`A` 位置是条件表达式，即循环条件，只要满足循环条件就会执行循环体。

例如：

```go
for i<5{
  // 循环体
}
```

##### break 关键词

break 放在循环体中，只要执行到 break，则会立马跳出所在最里循环（注意，是所在最里循环，若嵌套，则无法跳出更外层循环）

例如：

```go
for i:=1;i<4;i++{
  j := i
  for j<4{
    if j == 2{
      break
    }
  }
  fmt.Println("hello lanshan")
}
```

上面这个函数不需要去体会其中的意思，我只是举个例子。若执行到了break，则只会跳出条件是 j < 4 这个循环，依然会执行println打印

这是总的示例，可以回顾一下：

```go
package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println("loop")
		break // 跳出循环
	}
	
	// 打印7、8
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
			// 当n模2为0时不打印，进到下一次的循环
		}
		fmt.Println(n)
	}
	// 直到i>3
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
  // for 循环嵌套
  for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
```

### if

```go
if 条件表达式 {
	//当条件表达式结果为true时，执行此处代码   
}

if 条件表达式 {
    //当条件表达式结果为true时，执行此处代码  
} else {
    //当条件表达式结果为false时，执行此处代码  
}
```

```go
package main

import "fmt"

func main() {
	// 条件表达式为false，打印出"7 是奇数"
	if 7%2 == 0 {
		fmt.Println("7 是偶数")
	} else {
		fmt.Println("7 是奇数")
	}

	// 条件表达式为ture，打印出"8 被 4 整除"
	if 8%4 == 0 {
		fmt.Println("8 被 4 整除")
	}

	// 这是一个短声明，效果等效于
	//num := 9
	//if num < 0{
	//	...
	//}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```

### switch

当分支过多的时候，使用if-else语句会降低代码的可阅读性，这个时候，我们就可以考虑使用switch语句

- switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
- switch 语句在默认情况下 case 相当于自带 break 语句，匹配一种情况成功之后就不会执行其它的case，这一点和 c/c++ 不同
- 如果我们希望在匹配一条 case 之后，继续执行后面的 case ，可以使用 fallthrough

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		// 在此打印"two"并跳出
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	// 根据现在的时间判断是上午还是下午
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
```

### func

函数是指一段可以直接被另一段程序或代码引用的程序或代码，一个较大的程序一般应分为若干个程序块，每一个模块用来实现一个特定的功能。

1. **函数的声明和定义**：

   在Go语言中，函数的定义以 `func` 关键字开始，然后是函数名、参数列表、返回类型和函数体。以下是一个函数的典型定义：

   ```go
   func add(x int, y int) int {
       return x + y
   }
   ```

   这个函数名为 `add`，接受两个整数参数 `x` 和 `y`，并返回一个整数。

2. **函数的参数**：

   函数可以接受零个或多个参数，参数在参数列表中定义，并且需要指定参数的类型。例如：

   ```go
   func greet(name string) {
       fmt.Println("Hello, " + name)
   }
   ```

   这个函数接受一个字符串参数 `name`。

3. **函数的返回值**：

   函数可以返回一个或多个值，返回值的类型也需要在函数定义中指定。如果函数没有返回值，可以将返回类型留空。例如：

   ```go
   func addAndMultiply(x, y int) (int, int) {
       sum := x + y
       product := x * y
       return sum, product
   }
   ```

   这个函数返回两个整数值。

4. **函数的调用**：

   要调用函数，只需使用函数名并传递参数。例如：

   ```go
   result := add(3, 5)
   fmt.Println(result)
   ```

   这里我们调用了 `add` 函数，将参数 `3` 和 `5` 传递给它，并将返回值赋给 `result` 变量。

```go
package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
	result := add(3, 5)
	fmt.Println(result)
}
```

### fmt

fmt 库函数

```go
package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n) // hello 123
	fmt.Println(p)    // {1 2}

	fmt.Printf("s=%v\n", s)  // s=hello
	fmt.Printf("n=%v\n", n)  // n=123
	fmt.Printf("p=%v\n", p)  // p={1 2}
	fmt.Printf("p=%+v\n", p) // p={x:1 y:2}
	fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2}

	f := 3.141592653
	fmt.Println(f)          // 3.141592653
	fmt.Printf("%.2f\n", f) // 3.14
}
```

## 年轻人的第一个GoProject

## 作业

### LV1 

编写一个Go函数，接受两个整数作为参数，然后返回它们的和。在 `main` 函数中调用此函数并打印结果。

### LV2

编写一个Go函数，接受圆的半径作为参数，然后返回圆的面积。使用 `math` 包中的常数 Pi。在 `main` 函数中调用此函数并打印结果。

提示，引入 Pi 只需要写出`math.Pi`

### LV3

编写一个Go函数，接受一个整数作为参数，然后判断它是否为素数（质数）。在 `main` 函数中调用此函数并打印结果。提示：一个素数是只能被 1 和自身整除的正整数。

### LVX

编写一个Go函数，使用`rand`包随机选择一个1-100的数（必须每次执行的随机数都不一样），然后使用**二分法**找到这个数。

tips：rand 包的使用和二分法自行研究


作业完成后将作业 GitHub 地址发送至 **nihilism0@qq.email** ，若对 GitHub 的使用有问题，可以先网上寻找解决方法，实在不行可以私信我。

**提交格式（主题）：2023xxxx00-邓卓-LV3** (最后LV中写出你完成的最大等级)

**截止时间**：下一次上课之前