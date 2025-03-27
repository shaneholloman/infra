// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: teams_usersteams_join.sql

package queries

import (
	"context"

	"github.com/google/uuid"
)

const getTeamsWithUsersTeams = `-- name: GetTeamsWithUsersTeams :many
SELECT t.id, t.created_at, t.is_blocked, t.name, t.tier, t.email, t.is_banned, t.blocked_reason, ut.id, ut.user_id, ut.team_id, ut.is_default, ut.added_by
FROM "public"."teams" t
JOIN "public"."users_teams" ut ON ut.team_id = t.id
WHERE ut.user_id = $1
`

type GetTeamsWithUsersTeamsRow struct {
	Team      Team
	UsersTeam UsersTeam
}

func (q *Queries) GetTeamsWithUsersTeams(ctx context.Context, userID uuid.UUID) ([]GetTeamsWithUsersTeamsRow, error) {
	rows, err := q.db.Query(ctx, getTeamsWithUsersTeams, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTeamsWithUsersTeamsRow
	for rows.Next() {
		var i GetTeamsWithUsersTeamsRow
		if err := rows.Scan(
			&i.Team.ID,
			&i.Team.CreatedAt,
			&i.Team.IsBlocked,
			&i.Team.Name,
			&i.Team.Tier,
			&i.Team.Email,
			&i.Team.IsBanned,
			&i.Team.BlockedReason,
			&i.UsersTeam.ID,
			&i.UsersTeam.UserID,
			&i.UsersTeam.TeamID,
			&i.UsersTeam.IsDefault,
			&i.UsersTeam.AddedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
