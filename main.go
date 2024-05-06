package main

import (
	"encoding/json"
	"fmt"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Contact struct {
	ID int `json:"id"`
	FirstName  string   `json:"first_name"`
	MiddleName string   `json:"middle_name"`
	LastName   string   `json:"last_name"`
	Mobile     string      `json:"mobile"`
	Email      string   `json:"mail"`
	Company    string   `json:"company"`
	Location   string   `json:"location"`
	AddMobile  []string    `json:"extra_mobile"`
	AddEmails  []string `json:"extra_mails"`
	Phone
}

type Phone interface {
	Call()
	Message()
	Edit()
	Delete()
}

func (member *Contact) Call() {
	fmt.Printf("Calling to Mr. %s %s %s", member.FirstName, member.MiddleName, member.LastName)
}

func (member *Contact) Message(message string) {
}

func (member *Contact) Edit(name, mail string, phone int) {
	// update contact set first_name="Mr";
}

func (member *Contact) Delete() {
	// delete from contact where middle_name="Srinivasan";
}

func CreateContact(db *sql.DB) error {
	
	var member Contact

	_, err := db.Exec("create table if not exists contact (id int auto_increment primary key, first_name varchar(255) unique, middle_name varchar(255) unique, last_name varchar(255) unique, mobile varchar(255), mail varchar(255), company varchar(255), location varchar(255), extra_mobile varchar(255), extra_mails varchar(255));")
	if err != nil {
		fmt.Printf("\nError to create db: %s\n", err)
		return err
	}

	member.FirstName = "S"
	member.MiddleName = "Srinivasan"
	member.LastName = "S"
	member.Mobile = "9943080828"
	member.Email = "srinivasan@athinio.com"
	member.Company = "Neridio"
	member.Location = "Bangalore"
	member.AddMobile = append(member.AddMobile,"8220155210")
	member.AddEmails = append(member.AddEmails,"mrsrinivasanofficial@gmail.com")

	_, err = db.Exec("insert into contact (first_name, middle_name, last_name, mobile, mail, company, location, extra_mobile, extra_mails) values(?, ?, ?, ?, ?, ?, ?, ?, ?);", member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, member.AddMobile[0], member.AddEmails[0])
	if err != nil {
		fmt.Printf("\nError to insert values: %s\n", err)
		return err
	}

	return nil
}

func CallContact(db *sql.DB, name string) error {
	
	var member Contact
	var extraMobile, extraMail string

	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &extraMobile, &extraMail)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nNo rows found %s: %s", name, err)
		} else {
			fmt.Printf("\nError to query %s: %s", name, err)
			return err
		}
	}

	fmt.Println(member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, extraMobile, extraMail)
	member.Call()

	return nil
}

func SearchContact(db *sql.DB, name string) error {
	
	var member Contact
	var extraMobile, extraMail string

	err := db.QueryRow("select * from contact where first_name = ? or middle_name = ? or last_name = ? or mobile = ? or mail = ? or company = ? or location = ? or extra_mobile = ? or extra_mails = ?;", name, name, name, name, name, name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &extraMobile, &extraMail)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nNo rows found %s: %s", name, err)
		} else {
			fmt.Printf("\nError to query %s: %s", name, err)
			return err
		}
	}

	fmt.Println(member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, extraMobile, extraMail)
	return nil
}

func SendMessage(name, message string) {
	// create table if not exists message (contact varchar(20), message varchar(20));

}

func SearchMessage(message string) error {
	var member Contact
	data, err := os.ReadFile("message.json")
	if err != nil {
		fmt.Printf("\nError in read: %s", err)
		return err
	}
	err = json.Unmarshal(data, &member)
	if err != nil {
		fmt.Printf("\nError in json unmarshal: %s", err)
		return err
	}
	fmt.Println(member)
	return nil
}

func GetTop10Contact() error {
	var member Contact
	data, err := os.ReadFile("activity.json")
	if err != nil {
		fmt.Printf("\nError in read file: %s", err)
		return err
	}
	err = json.Unmarshal(data, &member)
	if err != nil {
		fmt.Printf("\nError in json unmarshal: %s", err)
		return err
	}
	fmt.Println(member)
	return nil
}

func GetCallHistory() error {
	var member Contact
	data, err := os.ReadFile("history.json")
	if err != nil {
		fmt.Printf("\nError in read file: %s", err)
		return err
	}
	err = json.Unmarshal(data, &member)
	if err != nil {
		fmt.Printf("\nError in json unmarshal: %s", err)
		return err
	}
	fmt.Println(member)
	return nil
}

func main() {
	
	db, err := sql.Open("mysql", "root:athinio@tcp(127.0.0.1:3306)/contacts")
	if err != nil {
		fmt.Printf("\nError to open db: %s", err)
		return
	}
	defer db.Close()

	CreateContact(db)
	SearchContact(db, "Srinivasan")
	CallContact(db, "S")

	return
}
