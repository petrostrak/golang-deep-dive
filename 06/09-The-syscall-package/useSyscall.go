package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("My process id is:", pid)

	uid, _, _ := syscall.Syscall(24, 0, 0, 0)
	fmt.Println("The User id is:", uid)

	msg := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	fd := 1
	syscall.Write(fd, msg)

	fmt.Println("Using syscall.Exec()")
	cmd := "/bin/ls"
	env := os.Environ()
	syscall.Exec(cmd, []string{"ls", "-a", "-x"}, env)
}
