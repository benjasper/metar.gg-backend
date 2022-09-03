// Code generated by ent, DO NOT EDIT.

package frequency

const (
	// Label holds the string label denoting the frequency type in the database.
	Label = "frequency"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the frequency in the database.
	Table = "frequencies"
)

// Columns holds all SQL columns for frequency fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
