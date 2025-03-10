package storage

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
	Type string
}


