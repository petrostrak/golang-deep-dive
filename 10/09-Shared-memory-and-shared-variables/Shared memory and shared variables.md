### Shared memory and shared variables
`Shared memory` and `shared variables` are the most common ways for UNIX threads to communicate with each other.

A `mutex` variable, which is an abbreviation of `mutual exclusion` variable, is mainly used for thread synchronization and for protecting shared data when multiple writes can occure at the same time. A mutex works like a buffered channel with a capacity of one, which allows at most one goroutine to access a shared variable at any given time. This means that there is no way for two or more goroutines to try to update that variable simultaneously.

A `critical section` of a concurrent program is the code that cannot be executed simultaneously by all processes, threads or in this case, goroutines. It is the code that need to be protected by mutexes. Therefore, identifying the critical sections of your code will make the whole programming process so much simpler that you should pay particular attention to this task.

A critical section cannot be embedded into another critical section when both critical sections use the same `sync.Mutex` or `sync.RWMutex` variable. Put simply, avoid at almost any cost spreading mutexes across functions because that makes it really hard to see whether you are embedded or not.

### The synx.RWMutex
The `sync.RWMutex` type is another kind of mutex - actually, it is an improved version of `sync.Mutex` which is defined in the `sync` directory as follows:
```
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // semaphore for writers to wait for completing readers
	readerSem   uint32 // semaphore for readers to wait for completing writers
	readerCount int32  // number of pending readers
	readerWait  int32  // number of departing readers
}
```

In other words, `sync.RWMutex` is based on `syncMutex` with the necessary additions and improvements. Now lets talk about how `sync.RWMutex` improves `sync.Mutex`. Although only one function is allowed to perform write operations using `sync.RWMutex` mutex, you can have multiple readers owning a `sync.RWMutex` mutex. However, there is one thing that you should be aware of: until all of the readers of a `sync.RWMutex` mutex unlock that mutex, you cannot lock it for writing, which is the small price you have to pay for allowing multiple readers.

The functions that can help you work with a `sync.RWMutex` mutex are `RLock()` and `RUnlock()`, which are used for locking and unlocking the mutex for reading purposes, respectively. The `Lock()` and `Unlock()` functions used in a `sync.Mutex` mutex should still be used when you want to lock and unlock a `sync.Mutex` mutex for writing purposes. Hence, an `RLock()` function call that locks for reading purposes should be paired with an `RUnlock()` function call. Finally, it should be apparent that you should not make changes to any shared variables inside an `RLock()` and `RUnlock()` block of code.

### The atomic package
An `atomic operation` is an operation that is completed in a single step relative to other threads or in this case to other goroutines. This means that an atomic operation cannot be interrupted in the middle of it.

The Go standard library offers the `atomic` package which in some cases can help you to avoid using mutex. However mutexes are more versitile that atomic operations. Using the `atomic` package, you can habe atomic counters accessed by multiple goroutines without synchronization issues and race conditions.

### Sharing memory using goroutines
Sharing memory using goroutines is a topic that describes how you can share data using a dedicated goroutine. Although shared memory is the traditional way that threads communicate with each other, Go comes with built-in synchronization features that allow a single goroutine to own a shared piece of data. This means that other goroutines must send messages to this single goroutine that owns the shared data, which prevents the corruption of the data. Such a goroutine is called a `monitor goroutine`. In Go terminology, this is `sharing by communication instead of communication by sharing`.

### Catching race conditions
A `data race condition` is a situation where two or more running elements, such as threads and goroutines, try to take control of or modify a shared resource or a variable of a program. Strictly speaking, a data race occurs when two or more instructions access the same memory address, where at least one of them performs a write operation. If all operations are read operations, then there is no race condition.

Using the `-race` flag when running or building a Go source file will turn on the Go `race detector`, which will make the compiler create a modified version of typical executable file. This modified version can record all access to shared variables as well as all synchronization events that take place, including calls to `sync.Mutex` and `sync.WaitGroup`. After analyzing the relevant events, the race detector prints a report that can help you to identify potential problems so that you can correct them.