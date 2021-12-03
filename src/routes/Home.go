package Server

import (
	"encoding/json"
	"net/http"
	"time"
)

type Home struct {
	Title    string    `json:"title"`
	DateTime time.Time `json:"date"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&Home{Title: "First Api Application With Golang", DateTime: time.Now()})
}
