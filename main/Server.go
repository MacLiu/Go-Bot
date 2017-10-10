package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

var auth Authentication;

func messageRecieved(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Message recieved");
	r.ParseForm();
	fmt.Println(r.Form);
	fmt.Println(r.Body);

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprint(w, "Hi there!");
}

func main() {
	auth = getAuthentication()

	http.HandleFunc("/message", messageRecieved) // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

