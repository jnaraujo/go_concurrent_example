package main

import (
	"fmt"

	"github.com/jnaraujo/go_concurrent_example/mapmutex"
)

type User struct {
	ID   int
	Name string
	tech []string
}

func main() {
	users := mapmutex.New[int, User]()

	users.Write(0, User{
		ID: 1, Name: "John Doe", tech: []string{"Go", "Python"},
	})

	users.Write(1, User{
		ID: 2, Name: "Jane Doe", tech: []string{"Go", "Python", "Java"},
	})

	users.Write(2, User{
		ID: 3, Name: "John Smith", tech: []string{"Go", "Python", "Java", "C++"},
	})

	users.Delete(1)

	user := users.Read(0)

	fmt.Printf("User: %v\n", user)

	println("Hello, World!")
}
