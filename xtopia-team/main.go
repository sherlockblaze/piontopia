package main

import "fmt"

type User struct {
	id   int
	name string
}

func main() {
	usera := User{}
	userb := &User{}
	// fmt.Printf("%T\n", usera)
	// fmt.Printf("%T\n", userb)

	test(&usera)
	testb(*userb)

	test(userb)
	testb(*userb)
}

func test(user *User) {
	fmt.Println(user.id)
	fmt.Printf("%T\n", user)
}

func testb(user User) {
	fmt.Println(user.id)
	fmt.Printf("%T\n", user)

}

func (user *User) changeID(user &User) {
	fmt.Print ()
}
