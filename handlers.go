package main

import (
	"fmt"
	// "github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type User_Data struct {
	User_id    string `bson:"user_id"`
	User_name  string `bson:"user_name"`
	User_email string `bson:"user_email"`
	Csrf       string `bson:"csrf"`
}

var store = sessions.NewCookieStore([]byte("living-on-the-edge"))

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
func Kill(w http.ResponseWriter, r *http.Request) {

}
func Login(w http.ResponseWriter, r *http.Request) {
	data, err1 := mgo.Dial("devdb.leaf.co.in:27017")
	if err1 != nil {
		panic(err1)
	}

	defer data.Close()
	data.SetMode(mgo.Monotonic, true)

	// vars := mux.Vars(r)
	// var session_id string
	// var err bool
	// session_id = r.PostFormValue("session_id")
	r.ParseForm() // Parses the request body
	session_id := r.FormValue("session_id")
	fmt.Println(session_id)

	// fmt.Println("session_id123", r.Form(session_id))
	session_id = "" + session_id + ""
	user := &User_Data{}
	user_sessions_DB := data.DB("beta").C("user_sessions")
	err2 := user_sessions_DB.FindId(bson.ObjectIdHex("5576bc995e62ac301ae5d096")).One(&user)
	if err2 != nil {
		log.Fatal(err2)
	}

	//GO TO LOGIN PAGE --> TAKE DETAILS -->CREATE SESSION -->STORE SESSSION
	// Set some session values.
	// session.Values["user_id"] = "54f4cf54b2a1fd3744dc6026"
	// session.Values["user_name"] = "Akshay"
	// session.Values["user_email"] = "akshayiitb.jain@gmail.com"
	// // Save it before we write to the response/return from the handler.
	// session.Save(r, w)
	fmt.Println("session2", user)

}
