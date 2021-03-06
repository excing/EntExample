// Code generated by entc, DO NOT EDIT.

package card

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldExpired holds the string denoting the expired field in the database.
	FieldExpired = "expired"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the card in the database.
	Table = "cards"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "cards"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_card"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNumber,
	FieldExpired,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Card type.
var ForeignKeys = []string{
	"user_card",
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

var (
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator func(string) error
)
