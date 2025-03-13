package memory

import (
	"fmt"
	"go-rdb/pkg/storage"
)

type MemoryDatabaseManager struct {
    databases map[string]*storage.Database
}

func NewDatabaseManager() *MemoryDatabaseManager {
    return &MemoryDatabaseManager{
        databases: make(map[string]*storage.Database),
    }
}

func (m *MemoryDatabaseManager) CreateDatabase(name string) error {
		if _, exists := m.databases[name]; exists {
			return fmt.Errorf("database '%s' already exists", name)
		}
		m.databases[name] = &storage.Database{
			Name: name,
			Tables: make(map[string]*storage.Table),
		}
    return nil
}

func (m *MemoryDatabaseManager) GetDatabase(name string) (*storage.Database, error) {
	db, exists := m.databases[name]
	if !exists {
		return nil, fmt.Errorf("database '%s' not found", name)
	}
	return db, nil
}

func (m *MemoryDatabaseManager) DropDatabase(name string) error {
    delete(m.databases, name)
    return nil
}

// ListDatabases returns a list of all database names.
// The order of the returned names is not guaranteed.
func (m *MemoryDatabaseManager) ListDatabases() []string {
    names := make([]string, 0, len(m.databases))
    for name := range m.databases {
        names = append(names, name)
    }
    return names
} 