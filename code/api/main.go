/* vim: foldmethod=marker foldmarker={{{,}}} foldlevel=2
*/
package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"strings"
	"net/http"
	"database/sql"
	"path/filepath"
	"github.com/hoisie/mustache"
	utils "github.com/ondrax/sympinator-be/code/utils"
)

// {{{ ENDPOINT /
func rootHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//- SERVE APP HERE
		http.ServeFile(w,r,"static/login.html");
	})
	// }}}
}
// }}}


// {{{ testing out serving the React app
// TODO: routing for react app and go app
// SUGGESTION: allow React to use all / but /static, /public, /img, /login, POST, GET, DELETE endpoints and so on
// ALTERNATELY: serve app from /app
func appHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving React app.")
		http.ServeFile(w,r,"frontend/index.html");
	})
}
// }}}
// {{{ testing out Mustache templates 
// func httpMustachedDir() http.HandlerFunc {
// func (w http.ResponseWriter, r *http.Request) {
//     p, err := loadPage(title)
//     if err != nil {
//         p = &Page{Title: title}
//     }
//     t, _ := mustache.RenderFiles("edit.html")
//     t.Execute(w, p)
// }
// }
func mustacheHandler(dirname string) http.HandlerFunc {
	// desired logic:
	// variables served from database's websites table -- $PAGE/... where $PAGE is filename without extension,
	// 	plus GLOBALS
	// 	plus array of POSTS in order from YOUNGEST to OLDEST
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var fnamesToDataBundles = map[string]map[string]string {}
	// TODO: get the following from database in future {{{
	fnamesToDataBundles["default"] = map[string]string{}
	fnamesToDataBundles["lipsum"] = map[string]string{}
	fnamesToDataBundles["helloworld"] = map[string]string{}
	fnamesToDataBundles["default"]["teststr"] = "default page"
	fnamesToDataBundles["lipsum"]["teststr"] = "ipsum"
	fnamesToDataBundles["helloworld"]["teststr"] = "world"
	// }}}
    fname := r.URL.Path[len(dirname):]
	name := fname[0:len(fname)-len(filepath.Ext(fname))]
	var usedMap = map[string]string{}
	if val, ok := fnamesToDataBundles[name] ; ok {
		usedMap = val
	} else {
		usedMap = fnamesToDataBundles["default"]
	}
		data := mustache.RenderFile("static/"+fname, usedMap)
		fmt.Fprintf(w,data)
	})}
// }}}
// {{{ ENDPOINT /make_user
func makeMessage(success bool, messtr string){
	log.Println(messtr)
	log.Println("WRITING MESSAGE")
	// TODO: decide on handling errors -- these just need to be displayed on the user creation page
}
func userMakeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			http.ServeFile(w,r,"static/make_user.html");
		case "POST":
			r.ParseForm()
			// validation
			if ( r.Form["user-pass"][0] != r.Form["user-pass-confirm"][0] ){ // this would ideally be done asynchronously during password entry on front-end, but this is easier
				makeMessage(false,"passwords don't match")
			}
			//e-mail: contains @ followed by at least one dot
			email := r.Form["user-email"][0]
			if (strings.Contains(email,"@")){
				splitEmail := strings.SplitAfter(email,"@")
				if(strings.Contains(splitEmail[1],".") && len(splitEmail[0]) > 1) {// todo: use some well-known regex to test for valid domain name perhaps
					makeMessage(false,"e-mail address invalid")
				}
			}

			hashedPass,err := saltedHash([]byte(r.Form["user-pass"][0]))
			log.Println(hashedPass)

			if ( err != nil){ // this would ideally be done asynchronously during password entry on front-end, but this is easier
				makeMessage(false,"error handling password: "+err.Error())
			}

			userData := UserData{
				r.Form["user-name"][0],
				r.Form["user-real-name"][0],
				r.Form["user-email"][0],
				time.Now(),
				r.Form["user-role"][0],
				hashedPass,
			}
			addNewUser(userData)
		default:
			log.Println("Trying to access /make_user by other than specified means.")
			fmt.Fprintf(w,"Trying to access /make_user by other than specified means.")
		}
	})
}
// }}}
// {{{ ENDPOINT /login
type LoginResponse struct {
	Token string
	Success bool
}
func loginHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			r.ParseForm()
			// check if user exists
			response, err := utils.QuerySQLConn(os.Getenv("DB_NAME"),"SELECT `pass_hash` FROM `users` where `name`=?",r.Form["user-name"][0])
			var hash string
			response.Next()
			err = response.Scan(&hash)

			if (err == sql.ErrNoRows) {
				log.Println("wrong user or password")
			}
			if (err != nil) {
				log.Println("error retrieving user: ",err.Error()) //TODO: write these properly
					response := LoginResponse{"",false}
				utils.WriteJSONResponse(w, response,418)
				return
			}

			err = checkPassWithHash(r.Form["user-pass"][0], hash)

			if (err != nil) {
					log.Println(err.Error())
				log.Println("wrong user or password")
					response := LoginResponse{"",false}
				utils.WriteJSONResponse(w, response,418)
				return
				// failure, let user know
				} else{
				// serve success, token
				// make token containing user name which lasts for 16 hours
				tokenString, err := makeTokenForUser(r.Form["user-name"][0],16)

				if (err != nil) {
					log.Println("ERR",err.Error())
				}
				log.Println(tokenString,"TOKEN")
				log.Println(tokenString,"? VALID")
					response := LoginResponse{tokenString,true}
				utils.WriteJSONResponse(w, response, 200)
				return
				}
			// http.ServeFile(w,r,"static/login.html");
		default:
			log.Println("Trying to access /login by other than specified means.")
			fmt.Fprintf(w,"Trying to access /login by other than specified means.")
		}
	})
}
// }}}

// {{{ FUTURE ENDPOINT /api/users
// requires authentication,
func handleUsers() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// get list of users with their roles
			// requires READ_USERLIST permissions
			// should return JSON {
			//       	users: [ ...
			//       		"name":"...",
			//       		"groups":"...",
			//       	... ]
			//       }
		case "POST":
			// this should do nothing -- userlist is a singleton -- error out
		case "DELETE":
			// requires ALTER_USERLIST permissions
			// requires higher status of current user than deleted user
			// remove a user
		case "PATCH":
			// requires ALTER_USERLIST permissions
			// expects user_id, object with subset of user's values
			// returns altered full user object for checking
			// IMPORTANT! you can't promote a user to higher status than your own
		case "PUT":
			// requires ALTER_USERLIST permissions
			// expects  object full set of user's values
			// returns created full user object for checking
			// IMPORTANT! you can't create a user with higher status than your own
		default:
			// error out
		}
	})
}
// }}}
//{{{ FUTURE ENDPOINT /api/hierarchy
 func handleHierarchy() http.HandlerFunc {
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 switch r.Method {
		 case "GET":
			 // get array of hierarchical positions in order of status
			 // optional parameter "lower_than"|"greater_than" -- only last one specified will be evaluated
		 case "POST":
			 // noop
		 case "DELETE":
			 // REQUIRES greater status than deleted
			 // delete singular hierarchy entry -- users get defaulted to NOBODY status?
		 case "PATCH":
			 // REQUIRES greater status
			 // change singular hierarchy entry's status
		 case "PUT":
			 // REQUIRES greater status
			 // create singular hierarchy entry's status
		 default:
		 }

	 })
 }
//  }}}

// {{{ FUTURE ENDPOINT /api/meta_schedule
 func scheduleHandler() http.HandlerFunc {
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 switch r.Method {
			 // ALWAYS expects: "is_editing": "org[anisation]|ev[en]t"
			 // 	this parameter tells us whether we
		 case "GET":
			 // gets days and timeslots along with associated ids, also gets phases of project -- when what can be submitted
		 case "POST":
			 // noop
		 case "DELETE":
			 // delete a day
		 case "PATCH":
			 // alter a day
		 case "PUT":
			 // create a day
		 default:
		 }

	 })
 }
//  }}}
// {{{ FUTURE ENDPOINT /api/paradigm
 func paradigmHandler() http.HandlerFunc {
	// working with days, classes, 
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 switch r.Method {
			 // ALWAYS expects: "is_editing": "org[anisation]|ev[en]t"
			 // 	this parameter tells us whether we
		 case "GET":
			 // gets days and timeslots along with associated ids, also gets phases of project -- when what can be submitted
		 case "POST":
			 // noop
		 case "DELETE":
			 // delete a day and all references to lectures within it
			 // expects id
			 // requires ALTER_PARADIGM
		 case "PUT":
			 // create a day or replace existing day if request contains day id
			 // optional parameter "day_id"
		 default:
			 // noop, throw error
		 }

	 })
 }
//  }}}

//{{{ FUTURE ENDPOINT /api/lectures
// requires authentication
 func handleLectures() http.HandlerFunc {
		 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			 switch r.Method {
			 case "GET":
			// get simple list of  lectures
			// should return JSON {
			//       	lectures: {...
			//       		"uid":{"name": "...",
			//       			"lectureName": "...",
			//       			"bio": "...",
			//       			"permalink": "...",
			//       			"canIEditThis": true|false},
			//       			...},
			//       		}
			 case "POST":
			 case "PUT":
				 // expects uid
				 // either: 
				 // requiring CREATE_LECTURES privilege
				 // creates a new lecture if no id specified (also
				 // or:
				 // if id specified
				 // requiring ALTER_LECTURES privilege if specified lecture isn't yours
				 // changes lecture info
			 case "DELETE":
				 // expects uid
				 // requires ALTER_LECTURES privilege
				 // removes lecture
			 default:
			 }

	 })
 }
//  }}}
// {{{ FUTURE ENDPOINT /api/timetable
// requires authentication,
func handleTimetable() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// read timetable
			// requires READ_TIMETABLE permissions
			// expects "day" entry with day id from datbase
			// should return JSON {
			//       	lectures: {
			//       		...
			//       		"unique_id":{"name": "...",
			//       			"lectureName": "...",
			//       			"bio": "...",
			//       			"permalink": "..." },
			//       		...
			//       			},
			//       	"rooms": [... room labels ...],
			//       	"hours": [... timeslot names ...],
			//       	"lectureTable": // NxM array; N number of rooms, M number of timeslots
			//       	[
			//       	[... unique_ids ...],
			//       	[... unique_ids ...],
			//       	[... unique_ids ...],
			//       	]
			//       }
		case "POST":
			// this should do nothing -- error out
		case "DELETE":
			// requires ALTER_TIMETABLE permissions
			// Since we're working on a per-record basis (batch editing is a NTH that I might implement in a very
			// optimistic version of the future), this should clean up a single cell
		case "PATCH":
			// requires ALTER_TIMETABLE permissions
		case "PUT":
			// requires ALTER_TIMETABLE permissions
		default:
			// error out
		}
	})
}
// }}}

//{{{ FUTURE ENDPOINT /api/messages
 func handleMessages() http.HandlerFunc {
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 // send messages to people, specify person, header, urgency
		 // uses POST to write, GET to read
	 })
 }
//  }}}
//{{{ FUTURE ENDPOINT /api/web
 func handleWebEdits() http.HandlerFunc {
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 // write fields into website database, maybe implement arrays
		 // GET 
	 })
 }
//  }}}
//{{{ FUTURE ENDPOINT DIRECTORY /web/
 func handleWeb() http.HandlerFunc {
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 // only GET -- serve templates provided by user, providing a map[]{} with values for given page + globals
	 })
 }
//  }}}

//{{{ FUTURE ENDPOINT /public/timetable
 func handlePublicTable() http.HandlerFunc {
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 switch r.Method {
		 case "GET":
			 // returns JSON {
			 // [
			 // {"day":"...",
			 // 	"timeslots":[...{"label":...,"from":"...","to":...,"id":...}...],
			 // 	"rooms":[...{"label":...,"id":...}...],
			 // 	"lectureRows":[...
			 // 		[... {LectureObject[N] or null} ...]
			 // 	...],
			 // 	ALTERNATELY "lectureRowObjects":[...
			 // 	{"timeslot":..., "roomN":{LectureObject[N] or null},
			 // 	...] ordered by timeslot
			 // }
			 // }
			 // ]
			 // }
		 default:
			 // noop
		 }

	 })
 }
//  }}}

// {{{ Writeup of database model requirements
// GET USER PERMISSIONS: from groups where user id = user id, concatenate all permission IDs
// GET USER HIERARCHY STATUS
// GET [LECTURE|SUGGESTION] CREATOR -- so that you can edit your own suggestions
// => ADD creator field (ref to user) to model
// MAKE UNIQUE IDs FOR LECTURES
// }}}
func main() {
	// ROUTING {{{
	// TODO: maybe this all schould be accessible under /api/endpoint and root should serve the app (or login if not logged in)
	// uncomment these as soon as we've got a meaningful basis to assign tokens on
	// http.HandleFunc("/", JWTAuthMiddleware(rootHandler()))
	// http.HandleFunc("/login", JWTAuthMiddleware(loginHandler()))
	// http.HandleFunc("/make_user", JWTAuthMiddleware(userMakeHandler()))
	http.HandleFunc("/", rootHandler())
	http.HandleFunc("/login", loginHandler()) // TODO: handle corresponding html using template
	http.HandleFunc("/api/login", loginHandler())
	http.HandleFunc("/api/make_user", userMakeHandler())
	// http.HandleFunc("/app/*", appHandler()) // all urls except those in dist get routed to index.html
	// http.HandleFunc("/app/dist/*", appDistHandler())
	// http.HandleFunc("/web/*", mustacheTestHandler()) // all urls except those in dist get routed to index.html
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/mustache/", mustacheHandler("/mustache/"))
	http.HandleFunc("/app/",appHandler());
	http.Handle("/app/dist/",http.StripPrefix("/app/dist/", http.FileServer(http.Dir("./frontend/dist"))));
	// In practice, the app would run on a subdomain (hopefully HTACCESS or DNS can do this somehow), / would be served from /web/, /api would be served from ???
	// http.HandleFunc("/app/dist/*", appDistHandler())
	// ROUTING }}}
	// TODO: get default 404 handling, not ugly error messages
	log.Println("listening on 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
