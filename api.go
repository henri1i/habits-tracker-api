package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
    listenAddr string
    repositories Repositories
}

func New(listenAddr string, repos Repositories) *Server {
    return &Server{
        listenAddr: listenAddr, 
        repositories: repos,
    }
}

func (server *Server) Listen() {
    router := mux.NewRouter()
    router.HandleFunc("/habit", handleFuncError(server.habitHandler))
    router.HandleFunc("/habit/{id}", handleFuncError(server.specificHabitHandler))

    log.Println("Server listen on: ", server.listenAddr)

    http.ListenAndServe(server.listenAddr, router)
}

func (server *Server) habitHandler(res http.ResponseWriter, req *http.Request) error {
    switch req.Method {
    case "GET":
        return server.listHabits(res, req)
    case "POST":
        return server.createHabit(res, req)
    }

    return fmt.Errorf("method not allowed %s", req.Method)
}

func (server *Server) specificHabitHandler(res http.ResponseWriter, req *http.Request) error {
    switch req.Method {
    case "GET":
        return server.findHabit(res, req)
    case "PATCH":
        return server.toggleHabitStatus(res, req)
    case "DELETE":
        server.deleteHabit(res, req)
    }

    return fmt.Errorf("method not allowed %s", req.Method)
} 

func (server *Server) listHabits(res http.ResponseWriter, req *http.Request) error {
    habits, err := server.repositories.habits.FetchAll()

    if err != nil {
        return err
    }

    return parseJson(res, 200, habits)
}

func (server *Server) createHabit(res http.ResponseWriter, req *http.Request) error {
    vars := mux.Vars(req)
    
    return server.repositories.habits.Create(
        vars["name"], 
        vars["status"], 
        vars["kind"],
    )
}

func (server *Server) findHabit(res http.ResponseWriter, req *http.Request) error {
    vars := mux.Vars(req)

    fmt.Println(vars)

    habit, err := server.repositories.habits.FetchOne(vars["id"])

    if err != nil {
        return err
    }

    return parseJson(res, 200, habit)
}

func (server *Server) toggleHabitStatus(res http.ResponseWriter, req *http.Request) error {
    vars := mux.Vars(req)

    return server.repositories.habits.Toggle(vars["id"])
}

func (server *Server) deleteHabit(res http.ResponseWriter, req *http.Request) error {
    vars := mux.Vars(req)

    return server.repositories.habits.Delete(vars["id"])
}

type funcHandler func (res http.ResponseWriter, req *http.Request) error;

type ServerError struct  {
    Error string
}

func handleFuncError(handler funcHandler) http.HandlerFunc {
    return func (res http.ResponseWriter, req *http.Request) {
        if err := handler(res, req); err != nil {
            parseJson(res, http.StatusBadRequest, ServerError{ Error: err.Error() })
        }
    }
}

func parseJson(res http.ResponseWriter, status int, value any) error {
    res.Header().Add("Content-Type", "application/json")
    res.WriteHeader(status)
    return json.NewEncoder(res).Encode(value)
}

