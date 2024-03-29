// Code generated by ent, DO NOT EDIT.

package country

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the country type in the database.
	Label = "country"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldImportID holds the string denoting the import_id field in the database.
	FieldImportID = "import_id"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldImportFlag holds the string denoting the import_flag field in the database.
	FieldImportFlag = "import_flag"
	// FieldLastUpdated holds the string denoting the last_updated field in the database.
	FieldLastUpdated = "last_updated"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldContinent holds the string denoting the continent field in the database.
	FieldContinent = "continent"
	// FieldWikipediaLink holds the string denoting the wikipedia_link field in the database.
	FieldWikipediaLink = "wikipedia_link"
	// FieldKeywords holds the string denoting the keywords field in the database.
	FieldKeywords = "keywords"
	// EdgeAirports holds the string denoting the airports edge name in mutations.
	EdgeAirports = "airports"
	// Table holds the table name of the country in the database.
	Table = "countries"
	// AirportsTable is the table that holds the airports relation/edge.
	AirportsTable = "airports"
	// AirportsInverseTable is the table name for the Airport entity.
	// It exists in this package in order to avoid circular dependency with the "airport" package.
	AirportsInverseTable = "airports"
	// AirportsColumn is the table column denoting the airports relation/edge.
	AirportsColumn = "country_airports"
)

// Columns holds all SQL columns for country fields.
var Columns = []string{
	FieldID,
	FieldImportID,
	FieldHash,
	FieldImportFlag,
	FieldLastUpdated,
	FieldCode,
	FieldName,
	FieldContinent,
	FieldWikipediaLink,
	FieldKeywords,
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

var (
	// DefaultImportFlag holds the default value on creation for the "import_flag" field.
	DefaultImportFlag bool
	// DefaultLastUpdated holds the default value on creation for the "last_updated" field.
	DefaultLastUpdated func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Continent defines the type for the "continent" enum field.
type Continent string

// Continent values.
const (
	ContinentAfrica       Continent = "AF"
	ContinentAntarctica   Continent = "AN"
	ContinentAsia         Continent = "AS"
	ContinentEurope       Continent = "EU"
	ContinentNorthAmerica Continent = "NA"
	ContinentSouthAmerica Continent = "SA"
	ContinentOceania      Continent = "OC"
)

func (c Continent) String() string {
	return string(c)
}

// ContinentValidator is a validator for the "continent" field enum values. It is called by the builders before save.
func ContinentValidator(c Continent) error {
	switch c {
	case ContinentAfrica, ContinentAntarctica, ContinentAsia, ContinentEurope, ContinentNorthAmerica, ContinentSouthAmerica, ContinentOceania:
		return nil
	default:
		return fmt.Errorf("country: invalid enum value for continent field: %q", c)
	}
}

// OrderOption defines the ordering options for the Country queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByImportID orders the results by the import_id field.
func ByImportID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImportID, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// ByImportFlag orders the results by the import_flag field.
func ByImportFlag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImportFlag, opts...).ToFunc()
}

// ByLastUpdated orders the results by the last_updated field.
func ByLastUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByContinent orders the results by the continent field.
func ByContinent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContinent, opts...).ToFunc()
}

// ByWikipediaLink orders the results by the wikipedia_link field.
func ByWikipediaLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWikipediaLink, opts...).ToFunc()
}

// ByAirportsCount orders the results by airports count.
func ByAirportsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAirportsStep(), opts...)
	}
}

// ByAirports orders the results by airports terms.
func ByAirports(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAirportsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAirportsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AirportsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AirportsTable, AirportsColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Continent) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Continent) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Continent(str)
	if err := ContinentValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Continent", str)
	}
	return nil
}
