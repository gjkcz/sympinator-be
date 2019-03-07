package main

import (
	"log"
	"os"
	"golang.org/x/crypto/bcrypt"
	"time"
	utils "github.com/ondrax/sympinator-be/code/utils"
)

type UserData struct{
	Name string
	RealName string
	Email string
	CreatedAt time.Time
	Role string // TODO: maybe make role a struct, load this
	PassHash string
}


func saltedHash(password []byte) (string,error) {

	hash,err := bcrypt.GenerateFromPassword(password,bcrypt.DefaultCost)

	if (err != nil) {
		return "",err
	}
	return string(hash),nil
}

func checkPassWithHash(storedHash string, password string) error { // this really is redundant -- used to be implemented differently, I want to keep this part of the "user" category of behaviour -- ideally I'd have a struct to handle all things user-related and this would be its method
	return bcrypt.CompareHashAndPassword([]byte(password),[]byte(storedHash))
}

func addNewUser( data UserData ) error {
// connect to db or use existing connection
// validate user credentials
// TODO
// send sql request
	dbname := os.Getenv("DB_NAME")
	result, err := utils.QuerySQLConn(dbname,"INSERT INTO users ( name, real_name, email, created_at, role, pass_hash) VALUES (?,?,?,?,?,?);",
	data.Name, data.RealName, data.Email, data.CreatedAt.Format("2006-01-02 15:04:05"), data.Role, data.PassHash)

		utils.PrintQueryResult(os.Stderr,result)
// ALSO DUMP USERS AFTER CREATION
log.Println("DUMPING EXISTING USERS FOR CHECKAGE")
log.Println("===================================")
	result, err = utils.QuerySQLConn(dbname,"SELECT * FROM users;")

		utils.PrintQueryResult(os.Stderr,result)

if (err != nil) {
	return err
}
	return nil // returning nil as error means success
}
