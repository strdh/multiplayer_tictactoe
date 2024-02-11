package main 

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

var usernameMap map[string]struct{}

type UsernameRequest struct{
	Username string `json:"username"`
}

type Response struct {
	Data interface{} `json:"data"`
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

func writeResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	response := Response{
		Data: data,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(jsonData)
}

func isValidUsername(username string) bool {
	pattern := "^[A-Za-z0-9_]+$"
	regexpPattern := regexp.MustCompile(pattern)
	return regexpPattern.MatchString(username)
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeResponse(w, r, http.StatusMethodNotAllowed, nil)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeResponse(w, r, http.StatusInternalServerError, nil)
		return
	}
	defer r.Body.Close()

	request := UsernameRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		writeResponse(w, r, http.StatusInternalServerError, nil)
		return
	}

	isValid := isValidUsername(request.Username)
	if !isValid {
		writeResponse(w, r, http.StatusBadRequest, "Invalid username: username should be A-Z a-z 0-9 and _")
		return
	}

	_, exist := usernameMap[request.Username]
	if exist {
		writeResponse(w, r, http.StatusBadRequest, "Username already exist")
		return
	}

	usernameMap[request.Username] = struct{}{}
	writeResponse(w, r, http.StatusOK, nil)
}