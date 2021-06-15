package main

import (
	"fmt"
	u "github.com/d-sauer/exploring-golang/hello-interface/user"
)

func main() {
	user := &u.User{"Neil", "Young"}
	fmt.Printf("First Name: %s, Last Name: %s \n", user.FirstName, user.LastName)
	fmt.Printf("First & Lase Name len: %d \n", u.GetFirstLastNameLength(user))
}
