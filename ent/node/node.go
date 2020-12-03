// Code generated by entc, DO NOT EDIT.

package node

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"

	// EdgePrev holds the string denoting the prev edge name in mutations.
	EdgePrev = "prev"
	// EdgeNext holds the string denoting the next edge name in mutations.
	EdgeNext = "next"

	// Table holds the table name of the node in the database.
	Table = "nodes"
	// PrevTable is the table the holds the prev relation/edge.
	PrevTable = "nodes"
	// PrevColumn is the table column denoting the prev relation/edge.
	PrevColumn = "node_next"
	// NextTable is the table the holds the next relation/edge.
	NextTable = "nodes"
	// NextColumn is the table column denoting the next relation/edge.
	NextColumn = "node_next"
)

// Columns holds all SQL columns for node fields.
var Columns = []string{
	FieldID,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Node type.
var ForeignKeys = []string{
	"node_next",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
