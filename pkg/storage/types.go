package storage

import "time"

type DataType int 

const (
	Integer DataType = iota
	String
	Float
	Boolean
	// TODO: Add more types
)

type Database struct {
	Name      string
	Tables    map[string]*Table
	CreatedAt time.Time
	Version   string
}

type Table struct {
	Name      string
	Columns   map[string]*Column
	Rows      []*Row
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Column struct {
	Name     string
	Type     DataType
	NotNull  bool    
	Default  any     
}

type Row struct {
	// Values maps column names to their values
	Values map[string]any
}


