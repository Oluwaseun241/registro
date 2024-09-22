package handler


import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "registro/kafka"
)

type EventRequest struct {
    Topic   string `json:"topic"`
    Message string `json:"message"`
}

func ProduceEvent(w http.ResponseWriter, r *http.Request) {
    var event EventRequest
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &event)

    err := kafka.ProduceEvent(event.Topic, []byte(event.Message))
    if err != nil {
        http.Error(w, "Failed to produce event", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Event produced successfully")
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "List of events...")
}

