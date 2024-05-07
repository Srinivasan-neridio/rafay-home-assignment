package main

import (
	"fmt"
	"time"
	"os"
	"bufio"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

type Contact struct {
	ID int              `json:"id"`
	FirstName  string   `json:"first_name"`
	MiddleName string   `json:"middle_name"`
	LastName   string   `json:"last_name"`
	Mobile     string   `json:"mobile"`
	Email      string   `json:"mail"`
	Company    string   `json:"company"`
	Location   string   `json:"location"`
	AddMobile  string   `json:"extra_mobile"`
	AddEmails  string   `json:"extra_mails"`
	Phone
}

type Phone interface {
	Call()
	Message()
	Edit()
	Delete()
}

func (member *Contact) Call(db *sql.DB) error {
	
	callHistory := fmt.Sprintf("%v : Called to Mr. %s %s %s [mobile: %s]", time.Now(), member.FirstName, member.MiddleName, member.LastName, member.Mobile)

	_, err := db.Exec("create table if not exists history (id int auto_increment primary key, call_history varchar(255));")
	if err != nil {
		fmt.Printf("\nError: Failed to create db in Call: %s\n", err)
		return err
	}

	_, err = db.Exec("insert into history (call_history) values (?);", callHistory)
	if err != nil {
		fmt.Printf("\nError: Failed to insert value in Call: %s\n", err)
		return err
	}

	fmt.Printf("\nSuccess: %s", callHistory)
	UpdateActivity(db, callHistory)
	return nil
}

func (member *Contact) Message(db *sql.DB, message string) error {
	
	messageHistory := fmt.Sprintf("%v : %s message sent to Mr. %s %s %s [mobile: %s]", time.Now(), message, member.FirstName, member.MiddleName, member.LastName, member.Mobile)

	_, err := db.Exec("create table if not exists message (id int auto_increment primary key, message_history varchar(255));")
	if err != nil {
		fmt.Printf("\nError: Failed to create db in Message: %s\n", err)
		return err
	}

	_, err = db.Exec("insert into message (message_history) values (?);", messageHistory)
	if err != nil {
		fmt.Printf("\nError: Failed to insert value in Message: %s\n", err)
		return err
	}

	fmt.Printf("\nSuccess: %s", messageHistory)
	UpdateActivity(db, messageHistory)
	return nil
}

func (member *Contact) Edit(db *sql.DB) error {
	
	_, err := db.Exec("update contact set first_name = ?, middle_name = ?, last_name = ?, mobile = ?, mail = ?, company = ?, location = ?, extra_mobile = ?, extra_mails = ? where id = ?;", member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, member.AddMobile, member.AddEmails, member.ID);
	if err != nil {
		fmt.Printf("\nError: Failed to update db in Edit: %s\n", err)
		return err
	}

	fmt.Printf("\nSuccess: Contact id = %d updated\n", member.ID)
	return nil
}

func (member *Contact) Delete(db *sql.DB) error {
	
	_, err := db.Exec("delete from contact where id = ?;", member.ID);
	if err != nil {
		fmt.Printf("\nError: Failed to update db in Delete: %s\n", err)
		return err
	}

	fmt.Printf("\nSuccess: Contact id = %d deleted\n", member.ID)
	return nil
}

func ValidateInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func CreateContact(db *sql.DB) error {
	
	var member Contact

	_, err := db.Exec("create table if not exists contact (id int auto_increment primary key, first_name varchar(255) unique, middle_name varchar(255) unique, last_name varchar(255) unique, mobile varchar(255), mail varchar(255), company varchar(255), location varchar(255), extra_mobile varchar(255), extra_mails varchar(255));")
	if err != nil {
		fmt.Printf("\nError: Failed to create db in CreateContact: %s\n", err)
		return err
	}

	fmt.Printf("\n\n**************** Create Contact ****************\n")
	fmt.Printf("\nNote: Press enter to skip the field\n")

	fmt.Printf("\nEnter your first name: ")
	member.FirstName = ValidateInput()
	fmt.Printf("\nEnter your middle name: ")
	member.MiddleName = ValidateInput()
	fmt.Printf("\nEnter your last name: ")
	member.LastName = ValidateInput()
	fmt.Printf("\nEnter your mobile number: ")
	member.Mobile = ValidateInput()
	fmt.Printf("\nEnter your email id: ")
	member.Email = ValidateInput()
	fmt.Printf("\nEnter your company name: ")
	member.Company = ValidateInput()
	fmt.Printf("\nEnter your location: ")
	member.Location = ValidateInput()
	fmt.Printf("\nEnter extra mobile number: ")
	member.AddMobile = ValidateInput()
	fmt.Printf("\nEnter extra email id: ")
	member.AddEmails = ValidateInput()

	_, err = db.Exec("insert into contact (first_name, middle_name, last_name, mobile, mail, company, location, extra_mobile, extra_mails) values(?, ?, ?, ?, ?, ?, ?, ?, ?);", member.FirstName, member.MiddleName, member.LastName, member.Mobile, member.Email, member.Company, member.Location, member.AddMobile, member.AddEmails)
	if err != nil {
		fmt.Printf("\nError: Failed to insert values in CreateContact: %s\n", err)
		return err
	}

	fmt.Printf("\nSuccess: Created the new contact with the respected values\n")
	return nil
}

func CallContact(db *sql.DB, name string) error {
	
	var member Contact

	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &member.AddMobile, &member.AddEmails)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nInfo: No data found with the %s contact name\n", name)
			return nil
		} else {
			fmt.Printf("\nError: Failed to query %s in CallContact: %s\n", name, err)
			return err
		}
	}

	member.Call(db)
	return nil
}

func SearchContact(db *sql.DB, name string) error {
	
	var member Contact

	err := db.QueryRow("select * from contact where first_name = ? or middle_name = ? or last_name = ? or mobile = ? or mail = ? or company = ? or location = ? or extra_mobile = ? or extra_mails = ?;", name, name, name, name, name, name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &member.AddMobile, &member.AddEmails)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nInfo: No data found with the %s contact name\n", name)
			return nil
		} else {
			fmt.Printf("\nError to query %s in SearchContact: %s\n", name, err)
			return err
		}
	}

	fmt.Printf("\n\n**************** Search Contact Result ****************\n")
	fmt.Printf("\nYour first name: %s", member.FirstName)
	fmt.Printf("\nYour middle name: %s", member.MiddleName)
	fmt.Printf("\nYour last name: %s", member.LastName)
	fmt.Printf("\nYour mobile number: %s", member.Mobile)
	fmt.Printf("\nYour email id: %s", member.Email)
	fmt.Printf("\nYour company name: %s", member.Company)
	fmt.Printf("\nYour location name: %s", member.Location)
	fmt.Printf("\nYour extra mobile number: %s", member.AddMobile)
	fmt.Printf("\nYour extra email id: %s", member.AddEmails)

	return nil
}

func SendMessage(db *sql.DB, name, message string) error {
	
	var member Contact

	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &member.AddMobile, &member.AddEmails)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nInfo: No data found with the %s contact name to send message\n", name)
			return nil
		} else {
			fmt.Printf("\nError to query %s in SendMessage: %s\n", name, err)
			return err
		}
	}

	member.Message(db, message)
	return nil
}

func SearchMessage(db *sql.DB, message string) error {
	
	var msg string
	var msgSearch []string

	datas, err := db.Query("select message_history from message;")
	if err != nil {
		fmt.Printf("\nError to query %s in SearchMessage: %s\n", message, err)
		return err
	}
	defer datas.Close()

	for datas.Next() {
		err = datas.Scan(&msg)
		if err != nil {
			fmt.Printf("\nError to scan in SearchMessage: %s\n", err)
		}
		msgSearch = append(msgSearch, msg)
	}

	err = datas.Err()
	if err != nil {
		fmt.Printf("\nError in row scan in SearchMessage: %s\n", err)
	}

	fmt.Printf("\n\n**************** Search result for: %s in messages ****************\n", message)
	for _, value := range msgSearch {
		fmt.Printf("\n%s", value)
	}

	return nil
}

func EditContact(db *sql.DB, name string) error {
	
	var value string
	var member Contact

	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &member.AddMobile, &member.AddEmails)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nInfo: No data found with the %s contact name to edit\n", name)
			return nil
		} else {
			fmt.Printf("\nError: Failed to query %s in EditContact: %s\n", name, err)
			return err
		}
	}

	fmt.Printf("\n\n**************** Edit the existing contact id: %d ****************\n", member.ID)
	fmt.Printf("\nNote: Press enter to skip the field\n")

	fmt.Printf("\nExisting first name: %s, Enter the new one: ", member.FirstName)
	if value = ValidateInput(); value != "" {
		member.FirstName = value
	}
	fmt.Printf("\nExisting middle name: %s, Enter the new one: ", member.MiddleName)
	if value = ValidateInput(); value != "" {
		member.MiddleName = value
	}
	fmt.Printf("\nExisting last name: %s, Enter the new one: ", member.LastName)
	if value = ValidateInput(); value != "" {
		member.LastName = value
	}
	fmt.Printf("\nExisting mobile number: %s, Enter the new one: ", member.Mobile)
	if value = ValidateInput(); value != "" {
		member.Mobile = value
	}
	fmt.Printf("\nExisting email id: %s, Enter the new one: ", member.Email)
	if value = ValidateInput(); value != "" {
		member.Email = value
	}
	fmt.Printf("\nExisting company name: %s, Enter the new one: ", member.Company)
	if value = ValidateInput(); value != "" {
		member.Company = value
	}
	fmt.Printf("\nExisting location name: %s, Enter the new one: ", member.Location)
	if value = ValidateInput(); value != "" {
		member.Location = value
	}
	fmt.Printf("\nExisting extra mobile number: %s, Enter the new one: ", member.AddMobile)
	if value = ValidateInput(); value != "" {
		member.AddMobile = value
	}
	fmt.Printf("\nExisting extra email id: %s, Enter the new one: ", member.AddEmails)
	if value = ValidateInput(); value != "" {
		member.AddEmails = value
	}
	
	member.Edit(db)
	return nil
}

func DeleteContact(db *sql.DB, name string) error {

	var member Contact

	fmt.Printf("\nIn DeleteContact() %s", name)
	err := db.QueryRow("select * from contact where id = ? or first_name = ? or middle_name = ? or last_name = ?;", name, name, name, name).Scan(&member.ID, &member.FirstName, &member.MiddleName, &member.LastName, &member.Mobile, &member.Email, &member.Company, &member.Location, &member.AddMobile, &member.AddEmails)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nInfo: No data found with the %s contact name to delete\n", name)
			return nil
		} else {
			fmt.Printf("\nError: Failed to query %s in DeleteContact: %s\n", name, err)
			return err
		}
	}

	member.Delete(db)
	return nil
}

func GetTop10Contact(db *sql.DB) error {
	
	var id int
	var activity string
	var top10activities []string

	datas, err := db.Query("select * from activity;")
	if err != nil {
		fmt.Printf("\nError: Failed to query in GetTop10Contact: %s\n", err)
		return err
	}
	defer datas.Close()

	for datas.Next() {
		err = datas.Scan(&id, &activity)
		if err != nil {
			fmt.Printf("\nError: Failed to scan in GetTop10Contact: %s\n", err)
		}
		top10activities = append(top10activities, activity)
	}

	err = datas.Err()
	if err != nil {
		fmt.Printf("\nError: Error occured in row scan in GetTop10Contact: %s\n", err)
	}

	for start, end := 0, len(top10activities)-1; start < end; start, end = start+1, end-1 {
		top10activities[start], top10activities[end] = top10activities[end], top10activities[start]
	}

	fmt.Printf("\n\n**************** Top 10 Contact based on Activites ****************\n")
	for count, top10 := range top10activities {
		fmt.Printf("\n%s", top10)
		if count == 9 {
			break
		}
	}

	return nil
}

func GetCallHistory(db *sql.DB) error {

	var call string
	var callHistory []string

	datas, err := db.Query("select call_history from history;")
	if err != nil {
		fmt.Printf("\nError: Failed to query in GetCallHistory: %s\n", err)
		return err
	}
	defer datas.Close()

	for datas.Next() {
		err = datas.Scan(&call)
		if err != nil {
			fmt.Printf("\nError: Failed to scan in GetCallHistory: %s\n", err)
		}
		callHistory = append(callHistory, call)
	}

	err = datas.Err()
	if err != nil {
		fmt.Printf("\nError: Error occured in row scan in GetCallHistory: %s\n", err)
	}

	fmt.Printf("\n\n**************** Call history ****************\n")
	for _, value := range callHistory {
		fmt.Printf("\n%s", value)
	}

	return nil
}

func UpdateActivity(db *sql.DB, content string) error {

	_, err := db.Exec("create table if not exists activity (id int auto_increment primary key, activities varchar(255))")
	if err != nil {
		fmt.Printf("\nError: Failed to create db in UpdateActivity: %s\n", err)
		return err
	}

	_, err = db.Exec("insert into activity (activities) values (?)", content)
	if err != nil {
		fmt.Printf("\nError: Failed to insert value in UpdateActivity: %s\n", err)
		return err
	}

	return nil
}

func main() {
	
	var name, message string

	db, err := sql.Open("mysql", "root:neridio@tcp(127.0.0.1:3306)/contacts")
	if err != nil {
		fmt.Printf("\nError: Failed to open db: %s\n", err)
		return
	}
	defer db.Close()

	rootCommand := &cobra.Command{
		Use:   "Contact Application",
		Short: "\nThis is a command line application, named as contact",
	}

	rootCommand.PersistentFlags().StringVarP(&name, "name","n", "", "Enter first name or middle name or last name")
	rootCommand.PersistentFlags().StringVarP(&message, "message","m", "", "Enter the message")

	subCommand1 := &cobra.Command{
		Use:   "CreateContact",
		Short: "To create a new contact",
		Run: func(cmd *cobra.Command, args []string) {
			CreateContact(db)
		},
	}

	subCommand2 := &cobra.Command{
		Use:   "CallContact",
		Short: "To make a call from the available contacts",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				fmt.Printf("\nInfo: Please provide the name to perform\n")
				return
			}
			CallContact(db, name)
		},
	}

	subCommand3 := &cobra.Command{
		Use:   "SearchContact",
		Short: "To search a contact from the available contacts",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				fmt.Printf("\nInfo: Please provide the name to perform\n")
				return
			}
			SearchContact(db, name)
		},
	}

	subCommand4 := &cobra.Command{
		Use:   "SendMessage",
		Short: "To send a message to the existing contact",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			message, _ := cmd.Flags().GetString("message")
			if name == "" || message == ""  {
				fmt.Printf("\nInfo: Please provide the name and message to perform\n")
				return
			}
			SendMessage(db, name, message)
		},
	}

	subCommand5 := &cobra.Command{
		Use:   "SearchMessage",
		Short: "To search a message from the message sent list",
		Run: func(cmd *cobra.Command, args []string) {
			message, _ := cmd.Flags().GetString("message")
			if message == "" {
				fmt.Printf("\nInfo: Please provide the message to perform\n")
				return
			}
			SearchMessage(db, message)
		},
	}

	subCommand6 := &cobra.Command{
		Use:   "EditContact",
		Short: "To edit the particular contact from the available contacts",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				fmt.Printf("\nInfo: Please provide the contact name to perform\n")
				return
			}
			EditContact(db, name)
		},
	}

	subCommand7 := &cobra.Command{
		Use:   "DeleteContact",
		Short: "To delete the particular contact from the available contacts",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				fmt.Printf("\nInfo: Please provide the contact name to perform\n")
				return
			}
			DeleteContact(db, name)
		},
	}

	subCommand8 := &cobra.Command{
		Use:   "GetTop10Contact",
		Short: "To get the recent top 10 contact based on call or message activity",
		Run: func(cmd *cobra.Command, args []string) {
			GetTop10Contact(db)
		},
	}

	subCommand9 := &cobra.Command{
		Use:   "GetCallHistory",
		Short: "To get the call history",
		Run: func(cmd *cobra.Command, args []string) {
			GetCallHistory(db)
		},
	}
	
	rootCommand.AddCommand(subCommand1)
	rootCommand.AddCommand(subCommand2)
	rootCommand.AddCommand(subCommand3)
	rootCommand.AddCommand(subCommand4)
	rootCommand.AddCommand(subCommand5)
	rootCommand.AddCommand(subCommand6)
	rootCommand.AddCommand(subCommand7)
	rootCommand.AddCommand(subCommand8)
	rootCommand.AddCommand(subCommand9)
	
	err = rootCommand.Execute()
	if err != nil {
		fmt.Printf("\nError: Error occured %s", err)
	}

	return
}
