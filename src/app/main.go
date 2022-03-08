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
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	restart, err := restartServer("pzserver")

	if err != nil {
		resp, _ := json.Marshal(err)
		w.Write(resp)
	} else {
		resp, json_err := json.Marshal(restart)

		if json_err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", json_err)
		}
		w.Write(resp)
	}

}
