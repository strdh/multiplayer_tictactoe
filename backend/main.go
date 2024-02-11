package main 

import (
	"fmt"
	"io/util"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

var usernameMap map[string]struct{}

type UsernameRequest struct{
	Username string `json:"username"`
}

func main() {
	usernameMap = map[string]struct{}{}
	r := mux.NewRouter()
	r.HandleFunc("/ttt/username/create", userCreateHandler).Methods("POST", "OPTIONS")

	server := http.Server{
		Addr: ":5000",
		Handler: r,
	}

	fmt.Println("Server is running....")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	request := UsernameRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		return
	}

	_, exist := usernameMap[request.Username]
	if exist {
		// username already exist
	} else {
		// insert username on the usernameMap
	}

	// make response
}