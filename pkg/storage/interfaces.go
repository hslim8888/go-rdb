package storage

// DatabaseManager handles database level operations
type DatabaseManager interface {
    CreateDatabase(name string) error
    DropDatabase(name string) error
    GetDatabase(name string) (Database, error)  // Database 인터페이스로 변경
    ListDatabases() []string
}

// Database represents a database instance
type Database interface {
    Name() string
    CreateTable(name string, columns []*Column) error
    DropTable(name string) error
    GetTable(name string) (Table, error)  // Table 인터페이스로 변경
    ListTables() []string
}

// Table represents a table instance
type Table interface {
    Name() string
    Columns() []*Column
    Insert(row *Row) error
    Select() ([]*Row, error)
    Update(newValues map[string]any) error
    Delete() error
} 