package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {

	// os.Getuid finds the user ID of the current user
	fmt.Println("User id:", os.Getuid())

	var u *user.User
	u, _ = user.Current()
	fmt.Println("Group ids:")

	// Find the group IDs to which the user belongs
	groupIDs, _ := u.GroupIds()
	for _, i := range groupIDs {
		fmt.Print(i, " ")
	}
	fmt.Println()

}
