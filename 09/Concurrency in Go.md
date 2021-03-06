### Concurrency in Go

### About processes, threads and goroutines
A `process` is an execution environment that contains instructions, user data and system data parts, as well as other types of resources that are obtained during runtime, whereas a program is a file that contains instructions and data that are used for initializing the instruction and user-data parts of the process.

A `thread` is a smaller and lighter entity than a process or a program. Threads are created by processes and have their own flow of control and stack. A quick and simplistic way to differentiate a thread from a process is to consider a process as the running binary file and a thread as a subset of a `process`.

A `goroutine` is the minimun Go entity that can be executed concurrently. The use of the word "minimum" is very important here, as goroutines are not autonomous entities like UNIX processes - goroutines live in UNIX `threads` that live in UNIX processes. The main advantage of goroutines is that they are extremely lightweight and running thousands or hundreds of thousands of them on a single machine is not a problem.

The good thing is that goroutines are lighter that threads, which, in turn are lighter than processes. In practice, this means that a process can have multiply threads as well as lots of goroutines, whereas a goroutine needs the environment of a process in order to exist. So, in order to create a goroutine you will need to have a process with at least one thread - UNIX takes care of the `process` and thread management, while Go and the developer need to take care of the goroutines.

### The Go scheduler
The UNIX kerner `scheduler` is responsible for the execution of the theads of a program. On the other hand, the Go runtime has its own scheduler, which is responsible for the execution of the goroutines using a technique known as `m:n scheduling`, where `m` goroutines are executed using `n` operating system threads using multiplexing. The Go scheduler is the Go compoment responsible for the way and the order in which the goroutines of a Go program get executed. This makes the Go scheduler a really important part of the Go programming language, as everything in a Go program is executed as a goroutine.

Be aware that as the Go scheduler only deals with the goroutines of a single program, its operation is much simpler, cheaper and faster that the operation of the kernel scheduler.

### Concurrency and parallelism
It is very common misconception that `concurrency` is the same as `parallelism` - this is not true. Parallelism is the simultaneous execution of multiple entities of some kind, whereas concurrency is a way of structuring your components so that they can be executed independently when possible.

In a valid concurrent design, adding concurrent entities makes the whole system run faster because more things can be executed in parallel. So, the desired perallelism comes from a better concurrent expression and implementation of the problem.

### Goroutines
You can define, create and execute a new goroutine using the `go` keyword followed by a function name of the full definition of an `anonymous function`. The `go` keyword makes the function call return immediately, while the function starts running in the background as a goroutine and the rest of the program continues its execution.

### Channels
A channel is acommunication mechanism that allows goroutines to exchange data among other things.
However, there are some specific rules. Firstly, each channel allows the exchange of a particular data type, which is also called the `element type` of the channel, and secondly, for a channel to operate properly, you will need someone to receive what is sent via the channel. You should declare a new channel using the `chan` keyword and you can close a channel using the `close()` function.

Finally, a very important detail: when you are using a channel as a function parameter, you can specify its direction; whetherit is going to be used for sending or receiveing. If you know the purpose of a channel in advance, you should use this capability because it will make your programs more robust, as well faster. You will not be able to send data accidentlly to a channel from wich you should only receive data.

### Pipelines
A `pipeline` is a virtual method for connecting goroutines and channels so that the output of one goroutine becomes the input of another goroutine using channels to transfer your data.

One of the benefits that you get from using pipelines is that there is a constant data flow in your program, as no goroutine and channel have to wait for everything to be completed in order to start their execution. Additionally, you see fewer variables and therefore less memory space because you do not have to save everything as a variable. Finally, the use of pipelines simplifies the design of the program and improves its maintainability.