// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// EdgeTeams holds the string denoting the teams edge name in mutations.
	EdgeTeams = "teams"
	// EdgeAccessTokens holds the string denoting the access_tokens edge name in mutations.
	EdgeAccessTokens = "access_tokens"
	// EdgeCreatedAPIKeys holds the string denoting the created_api_keys edge name in mutations.
	EdgeCreatedAPIKeys = "created_api_keys"
	// EdgeUsersTeams holds the string denoting the users_teams edge name in mutations.
	EdgeUsersTeams = "users_teams"
	// AccessTokenFieldID holds the string denoting the ID field of the AccessToken.
	AccessTokenFieldID = "access_token"
	// TeamAPIKeyFieldID holds the string denoting the ID field of the TeamAPIKey.
	TeamAPIKeyFieldID = "api_key"
	// Table holds the table name of the user in the database.
	Table = "users"
	// TeamsTable is the table that holds the teams relation/edge. The primary key declared below.
	TeamsTable = "users_teams"
	// TeamsInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamsInverseTable = "teams"
	// AccessTokensTable is the table that holds the access_tokens relation/edge.
	AccessTokensTable = "access_tokens"
	// AccessTokensInverseTable is the table name for the AccessToken entity.
	// It exists in this package in order to avoid circular dependency with the "accesstoken" package.
	AccessTokensInverseTable = "access_tokens"
	// AccessTokensColumn is the table column denoting the access_tokens relation/edge.
	AccessTokensColumn = "user_id"
	// CreatedAPIKeysTable is the table that holds the created_api_keys relation/edge.
	CreatedAPIKeysTable = "team_api_keys"
	// CreatedAPIKeysInverseTable is the table name for the TeamAPIKey entity.
	// It exists in this package in order to avoid circular dependency with the "teamapikey" package.
	CreatedAPIKeysInverseTable = "team_api_keys"
	// CreatedAPIKeysColumn is the table column denoting the created_api_keys relation/edge.
	CreatedAPIKeysColumn = "created_by"
	// UsersTeamsTable is the table that holds the users_teams relation/edge.
	UsersTeamsTable = "users_teams"
	// UsersTeamsInverseTable is the table name for the UsersTeams entity.
	// It exists in this package in order to avoid circular dependency with the "usersteams" package.
	UsersTeamsInverseTable = "users_teams"
	// UsersTeamsColumn is the table column denoting the users_teams relation/edge.
	UsersTeamsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
}

var (
	// TeamsPrimaryKey and TeamsColumn2 are the table columns denoting the
	// primary key for the teams relation (M2M).
	TeamsPrimaryKey = []string{"team_id", "user_id"}
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
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByTeamsCount orders the results by teams count.
func ByTeamsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamsStep(), opts...)
	}
}

// ByTeams orders the results by teams terms.
func ByTeams(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAccessTokensCount orders the results by access_tokens count.
func ByAccessTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAccessTokensStep(), opts...)
	}
}

// ByAccessTokens orders the results by access_tokens terms.
func ByAccessTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccessTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedAPIKeysCount orders the results by created_api_keys count.
func ByCreatedAPIKeysCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedAPIKeysStep(), opts...)
	}
}

// ByCreatedAPIKeys orders the results by created_api_keys terms.
func ByCreatedAPIKeys(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedAPIKeysStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUsersTeamsCount orders the results by users_teams count.
func ByUsersTeamsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersTeamsStep(), opts...)
	}
}

// ByUsersTeams orders the results by users_teams terms.
func ByUsersTeams(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersTeamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTeamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TeamsTable, TeamsPrimaryKey...),
	)
}
func newAccessTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccessTokensInverseTable, AccessTokenFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AccessTokensTable, AccessTokensColumn),
	)
}
func newCreatedAPIKeysStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedAPIKeysInverseTable, TeamAPIKeyFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedAPIKeysTable, CreatedAPIKeysColumn),
	)
}
func newUsersTeamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersTeamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UsersTeamsTable, UsersTeamsColumn),
	)
}
