package main

import (
	"fmt"
	"time"
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
	AddMobile  string    `json:"extra_mobile"`
	AddEmails  string `json:"extra_mails"`
	Phone
}

type Phone interface {
	Call()
	Message()
	Edit()
	Delete()
}

func (member *Contact) Call(db *sql.DB) error {
	
	fmt.Printf("\nIn Call()")

	call := fmt.Sprintf("%v : Called to Mr. %s %s %s [mobile: %s]", time.Now(), member.FirstName, member.MiddleName, member.LastName, member.Mobile)

	_, err := db.Exec("create table if not exists history (id int auto_increment primary key, call_history varchar(255), message_history varchar(255))")
	if err != nil {
		fmt.Printf("\nError to create db: %s", err)
		return err
	}

	_, err = db.Exec("insert into history (call_history) values (?)", call)
	if err != nil {
		fmt.Printf("\nError to insert value: %s", err)
		return err
	}

	return nil
}

func (member *Contact) Message(db *sql.DB, message string) error {
	
	fmt.Printf("\nIn Message() message %s", message)

	messageHis := fmt.Sprintf("%v : %s message sent to Mr. %s %s %s [mobile: %s]", time.Now(), message, member.FirstName, member.MiddleName, member.LastName, member.Mobile)

	_, err := db.Exec("create table if not exists history (id int auto_increment primary key, call_history varchar(255), message_history varchar(255))")
	if err != nil {
		fmt.Printf("\nError to create db: %s", err)
		return err
	}

	_, err = db.Exec("insert into history (message_history) values (?)", messageHis)
	if err != nil {
		fmt.Printf("\nError to insert value: %s", err)
		return err
	}

	return nil
}

func (member *Contact) Edit(db *sql.DB) error {
	
	fmt.Printf("\nIn Edit()")
	
	_, err := db.Exec("update contact set id = ?, first_name = ?, middle_name = ?, last_name = ?, mobile = ?, mail = ?, company = ?, location = ?, extra_mobile = ?, extra_mails = ?", member.ID, member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, member.AddMobile, member.AddEmails);
	if err != nil {
		fmt.Printf("\nError to update db: %s", err)
		return err
	}

	return nil
}

func (member *Contact) Delete(db *sql.DB) error {
	
	fmt.Printf("\nIn Delete()")
	
	_, err := db.Exec("delete from contact where id = ?", member.ID);
	if err != nil {
		fmt.Printf("\nError to update db: %s", err)
		return err
	}

	return nil
}

func CreateContact(db *sql.DB) error {
	
	var member Contact

	fmt.Printf("\nIn CreateContact()")
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
	member.AddMobile = "8220155210"
	member.AddEmails = "mrsrinivasanofficial@gmail.com"

	_, err = db.Exec("insert into contact (first_name, middle_name, last_name, mobile, mail, company, location, extra_mobile, extra_mails) values(?, ?, ?, ?, ?, ?, ?, ?, ?);", member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, member.AddMobile, member.AddEmails)
	if err != nil {
		fmt.Printf("\nError to insert values: %s\n", err)
		return err
	}

	return nil
}

func CallContact(db *sql.DB, name string) error {
	
	var member Contact

	fmt.Printf("\nIn CallContact() %s", name)
	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &member.AddMobile, &member.AddEmails)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nNo rows found %s: %s", name, err)
		} else {
			fmt.Printf("\nError to query %s: %s", name, err)
			return err
		}
	}

	fmt.Println("\n", member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, member.AddMobile, member.AddEmails)
	member.Call(db)

	return nil
}

func SearchContact(db *sql.DB, name string) error {
	
	var member Contact

	fmt.Printf("\nIn SearchContact() %s", name)
	err := db.QueryRow("select * from contact where first_name = ? or middle_name = ? or last_name = ? or mobile = ? or mail = ? or company = ? or location = ? or extra_mobile = ? or extra_mails = ?;", name, name, name, name, name, name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &member.AddMobile, &member.AddEmails)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nNo rows found %s: %s", name, err)
		} else {
			fmt.Printf("\nError to query %s: %s", name, err)
			return err
		}
	}

	fmt.Println("\n", member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, member.AddMobile, member.AddEmails)
	return nil
}

func SendMessage(db *sql.DB, name, message string) error {
	
	var member Contact

	fmt.Printf("\nIn SendMessage() %s %s", name, message)
	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &member.AddMobile, &member.AddEmails)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nNo rows found %s: %s", name, err)
		} else {
			fmt.Printf("\nError to query %s: %s", name, err)
			return err
		}
	}

	fmt.Println("\n", member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, member.AddMobile, member.AddEmails)
	member.Message(db, message)

	return nil
}

func SearchMessage(db *sql.DB, message string) error {
	
	var msg string

	fmt.Printf("\nIn SearchMessage() %s", message)
	datas, err := db.Query("select message_history from history;")
	if err != nil {
		fmt.Printf("\nError to query %s: %s", message, err)
		return err
	}
	defer datas.Close()

	for datas.Next() {
		err = datas.Scan(&msg)
		if err != nil {
			fmt.Printf("\nError to scan: %s", err)
		}
		fmt.Printf("\nmsg %s", msg)
	}

	err = datas.Err()
	if err != nil {
		fmt.Printf("\nError in row scan: %s", err)
	}

	return nil
}

func EditContact(db *sql.DB, name string) error {
	
	var oldmember, newmember Contact

	fmt.Printf("\nIn EditContact() %s", name)
	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&oldmember.ID, &oldmember.FirstName, &oldmember.MiddleName, &oldmember.LastName, &oldmember.Mobile, &oldmember.Email, &oldmember.Company, &oldmember.Location, &oldmember.AddMobile, &oldmember.AddEmails)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nNo rows found %s: %s", name, err)
		} else {
			fmt.Printf("\nError to query %s: %s", name, err)
			return err
		}
	}

	newmember.ID = oldmember.ID
	fmt.Printf("\nExisting first name: %s, Enter the new one: ", oldmember.FirstName)
	fmt.Scan(&newmember.FirstName)
	fmt.Printf("\nExisting middle name: %s, Enter the new one: ", oldmember.MiddleName)
	fmt.Scan(&newmember.MiddleName)
	fmt.Printf("\nExisting last name: %s, Enter the new one: ", oldmember.LastName)
	fmt.Scan(&newmember.LastName)
	fmt.Printf("\nExisting mobile number: %s, Enter the new one: ", oldmember.Mobile)
	fmt.Scan(&newmember.Mobile)
	fmt.Printf("\nExisting email name: %s, Enter the new one: ", oldmember.Email)
	fmt.Scan(&newmember.Email)
	fmt.Printf("\nExisting company name: %s, Enter the new one: ", oldmember.Company)
	fmt.Scan(&newmember.Company)
	fmt.Printf("\nExisting location name: %s, Enter the new one: ", oldmember.Location)
	fmt.Scan(&newmember.Location)
	fmt.Printf("\nExisting extra mobile name: %s, Enter the new one: ", oldmember.AddMobile)
	fmt.Scan(&newmember.AddMobile)
	fmt.Printf("\nExisting extra email: %s, Enter the new one: ", oldmember.AddEmails)
	fmt.Scan(&newmember.AddEmails)
	
	newmember.Edit(db)

	return nil
}

func DeleteContact(db *sql.DB, name string) error {

	var member Contact

	fmt.Printf("\nIn DeleteContact() %s", name)
	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &member.AddMobile, &member.AddEmails)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nNo rows found %s: %s", name, err)
		} else {
			fmt.Printf("\nError to query %s: %s", name, err)
			return err
		}
	}

	fmt.Printf("\nDo you want to delete this %s %s %s contact ?", member.FirstName, member.MiddleName, member.LastName)
	member.Delete(db)

	return nil
}

func GetTop10Contact(db *sql.DB) error {
	
	var id int
	var call, msg string

	fmt.Printf("\nIn GetTop10Contact()")
	datas, err := db.Query("select * from history;")
	if err != nil {
		fmt.Printf("\nError to query: %s", err)
		return err
	}
	defer datas.Close()

	for datas.Next() {
		err = datas.Scan(&id, &call, &msg)
		if err != nil {
			fmt.Printf("\nError to scan: %s", err)
		}
		fmt.Printf("\nid %d, call %s, message %v", id, call, msg)
	}

	err = datas.Err()
	if err != nil {
		fmt.Printf("\nError in row scan: %s", err)
	}

	return nil
}

func GetCallHistory(db *sql.DB) error {

	var call string

	fmt.Printf("\nIn GetCallHistory()")
	datas, err := db.Query("select call_history from history;")
	if err != nil {
		fmt.Printf("\nError to query: %s", err)
		return err
	}
	defer datas.Close()

	for datas.Next() {
		err = datas.Scan(&call)
		if err != nil {
			fmt.Printf("\nError to scan: %s", err)
		}
		fmt.Printf("\ncall history: %s", call)
	}

	err = datas.Err()
	if err != nil {
		fmt.Printf("\nError in row scan: %s", err)
	}

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
	CallContact(db, "S")
	SearchContact(db, "Srinivasan")
	SendMessage(db, "Srinivasan", "Hey! what's up ?")
	// SearchMessage(db, "Hey! what's up ?")
	EditContact(db, "Srinivasan")
	DeleteContact(db, "Devil")

	return
}
