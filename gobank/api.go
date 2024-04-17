package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc(`/account`, makeHttpHandleFunc(s.handleAccount))
	router.HandleFunc(`/account/{id}`, makeHttpHandleFunc(s.handleAccountById))
	router.HandleFunc("/transfer", makeHttpHandleFunc(s.handleTransfer))

	log.Println("Starting server on", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccountById(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAccountById(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("invalid method %s", r.Method)

}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	log.Println("Request received ")

	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusAccepted, accounts)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {

	id, err := getId(r)

	if err != nil {
		return err
	}

	account, err := s.store.GetAccountByID(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	log.Println("Create Account Request received")

	createAccountRequest := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
		return err
	}

	account := NewAccount(createAccountRequest.FirstName, createAccountRequest.LastName)

	id, err := s.store.CreateAccount(account)

	if err != nil {
		log.Fatalf("Error creating account: %v", err)
		return err
	}

	account.ID = id

	return WriteJSON(w, http.StatusAccepted, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	id, err := getId(r)

	if err != nil {
		return err
	}

	_, err = s.store.DeleteAccount(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusAccepted, map[string]int{"deleted": id})
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferRequest := new(TransferRequest)

	if err := json.NewDecoder(r.Body).Decode(transferRequest); err != nil {
		return err
	}

	defer r.Body.Close()

	return WriteJSON(w, http.StatusAccepted, transferRequest)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func getId(r *http.Request) (int, error) {

	idString := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idString)

	if err != nil {
		return id, fmt.Errorf("invalid id provided %s", idString)
	}

	return id, nil

}
