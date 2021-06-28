# Golang deep dive
Create Golang production applications using network libraries, concurrency, machine learning and advanced data structures.

* Clear guidance on using Go for production systems
* Detailed explanations of how Go internals work, the disign choices behind the language, and how to optimize your Go code.
* A full guide to all Go data types, composite types and data structures
* Master packages, reflection, and interfaces for effective Go programming
* Build high-performance systems networking code, including server and client-side applications
* Interface with other systems using WebAssembly, JSON and gRPC
* Write reliable, high-performance concurrent code
* Build machine learning systems in Go, from simple statistical regression to complex neural networks

### UNIX stdin, stdout, and stderr
Every UNIX OS has three files open all the time for its processes. Remember that UNIX considers everything, even a printer or your mouse, a file. UNIX uses `file descriptors`, which are positive integer values, as an internal representation for accessing all of its open files, which is much prettier than using long paths. So, by default, all UNIX systems support three special and standard filenames: `/dev/stdin` , `/dev/stdout` , and `/dev/stderr` , which can also be accessed using file descriptors `0`, `1`, and `2`, respectively. These three file descriptors are also called `standard input`, `standard output`, and `standard error`, respectively. Additionally, file descriptor 0 can be accessed as both `/dev/fd/0` and `/dev/pts/0` on a Debian Linux machine.
