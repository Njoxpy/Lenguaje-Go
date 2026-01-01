package main

import "fmt"

type User struct {
	Name string
	Age  int
	Role string
}

// user: login, logout, register
type Product struct {
	ProductName string
	Quantity    int
	Price       float64
}

// product: buy, sell

type UserActions interface {
	Login()
	Logout()
}

type ProductActions interface {
	GetPrice()
	AddProduct()
}

func (u *User) Login() {
	fmt.Println(u.Name, "logged in")
}

func (u *User) Logout() {
	fmt.Println(u.Name, "logged out")
}

func main() {
	fmt.Println("Golang exercise")
}

/*
Make small structs (User, Product)

Implement interfaces (UserActions, AdminActions)

Write functions that accept interfaces

Create a generic repository for multiple types

Practice error handling with error and custom types
*/
