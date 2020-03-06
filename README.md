# go-demos

## goRoutine

### example1
#### `runtime.GOMAXPROCS`
设置同时运行的协程数
####`sync.waitGroup`
协调不同协程运行结束，类似于`CountDownLatch`。有`Add()`方法初始化数量，具体运行时调用`Done()`方法“CountDown”，主协程调用`Wait()`方法等待结束。
#### `runtime.Gosched()`
让出当前时间片给其他协程使用。
#### `defer`
后面必须跟函数调用语句，表示在当前函数执行结束后调用（若希望它在函数中间执行，需要定义一个匿名函数在内部调用）。

常用来关闭打开的资源或者执行必须执行的内容，例如实例中协程运行结束后对waitGroup()倒数：`defer wg.Done()`。或者`mutx.Lock()`调用后必须释放时加上`defer mutx.unlock()`、文件打开后自动关闭等。

执行顺序为后定义先执行

``` Go
defer wg.Done()
defer mutx.Unlock()
// 先执行解锁，后执行wg倒数
```

`defer`计算时间为定义时间，不是实际执行时间。
若需要进行多项工作，可以定义到一个函数中，集中调用。

参考https://www.jianshu.com/p/5b0b36f398a2
### example2
#### 竞态
Go语言竞态检测工具： `go build -race main.go`
再调用`./main`可以看到输出中有标识出的竞态语句
``` bash
==================
WARNING: DATA RACE
Read at 0x000000653278 by goroutine 8:
  main.addCount()
      /home/liuzhx/Desktop/go-demos/goRoutine/example2/main.go:17 +0x79

Previous write at 0x000000653278 by goroutine 7:
  main.addCount()
      /home/liuzhx/Desktop/go-demos/goRoutine/example2/main.go:20 +0x9a

Goroutine 8 (running) created at:
  main.main()
      /home/liuzhx/Desktop/go-demos/goRoutine/example2/main.go:27 +0x77

Goroutine 7 (running) created at:
  main.main()
      /home/liuzhx/Desktop/go-demos/goRoutine/example2/main.go:26 +0x5f
==================
counter: 4
Found 1 data race(s)
```
#### 处理竞态的方法
- 原子函数操作数据
- 互斥锁锁住临界区
- 使用管道`chan`

