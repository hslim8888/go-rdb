package storage

// DatabaseManager handles database level operations
type DatabaseManager interface {
    CreateDatabase(name string) error
    DropDatabase(name string) error
    GetDatabase(name string) (*Database, error)
    ListDatabases() []string
}

// TableManager handles table level operations
type TableManager interface {
    CreateTable(name string, columns []*Column) error
    DropTable(name string) error
    GetTable(name string) (*Table, error)
    ListTables() []string
}

// TableOperator handles data operations on a table
type TableOperator interface {
    Insert(row *Row) error
    Select() ([]*Row, error)  // 모든 row 반환
    Update(newValues map[string]any) error  // 모든 row 업데이트
    Delete() error  // 모든 row 삭제
} 