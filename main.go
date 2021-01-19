package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Ticket struct {
  ID         string `json:"id,omitempty"`
  User       string `json:"firstname,omitempty"`
  Creation   string `json:"creation,omitempty"`
  Update     string `json:"update,omitempty"`
}



var Tickets []Ticket

// EndPoints
func GetTicketEndpoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  for _, item := range Tickets {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Ticket{})
}

func GetTicketsEndpoint(w http.ResponseWriter, req *http.Request){
  json.NewEncoder(w).Encode(Tickets)
}

func CreateTicketEndpoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  var Ticket Ticket
  _ = json.NewDecoder(req.Body).Decode(&Ticket)
  Ticket.ID = params["id"]
  Tickets = append(Tickets, Ticket)
  json.NewEncoder(w).Encode(Tickets)

}

func UpdateTicketEndpoint(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for index, item := range Tickets {
    if item.ID == params["id"] {
      Tickets = append(Tickets[:index], Tickets[index + 1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(Tickets)
}

func DeleteTicketEndpoint(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for index, item := range Tickets {
    if item.ID == params["id"] {
      Tickets = append(Tickets[:index], Tickets[index + 1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(Tickets)
}

func main() {
  router := mux.NewRouter()
  
  // adding example data
  //Tickets = append(Tickets, Ticket{ID: "1", FirstName:"Ryan", LastName:"Ray", Address: &Address{City:"Dubling", State:"California"}})
  //Tickets = append(Tickets, Ticket{ID: "2", FirstName:"Maria", LastName:"Ray"})

  // endpoints
  router.HandleFunc("/Tickets", GetTicketsEndpoint).Methods("GET")
  router.HandleFunc("/Tickets/{id}", GetTicketEndpoint).Methods("GET")
  router.HandleFunc("/Tickets/{id}", CreateTicketEndpoint).Methods("POST")
  router.HandleFunc("/Tickets/{id}", DeleteTicketEndpoint).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":3000", router))
}