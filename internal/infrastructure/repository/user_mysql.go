package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

type userRepositoryMySQL struct {
	users *sql.DB
}

var _ interfaces.IUserRepository = (*userRepositoryMySQL)(nil)

func NewUserRepositoryMySQL(db *sql.DB) *userRepositoryMySQL {
	return &userRepositoryMySQL{
		users: db,
	}
}

func (r *userRepositoryMySQL) Add(ctx context.Context, user entity.UserExtend) error {
	_, errInsertUser := r.users.ExecContext(ctx,
		"INSERT INTO users (`username`, `password`, `id`) VALUES (?, ?, ?)",
		user.Username, user.Password, user.ID)
	if errInsertUser != nil {
		return fmt.Errorf("[userRepositoryMySQL.Add]: %w", errInsertUser)
	}

	return nil
}

func (r *userRepositoryMySQL) Get(ctx context.Context, username string) (entity.UserExtend, error) {
	user := entity.UserExtend{}

	row := r.users.QueryRowContext(ctx, "SELECT `username`, `password`, `id` FROM users WHERE `username` = ?", username)
	errScan := row.Scan(&user.Username, &user.Password, &user.ID)
	if errScan == sql.ErrNoRows {
		return entity.UserExtend{}, fmt.Errorf("[userRepositoryMySQL.Get]: %w", interfaces.ErrUserNotExists)
	}
	if errScan != nil {
		return entity.UserExtend{}, fmt.Errorf("[userRepositoryMySQL.Get]: %w", errScan)
	}

	return user, nil
}

func (r *userRepositoryMySQL) Contains(ctx context.Context, username string) (bool, error) {
	user := entity.UserExtend{}

	row := r.users.QueryRowContext(ctx, "SELECT `username`, `password`, `id` FROM users WHERE `username` = ?", username)
	errScan := row.Scan(&user.Username, &user.Password, &user.ID)
	if errors.Is(errScan, sql.ErrNoRows) {
		return false, nil
	}
	if errScan != nil {
		return false, fmt.Errorf("[userRepositoryMySQL.Contains]: %w", errScan)
	}

	return true, nil
}
