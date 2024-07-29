package users

import (
	"context"

	"github.com/imperatorofdwelling/Auth/internal/models/entity"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	InsertUser(ctx context.Context, user *entity.User, hashedPassword string) error

	GetUserByID(ctx context.Context, id entity.UserID) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)

	UserUpdatePhone(ctx context.Context, userID entity.UserID, newPhone string) error
	UserUpdateEmail(ctx context.Context, userID entity.UserID, newEmail string) error
	UserUpdateNickname(ctx context.Context, userID entity.UserID, newNickname string) error

	DeleteUserByID(ctx context.Context, id entity.UserID) error
}

type database struct {
	conn sqlx.Conn
	// TODO Redis?
}

func NewDatabase(conn sqlx.Conn) Database {
	return &database{
		conn: conn,
	}
}

// Insert

func (d *database) InsertUser(ctx context.Context, user *entity.User, hashedPassword string) error {
	_, err := d.conn.ExecContext(ctx, insertNewUserStmt,
		user.ID,
		user.Email,
		user.Name,
		user.Surname,
		hashedPassword,
		user.Username,
		user.Phone,
		user.Role,
		user.Country,
		user.Rating,
	)

	return err
}

// Select

func (d *database) GetUserByID(ctx context.Context, id entity.UserID) (*entity.User, error) {
	row := d.conn.QueryRowxContext(ctx, selectByIDStmt, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	model, err := ScanRow(row)
	if err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}
func (d *database) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	row := d.conn.QueryRowxContext(ctx, selectByEmailStmt, email)
	if row.Err() != nil {
		return nil, row.Err()
	}

	model, err := ScanRow(row)
	if err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}
func (d *database) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	row := d.conn.QueryRowxContext(ctx, selectByPhoneStmt, phone)
	if row.Err() != nil {
		return nil, row.Err()
	}

	model, err := ScanRow(row)
	if err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}
func (d *database) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	row := d.conn.QueryRowxContext(ctx, selectByUsernameStmt, username)
	if row.Err() != nil {
		return nil, row.Err()
	}

	model, err := ScanRow(row)
	if err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

// Update

func (d *database) UserUpdatePhone(ctx context.Context, userID entity.UserID, newPhone string) error {
	result, err := d.conn.ExecContext(ctx, updatePhoneStmt, userID, newPhone)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	return err
}
func (d *database) UserUpdateEmail(ctx context.Context, userID entity.UserID, newEmail string) error {
	result, err := d.conn.ExecContext(ctx, updateByEmailStmt, userID, newEmail)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	return err
}
func (d *database) UserUpdateNickname(ctx context.Context, userID entity.UserID, newNickname string) error {
	result, err := d.conn.ExecContext(ctx, updateByNicknameStmt, userID, newNickname)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	return err
}

// Delete

func (d *database) DeleteUserByID(ctx context.Context, id entity.UserID) error {
	result, err := d.conn.ExecContext(ctx, deleteByIDStmt, id)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	return err
}
