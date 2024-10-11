package main

import (
	"encoding/json"
	"log"
	"main/structs"
	"net/http"
)

func main() {
	commander := structs.NewCommander()
	http.HandleFunc("/execute", handleCommand(commander))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCommand(cmdr structs.Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req structs.CommandRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		var res structs.CommandResponse
		switch req.Type {
		case "ping":
			result, err := cmdr.Ping(req.Payload)
			res = structs.CommandResponse{Success: err == nil, Data: result, Error: errToString(err)}
		case "sysinfo":
			result, err := cmdr.GetSystemInfo()
			res = structs.CommandResponse{Success: err == nil, Data: result, Error: errToString(err)}
		default:
			res = structs.CommandResponse{Success: false, Error: "Invalid command type"}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func errToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
