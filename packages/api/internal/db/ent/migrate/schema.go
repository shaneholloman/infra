// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccessTokensColumns holds the columns for the "access_tokens" table.
	AccessTokensColumns = []*schema.Column{
		{Name: "access_token", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
	}
	// AccessTokensTable holds the schema information for the "access_tokens" table.
	AccessTokensTable = &schema.Table{
		Name:       "access_tokens",
		Columns:    AccessTokensColumns,
		PrimaryKey: []*schema.Column{AccessTokensColumns[0]},
	}
	// EnvsColumns holds the columns for the "envs" table.
	EnvsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "dockerfile", Type: field.TypeString},
		{Name: "public", Type: field.TypeBool},
		{Name: "build_id", Type: field.TypeUUID},
		{Name: "build_count", Type: field.TypeInt, Default: 1},
		{Name: "team_id", Type: field.TypeUUID},
	}
	// EnvsTable holds the schema information for the "envs" table.
	EnvsTable = &schema.Table{
		Name:       "envs",
		Columns:    EnvsColumns,
		PrimaryKey: []*schema.Column{EnvsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "envs_teams_envs",
				Columns:    []*schema.Column{EnvsColumns[7]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// EnvAliasColumns holds the columns for the "env_alias" table.
	EnvAliasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "alias", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "env_id", Type: field.TypeString},
	}
	// EnvAliasTable holds the schema information for the "env_alias" table.
	EnvAliasTable = &schema.Table{
		Name:       "env_alias",
		Columns:    EnvAliasColumns,
		PrimaryKey: []*schema.Column{EnvAliasColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "env_alias_envs_env_aliases",
				Columns:    []*schema.Column{EnvAliasColumns[2]},
				RefColumns: []*schema.Column{EnvsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "is_default", Type: field.TypeBool},
		{Name: "name", Type: field.TypeString},
		{Name: "is_blocked", Type: field.TypeBool},
		{Name: "tier", Type: field.TypeString, Size: 2147483647},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "teams_tiers_teams",
				Columns:    []*schema.Column{TeamsColumns[5]},
				RefColumns: []*schema.Column{TiersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TeamAPIKeysColumns holds the columns for the "team_api_keys" table.
	TeamAPIKeysColumns = []*schema.Column{
		{Name: "api_key", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "team_id", Type: field.TypeUUID},
	}
	// TeamAPIKeysTable holds the schema information for the "team_api_keys" table.
	TeamAPIKeysTable = &schema.Table{
		Name:       "team_api_keys",
		Columns:    TeamAPIKeysColumns,
		PrimaryKey: []*schema.Column{TeamAPIKeysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "team_api_keys_teams_team_api_keys",
				Columns:    []*schema.Column{TeamAPIKeysColumns[2]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TiersColumns holds the columns for the "tiers" table.
	TiersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "vcpu", Type: field.TypeInt64},
		{Name: "ram_mb", Type: field.TypeInt64},
		{Name: "disk_mb", Type: field.TypeInt64},
	}
	// TiersTable holds the schema information for the "tiers" table.
	TiersTable = &schema.Table{
		Name:       "tiers",
		Columns:    TiersColumns,
		PrimaryKey: []*schema.Column{TiersColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString},
		{Name: "access_token_users", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_access_tokens_users",
				Columns:    []*schema.Column{UsersColumns[2]},
				RefColumns: []*schema.Column{AccessTokensColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersTeamsColumns holds the columns for the "users_teams" table.
	UsersTeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "team_id", Type: field.TypeUUID},
	}
	// UsersTeamsTable holds the schema information for the "users_teams" table.
	UsersTeamsTable = &schema.Table{
		Name:       "users_teams",
		Columns:    UsersTeamsColumns,
		PrimaryKey: []*schema.Column{UsersTeamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_teams_users_users",
				Columns:    []*schema.Column{UsersTeamsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "users_teams_teams_teams",
				Columns:    []*schema.Column{UsersTeamsColumns[2]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usersteams_team_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{UsersTeamsColumns[2], UsersTeamsColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccessTokensTable,
		EnvsTable,
		EnvAliasTable,
		TeamsTable,
		TeamAPIKeysTable,
		TiersTable,
		UsersTable,
		UsersTeamsTable,
	}
)

func init() {
	EnvsTable.ForeignKeys[0].RefTable = TeamsTable
	EnvAliasTable.ForeignKeys[0].RefTable = EnvsTable
	TeamsTable.ForeignKeys[0].RefTable = TiersTable
	TeamAPIKeysTable.ForeignKeys[0].RefTable = TeamsTable
	UsersTable.ForeignKeys[0].RefTable = AccessTokensTable
	UsersTeamsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTeamsTable.ForeignKeys[1].RefTable = TeamsTable
}
