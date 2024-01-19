package db

import (
	"context"
	"time"

	"github.com/ATtendev/share/internal/log"
	"github.com/ATtendev/share/store/db/ent/token"
	"github.com/ATtendev/share/store/db/ent/user"
	"github.com/google/uuid"
)

type Token struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
	UserID    uuid.UUID
	Token     string
}

type FindToken struct {
	ID        *uuid.UUID
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeleteAt  *time.Time
	UserID    *uuid.UUID
	Token     *string
}

func (s *Store) IsTokenAvailable(ctx context.Context, in *FindToken) bool {
	exist, err := s.ent.Token.Query().Where(token.TokenEQ(*in.Token), token.DeleteAtIsNil()).
		QueryUsers().Where(user.IDEQ(*in.UserID), user.DeleteAtIsNil()).
		Exist(ctx)
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return exist
}

func (s *Store) CreateToken(ctx context.Context, in *Token) error {
	return s.ent.Token.Create().SetToken(in.Token).SetUserID(in.UserID).Exec(ctx)
}

func (s *Store) DeleteToken(ctx context.Context, in FindToken) error {
	return s.ent.Token.Update().Where(
		token.UserIDEQ(*in.UserID),
		token.TokenEQ(*in.Token)).
		SetDeleteAt(time.Now()).
		Exec(ctx)
}
