package db

import (
	"context"
	"time"

	"github.com/ATtendev/share/store/db/ent"
	"github.com/ATtendev/share/store/db/ent/user"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	UserName     string
	Email        string
	AvatarURL    string
	PasswordHash string
	CreatedAt    time.Time
	UpdateAt     time.Time
}

type UpdateUser struct {
	ID           uuid.UUID
	UserName     *string
	Email        *string
	AvatarURL    *string
	PasswordHash *string
}

type FindUser struct {
	ID           *uuid.UUID
	UserName     *string
	Email        *string
	AvatarURL    *string
	PasswordHash *string
}

func (s *Store) GetUser(ctx context.Context, in *FindUser) (*ent.User, error) {
	return s.ent.User.Query().Where(
		user.UserNameEQ(*in.UserName),
		user.DeleteAtIsNil(),
	).First(ctx)
}

func (s *Store) CreateUser(ctx context.Context, in *User) (*ent.User, error) {
	return s.ent.User.Create().
		SetUserName(in.UserName).
		SetPasswordHash(in.PasswordHash).
		SetEmail(in.Email).
		Save(ctx)
}
