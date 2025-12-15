package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/surya/user-age-api/db/sqlc"
)

type UserRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		q: sqlc.New(db),
	}
}

func toPgDate(t time.Time) pgtype.Date {
	return pgtype.Date{
		Time:  t,
		Valid: true,
	}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (sqlc.User, error) {
	return r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  toPgDate(dob),
	})
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int32) (sqlc.User, error) {
	return r.q.GetUserByID(ctx, id)
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]sqlc.User, error) {
	return r.q.ListUsers(ctx)
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	id int32,
	name string,
	dob time.Time,
) (sqlc.User, error) {
	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  toPgDate(dob),
	})
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.q.DeleteUser(ctx, id)
}
