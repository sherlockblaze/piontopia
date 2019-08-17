package main

import "fmt"

// User 开放的结构体需要有注释
type User struct {
	id   int
	name string
}

func main() {
	usera := User{0, "ceshi_a"}
	userb := &User{0, "ceshi_b"}
	// fmt.Printf("%T\n", usera)
	// fmt.Printf("%T\n", userb)

	fmt.Println("usera, which initialized using no &:")
	test(&usera)
	testb(usera)

	fmt.Println("userb, which initialized using &:")
	test(userb)
	testb(*userb)

	fmt.Println("change using *User binder:")
	usera.changeIDp("t")
	fmt.Println(usera.id)

	fmt.Println("change using User binder:")
	userb.changeID("t")
	fmt.Println(userb.id)

}

func test(user *User) {
	fmt.Println(user.name)
	fmt.Printf("%T\n", user)
}

func testb(user User) {
	fmt.Println(user.name)
	fmt.Printf("%T\n", user)

}

func (user *User) changeIDp(a string) {
	user.name = a
}

func (user User) changeID(a string) {
	user.name = a
}
