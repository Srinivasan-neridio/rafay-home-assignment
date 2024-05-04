package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	FirstName  string   `json:"first_name"`
	MiddleName string   `json:"middle_name"`
	LastName   string   `json:"last_name"`
	Mobile     int      `json:"mobile"`
	Email      string   `json:"mail"`
	Company    string   `json:"company"`
	Location   string   `json:"location"`
	AddMobile  []int    `json:"extra mobile"`
	AddEmails  []string `json:"extra mails"`
	phone
}

type phone interface {
	Call()
	Message()
	Edit()
	Delete()
}

func (member *Contact) Call() {
	fmt.Printf("Calling Mr %s %s %s", member.FirstName, member.MiddleName, member.LastName)
}

func (member *Contact) Message(message string) {
}

func (member *Contact) Edit(name, mail string, phone int) {
}

func (member *Contact) Delete() {
}

func CreateContact() error {
	var member Contact
	fmt.Printf("\nEnter your first name: ")
	fmt.Scanf("%s", &member.FirstName)
	fmt.Printf("\nEnter your middled name: ")
	fmt.Scanf("\n%s", &member.MiddleName)
	fmt.Printf("\nEnter your last name: ")
	fmt.Scanf("\n%s", &member.LastName)
	fmt.Printf("\nEnter your mobile number name: ")
	fmt.Scanf("\n%d", &member.Mobile)
	fmt.Printf("\nEnter your email id: ")
	fmt.Scanf("\n%s", &member.Email)
	fmt.Printf("\nEnter your company name: ")
	fmt.Scanf("\n%s", &member.Company)
	data, err := json.Marshal(member)
	if err != nil {
		fmt.Printf("\nJson marshalling failed: %s", err)
		return err
	}
	err = os.WriteFile("contact.json", data, 0777)
	if err != nil {
		fmt.Printf("\nWrite failed: %s", err)
		return err
	}
	return nil
}

func main() {
	CreateContact()
}
