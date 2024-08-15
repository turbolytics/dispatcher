package ms

import (
	"errors"
	"github.com/google/uuid"
	"github.com/turbolytics/dispatcher/internal"
	"sync"
)

type Memory struct {
	mu            sync.Mutex
	organizations map[uuid.UUID]internal.Organization
}

// NewMemory creates and returns a new Memory
func NewMemory() *Memory {
	return &Memory{
		organizations: make(map[uuid.UUID]internal.Organization),
	}
}

// GetOrganization retrieves an organization by its ID
func (m *Memory) GetOrganization(id uuid.UUID) (internal.Organization, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	org, exists := m.organizations[id]
	if !exists {
		return internal.Organization{}, errors.New("organization not found")
	}

	return org, nil
}

// CreateOrganization adds a new organization to the m
func (m *Memory) CreateOrganization(org internal.Organization) (internal.Organization, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.organizations[org.ID]; exists {
		return internal.Organization{}, errors.New("organization already exists")
	}

	org.ID = uuid.New() // Assign a new UUID if not already set
	m.organizations[org.ID] = org

	return org, nil
}

// DeleteOrganization removes an organization from the m by its ID
func (m *Memory) DeleteOrganization(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.organizations[id]; !exists {
		return errors.New("organization not found")
	}

	delete(m.organizations, id)
	return nil
}
