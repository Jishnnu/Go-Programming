package main

import (
	"fmt"	
	"net/http"
)

func main() {
	var port string
	fmt.Print("\nENTER A PORT NUMBER : ")
	fmt.Scanln(&port)

	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", handler)
	http.ListenAndServe(":" + port, multiplexer) // ACTIVATES THE SERVER
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World")
	w.Write([]byte("<h1><center>GO PROGRAMMING</center></h1>"))
	w.Write([]byte("<p><center>Welcome to Sem VII Go Elective</center></p>"))
}
