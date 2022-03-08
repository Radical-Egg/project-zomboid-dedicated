package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/restart", RestartHandler)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func RestartHandler(w http.ResponseWriter, r *http.Request) {
	restart, err := restartServer("pzserver")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintln(err)
		w.Write([]byte(msg))
	} else {
		resp, json_err := json.Marshal(restart)

		if json_err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintln(json_err)
			w.Write([]byte(msg))
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}
}
