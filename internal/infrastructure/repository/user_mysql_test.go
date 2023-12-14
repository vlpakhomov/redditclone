package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestAddSucces(t *testing.T) {
	ctx := context.Background()

	db, mock, errMockNew := sqlmock.New()
	if errMockNew != nil {
		t.Fatalf("cant create mock: %s", errMockNew)
	}
	defer db.Close()

	userRepo := NewUserRepositoryMySQL(db)

	username := "username"
	password := "password"
	id := "id"
	mock.
		ExpectExec("INSERT INTO users").
		WithArgs(username, password, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	user := entity.UserExtend{
		User: entity.User{
			Username: username,
			Password: password,
		},
		ID: id,
	}
	errAdd := userRepo.Add(ctx, user)
	require.Nilf(t, errAdd, "[test.AddUser]: userRepo.Add(ctx, user) failed: %v", errAdd)

	errExpectationsMet := mock.ExpectationsWereMet()
	require.Nilf(t, errExpectationsMet, "[test.AddUser]: mock.ExpectationsWereMet() failed: %v", errExpectationsMet)

}

func TestAddFailure(t *testing.T) {
	ctx := context.Background()

	db, mock, errMockNew := sqlmock.New()
	if errMockNew != nil {
		t.Fatalf("cant create mock: %s", errMockNew)
	}
	defer db.Close()

	userRepo := NewUserRepositoryMySQL(db)

	username := "username"
	password := "password"
	id := "id"
	mock.
		ExpectExec("INSERT INTO users").
		WithArgs(username, password, id).
		WillReturnError(errors.New("insert error"))

	user := entity.UserExtend{
		User: entity.User{
			Username: username,
			Password: password,
		},
		ID: id,
	}
	errAdd := userRepo.Add(ctx, user)
	require.NotNilf(t, errAdd, "[test.AddUser]: expected error from userRepo.Add(ctx, user) : %v", errAdd)

	errExpectationsMet := mock.ExpectationsWereMet()
	require.Nilf(t, errExpectationsMet, "[test.AddUser]: mock.ExpectationsWereMet() failed: %v", errExpectationsMet)

}

func TestGetSuccess(t *testing.T) {
	ctx := context.Background()

	db, mock, errMockNew := sqlmock.New()
	if errMockNew != nil {
		t.Fatalf("cant create mock: %s", errMockNew)
	}
	defer db.Close()

	userRepo := NewUserRepositoryMySQL(db)

	rows := sqlmock.NewRows([]string{"username", "password", "id"})
	expect := []entity.UserExtend{
		{User: entity.User{
			Username: "username",
			Password: "password",
		},
			ID: "id"},
	}
	for _, user := range expect {
		rows = rows.AddRow(user.Username, user.Password, user.ID)
	}
	username := "username"
	mock.
		ExpectQuery("SELECT `username`, `password`, `id` FROM users WHERE").
		WithArgs(username).
		WillReturnRows(rows)

	user, errGet := userRepo.Get(ctx, username)
	require.Nilf(t, errGet, "[test.GetByUsername]: userRepo.Get(ctx, row.Username) failed: %v", errGet)

	errExpectationsMet := mock.ExpectationsWereMet()
	require.Nilf(t, errExpectationsMet, "[test.GetByUsername]: mock.ExpectationsWereMet() failed: %v", errExpectationsMet)

	if !reflect.DeepEqual(user, expect[0]) {
		t.Errorf("[test.GetByUsername]: results not match, want %v, have %v", expect[0], user)
		return
	}
}

func TestGetFailure(t *testing.T) {
	ctx := context.Background()

	db, mock, errMockNew := sqlmock.New()
	if errMockNew != nil {
		t.Fatalf("cant create mock: %s", errMockNew)
	}
	defer db.Close()

	userRepo := NewUserRepositoryMySQL(db)

	// ErrUserNotExists
	rowsInvalidUsername := sqlmock.NewRows([]string{"username", "password", "id"})
	invalidUsername := "invalidUsername"
	mock.
		ExpectQuery("SELECT `username`, `password`, `id` FROM users WHERE").
		WithArgs(invalidUsername).
		WillReturnRows(rowsInvalidUsername)

	_, errGetUsername := userRepo.Get(ctx, invalidUsername)
	require.NotNil(t, errGetUsername, "[test.GetByUsernameErrUserNotExists]: expected error from userRepo.Get(ctx, row.Username)")

	errExpectationsMetInvalid := mock.ExpectationsWereMet()
	require.Nil(t, errExpectationsMetInvalid, "[test.GetByUserNameErrUserNotExists]: mock.ExpectationsWereMet() failed: %v", errExpectationsMetInvalid)

	require.Truef(t, errors.Is(errGetUsername, interfaces.ErrUserNotExists), "[test.GetByUsernameErrUserNotExists]: error not match, want %v, have %v", interfaces.ErrUserNotExists, errGetUsername)

	// ErrScan
	rowsScanErr := sqlmock.NewRows([]string{"username", "password"}).AddRow("username", "password")
	username := "username"
	mock.
		ExpectQuery("SELECT `username`, `password`, `id` FROM users WHERE").
		WithArgs(username).
		WillReturnRows(rowsScanErr)
	_, errGetScan := userRepo.Get(ctx, username)
	require.NotNil(t, errGetScan, "[test.GetByUsernameErrScan]: expected error from userRepo.Get(ctx, row.Username)")

	errExpectationsMetScan := mock.ExpectationsWereMet()
	require.Nil(t, errExpectationsMetScan, "[test.GetByUsernameErrScan]: mock.ExpectationsWereMet() failed: %v", errExpectationsMetScan)
}

func TestContainsSuccess(t *testing.T) {
	ctx := context.Background()

	db, mock, errMockNew := sqlmock.New()
	if errMockNew != nil {
		t.Fatalf("cant create mock: %s", errMockNew)
	}
	defer db.Close()

	userRepo := NewUserRepositoryMySQL(db)

	// Exists
	rowsExists := sqlmock.NewRows([]string{"username", "password", "id"})
	expectExists := []entity.UserExtend{
		{User: entity.User{
			Username: "username",
			Password: "password",
		},
			ID: "id"},
	}
	for _, user := range expectExists {
		rowsExists = rowsExists.AddRow(user.Username, user.Password, user.ID)
	}
	username := "username"
	mock.
		ExpectQuery("SELECT `username`, `password`, `id` FROM users WHERE").
		WithArgs(username).
		WillReturnRows(rowsExists)

	exists, errContainsExists := userRepo.Contains(ctx, username)
	require.Nilf(t, errContainsExists, "[test.ContainsExists]: userRepo.Contains(ctx, username) failed: %v", errContainsExists)

	errExpectationsMetExists := mock.ExpectationsWereMet()
	require.Nilf(t, errExpectationsMetExists, "[test.ContainsExists]: mock.ExpectationsWereMet() failed: %v", errExpectationsMetExists)

	require.Truef(t, exists, "[test.ContainsExists]: expected true from userRepo.Contains(ctx, username)")

	// Not exists
	rowsNotExists := sqlmock.NewRows([]string{"username", "password", "id"})
	usernameNotExists := "usernameNotExists"
	mock.
		ExpectQuery("SELECT `username`, `password`, `id` FROM users WHERE").
		WithArgs(usernameNotExists).
		WillReturnRows(rowsNotExists)
	exists, errContainsNotExists := userRepo.Contains(ctx, usernameNotExists)
	require.Nilf(t, errContainsNotExists, "[test.ContainsNotExists]: userRepo.Contains(ctx, username) failed: %v", errContainsNotExists)

	errExpectationsMetNotExists := mock.ExpectationsWereMet()
	require.Nilf(t, errExpectationsMetNotExists, "[test.ContainsNotExists]: mock.ExpectationsWereMet() failed: %v", errExpectationsMetNotExists)

	require.False(t, exists, "[test.ContainsNotExists]: expected true from userRepo.Contains(ctx, username)")
}

func TestContainsFailure(t *testing.T) {
	ctx := context.Background()

	db, mock, errMockNew := sqlmock.New()
	if errMockNew != nil {
		t.Fatalf("cant create mock: %s", errMockNew)
	}
	defer db.Close()

	userRepo := NewUserRepositoryMySQL(db)

	// ErrScan
	rowsScanErr := sqlmock.NewRows([]string{"username", "password"}).AddRow("username", "password")
	username := "username"
	mock.
		ExpectQuery("SELECT `username`, `password`, `id` FROM users WHERE").
		WithArgs(username).
		WillReturnRows(rowsScanErr)
	_, errGetScan := userRepo.Contains(ctx, username)
	require.NotNil(t, errGetScan, "[test.ContainsErrScan]: expected error from userRepo.Contains(ctx, row.Username)")

	errExpectationsMetScan := mock.ExpectationsWereMet()
	require.Nil(t, errExpectationsMetScan, "[test.ContainsErrScan]: mock.ExpectationsWereMet() failed: %v", errExpectationsMetScan)
}
