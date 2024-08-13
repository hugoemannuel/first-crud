package main

import (
	"first-crud/routes"
	"fmt"
	"net/http"
)


func main(){
	fmt.Println("server started at http://localhost:3000")	
	server:= http.Server{
		Addr: ":3000",
		Handler:routes.User(),
	}
	server.ListenAndServe()
}


