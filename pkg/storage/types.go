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
	Rows []*Row
}

type Column struct {
	Name string
	Type DataType
}

type Value struct {
	Type DataType
	Data interface{}
}

type Row struct {
	Values map[string]Value
}


