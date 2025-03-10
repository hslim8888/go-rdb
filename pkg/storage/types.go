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

type Row struct {
	// Values maps column names to their values
	Values map[string]any
}


