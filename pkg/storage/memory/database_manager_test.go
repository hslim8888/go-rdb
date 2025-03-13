package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseManager(t *testing.T) {
    t.Run("CreateDatabase", func(t *testing.T) {
        // Given
        manager := NewDatabaseManager()
        dbName := "testdb"

        // When
        err := manager.CreateDatabase(dbName)

        // Then
        assert.NoError(t, err)
        
        // Verify database exists
        db, err := manager.GetDatabase(dbName)
        assert.NoError(t, err)
        assert.NotNil(t, db)
        assert.Equal(t, dbName, db.Name)
    })

    t.Run("CreateDatabase_AlreadyExists", func(t *testing.T) {
        // Given
        manager := NewDatabaseManager()
        dbName := "testdb"
        _ = manager.CreateDatabase(dbName)

        // When
        err := manager.CreateDatabase(dbName)

        // Then
        assert.Error(t, err)
    })

    t.Run("GetDatabase_NotFound", func(t *testing.T) {
        // Given
        manager := NewDatabaseManager()
        dbName := "nonexistent"

        // When
        db, err := manager.GetDatabase(dbName)

        // Then
        assert.Error(t, err)
        assert.Nil(t, db)
    })

    t.Run("ListDatabases", func(t *testing.T) {
        // Given
        manager := NewDatabaseManager()
        dbNames := []string{"db1", "db2", "db3"}
        for _, name := range dbNames {
            _ = manager.CreateDatabase(name)
        }

        // When
        list := manager.ListDatabases()

        // Then
        assert.Equal(t, len(dbNames), len(list))
        for _, name := range dbNames {
            assert.Contains(t, list, name)
        }
    })

    t.Run("DropDatabase", func(t *testing.T) {
        // Given
        manager := NewDatabaseManager()
        dbName := "testdb"
        _ = manager.CreateDatabase(dbName)

        // When
        err := manager.DropDatabase(dbName)

        // Then
        assert.NoError(t, err)
        
        // Verify database no longer exists
        db, err := manager.GetDatabase(dbName)
        assert.Error(t, err)
        assert.Nil(t, db)
    })

    t.Run("DropDatabase_NotFound", func(t *testing.T) {
        // Given
        manager := NewDatabaseManager()
        dbName := "nonexistent"

        // When
        err := manager.DropDatabase(dbName)

        // Then
        assert.NoError(t, err)  // 없는 DB 삭제 시도시 에러 없음
    })
} 