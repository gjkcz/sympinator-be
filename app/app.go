package main
// Temporary root app file for testing basic api functionality
import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"time"
	"net/http"
	"os"
	"encoding/json"
	"github.com/gomodule/redigo/redis" // cache api access token
	"strconv" // convert number to string
	"github.com/satori/go.uuid" // convert number to string
)

var cache redis.Conn
const tokenExpireTime = 360

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "application root")
}
func loginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func main() {

	initCache()
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/login",loginHandler);
	router.GET("/",indexHandler);

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initCache(){
	conn,err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	cache = conn
}

// TEMPORARY list of admissible users
var users = map[string]string{
	"ahojky": "hranolky",
	"nazdar": "kvazar",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func ValidatePass(received string,expected string) bool{
	// TEMP: here for easy implementation of at least hypothetically secure password authentication
	return (received == expected)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || !ValidatePass(expectedPassword,creds.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// generate unique session token
	sessionToken, uerr := uuid.NewV4()
	_, err = cache.Do("SETEX", sessionToken.String(),strconv.Itoa( tokenExpireTime ), creds.Username)
	if err != nil || uerr != nil{
		log.Println("Error setting cache or generating uuid.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken.String(),
		Expires: time.Now().Add(tokenExpireTime*time.Second),
	})
}
