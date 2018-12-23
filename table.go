package table

// Val ... datastructure for storing a Value
type Val interface {
}

// StringVal ... string
type StringVal string

// IntVal ... int
type IntVal int

const (
	_ = iota
	// IntegerCol ... ast
	IntegerCol
	// FloatCol ... asdf
	FloatCol
	// StringCol ... astd
	StringCol
)

// Col ... column specification
type Col struct {
	Type int
	Name string
}

// Table ... datastructure for storing a table.Table
type Table struct {
	Columns []Col
	Values  [][]Val
}
