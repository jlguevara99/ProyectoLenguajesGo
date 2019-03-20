package main

import (
	"html/template"
  "log"
  "net/http"
  "fmt"
)

type PageVariables struct {
	Email   string
	Subject string
	Message string
}


func main(){
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request){

	
	//p := [5]int{0,0,0,0,0}
	//name:= r.PostFormValue("Name")

	HomePageVars := PageVariables{ //store the date and time in a struct
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	fmt.Println(HomePageVars.Email)

	t, err := template.ParseFiles("index.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
