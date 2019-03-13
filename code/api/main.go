/* vim: foldmethod=marker foldmarker={{{,}}} foldlevel=2
*/
package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"strconv"
	"io/ioutil"
	"net/http"
	"database/sql"
	"encoding/json"
	"path/filepath"
	"github.com/hoisie/mustache"
	utils "github.com/ondrax/sympinator-be/code/utils"
)

const DBNAME string = "internal"

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
		// TEMPORARY mappings for contents of specific pages
		//TODO: get the following from database in future {{{
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
	// {{{ ENDPOINT /api/users
	func usersHandler() http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				path := r.URL.Path
				lastBit := path[strings.LastIndex(path,"/")+1:]
				var err error

				if lastBitInt, interr := strconv.Atoi(lastBit); interr == nil { // integer as last bit -- requesting particular user
					dat,err := GetUser(lastBitInt)
					if err == nil{
						utils.WriteJSONResponse(w, dat, 200)
					} else {
						utils.WriteJSONResponse(w, map[string]interface{}{"Success":false,"message":"No user found with id: "+strconv.Itoa(lastBitInt)}, 200)
					}
				} else {
					dat,err := GetUsers()
					if err == nil{
						utils.WriteJSONResponse(w, dat, 200)
						return
					}
				}
				if err != nil {
					utils.WriteJSONResponse(w, AuthResponse{false,"couldn't get user/s: "+err.Error()}, 418)
					return
				}
				return
			case "POST":

				// POST HANDLING {{{
				// var ld map[string]string
				var user NewUserData
				// body, err := ioutil.ReadAll(r.Body)
				// if err != nil {
				//     log.Printf("Error reading body: %v", err)
				// }
				err := json.NewDecoder(r.Body).Decode(&user)
				if err != nil {
					panic(err)
				}
				err = user.Valid()
				if err != nil{
					response:=AuthResponse{false,err.Error()+" -- invalid user credentials supplied"}
					utils.WriteJSONResponse(w, response, 406)
					return
				}
				err = user.SaveToDB()
				if err != nil{
					panic(err.Error())
				}else{
					utils.WriteJSONResponse(w, user,200)
					return
				}
				// username := ld["user-name"]
				// userrealname := ld["user-real-name"]
				// useremail := ld["user-e-mail"]
				// userpass := ld["user-pass"]
				// userrole := ld["user-role"]
				// user := NewUserData{username,userrealname,useremail,userpass,int(userrole)}

				// THE QUERY
				// response, err := db.Query("SELECT pass_hash FROM users WHERE name=?",username)
				// var hashResult string
				// ok := response.Next()
				// if !ok {
				//     response := LoginResponse{"",false,"Wrong username or password."}
				// 	utils.WriteJSONResponse(w, response,418)
				// 	return
				// }
				// err = response.Scan(&hashResult)
				// hash := hashResult[1:]
				// if (!ok || hash=="") {
				//     response := LoginResponse{"",false,"Wrong username or password."}
				// 	utils.WriteJSONResponse(w, response,418)
				// 	return
				// }
				// }}}

			default:
				log.Println("Trying to access /api/user by other than specified means.")
				http.Error(w,"Trying to access /api/user by other than specified means.",405);
			}
		})
	}
	// }}}
	// {{{ ENDPOINT /login
	type LoginResponse struct {
		Token string
		Name string
		ID string
		Success bool
		Message string
	}
	func loginHandler() http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "POST":

				var ld map[string]string //TODO: make  form-name safe using a struct, but eh
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Printf("Error reading body: %v", err)
					// http.Error(w, "can't read body", http.StatusBadRequest)
				}
				err = json.Unmarshal(body,&ld)
				log.Println(ld)
				if err != nil {
					panic(err)
				}
				username := ld["user-name"]
				userpass := ld["user-pass"]

				// check if user exists
				// response, err := utils.QuerySQLConn("internal","SELECT `pass_hash` FROM `users` where `NAME`=?",
				// ld.Name)
				// {{{ getting data from database BOILERPLATE BEGIN
				hostname := os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")
				if os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" {
					log.Println("could not find host or port env variable")
					hostname = "localhost:3306"
				} else {
					log.Println("HOSTNAME",hostname)
				}
				// TODO: figure out how to securely connect to my own created database (no root/root password)
				db, err := sql.Open("mysql", "root:root@tcp("+hostname+")/"+DBNAME)
				if err != nil {
					log.Println("ERROR OPENING DATABASE CONNECTION")
					panic(err.Error())
				}
				err = db.Ping()
				if err != nil {
					log.Println("ERROR CONNECTING TO DATABASE")
					panic(err.Error())
				}
				// }}}
				// THE QUERY
				response, err := db.Query("SELECT name,id,pass_hash FROM users WHERE name=?",username)
				var hashResult string
				var name string
				var id string
				ok := response.Next()
				if !ok {
					response := LoginResponse{"","","",false,"Wrong username or password."}
					utils.WriteJSONResponse(w, response,418)
					return
				}
				err = response.Scan(&name,&id,&hashResult)
				if (!ok || hashResult=="") {
					response := LoginResponse{"","","",false,"Wrong username or password."}
					utils.WriteJSONResponse(w, response,418)
					return
				}
				err = checkPassWithHash(hashResult, userpass)

				if (err != nil) {
					log.Println("Possibly malformed database entry, check below error")
					log.Println(err)
					log.Println()
					response := LoginResponse{"","","",false,"Wrong username or password."}
					utils.WriteJSONResponse(w, response,418)
					return
					// failure, let user know
				} else{
					// serve success, token
					// make token containing user Name which lasts for 24 hours
					tokenString, err := makeTokenForUser(username,24)

					if (err != nil) {
						log.Println("ERR",err.Error())
					}
					log.Println(tokenString,"TOKEN")
					response := LoginResponse{tokenString,name,id,true,"Login successful."}
					utils.WriteJSONResponse(w, response, 200)
					return
				}
				// http.ServeFile(w,r,"static/login.html");
			default:
				log.Println("Trying to access /login by other than specified means.")
				http.Error(w,"Trying to access /login by other than specified means.",405)
			}
		})
	}
	// }}}


	//{{{ ENDPOINT /api/permissions
	func handleUserPermissions() http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method ==  "GET" {
				// {{{ query get last numerical bit copypasta
				path := r.URL.Path
				lastBit := path[strings.LastIndex(path,"/")+1:]
				if lastBitInt, interr := strconv.Atoi(lastBit); interr == nil { // integer as last bit -- requesting particular user
					query := `SELECT DISTINCT u.name as Person_Name, u.id as Person, p.id as Permission, p.name as Permission_name FROM users u
					INNER JOIN collections_by_users r1 ON r1.users_id = u.id
					INNER JOIN permissions_by_collections r2 ON r1.collections_id = r2.collections_id
					INNER JOIN user_permissions p ON r2.perm_id = p.id
					WHERE u.id=?`;
					// }}}
					result,err := utils.QuerySQLConn(DBNAME,query,lastBitInt)
					if err != nil {
						// handle error with http stuff
					}
					utils.PrintQueryResult(w,result); // TODO: jsonify
				} else {
					result,err := utils.QuerySQLConn(DBNAME, `SELECT * FROM user_permissions`);
					utils.PrintQueryResult(w,result);
					if err != nil {
						// handle error with http stuff
					}
				}
		} else {
			log.Println("Trying to access /api/permissions/user/ by other than specified means.")
			http.Error(w,"Trying to access /api/permissions/user/ by other than specified means.",405)

		}
	})
}
	func handleGroupPermissions() http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				// {{{ query get last numerical bit copypasta
				path := r.URL.Path
				lastBit := path[strings.LastIndex(path,"/")+1:]
				if lastBitInt, interr := strconv.Atoi(lastBit); interr == nil { // integer as last bit -- requesting particular user
					// }}}
					query := `SELECT c.id as Collection, c.collection_name as Collection_Name,
					p.id as Privilege, p.name as Privilege_name FROM collections c
					INNER JOIN permissions_by_collections r ON r.collections_id = c.id
					INNER JOIN user_permissions p ON r.collections_id = p.id
					WHERE c.id = ?`;
					result,err := utils.QuerySQLConn(DBNAME,query,lastBitInt)
					if err != nil {
						// handle error with http stuff
					}
					utils.PrintQueryResult(w,result); // TODO: jsonify
				} else {
					result,err := utils.QuerySQLConn(DBNAME, `SELECT * FROM user_permissions`);
					utils.PrintQueryResult(w,result);
					if err != nil {
						// handle error with http stuff
					}
				}
			case "POST":
			case "DELETE":
			case "PUT":
			default:
			}
		})
	}
	//  }}}


	//{{{ ENDPOINT /api/groups
	func groupsHandler() http.HandlerFunc {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				// {{{ query get last numerical bit copypasta
				path := r.URL.Path
				lastBit := path[strings.LastIndex(path,"/")+1:]
				if lastBitInt, interr := strconv.Atoi(lastBit); interr == nil { // integer as last bit -- requesting particular user
					// }}}
					result,err := utils.QuerySQLConn(DBNAME, `SELECT u.name as Person_Name, u.id as Person, c.id as Collection, c.collection_name as Collection_Name FROM users u INNER JOIN collections_by_users r ON r.users_id = u.id INNER JOIN collections c ON r.collections_id = c.id WHERE u.id = ?`,lastBitInt);
					if err != nil {
						// handle error with http stuff
					}
					utils.PrintQueryResult(w,result); // TODO: jsonify
				} else {
					result,err := utils.QuerySQLConn(DBNAME, `SELECT * FROM collections`);
					utils.PrintQueryResult(w,result);
					if err != nil {
						// handle error with http stuff
					}
				}
			case "POST":
			case "DELETE":
			case "PUT":
			default:
			}

		})
	}
	//  }}}

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
				// expects user_id, object with subset of users values
				// returns altered full user object for checking
				// IMPORTANT! you cant promote a user to higher status than your own
			case "PUT":
				// requires ALTER_USERLIST permissions
				// expects  object full set of users values
				// returns created full user object for checking
				// IMPORTANT! you cant create a user with higher status than your own
			default:
				// error out
			}
		})
	}
	// }}}
	//{{{ FUTURE ENDPOINT /api/hierarchy

	//{{{ ENDPOINT api/hierarchy/mine
	func hierarchyByUser() http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// {{{ query get last numerical bit copypasta
			path := r.URL.Path
			lastBit := path[strings.LastIndex(path,"/")+1:]
			if lastBitInt, interr := strconv.Atoi(lastBit); interr == nil { // integer as last bit -- requesting particular user
				// }}}
				query := `SELECT u.name as Person_Name, u.id as Person, c.id as Collection, c.collection_name as Collection_Name, c.hierarchy_pos as Status FROM users u
				INNER JOIN collections_by_users r ON r.users_id = u.id
				INNER JOIN collections c ON r.collections_id = c.id
				ORDER BY c.hierarchy_pos LIMIT 1
				WHERE u.id=?`
				result,err := utils.QuerySQLConn(DBNAME,query,lastBitInt)
				if err != nil {
					// handle error with http stuff
				}
				utils.PrintQueryResult(w,result); // TODO: jsonify
			} else {
				result,err := utils.QuerySQLConn(DBNAME, `SELECT * FROM user_permissions`);
				utils.PrintQueryResult(w,result);
				if err != nil {
					// handle error with http stuff
				}
			}

		})
	}
	//  }}}

	func hierarchyHandler(sign int) http.HandlerFunc {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				var query string;
				if sign > 0 {
					query = `SELECT u.id as Person, u.name as Person_name FROM users u
					WHERE u.id NOT IN (
						SELECT u.id FROM users u
						INNER JOIN collections_by_users r ON r.users_id = u.id
						INNER JOIN collections c ON r.collections_id = c.id
						WHERE c.hierarchy_pos <= (
							SELECT MIN(c.hierarchy_pos) FROM users u
							INNER JOIN collections_by_users r ON r.users_id = u.id
							INNER JOIN collections c ON r.collections_id = c.id
							WHERE u.id = 4
						)
					);`;
				} else if sign < 0 {
					query = `SELECT u.id as Person, u.name as Person_name FROM users u
					WHERE u.id NOT IN (
						SELECT u.id FROM users u
						INNER JOIN collections_by_users r ON r.users_id = u.id
						INNER JOIN collections c ON r.collections_id = c.id
						WHERE c.hierarchy_pos >= (
							SELECT MIN(c.hierarchy_pos) FROM users u
							INNER JOIN collections_by_users r ON r.users_id = u.id
							INNER JOIN collections c ON r.collections_id = c.id
							WHERE u.id = 4
						)
					);`;
				} else { // just show all in order
					query = `SELECT u.id as Person, u.name as Person_name, MIN(ANY_VALUE(c.hierarchy_pos)) as Status FROM users u
					INNER JOIN collections_by_users r ON r.users_id = u.id
					INNER JOIN collections c ON r.collections_id = c.id
					GROUP BY u.id, u.name
					ORDER BY Status;`;
				}

				fmt.Fprintf(w,query);
				if(sign < 0) {

				}else{
				}
			case "POST":
				// should create role, can only make it lower than status of currently active user
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



//	{{{ ENDPOINT /api/lecture_suggestions
	 func handleSuggestion() http.HandlerFunc {
		 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "POST":
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Printf("Error reading body: %v", err)
					// http.Error(w, "can't read body", http.StatusBadRequest)
				}
				ld := NewLectureSuggestionData{}
				err = json.Unmarshal(body,&ld)
				newSugg,err:= ld.SaveToDB()
				if err != nil {
					log.Printf("Error saving suggestion body:", err)
					// http.Error(w, "can't read body", http.StatusBadRequest)
				}
				utils.WriteJSONResponse(w,map[string]interface{}{"dataSent":ld,"newSugg":newSugg},200);
				case "GET":
					suggestions,err := GetLectureSuggestions()
					if err == nil{
						utils.WriteJSONResponse(w, suggestions, 200)
						return
					}
				case "PUT":
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Printf("Error reading body: %v", err)
					// http.Error(w, "can't read body", http.StatusBadRequest)
				}
				ld := map[string]interface{}{} // TODO: validatificate
				err = json.Unmarshal(body,&ld)
				if val, ok := ld["idToPut"]; ok {
					// idToMove is specified, we can go on  move
					_,err:=utils.QuerySQLConn(DBNAME,`INSERT INTO lectures (lecture_name,speaker_name,speaker_bio,lecture_desc,from_nonprague,preferences) SELECT lecture_name,speaker_name,speaker_bio,lecture_desc,from_nonprague,preferences FROM lecture_suggestion WHERE id = ?;
					DELETE FROM lecture_suggestion WHERE id=?;`,val,val)
					if err != nil {
						utils.WriteJSONResponse(w, map[string]string{"Success":"false","message":"Some error moving lecture, possibly failed unique key lecture name; read here: "+err.Error()}, 501)
				}
				}
				default:
				log.Println("Trying to access /api/lecture_suggestions by other than specified means.")
				http.Error(w,"Trying to access /api/lecture_suggestions by other than specified means.",405);
			 }
		 })
	 }
//	  }}}

	//{{{ FUTURE ENDPOINT /api/lectures
	// requires authentication
	func handleLectures() http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				lectures, err := GetLectures()
				if err != nil {
					log.Println("ERROR handling lectures: "+err.Error())
				}
				returnObject := map[string]interface{}{"lectureList":lectures}
				utils.WriteJSONResponse(w,returnObject,200)
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
				var dayId int = 1
				lectures, err := GetLecturesAsIndexedObject()
				if err != nil {
					log.Println("ERROR handling lectures: "+err.Error())
				}
				columnNames, err := GetClassLabels()
				if err != nil {
					log.Println("ERROR handling column names: "+err.Error())
				}
				rowNames, err := GetTimeslotLabels()
				if err != nil {
					log.Println("ERROR handling row names: "+err.Error())
				}
				fullRefList, err := GetTimetableRefList(1)
				var dayName string
				result, err := utils.QuerySQLConn(DBNAME,"SELECT name FROM days WHERE id=?",dayId)
				if err != nil {
					log.Println("error executing day name query: "+err.Error())
				}
				result.Scan(&dayName)
				thingy := map[string]interface{}{"lectures":lectures,"day":dayName,"rooms":columnNames,"hours":rowNames,"lectureTable":fullRefList}
				response := map[string]map[string]interface{} {"tableData":thingy}
				utils.WriteJSONResponse(w,response,200);
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
	func toApp(w http.ResponseWriter, r *http.Request){
		http.Redirect(w, r, "/app", 301)
	}
	//{{{ STUB ENDPOINT /api/is-authenticated
	func serveSuccess(message string) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if(r.Method == "GET"){
				log.Println("r.Header['header']");
				log.Println(r.Header);
				// response := AuthResponse{true,"token valid"}
				// utils.WriteJSONResponse(w,response,200) // TODO: look up non-arbitrary server error numbers
			} else {
				http.Error(w,"Only GET with header allowed for this stub endpoint.",405);
			}
			// this should do nothing -- userlist is a singleton -- error out
			// uses POST to write, GET to read
		})
	}

	//  }}}
	func main() {
		// ROUTING {{{
		// TODO: maybe this all schould be accessible under /api/endpoint and root should serve the app (or login if not logged in)
		// uncomment these as soon as we've got a meaningful basis to assign tokens on
		// http.HandleFunc("/", JWTAuthMiddleware(rootHandler()))
		// http.HandleFunc("/login", JWTAuthMiddleware(loginHandler()))
		// http.HandleFunc("/make_user", JWTAuthMiddleware(userMakeHandler()))
		http.HandleFunc("/login", loginHandler()) // TODO: handle corresponding html using template
		http.HandleFunc("/api/login", loginHandler())
		http.HandleFunc("/api/timetable/", handleTimetable())
		http.HandleFunc("/api/lectures/", handleLectures())
		http.HandleFunc("/api/users/", usersHandler())
		http.HandleFunc("/api/groups/", groupsHandler())
		http.HandleFunc("/api/hierarchy/higher", hierarchyHandler(1))
		http.HandleFunc("/api/hierarchy/lower", hierarchyHandler(-1))
		http.HandleFunc("/api/hierarchy/mine/", hierarchyByUser())
		http.HandleFunc("/api/permissions/user/", handleUserPermissions())
		http.HandleFunc("/api/permissions/group/", handleGroupPermissions())
		http.HandleFunc("/api/lecture_suggestions", handleSuggestion())
		// // http.HandleFunc("/app/*", appHandler()) // all urls except those in dist get routed to index.html
		// // http.HandleFunc("/app/dist/*", appDistHandler())
		// // http.HandleFunc("/web/*", mustacheTestHandler()) // all urls except those in dist get routed to index.html
		// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
		http.HandleFunc("/mustache/", mustacheHandler("/mustache/"))
		// // http.HandleFunc("/api/is-authenticated",JWTAuthMiddleware(serveSuccess("token valid")));
		// http.HandleFunc("/api/is-authenticated",serveSuccess("token valid"));
		// // http.HandleFunc("/api/hierarchy/lower",JWTAuthMiddleware(serveSuccess("token valid")));
		// // http.HandleFunc("/api/hierarchy/higher",JWTAuthMiddleware(serveSuccess("token valid")));
		// http.HandleFunc("/api/hierarchy/lower",handleHierarchy(1))
		// http.HandleFunc("/api/hierarchy/higher",handleHierarchy(-1))
		// http.HandleFunc("/api/hierarchy",handleHierarchy(0))
		http.HandleFunc("/app/",appHandler());
		http.Handle("/app/dist/",http.StripPrefix("/app/dist/", http.FileServer(http.Dir("./frontend/dist"))));
		// http.HandleFunc("/", toApp) // TEMPORARY
		// In practice, the app would run on a subdomain (hopefully HTACCESS or DNS can do this somehow), / would be served from /web/, /api would be served from ???
		// http.HandleFunc("/app/dist/*", appDistHandler())
		// ROUTING }}}
		v := reflect.ValueOf(http.DefaultServeMux).Elem()
		fmt.Printf("routes: %v\n", v.FieldByName("m"))
		// TODO: get default 404 handling, not ugly error messages
		log.Println("listening on 8000")
		log.Fatal(http.ListenAndServe(":8000", nil))
	}

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
