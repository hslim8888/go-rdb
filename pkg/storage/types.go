package storage

// DataType represents the type of a column
type DataType int 

const (
	Integer DataType = iota
	String
	Float
	Boolean
	// TODO: Add more types
)

// Column represents a table column definition
type Column struct {
	Name     string
	Type     DataType
	NotNull  bool    
	Default  any     
}

// Row represents a single row of data
type Row struct {
	// Values maps column names to their values
	Values map[string]any
}


