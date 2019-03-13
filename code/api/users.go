package main

import (
	"log"
	"os"
	"golang.org/x/crypto/bcrypt"
	"time"
	"errors"
	"regexp"
	utils "github.com/ondrax/sympinator-be/code/utils"
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
)

type NewUserData struct{
	UserName string
	UserRealName string
	UserEMail string
	UserPass string
	UserRole string
}

type UserData struct{
	ID int
	Name string
	RealName string
	Email string
	CreatedAt time.Time
	passHash string
}
func (this UserData) GetPassHash() string{ // pass hash can't be public lest it appear in the api resposne
	return this.passHash
}
func (this NewUserData) Valid() error {

if len(this.UserPass) > 128 || len(this.UserPass) < 8 {
return errors.New("USER_PASSWORD_WRONG_LENGTH")
}

if len(this.UserName) > 32 || len(this.UserName) < 5 {
return errors.New("USER_NAME_WRONG_LENGTH")
}

// var multipleNames id
results, err:= utils.QuerySQLConn(DBNAME,"SELECT id,name FROM users WHERE name=?",this.UserName) // test for redundant entries -- name needs to be unique
if err != nil {
	return errors.New("ERROR DURING VALIDATION -- could not connect to database: "+err.Error())
}
		rowsExist := results.Next()
		if(rowsExist) {
	return errors.New("NAME_ALREADY_IN_DATABASE")
		}


results, err = utils.QuerySQLConn(DBNAME,"SELECT id FROM users WHERE email=?",this.UserEMail) // test for redundant entries -- name needs to be unique
if err != nil {
	return errors.New("ERROR DURING VALIDATION -- could not connect to database: "+err.Error())
}

		rowsExist = results.Next()
		if(rowsExist) {
	return errors.New("EMAIL_ALREADY_IN_DATABASE")
		}


if len(this.UserName) > 32 || len(this.UserName) < 5 {
return errors.New("USER_NAME_WRONG_LENGTH")
}

	  //{{{ RFC 5322 e-mail regex from gist by jatin69
	  var eMailTest = regexp.MustCompile(`(?m)^(((((((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?([A-Za-z0-9!#-'*+\/=?^_\x60{|}~-])+((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?)|(((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?"((\s? +)?(([!#-[\]-~])|(\\([ -~]|\s))))*(\s? +)?"))?)?(((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?<(((((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?(([A-Za-z0-9!#-'*+\/=?^_\x60{|}~-])+(\.([A-Za-z0-9!#-'*+\/=?^_\x60{|}~-])+)*)((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?)|(((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?"((\s? +)?(([!#-[\]-~])|(\\([ -~]|\s))))*(\s? +)?"))@((((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?(([A-Za-z0-9!#-'*+\/=?^_\x60{|}~-])+(\.([A-Za-z0-9!#-'*+\/=?^_\x60{|}~-])+)*)((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?)|(((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?\[((\s? +)?([!-Z^-~]))*(\s? +)?\]((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?)))>((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?))|(((((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?(([A-Za-z0-9!#-'*+\/=?^_\x60{|}~-])+(\.([A-Za-z0-9!#-'*+\/=?^_\x60{|}~-])+)*)((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?)|(((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?"((\s? +)?(([!#-[\]-~])|(\\([ -~]|\s))))*(\s? +)?"))@((((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?(([A-Za-z0-9!#-'*+\/=?^_\x60{|}~-])+(\.([A-Za-z0-9!#-'*+\/=?^_\x60{|}~-])+)*)((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?)|(((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?\[((\s? +)?([!-Z^-~]))*(\s? +)?\]((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?))))$`)
// }}}
if !eMailTest.MatchString(this.UserEMail) {
return errors.New("E_MAIL_MALFORMED")
}
	  // at least two words regex -- should not be 
	  // susceptible to catastrophic backtracking
var realNameTest = regexp.MustCompile(`^\S.*\s.*\S$`)
if !realNameTest.MatchString(this.UserRealName) {
return errors.New("REAL_NAME_WRONG_WORD_COUNT")
}
	  return nil
}

func (this NewUserData) SaveToDB() error {

			passHash,err := saltedHash([]byte(this.UserPass))
			log.Println(passHash,"PASS HASH")

			if ( err != nil){
				return errors.New("error handling password: "+err.Error())
			}

	result, err := utils.QuerySQLConn(DBNAME,"INSERT INTO users ( name, real_name, email, created_at, pass_hash) VALUES (?,?,?,?,?);",
	this.UserName, this.UserRealName, this.UserEMail, time.Now().Format("2006-01-02 15:04:05"), passHash)
	var lastId int
	result,err = utils.QuerySQLConn(DBNAME,"SELECT id FROM users ORDER BY id DESC LIMIT 1;") // get just inserted item's id
	result.Scan(&lastId)
	result,err = utils.QuerySQLConn(DBNAME,"INSERT INTO collections_by_users (collection_id,user_id) VALUES ?,?;",this.UserRole,lastId) // get just inserted item's id
		utils.PrintQueryResult(os.Stderr,result)
		// RETURN INSERTED ROW TO CHECK SUCCESS, also get id for future use

	// result, err = utils.QuerySQLConn(DBNAME,"INSERT INTO collections_by_users (collection_id,user_id) VALUES ?,?",this.UserRole,newUser.ID)
		// utils.PrintQueryResult(os.Stderr,result)
// ALSO DUMP USERS AFTER CREATION
log.Println("DUMPING EXISTING USERS FOR CHECKAGE")
log.Println("===================================")
result, _ = utils.QuerySQLConn(DBNAME,"SELECT * FROM users;")

		// utils.PrintQueryResult(os.Stderr,result)

	return err
}

func saltedHash(password []byte) (string,error) {

	hash,err := bcrypt.GenerateFromPassword(password,bcrypt.DefaultCost)

	if (err != nil) {
		return "",err
	}
	return string(hash),nil
}

func GetUser(id int) (UserData,error) {
	rows, err := utils.QuerySQLConn(DBNAME,"SELECT id, name, real_name, email, created_at, pass_hash FROM users WHERE id=?",id)
	if err != nil {
		log.Println(err.Error())
	}
	// HERE
	newUser := UserData{}
	ok:= rows.Next()
	rows.Scan(&newUser.ID,&newUser.Name,&newUser.RealName,&newUser.Email,&newUser.CreatedAt,&newUser.passHash)
	log.Println(newUser)

	if !ok {
		return newUser,errors.New("NO-SUCH-USER")
	}

	return newUser,err
}

func GetUsers() ([]UserData,error) {
	rows, err := utils.QuerySQLConn(DBNAME,"SELECT id, name, real_name, email, created_at, pass_hash FROM users ORDER BY id")
	if err != nil {
		log.Println(err.Error())
	}
	list := []UserData{}
	for rows.Next() {
		var i UserData
		err := rows.Scan(&i.ID,&i.Name,&i.RealName,&i.Email,&i.CreatedAt,&i.passHash)
		if err != nil {
			log.Println(err.Error())
		}
		list = append(list,i)
	}

	return list,err

}


func checkPassWithHash(password string, storedHash string) error { // this really is redundant -- used to be implemented differently, I want to keep this as part of the "user" category of behaviour -- ideally I'd have a struct to handle all things user-related and this would be its method
	return bcrypt.CompareHashAndPassword([]byte(password),[]byte(storedHash))
}

// func addNewUser( data UserData ) error {
// // connect to db or use existing connection
// // validate user credentials
// // TODO
// // send sql request
// 	dbname := os.Getenv("DB_NAME")
// 	result, err := utils.QuerySQLConn(dbname,"INSERT INTO users ( name, real_name, email, created_at, pass_hash) VALUES (?,?,?,?,?);",
// 	data.Name, data.RealName, data.Email, data.CreatedAt.Format("2006-01-02 15:04:05"), data.PassHash)
//
// 		utils.PrintQueryResult(os.Stderr,result)
// // ALSO DUMP USERS AFTER CREATION
// log.Println("DUMPING EXISTING USERS FOR CHECKAGE")
// log.Println("===================================")
// 	result, err = utils.QuerySQLConn(dbname,"SELECT * FROM users;")
//
// 		utils.PrintQueryResult(os.Stderr,result)
//
// if (err != nil) {
// 	return err
// }
// 	return nil // returning nil as error means success
// }
