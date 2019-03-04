package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	utils "utils"
)

		// {{{ ENDPOINT /
		func rootHandler() http.HandlerFunc {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// {{{ DATABASE GET EXAMPLE

				dbname := os.Getenv("DB_NAME")
				hostname := os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")
				// connect to database
				db, err := sql.Open("mysql", "root:root@tcp("+hostname+")/"+dbname)
				// if there is an error opening the connection, handle it
				if err != nil {
					panic(err.Error())
				}
				//// sample query
				rows, err := db.Query("SELECT * FROM `users`")

				if err != nil {
					panic(err.Error())
				}

				utils.PrintQueryResult(rows);

				// serve login html
				http.ServeFile(w,r,"static/login.html");
			})
			// }}}
		}
		// }}}

		// {{{ ENDPOINT /make_user
		func userMakeHandler() http.HandlerFunc {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "GET":
					http.ServeFile(w,r,"static/make_user.html");
				case "POST":
					r.ParseForm()
					for key,value := range r.Form {
						fmt.Fprintf(w,"%s -> %s\n", key, value)
					}
				default:
					log.Println("Trying to access /make_user by other than specified means.")
					fmt.Fprintf(w,"Trying to access /make_user by other than specified means.")
				}
			})
		}
		// }}}

		// {{{ ENDPOINT /login
		func loginHandler() http.HandlerFunc {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "POST":
					http.ServeFile(w,r,"static/login.html");
				default:
					log.Println("Trying to access /login by other than specified means.")
					fmt.Fprintf(w,"Trying to access /login by other than specified means.")
				}
			})
		}
		// }}}

		func main() {
			// ROUTING {{{
			// TODO: maybe this all schould be accessible under /api/endpoint and root should serve the app (or login if not logged in)
			// uncomment these as soon as we've got a meaningful basis to assign tokens on
			// http.HandleFunc("/", JWTAuthMiddleware(rootHandler()))
			// http.HandleFunc("/login", JWTAuthMiddleware(loginHandler()))
			// http.HandleFunc("/make_user", JWTAuthMiddleware(userMakeHandler()))
			http.HandleFunc("/", rootHandler())
			http.HandleFunc("/login", loginHandler())
			http.HandleFunc("/make_user", userMakeHandler())
			// ROUTING }}}
			log.Println("listening on 8000")
			log.Fatal(http.ListenAndServe(":8000", nil))
		}
