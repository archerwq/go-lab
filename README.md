# Golang Notes

## Four types of Go data

* basic type: bool, number(byte, int, float...), string
* composite: array, struct
* reference: pointer, slice, map, chan, function
* interface

## Choosing a value or pointer receiver

There are two reasons to use a pointer receiver.  
The first is so that the method can modify the value that its receiver points to.  
The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.  
In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

## Pass by Value

Go is very consistently pass-by-value, but people sometimes get confused by the slice and channel types. Those look like they're pass-by-reference since they have the same backing data even after being passed, but conceptually, you should think of a slice as a struct with an int length, int capacity, and a pointer to the backing array. Once you think of it that way, it all makes sense.
More...

* [Things I Wish Someone Had Told Me About Golang](http://openmymind.net/Things-I-Wish-Someone-Had-Told-Me-About-Go/)

## Interface Pointer

Maybe it's just me, but it took me a while to understand this method from Go's net/http package:
func ServeHTTP(res http.ResponseWriter, req *http.Request) {
  ...
}
Why is req passed as a reference, but res passed as a value? Why, I wondered, are they making me pay the performance penalty of passing a copy of an object?! That's not what's happening though. http.ResponseWriter is an interface and either a value-type or a reference-type can satisfy an interface. So, technically, you don't know whether the value being passed is a copy of a pointer or a copy of a value, but it's probably the former.
More...

* [Things I Wish Someone Had Told Me About Golang](http://openmymind.net/Things-I-Wish-Someone-Had-Told-Me-About-Go/)

## Go Concurrency

Two way to achieve concurrency in Go:

* CSP (Comunicating Sequential Process, a morden concurrency model), with goroutine and channel. A.K.A Don't communicate by sharing memory, share memory by communicating.
* Multiple threads sharing memory, with goroutine and utilities in sync packages (e.g. sync.RWMutex).

Goroutine is an independently executing function, launched by go statement.

* It has its own call stack, which grows and shrinks as required.
* It's very cheap, it's practical to have thousands, even hundreds of thousands of goroutines.
* It's not a thread. There might be only one thread in a program with thousands of goroutines.

More...

* [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)

## Go GC

To create a garbage collector for the next decade, we turned to an algorithm from decades ago. Go's new garbage collector is a concurrent, tri-color, mark-sweep collector, an idea first proposed by Dijkstra in 1978. This is a deliberate divergence from most "enterprise" grade garbage collectors of today, and one that we believe is well suited to the properties of modern hardware and the latency requirements of modern software.  
In a tri-color collector, every object is either white, grey, or black and we view the heap as a graph of connected objects. At the start of a GC cycle all objects are white. The GC visits all roots, which are objects directly accessible by the application such as globals and things on the stack, and colors these grey. The GC then chooses a grey object, blackens it, and then scans it for pointers to other objects. When this scan finds a pointer to a white object, it turns that object grey. This process repeats until there are no more grey objects. At this point, white objects are known to be unreachable and can be reused.  
This all happens concurrently with the application, known as the mutator, changing pointers while the collector is running. Hence, the mutator must maintain the invariant that no black object points to a white object, lest the garbage collector lose track of an object installed in a part of the heap it has already visited. Maintaining this invariant is the job of the write barrier, which is a small function run by the mutator whenever a pointer in the heap is modified. Go’s write barrier colors the now-reachable object grey if it is currently white, ensuring that the garbage collector will eventually scan it for pointers.

More...  

* [Go GC: Prioritizing low latency and simplicity](https://blog.golang.org/go15gc)
* [Golang’s Real-time GC in Theory and Practice](https://making.pusher.com/golangs-real-time-gc-in-theory-and-practice/)

## Go Vendor

When using the go tools such as go build or go run, they first check to see if the dependencies are located in ./vendor/. If so, use it. If not, revert to the $GOPATH/src/ directory.

* [How should I use vendor in Go 1.6?](https://stackoverflow.com/questions/37237036/how-should-i-use-vendor-in-go-1-6)

## Go Closure

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

## Go Compiler

Programming languages aren't programs, hence they're not "written" in any language. They are often described by formal grammars (e.g. BNF).
Interpreters and compilers for programming languages are programs and so must be written in some kind of programming language.
Go has at least two compilers, gc and gccgo. The former was written in C, but is now written in Go itself. While the latter is a gcc frontend written mainly in C++. Go's libraries are written in Go.

[What language is the Go programming language written in?](https://stackoverflow.com/questions/3327676/what-language-is-the-go-programming-language-written-in)

Generally the first version of the compiler is written in a different language, and then each subsequent version is written in that language and compiled with the older version. Once you've compiled version x with version x-1, you can use the newly built version x to recompile itself, taking advantage of any new optimizations that version introduces; GCC does its releases that way.

[How can a language's compiler be written in that language?](https://stackoverflow.com/questions/2998768/how-can-a-languages-compiler-be-written-in-that-language)

## Go Runtime

Go does have an extensive library, called the runtime, that is part of every Go program. The runtime library implements garbage collection, concurrency, stack management, and other critical features of the Go language. Although it is more central to the language, Go's runtime is analogous to libc, the C library.

It is important to understand, however, that Go's runtime does not include a virtual machine, such as is provided by the Java runtime. Go programs are compiled ahead of time to native machine code (or JavaScript or WebAssembly, for some variant implementations). Thus, although the term is often used to describe the virtual environment in which a program runs, in Go the word runtime is just the name given to the library providing critical language services.

[How does the Go runtime work?](https://www.quora.com/How-does-the-Go-runtime-work-What-does-it-consist-of-What-functionalities-does-it-provide-and-what-can-be-expected-from-a-developer-perspective)

## Go Embedding

The theory behind embedding is pretty straightforward: by including a type as a nameless parameter within another type, the exported parameters and methods defined on the embedded type are accessible through the embedding type. The compiler decides on this by using a technique called promotion: the exported properties and methods of the embedded type are promoted to the embedding type.

[Type embedding in Go](https://travix.io/type-embedding-in-go-ba40dd4264df)

## Go函数调用栈

一般编程语言的函数调用栈大小是固定值，递归调用时如果层级比较多，会出现栈溢出的情况，而Go的函数调用栈是动态变化的，所以不用担心栈溢出。

## Go关闭系统资源

Go的内存回收不会回收打开的系统资源，比如文件、网络连接等，需要显式地去关闭。

## Go错误处理

* 抛出错误：一般被调用函数返回的错误信息包含调用参数和错误描述，调用方如果处理不了可以添加额外信息后抛出。

```go
return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
```

* 重试：要设置timeout机制以及exponential back-off机制

* 记录日志并退出程序，只应在main中使用

```go
if err := waitForServer(url); err != nil {
  log.Fatalf("site is down: %v\n", err)
}
```

* 只记录日志

* 完全忽略错误，比如删除在/tmp创建的临时文件出错，直接忽略是安全的，因为系统会定期删除临时文件，但当我们忽略一个错误时要明确说明原因

## Go函数值

可以使得我们不仅可以用数据参数化函数，而且可以用行为参数化函数。

```go
func add1(r rune) rune {return r + 1}
fmt.Println(Strings.Map(add1, "HAL-9000")) // "IBM.:111"
```

## Go匿名函数

```go
fmt.Println(Strings.Map(func(r rune) rune {return r + 1}, "HAL-9000"))
```

## Go闭包

```go
func add(x int) func(int) int {
  return func(y int) int {
    return x + y
  }
}

func main() {
  add1 := add(1)
  add4 := add(4)
  fmt.Println(add1(3))
  fmt.Println(add4(3))
}
```

[Closure](https://en.wikipedia.org/wiki/Closure_(computer_programming))
In programming languages, a closure (also lexical closure or function closure) is a technique for implementing lexically scoped name binding in a language with first-class functions. Operationally, a closure is a record storing a function[a] together with an environment.[1] The environment is a mapping associating each free variable of the function (variables that are used locally, but defined in an enclosing scope) with the value or reference to which the name was bound when the closure was created.[b] A closure unlike a plain function allows the function to access those captured variables through the closure's copies of their values or references, even when the function is invoked outside their scope.

闭包的一个坑，不要在闭包里引用循环变量

```go
var rmdirs []func()
for _, d := range tempDirs {
  dir := d // 闭包引用的是变量，而不是特定时刻变量的值
  os.MkdirAll(dir, 0755)
  rmdirs = append(rmdirs, func(){
    os.RemoveAll(dir)
  })
  // do some work ...
  for _, rmdir := range rmdirs {
    rmdir()
  }
}
```

类似的，defer和go语句 + 闭包里引用循环变量也有同样的问题，因为这两个都是循环结束后执行

```go
func main() {
  for i := 0; i < 5; i++ {
    // i := i
    go func() {
      fmt.Println(i) // 输出4个3和1个5
    }()
    if i == 3 {
      time.Sleep(3 * time.Second)
    }
  }
  fmt.Scanln()
}
```

## Go可变长度参数

```go
func sum(values ...int) int {
  total := 0
  for _, v := range values {
    total += v
  }
  return total
}

func main() {
  fmt.Println(sum(1, 2, 3)) // 调用者隐式创建一个数组，并将原始参数复制到数组中，再把数组的一个切片作为参数传入
  values := []int{1, 2, 3}
  fmt.Println(sum(values...)) // 如果原始参数已经是一个切片，只需要在后面加上...
}
```

## nil slice vs. empty slice

The zero value of a slice is nil. A nil slice has a length and capacity of 0 and has no underlying array. It behaves all the same with an empty slice.

[nil slice vs. empty slice](https://stackoverflow.com/questions/44305170/nil-slices-vs-non-nil-slices-vs-empty-slices-in-go-language)

## 方法的接收器

接收器是指针类型，可以用非指针类型去调用，编译器会自动加上取地址操作，但是前提是得是能取到地址的变量，临时变量不可以的；接收器是非指针类型，可以用指针去调用，编译器会自动加上解引用操作。

## 接口值

* 一个接口的值由动态类型和动态值组成，接口值为nil指的是动态类型为空且动态值为空
* 两个接口相等的条件是：都为nil，或者动态类型一样且动态值相等
* 如果动态值不可比较（如slice, map, function)，接口就不可比较，`var x interface{} = []int{1, 2, 3}`

```go
func main() {
  var buf *bytes.Buffer
  f(buf) // out == nil ? false
}

func f(out io.Writer) {
  fmt.Println("out == nil ?", out == nil)
}
```

## Channel

一个无缓存的channel的发送操作导致发送者goroutine阻塞，直到另一个goroutine在相同的channel上执行接收操作，当发送的值通过channel成功传输之后，两个goroutine继续执行后面的语句。反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到另一个goroutine在相同的channel上执行发送操作。

接收数据在唤醒发送者之前发生(happens before).

当一个channel作为参数时，他一般总是被用作只发送或只接收，为了表明这种意图或者防止被滥用，go提供了单向channel类型，分别用于只发送(chan<- int)或只接收(<-chan int)。参数传递时存在隐式类型转换（无方向转有方向，有方向之间互转，但不能有方向转无方向）。

goroutine泄露：如果一个goroutine写一个无缓存channel，而这个channel因为没有人去读导致写的goroutine一直阻塞在那里。

带缓存的channel: 缓存队列满导致发送者goroutine阻塞，缓存空导致接收者goroutine阻塞，不满不空时，发送接收都正常进行不阻塞。

有时候增加缓存容量并不能解决性能问题，因为发送和接收速率不匹配，这个时候需要增加相应goroutine数量来解决速率不匹配问题。

Once a channel has been closed, you cannot send a value on this channel, but you can still receive from the channel. `v, ok := <- ch` return zero value for `ch` type and `false` indicating the channel has been closed, and will return the same values if it's called repeatedly. As soon as the finish channel is closed, it becomes ready to receive. This powerful idiom allows you to use a channel to send a signal to an unknown number of goroutines, without having to know anything about them, or worrying about deadlock. [A closed channel never blocks](https://dave.cheney.net/2013/04/30/curious-channels)

## Interface values with nil underlying values
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
Note that an interface value that holds a nil concrete value is itself non-nil.
