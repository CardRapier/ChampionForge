// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: tournamentScheduledQueries.sql

package db

import (
	"context"
)

const getScheduledTournament = `-- name: GetScheduledTournament :one
SELECT id, name, entry_fee, start_time, recurrence_pattern, recurrence_start_timestamp, recurrence_end_timestamp, must_renew FROM scheduled_tournaments WHERE id = $1
`

func (q *Queries) GetScheduledTournament(ctx context.Context, id int64) (ScheduledTournament, error) {
	row := q.db.QueryRow(ctx, getScheduledTournament, id)
	var i ScheduledTournament
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.EntryFee,
		&i.StartTime,
		&i.RecurrencePattern,
		&i.RecurrenceStartTimestamp,
		&i.RecurrenceEndTimestamp,
		&i.MustRenew,
	)
	return i, err
}

const getScheduledTournaments = `-- name: GetScheduledTournaments :many
SELECT id, name, entry_fee, start_time, recurrence_pattern, recurrence_start_timestamp, recurrence_end_timestamp, must_renew FROM scheduled_tournaments
`

func (q *Queries) GetScheduledTournaments(ctx context.Context) ([]ScheduledTournament, error) {
	rows, err := q.db.Query(ctx, getScheduledTournaments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ScheduledTournament
	for rows.Next() {
		var i ScheduledTournament
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.EntryFee,
			&i.StartTime,
			&i.RecurrencePattern,
			&i.RecurrenceStartTimestamp,
			&i.RecurrenceEndTimestamp,
			&i.MustRenew,
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
