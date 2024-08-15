package http

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/turbolytics/dispatcher/internal"
	stores "github.com/turbolytics/dispatcher/internal/stores"
	"net/http"
)

type Store interface {
	GetOrganization(id uuid.UUID) (internal.Organization, error)
	CreateOrganization(org internal.Organization) (internal.Organization, error)
	DeleteOrganization(id uuid.UUID) error
}

type Server struct {
	store Store
}

// CreateOrganizationHandler handles the creation of a new organization
func (s *Server) CreateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	var org internal.Organization
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdOrg, err := s.store.CreateOrganization(org)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdOrg)
}

// ReadOrganizationHandler retrieves an organization by its ID
func (s *Server) ReadOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	org, err := s.store.GetOrganization(id)
	if err != nil {
		http.Error(w, "Organization not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(org)
}

// DeleteOrganizationHandler deletes an organization by its ID
func (s *Server) DeleteOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	err = s.store.DeleteOrganization(id)
	if err != nil {
		http.Error(w, "Organization not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewServer() (*Server, error) {
	s := stores.NewMemory()
	server := &Server{
		store: s,
	}

	return server, nil
}

func NewRoutes(server *Server) chi.Router {
	r := chi.NewRouter()

	r.Post("/organizations", server.CreateOrganizationHandler)        // Create
	r.Get("/organizations/{id}", server.ReadOrganizationHandler)      // Read
	r.Delete("/organizations/{id}", server.DeleteOrganizationHandler) // Delete

	return r
}
