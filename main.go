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

func (member *Contact) Edit(name, mail string, phone int) {
	// update contact set first_name="Mr";
}

func (member *Contact) Delete() {
	// delete from contact where middle_name="Srinivasan";
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

	fmt.Printf("\nIn CallContact() %s", name)
	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &extraMobile, &extraMail)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nNo rows found %s: %s", name, err)
		} else {
			fmt.Printf("\nError to query %s: %s", name, err)
			return err
		}
	}

	fmt.Println("\n", member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, extraMobile, extraMail)
	member.Call(db)

	return nil
}

func SearchContact(db *sql.DB, name string) error {
	
	var member Contact
	var extraMobile, extraMail string

	fmt.Printf("\nIn SearchContact() %s", name)
	err := db.QueryRow("select * from contact where first_name = ? or middle_name = ? or last_name = ? or mobile = ? or mail = ? or company = ? or location = ? or extra_mobile = ? or extra_mails = ?;", name, name, name, name, name, name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &extraMobile, &extraMail)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nNo rows found %s: %s", name, err)
		} else {
			fmt.Printf("\nError to query %s: %s", name, err)
			return err
		}
	}

	fmt.Println("\n", member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, extraMobile, extraMail)
	return nil
}

func SendMessage(db *sql.DB, name, message string) error {
	
	var member Contact
	var extraMobile, extraMail string

	fmt.Printf("\nIn SendMessage() %s %s", name, message)
	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &extraMobile, &extraMail)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nNo rows found %s: %s", name, err)
		} else {
			fmt.Printf("\nError to query %s: %s", name, err)
			return err
		}
	}

	fmt.Println("\n", member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, extraMobile, extraMail)
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
	SearchMessage(db, "Hey! what's up ?")
	// GetCallHistory(db)

	return
}
