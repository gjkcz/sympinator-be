package main

import (
	"log"
	"os"
	"golang.org/x/crypto/bcrypt"
	"time"
	utils "utils"
)

type UserData struct{
	Name string
	RealName string
	Email string
	CreatedAt time.Time
	Role int // TODO: maybe make role a struct, load this
	PassHash string
}


func saltedHash(password []byte) string {

	hash,err := bcrypt.GenerateFromPassword(password,bcrypt.DefaultCost)

	if (err == nil) {
		return string(hash)
	} else{
		panic(err.Error())
	}
}

func checkPassWithHash(password []byte, storedHash string) bool {
	hashBytes := []byte(storedHash)
	err := bcrypt.CompareHashAndPassword(password,hashBytes)
	if (err == nil) {
		return true
	} else{
		log.Println(err.Error())
		return false
	}
}

func addNewUser( data UserData ) error {
// connect to db or use existing connection
// validate user credentials
// TODO
// send sql request
	dbname := os.Getenv("DB_NAME")
	result, err := utils.QuerySQLConn(dbname,"INSERT INTO users ( name, real_name, email, created_at, role_id, pass_hash) VALUES (?,?,?,?,?,?)",
		data.Name, data.RealName, data.Email, data.CreatedAt, data.Role, data.PassHash)

		utils.PrintQueryResult(result)

if (err != nil) {
	return err
}
	return nil // returning nil as error means success
}
