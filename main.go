package main

import (
	"fmt"
)

type Contact struct {
	mobile int
	name   string
	email  string
	phone
}

type phone interface {
	call()
	message()
	edit()
	delt()
	print()
}

func (member *Contact) call() {
	fmt.Println("Calling the person ", member.name)
	fmt.Println("Calling the mail ", member.email)
	fmt.Println("Calling the mobile ", member.mobile)
}

func (member *Contact) message(message string) {
	fmt.Printf("The message is %s\n", message)
}

func (member *Contact) edit(name, mail string, phone int) {
	member.name = name
	member.email = mail
	member.mobile = phone
	member.print()
}

func (member *Contact) delt() {
	member.name = ""
	member.email = ""
	member.mobile = 0
}

func (member *Contact) print() {
	fmt.Println("\nName: ", member.name)
	fmt.Println("Email: ", member.email)
	fmt.Println("Mobile: ", member.mobile)
}

func main() {
	var member Contact
	memberPtr := &member
	memberPtr.name = "Srinivasan"
	memberPtr.email = "test@mail.com"
	memberPtr.mobile = 12345
	memberPtr.call()
	memberPtr.message("Hello how are you ?")
	newname := "Raju"
	newmail := "raju@mail.com"
	newphone := 54321
	memberPtr.edit(newname, newmail, newphone)
	memberPtr.call()
	memberPtr.message("Hey, what's up ?")
	memberPtr.delt()
	memberPtr.print()
}
