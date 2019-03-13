/* vim: foldmethod=marker foldmarker={{{,}}} foldlevel=2
	if( UserData == null){
	}
*/
package main
//
// import (
// 	"log"
// 	_ "os"
// 	_ "golang.org/x/crypto/bcrypt"
// 	"errors"
// 	_ "regexp"
// 	utils "github.com/ondrax/sympinator-be/code/utils"
// )
// // {{{ TIMETABLE META-RELATED DATA MODELS
// type RowHeaderField struct{
// 	ID int
// 	Label string
// 	From int
// 	To int
// }
// type ColumnHeaderField struct{
// 	ID int
// 	Label string
// }
// // }}}
// // {{{ GET LABELS FOR ROW HEADERS (timeslots)
// func GetTimeslotLabels() ([]RowHeaderField ,error) {
// 	rows, err := utils.QuerySQLConn(DBNAME,"SELECT id, label, begins_at, ends_at FROM times")
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	// HERE
// 	list := []RowHeaderField{}
// 	for rows.Next() {
// 		var i RowHeaderField
// 		rows.Scan(&i.ID,&i.Label,&i.From,&i.To);
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		list = append(list,i)
// 	}
//
// 	return list,err
// }
// // }}}
// // {{{ GET LABELS FOR COLUMN HEADERS (classes)
// func GetClassLabels() ([]ColumnHeaderField ,error) {
// 	rows, err := utils.QuerySQLConn(DBNAME,"SELECT id, name FROM rooms")
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	// HERE
// 	list := []ColumnHeaderField{}
// 	for rows.Next() {
// 		var i ColumnHeaderField
// 		rows.Scan(&i.ID,&i.Label);
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		list = append(list,i)
// 	}
//
// 	return list,err
// }
// // }}}
//
// // {{{ LECTURE-RELATED DATA MODELS
// type NewLectureSuggestionData struct{
// 	LectureName string
// 	SpeakerName string
// 	SpeakerBio string
// 	LectureDesc string
// 	FromNonprague bool
// 	Preferences string
// }
//
// type LectureSuggestionData struct{
// 	ID int
// 	LectureName string
// 	SpeakerName string
// 	SpeakerBio string
// 	LectureDesc string
// 	FromNonprague bool
// 	Preferences string
// }
//
// type LectureData struct{
// 	ID int
// 	LectureName string
// 	SpeakerName string
// 	SpeakerBio string
// 	LectureDesc string
// 	FromNonprague bool
// 	Preferences string
// }
// func (this NewLectureSuggestionData) Valid() error {
// 	// is always valid -- should contain at least one field, really, but meh
// 	return nil
// }
//
// // }}}
//
// // {{{  SAVE LECTURE SUGGESTION TO DB
// func (this NewLectureSuggestionData) SaveToDB() (LectureSuggestionData,error) {
//
// 	// PREPARE ANY DATA WHICH NEEDS SERVER-SIDE PROCESSING
// 	// HERE
//
// 	// INSERT THE DATA
// 	result, err := utils.QuerySQLConn(DBNAME,"INSERT INTO lecture_suggestion (id,lecture_name,speaker_name,speaker_bio,lecture_desc,`from_nonprague`,preferences) VALUES (DEFAULT,?,?,?,?,?,?);",this.LectureName,this.SpeakerName,this.SpeakerBio,this.LectureDesc,this.FromNonprague,this.Preferences);
// 	if ( err != nil){
// 		return LectureSuggestionData{},errors.New("INSERT QUERY ERROR"+err.Error())
// 	}
// 	var lastId int
// 	// GET Last inserted item's ID for further manipulation
// 	// MySQL's LAST_INSERT_ID should be usable, but since a new connection is opened every time 
// 	result,err = utils.QuerySQLConn(DBNAME,"SELECT id FROM lecture_suggestion ORDER BY id DESC LIMIT 1;") // get just inserted item's id --there's a mysql function for this but it's no good for some reason
// 	if ( err != nil){
// 		return LectureSuggestionData{},errors.New("GETTING ID BACK ERROR"+err.Error())
// 	}
// 	result.Scan(&lastId)
// 	sugg,err:=GetLectureSuggestion(lastId)
// 	// Work with this ID -- ideally return just inserted db object to make sure
//
//
// 	return sugg,err
// }
// // }}}
// // {{{  MOVE LECTURE SUGGESTION TO LECTURES
// func (this NewLectureSuggestionData) MoveToLectureDBs() error {
//
// 	// PREPARE ANY DATA WHICH NEEDS SERVER-SIDE PROCESSING
// 	// HERE
//
// 	// INSERT THE DATA
// 	result, err := utils.QuerySQLConn(DBNAME,"INSERT INTO lecture_suggestion (id,lecture_name,speaker_name,speaker_bio,lecture_desc,`from_nonprague`,preferences) VALUES (DEFAULT,?,?,?,?,?,?);",this.LectureName,this.SpeakerName,this.SpeakerBio,this.LectureDesc,this.FromNonprague,this.Preferences);
// 	if ( err != nil){
// 		return errors.New("INSERT QUERY ERROR"+err.Error())
// 	}
// 	var lastId int
// 	// GET Last inserted item's ID for further manipulation
// 	// MySQL's LAST_INSERT_ID should be usable, but since a new connection is opened every time 
// 	result,err = utils.QuerySQLConn(DBNAME,"SELECT id FROM lecture_suggestion ORDER BY id DESC LIMIT 1;") // get just inserted item's id
// 	if ( err != nil){
// 		return errors.New("GETTING ID BACK ERROR"+err.Error())
// 	}
// 	result.Scan(&lastId)
// 	// Work with this ID -- ideally return just inserted db object to make sure
//
//
// 	return err
// }
// // }}}
//
// func GetLectureSuggestion(id int) (LectureSuggestionData,error) {
// 	rows, err := utils.QuerySQLConn(DBNAME,"SELECT id,lecture_name,speaker_name,speaker_bio,lecture_desc,`from_nonprague`,preferences FROM lecture_suggestions WHERE id=?",id)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	// HERE
// 	this := LectureSuggestionData{}
// 	rows.Next() //TODO: meaningful error when returning nothing
// 	rows.Scan(&this.ID,&this.LectureName,&this.SpeakerName,&this.SpeakerBio,&this.LectureDesc,&this.FromNonprague,&this.Preferences);
// 	return this,err
// }
//
// // {{{ FULL TIMETABLE GET METHOD
// func GetTimetableRefList(dayId int) ([][]int,error) {
// 	result, err := utils.QuerySQLConn(DBNAME,"SELECT lecture_id,timeslots_id,rooms_id FROM days_x_timeslots_x_rooms_to_lectures WHERE day_id=?",dayId)
//
// 	cols,err := GetClassLabels()
// 	if err != nil{
// 		log.Println(err.Error())
// 	}
// 	rows,err := GetTimeslotLabels()
// 	if err != nil{
// 		log.Println(err.Error())
// 	}
//
// 	table := make([][]int, len(rows))
// 	for i := range table {
// 		table[i] = make([]int, len(cols))
// 	}
// 	// We've made an m x n table to store the lectures
// 	for result.Next() {
// 		var timeslotID int
// 		var roomID int
// 		var lectureID int
// 		err := result.Scan(&lectureID,&timeslotID,&roomID)
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		log.Println("T,R",timeslotID,roomID,len(rows),len(cols))
// 		table[timeslotID-1][roomID-1] = lectureID
// 	}
//
// 	return table,err
// }
// // TODO: alternatively return full object (would have to figure out how to build it)
// // }}}
// func GetLectureSuggestions() ([]LectureSuggestionData,error) {
// 	rows, err := utils.QuerySQLConn(DBNAME,"SELECT id,lecture_name,speaker_name,speaker_bio,lecture_desc,`from_nonprague`,preferences FROM lecture_suggestions")
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	// HERE
// 	list := []LectureSuggestionData{}
// 	for rows.Next() {
// 		var i LectureSuggestionData
// 		rows.Scan(&i.ID,&i.LectureName,&i.SpeakerName,&i.SpeakerBio,&i.LectureDesc,&i.FromNonprague,&i.Preferences);
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		list = append(list,i)
// 	}
// 	return list, err
// }
//
//
// func GetLecture(id int) (LectureData,error) {
// 	rows, err := utils.QuerySQLConn(DBNAME,"SELECT id,lecture_name,speaker_name,speaker_bio,lecture_desc,`from_nonprague`,preferences FROM lectures WHERE id=?",id)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	// HERE
// 	this := LectureData{}
// 	rows.Next() //TODO: meaningful error when returning nothing
// 	rows.Scan(&this.ID,&this.LectureName,&this.SpeakerName,&this.SpeakerBio,&this.LectureDesc,&this.FromNonprague,&this.Preferences);
// 	return this,err
// }
//
// func GetLecturesAsIndexedObject() (map[int]LectureData,error) {
// 	rows, err := utils.QuerySQLConn(DBNAME,"SELECT id,lecture_name,speaker_name,speaker_bio,lecture_desc,`from_nonprague`,preferences FROM lectures")
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	// HERE
// 	list := map[int]LectureData{}
// 	for rows.Next() {
// 		var i LectureData
// 		rows.Scan(&i.ID,&i.LectureName,&i.SpeakerName,&i.SpeakerBio,&i.LectureDesc,&i.FromNonprague,&i.Preferences);
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		list[i.ID] = i
// 	}
// 	return list, err
// }
// func GetLectures() ([]LectureData,error) {
// 	rows, err := utils.QuerySQLConn(DBNAME,"SELECT id,lecture_name,speaker_name,speaker_bio,lecture_desc,`from_nonprague`,preferences FROM lectures")
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	// HERE
// 	list := []LectureData{}
// 	for rows.Next() {
// 		var i LectureData
// 		rows.Scan(&i.ID,&i.LectureName,&i.SpeakerName,&i.SpeakerBio,&i.LectureDesc,&i.FromNonprague,&i.Preferences);
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		list = append(list,i)
// 	}
// 	return list, err
// }
// // func GetFullTimetable() ([]UserData,error) {
// // 	rows, err := utils.QuerySQLConn(DBNAME,"SELECT id, name FROM users ORDER BY id")
// // 	if err != nil {
// // 		log.Println(err.Error())
// // 	}
// // 	list := []<++>Data{}
// // 	for rows.Next() {
// // 		var i <++>Data
// // 		err := rows.Scan()
// // 		if err != nil {
// // 			log.Println(err.Error())
// // 		}
// // 		list = append(list,i)
// // 	}
// //
// // 	return list,err
// //
// // }
// //
// //
// // func checkPassWithHash(password string, storedHash string) error { // this really is redundant -- used to be implemented differently, I want to keep this as part of the "user" category of behaviour -- ideally I'd have a struct to handle all things user-related and this would be its method
// // 	return bcrypt.CompareHashAndPassword([]byte(password),[]byte(storedHash))
// // }
// //
// // // func addNewUser( data UserData ) error {
// // // // connect to db or use existing connection
// // // // validate user credentials
// // // // TODO
// // // // send sql request
// // // 	dbname := os.Getenv("DB_NAME")
// // // 	result, err := utils.QuerySQLConn(dbname,"INSERT INTO users ( name, real_name, email, created_at, pass_hash) VALUES (?,?,?,?,?);",
// // // 	data.Name, data.RealName, data.Email, data.CreatedAt.Format("2006-01-02 15:04:05"), data.PassHash)
// // //
// // // 		utils.PrintQueryResult(os.Stderr,result)
// // // // ALSO DUMP USERS AFTER CREATION
// // // log.Println("DUMPING EXISTING USERS FOR CHECKAGE")
// // // log.Println("===================================")
// // // 	result, err = utils.QuerySQLConn(dbname,"SELECT * FROM users;")
// // //
// // // 		utils.PrintQueryResult(os.Stderr,result)
// // //
// // // if (err != nil) {
// // // 	return err
// // // }
// // // 	return nil // returning nil as error means success
// // // }
