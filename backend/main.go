package main 

import (
	"fmt"
	"time"
	"regexp"
	"context"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Room struct{
	RoomId string
	Username1 string
	Username2 string
	Username1Score int
	Username2Score int
	Round int
	Status bool
}

type UsernameRequest struct{
	Username string `json:"username"`
}

type Response struct {
	Data interface{} `json:"data"`
}

var usernameMap map[string]struct{}
var userStatus map[string]struct{}
var roomMap map[string]Room
var privateRoomMap map[string]Room

func main() {
	usernameMap = map[string]struct{}{}
	r := mux.NewRouter()
	r.Use(corsOptions)

	r.HandleFunc("/ttt/username/create", userCreateHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/ttt/room/join", joinRoomHandler).Methods("POST", "OPTIONS")

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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
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

func corsOptions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			writeResponse(w, r, http.StatusOK, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
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

func randomMatchHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10 * time.Second)
	defer cancel()

	var isMatch bool

	loop:
	for {
		select {
		case <-ctx.Done():
			http.Error(w, "Cannot find partner", http.StatusRequestTimeout)
			return
		default:
			if match {
				isMatch = true
				break loop
			}
		}
	}

	if isMatch {
		ulid := ulid.Make()
		finalId := ulid.String()

		playroom := Room{
			RoomId: finalId,
			Username1 string
			Username2 string
			Username1Score int
			Username2Score int
			Round int
			Status bool
		}
		writeResponse(w, r, http.StatusOK)
	} else {
		writeResponse(w, r, http.StatusBadRequest)
	}
}

func joinRoomHandler(w http.ResponseWriter, r *http.Request) {
	
}