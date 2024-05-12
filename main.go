package main

import (
	"embedded-nextjs-poc/memstore"
	"embedded-nextjs-poc/models"
	//"embedded-nextjs-poc/models"
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	h := UserHandler{}
	h.store = *memstore.NewUserMemStore()
	h.store.Add(models.User{
		Name:  "clay",
		Email: "clay@ih",
	})
	h.store.Add(models.User{
		Name:  "Sheila",
		Email: "sheila@ih",
	})
	h.store.Add(models.User{
		Name:  "Marcelo",
		Email: "marcelo@ih",
	})
	log.Println("Configuring...")
	mux.HandleFunc("GET /", h.Home)
	mux.HandleFunc("GET /api/user/{id}", h.GetUser)
	mux.HandleFunc("GET /api/user", h.GetUsers)
	mux.HandleFunc("POST /api/user", h.AddUser)
	mux.HandleFunc("PUT /api/user/{id}", h.SaveUser)
	mux.HandleFunc("DELETE /api/user/{id}", h.DeleteUser)

	// Run the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server

}

func (h *UserHandler) errorResponse(w http.ResponseWriter, statusCode int, errorString string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encodingError := json.NewEncoder(w).Encode(UserError{
		StatusCode: statusCode,
		Error:      errorString,
	})
	if encodingError != nil {
		http.Error(w, encodingError.Error(), http.StatusInternalServerError)
	}
}

type UserError struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

func (h *UserHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.List()
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user, err := h.store.Get(id)
	if err != nil {
		h.errorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)

}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err.Error())
	}
	newUser := h.store.Add(user)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newUser)
}

func (h *UserHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err.Error())
	}
	user, err = h.store.Update(id, user)
	if err != nil {
		h.errorResponse(w, http.StatusBadRequest, err.Error())
	}
	err = json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.store.Remove(id)
	if err != nil {
		h.errorResponse(w, http.StatusNotFound, err.Error())
	}
}

type UserHandler struct {
	store memstore.UserMemStore
}
