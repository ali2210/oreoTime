package main 

import (
	//"fmt"
	"log"
	"net/http"
	"text/template"
	)
func HomePage(w http.ResponseWriter , r * http.Request) {
	temp,err := template.ParseFiles("index.html")
	if err != nil {
	  log.Print("no page found", err)
	}
 	err = temp.Execute(w, HomePage)
	if err != nil {
	  log.Print("Error 404", err)
	}
}
func main(){
 	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8080", nil)
}
	
