package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int    `json:"id"`
}

func userRouter(w http.ResponseWriter, r *http.Request) {
	ourUser := User{}
	ourUser.Name = "Densil Johnson"
	ourUser.Email = "Densil.Johnson@kbc.com"
	ourUser.ID = 48849
	output, _ := json.Marshal(&ourUser)
	fmt.Fprintln(w, string(output))
}

func main() {
	fmt.Println("Starting JSON server")
	http.HandleFunc("/user", userRouter)
	http.ListenAndServe(":8080", nil)
}
