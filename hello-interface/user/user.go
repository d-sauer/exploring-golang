package user

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

type userClient interface {
	FirstLastName() string
}

func (user *User) FirstLastName() string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}

func GetFirstLastNameLength(client userClient) int {
	return len(client.FirstLastName())
}
