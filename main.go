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

	
	p := make([]string,5)
	//name:= r.PostFormValue("Name")

	HomePageVars := PageVariables{ //store the date and time in a struct
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
		
	}
	p[0] = HomePageVars.Email
	p[1] = HomePageVars.Subject
	p[2] = HomePageVars.Message

	fmt.Println(p)

	t, err := template.ParseFiles("index.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func pagar(lista []int) string {
	var StringMaster, datos, s= "", "", ""
	var lleno = true
	var precioMaster, precio, count = 0.0, 0.0, 0
	for lleno{
		count = 0
		datos = ""
		precio = 0
		for i := 0; i<5; i++{
			if(lista[i] != 0){
				count++
				lista[i]--
				precio += 800
				if (i == 0){
					datos+= "Lenguajes de Programacion - Ruby ------ 800\n"
				}
				if (i == 1){
					datos+= "Lenguajes de Programacion - Prolog ---- 800\n"
				}
				if (i == 2){
					datos+= "Lenguajes de Programacion - Scala ----- 800\n"					
				}
				if (i == 3){
					datos+= "Lenguajes de Programacion - Golang ---- 800\n"					
				}
				if (i == 4){
					datos+= "Lenguajes de Programacion - Kotlin ---- 800\n"
				}
			}
		}
		
		if (count == 2){
			precio*=0.95
		}
		if (count == 3){
			precio*=0.9
		}
		if (count == 4){
			precio*=0.8
		}
		if (count == 5){
			precio*=0.75
		}
		precioMaster += precio
		s = fmt.Sprintf("%f",precio)
		StringMaster += datos + "                 Subtotal = " + s + " Lps.\n"
		s = ""
		if(lista[0] == 0 && lista[1] == 0 && lista[2] == 0 && lista[3] == 0 && lista[4] == 0){
			lleno = false
		} 
	}
	s = fmt.Sprintf("%f", precioMaster)
	StringMaster += ("Total = " + s + " Lps.\n")
	return StringMaster
}
