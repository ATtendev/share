// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// FieldIsFinished holds the string denoting the is_finished field in the database.
	FieldIsFinished = "is_finished"
	// FieldIsShared holds the string denoting the is_shared field in the database.
	FieldIsShared = "is_shared"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the session in the database.
	Table = "sessions"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "sessions"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleteAt,
	FieldUserID,
	FieldDescription,
	FieldTitle,
	FieldPosition,
	FieldIsFinished,
	FieldIsShared,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultIsFinished holds the default value on creation for the "is_finished" field.
	DefaultIsFinished bool
	// DefaultIsShared holds the default value on creation for the "is_shared" field.
	DefaultIsShared bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Session queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeleteAt orders the results by the delete_at field.
func ByDeleteAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteAt, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByIsFinished orders the results by the is_finished field.
func ByIsFinished(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFinished, opts...).ToFunc()
}

// ByIsShared orders the results by the is_shared field.
func ByIsShared(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsShared, opts...).ToFunc()
}

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UsersTable, UsersColumn),
	)
}
