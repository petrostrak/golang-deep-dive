### Shared memory and shared variables
`Shared memory` and `shared variables` are the most common ways for UNIX threads to communicate with each other.

A `mutex` variable, which is an abbreviation of `mutual exclusion` variable, is mainly used for thread synchronization and for protecting shared data when multiple writes can occure at the same time. A mutex works like a buffered channel with a capacity of one, which allows at most one goroutine to access a shared variable at any given time. This means that there is no way for two or more goroutines to try to update that variable simultaneously.

A `critical section` of a concurrent program is the code that cannot be executed simultaneously by all processes, threads or in this case, goroutines. It is the code that need to be protected by mutexes. Therefore, identifying the critical sections of your code will make the whole programming process so much simpler that you should pay particular attention to this task.

A critical section cannot be embedded into another critical section when both critical sections use the same `sync.Mutex` or `sync.RWMutex` variable. Put simply, avoid at almost any cost spreading mutexes across functions because that makes it really hard to see whether you are embedded or not.