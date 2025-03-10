package storage

type DataType int 

const (
	Integer DataType = iota
	String
	Float
	Boolean
	// TODO: Add more types
)

type Database struct {
	Name string
	Tables map[string]*Table
}

type Table struct {
	Name string
	Columns map[string]*Column
	Rows []map[string]interface{}
}

type Column struct {
	Name string
	Type DataType
}


