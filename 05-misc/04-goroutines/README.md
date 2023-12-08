# Goroutines and channels

A goroutine is a lightweight thread managed by the Go runtime.

Channels is how you communicate between routines.

## Introduction

In this chapter you will:

- Understand the difference between concurrency and parallelism
- Use Goroutines to run your functions
- Create and use channels to communicate between your Goroutines
- Apply Goroutines to an app that searches files for faster execution.

## Concurrency, what's the benefit

Concurrency is the task of running and managing the multiple computations at the same time. While _parallelism_ is the task of running multiple computations simultaneously.

So what are some benefits:

- **Faster processing**. The benefit is getting tasks done faster. Imagine that you are searching a computer for files, or processing data, if it's possible to work on these workloads in parallel, you end up getting the response back faster.
- **Responsive apps** Another benefit is getting more responsive apps. If you have an app with a UI, imagine it would be great if you can perform some background work without interrupting the responsiveness of the UI.

## Goroutines

A goroutine is a lightweight thread managed by the Go runtime. What you do is to add the keyword `go` in front of a function. Here's an example:

```go
go myFunction()
```

Imagine the following code running, what would happen?

```go
package main

import "fmt"

func myFunction() {
 for i := 0; i < 3; i++ {
  fmt.Println("my function: ", i)
 }
}
func anotherFunction() {
 for i := 4; i <7; i++ {
  fmt.Println("another function: ", i)
 }
}

func main() {
 go myFunction()
 anotherFunction()
}
```

It would only print the result from `anotherFunction()` as it takes a short while for the go routine to start up. You can have the go routine execute as well by adding a little delay, like so:

```go
func main() {
 go myFunction()
 anotherFunction()
 time.Sleep(1 * time.Second)
}
```

The result is now the following:

```output
another function:  4
another function:  5
another function:  6
my function:  0
my function:  1
my function:  2
```

The function with the go routine finishes last. Lets modify the code slightly and have the two functions use a delay, so we simulate workloads taking different time to finish:

```go
func myFunction() {
 time.Sleep(1500 * time.Millisecond)
 for i := 0; i < 3; i++ {
  fmt.Println("my function: ", i)
 }
}
func anotherFunction() {
 time.Sleep(500 * time.Millisecond)
 for i := 4; i < 7; i++ {
  fmt.Println("another function: ", i)
 }
}

func main() {
 go myFunction()
 go anotherFunction()
 time.Sleep(2 * time.Second)
}
```

at this point, `anotherFunction()` finishes first as it has the shortest delay, which is to be expected. Here's what the output looks like now:

```output
another function:  4
another function:  5
another function:  6
my function:  0
my function:  1
my function:  2
```

### Use case - a file search

Imagine you have case where you need to find a file on disk. If you write a function like so, it will search a directory and report back the result if the file is found:

```go
func SearchFiles(dir string, lookFor string) string {
 log.Println("[SEARCHING] ", dir)
 files, err := ioutil.ReadDir(dir)
 if err != nil {
  log.Fatal(err)
 }

 for _, file := range files {
  log.Println(dir+file.Name(), file.IsDir())
  if file.Name() == lookFor {
   return "[FOUND] " + filepath.Join(dir, file.Name())
  }
 }
 return "[NOT FOUND] " + dir
}
```

Imagine you now run this code like so, to search many directories:

```go
result := make([]string, 0)
append(result, SearchFile("./tmp", "myfile.txt"))
append(result, SearchFile("./tmp2", "myfile.txt"))
append(result, SearchFile("./tmp3", "myfile.txt"))
append(result, SearchFile("./tmp4", "myfile.txt"))

for i := 0 i< len(result); i++ {
  fmt.Println(result[i])
}
```

If found, you will get an output similar to the below, depending on whether _myfile.txt_ is found in any of the searched directories:

```go
[FOUND] ./tmp/myfile.txt
[NOT FOUND] ./tmp2/myfile.txt
[NOT FOUND] ./tmp3/myfile.txt
[NOT FOUND] ./tmp4/myfile.txt
```

Now to speed up this process, it would be great if you are able to search many directories at once, so you could type something like so:

```go
go SearchFile("./tmp", "myfile.txt")
go SearchFile("./tmp2", "myfile.txt")
go SearchFile("./tmp3", "myfile.txt")
go SearchFile("./tmp4", "myfile.txt")
```

This works, it now searches all directories, in parallel. However, now we don't have a way to get the response back as we can't write like so:

```go
result := make([]string, 0)
go append(result,SearchFile("./tmp", "myfile.txt")) // won't compile, says "go discards results"
```

So how can we get the result from a go routine, the answer is by using channels, so lets discuss those next.

## Channels

A channel is how we can communicate cross go routines but also between go routines and the part of our code not using a go routine.

The idea is to send a value to a channel, and have part of our code listen to values from a channel.

### Creating a channel

To create a channel, you need the keyword `chan` and the data type of the messages you are about to send into it. Here's an example:

```go
ch := make(chan int)
```

In the above example, a channel `ch` will be created that accepts messages of type `int`.

### Sending a value to a channel

To send to a channel, you need to use this operator `<-`, it look like a left pointing arrow and is meant to be read as the direction something is sent. Here's an example of sending a message to a channel:

```go
ch <- 2
```

In the above code, the number 2 is sent into the channel `ch`.

### Listening to a channel

To listen to a channel, you again use the arrow `<-`, but this time you need a receiving variable on the left side and the channel on the right side, like so:

```go
value := <- ch
```

### Matching sending and receiving

Let's say you have the following code:

```go
package main

import "fmt"

func produceResults(ch chan int) {
 ch <- 1
 ch <- 2
}

func main() {
 ch := make(chan int)
 go produceResults(ch)

 var result int
 result = <-ch
 fmt.Println(result)
 result = <-ch
 fmt.Println(result)
}
```

You are invoking `produceResults()` and it sends messages to the channel twice:

```go
ch <- 1
ch <- 2
```

in `main()`, you receive the results:

```go
var result int
result = <-ch
fmt.Println(result)
result = <-ch
fmt.Println(result)
```

So what happens if you produce more values than you receive like so?

```go
ch <- 1
ch <- 2
ch <- 3
```

answer: you will miss out on the extra value.

What if it's the opposite, you try to receive one more value than you actually get?

```go
var result int
result = <-ch
fmt.Println(result)
result = <-ch
fmt.Println(result)
result = <-ch
fmt.Println(result)
```

At this point, your code will deadlock, like so: **fatal error: all goroutines are asleep - deadlock!**. Your code will never finish as that value will never arrive.

The lesson here is that you need to keep track of how many results you might get and only try to receive that many.

There's another way to receive values, and that's by using a `select` like so:

```go
for i := 0; i < 2; i++ {
  select {
  case x, ok := <-ch:
   if ok {
    fmt.Println(x)
   }
  }
 }
```

The idea is to _match_ the receiving of a value like so:

```go
case x, ok := <-ch:
```

What you are getting is two things, the value itself `x` and `bool` we name `ok`. If we managed to get a value ok, then `ok` holds the value `true`. What happens if it's not ok then? It would be `false` if the channel is closed and can no longer produce any more values, so lets discuss that next.

### Closing a channel

A channel is open until you close it. You can actively close it by calling `close()` with the channel as an input parameter:

```go
close(ch)
```

However, when we close a channel, we need to test for it. If we attempt to receive a value from a closed channel, it will cause a crash. To test whether the channel is open or not, we can use the `select` we just wrote:

```go
  select {
  case x, ok := <-ch:
   if ok {
    fmt.Println(x)
   } else {
     break // channel is closed
   }
  }
```

The value of `ok` is now false.

To apply the concept of closing a channel, we add `close()` to `produceResults()` and we have our for loop run one more time than there's values, like so:

```go
package main

import (
 "fmt"
)

func produceResults(ch chan int) {
 ch <- 1
 ch <- 2
 // ch <- 3
 close(ch)
}

func main() {
 ch := make(chan int)
 go produceResults(ch)
 // time.Sleep(1 * time.Second)

 for i := 0; i < 3; i++ {
  select {
  case x, ok := <-ch:
   if ok {
    fmt.Println(x)
   } else {
    fmt.Println("channel closed")
   }
  }
 }
}
```

The output of running said code is:

```output
1
2
channel closed
```

We can see how the `else` clause is matched on the third iteration.

Now, we might have more long running tasks, at which point we need to sit and wait until the channel tells us it closed. Here's code to handle that:

```go
label:
 for {
  select {
  case x, ok := <-ch:
   if ok {
    fmt.Println(x)
   } else {
    fmt.Println("channel closed")
    break label
   }
  }
 }
```

What's happening here is that we set up a for loop that runs forever, until closed. To ensure we break out of the for loop and not just the `select`, we add `label:`

### Using Range over Channel:

You can use range over the channel as well. To use range over a channel, we first need to create a channel and send some values to it. We can then use the range keyword to receive values from the channel.

```go
package main

import "fmt"

func main() {
  ch := make(chan int)
  go func() {
    ch <- 1
    ch <- 2
    // ch <- 3
    close(ch)
    fmt.Println("channel closed")
  }()
  // time.Sleep(1 * time.Second)

  // Use range to receive values from the channel
  for x := range ch {
    fmt.Println(x)
  }
}

```

### Using Range over Buffered Channel

We can also use range over buffered channels in Go. Buffered channels have a buffer that can hold a certain number of values. When we use range over a buffered channel, we will receive values until the buffer is empty.

Here's an example:

```go
package main

import "fmt"

func main() {
    // Create a buffered channel
    ch := make(chan int, 2)

    // Send some values to the channel
    ch := make(chan int, 2)
    ch_message := []int{1, 2, 3}

    // Send some values to the channel
    for x := range ch_message {
      if len(ch) == 2 {
        close(ch)
        break
      } else {
        ch <- x
      }
    }

    // Use range to receive values from the channel
    for i := range ch {
        fmt.Println(i)
    }
}
```

In this example, we create a buffered channel of type int with a buffer size of 2 using the make function. We then send two values, 1 and 2, to the channel.

~~We then use the range keyword to receive values from the channel. Since the channel has a buffer size of 2, we will receive the first two values from the channel. When we receive the third value, the channel is empty and the range loop terminates.~~ It turns out that range still loop over `ch` eventhough channel is already empty and cause fatal error with message _fatal error: all goroutines are asleep - deadlock!_

The output of this program will be:

```terminal
1
2
```

## Assignment - `SearchFiles()` with channels

Let's take all our learning and add channels to the program we wrote containing a file searcher.

## Challenge

## Solution

```go
package main

import (
 "io/ioutil"
 "log"
 "path/filepath"
)

func SearchFiles(dir string, lookFor string, ch chan string) {
 log.Println("[SEARCHING] ", dir)
 files, err := ioutil.ReadDir(dir)
 if err != nil {
  log.Fatal(err)
 }

 for _, file := range files {
  log.Println(dir+file.Name(), file.IsDir())
  if file.Name() == lookFor {
   ch <- "[FOUND] " + filepath.Join(dir, file.Name())
   return
  }
 }
 ch <- "[NOT FOUND] " + dir
}

func main() {
 ch := make(chan string)

 go SearchFiles("./test/", "test2.txt", ch)
 go SearchFiles("./other/", "test2.txt", ch)

 var res = ""
 for i := 0; i < 2; i++ {
  res = <-ch
  log.Println(res)
 }
}
```
