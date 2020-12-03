// Code generated by entc, DO NOT EDIT.

package group

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldMin holds the string denoting the min field in the database.
	FieldMin = "min"
	// FieldMax holds the string denoting the max field in the database.
	FieldMax = "max"
	// FieldRange holds the string denoting the range field in the database.
	FieldRange = "range"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldLog holds the string denoting the log field in the database.
	FieldLog = "log"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"

	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"

	// Table holds the table name of the group in the database.
	Table = "groups"
	// UsersTable is the table the holds the users relation/edge. The primary key declared below.
	UsersTable = "group_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNickname,
	FieldCount,
	FieldCode,
	FieldIndex,
	FieldMin,
	FieldMax,
	FieldRange,
	FieldNote,
	FieldLog,
	FieldUsername,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"group_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// CountValidator is a validator for the "count" field. It is called by the builders before save.
	CountValidator func(int) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(int) error
	// IndexValidator is a validator for the "index" field. It is called by the builders before save.
	IndexValidator func(int) error
	// MinValidator is a validator for the "min" field. It is called by the builders before save.
	MinValidator func(int) error
	// MaxValidator is a validator for the "max" field. It is called by the builders before save.
	MaxValidator func(int) error
	// RangeValidator is a validator for the "range" field. It is called by the builders before save.
	RangeValidator func(int) error
	// NoteValidator is a validator for the "note" field. It is called by the builders before save.
	NoteValidator func(string) error
	// LogValidator is a validator for the "log" field. It is called by the builders before save.
	LogValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
)
