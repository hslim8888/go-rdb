package memory

import (
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
    return nil
}

func (m *MemoryDatabaseManager) DropDatabase(name string) error {
    return nil
}

func (m *MemoryDatabaseManager) GetDatabase(name string) (*storage.Database, error) {
    return nil, nil
}

func (m *MemoryDatabaseManager) ListDatabases() []string {
    return nil
} 