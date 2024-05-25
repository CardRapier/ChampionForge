// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: tournamentQueries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getTournament = `-- name: GetTournament :one
SELECT id, name, entry_fee, start_time FROM tournaments WHERE id = $1
`

func (q *Queries) GetTournament(ctx context.Context, id int64) (Tournament, error) {
	row := q.db.QueryRow(ctx, getTournament, id)
	var i Tournament
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.EntryFee,
		&i.StartTime,
	)
	return i, err
}

const getTournaments = `-- name: GetTournaments :many
SELECT id, name, entry_fee, start_time FROM tournaments WHERE start_time >= $1 and start_time <= $2
`

type GetTournamentsParams struct {
	StartTime   pgtype.Timestamp
	StartTime_2 pgtype.Timestamp
}

func (q *Queries) GetTournaments(ctx context.Context, arg GetTournamentsParams) ([]Tournament, error) {
	rows, err := q.db.Query(ctx, getTournaments, arg.StartTime, arg.StartTime_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tournament
	for rows.Next() {
		var i Tournament
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.EntryFee,
			&i.StartTime,
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
