package main

import "net/http"

func write(w http.ResponseWriter, msg string){
	w.Write([]byte(msg))
}

func h1(w http.ResponseWriter, request *http.Request){
	write(w, "<h1><center>WELCOME TO GO PROGRAMMING</center></h1>")
}

func h2(w http.ResponseWriter, request *http.Request){
	write(w, "<h1><center>GO CONCEPT LECTURE - NET/HTTP PACKAGE</center></h1>")
}

func h3(w http.ResponseWriter, request *http.Request){
	write(w, "<h1><center>GO NET/HTTP - WITH CUSTOM URL</center></h1>")
}

func main()  {
	http.HandleFunc("/h1", h1)
	http.HandleFunc("/h2", h2)
	http.HandleFunc("/h3", h3)
	http.ListenAndServe(":8080", nil)
}
