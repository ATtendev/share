package db

import (
	"context"
	"time"

	"github.com/ATtendev/share/store/db/ent"
	"github.com/ATtendev/share/store/db/ent/schema"
	"github.com/ATtendev/share/store/db/ent/session"
	"github.com/google/uuid"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	T int64   `json:"t"`
}

type Session struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Description *string
	Title       *string
	Position    []Point
	IsFinished  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeleteAt    *time.Time
}

type UpdateSession struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Description *string
	Title       *string
	IsFinished  *bool
}

type FinishSession struct {
	ID          *uuid.UUID
	UserID      *uuid.UUID
	Description *string
	Title       *string
	IsFinished  *bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeleteAt    *time.Time
}

type UpdatePosition struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Position []Point
}

func (s *Store) CreateSession(ctx context.Context, in *Session) (*ent.Session, error) {
	points := []schema.Point{}
	for _, p := range in.Position {
		points = append(points, schema.Point{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			T: p.T,
		})
	}
	return s.ent.Session.Create().
		SetNillableTitle(in.Title).
		SetUserID(in.UserID).
		SetNillableDescription(in.Description).
		SetPosition(points).
		Save(ctx)
}

func (s *Store) UpdateSession(ctx context.Context, in *UpdateSession) error {
	return s.ent.Session.Update().Where(session.IDEQ(in.ID)).
		SetNillableDescription(in.Description).
		SetNillableTitle(in.Title).
		SetNillableIsFinished(in.IsFinished).
		Exec(ctx)
}

func (s *Store) UpdateSessionPosition(ctx context.Context, in *UpdatePosition) error {
	points := []schema.Point{}
	for _, p := range in.Position {
		points = append(points, schema.Point{
			X: p.X,
			Y: p.Y,
			Z: p.Z,
			T: p.T,
		})
	}
	return s.ent.Session.Update().Where(session.IDEQ(in.ID), session.IsFinished(false)).AppendPosition(points).Exec(ctx)
}

func (s *Store) DeleteSession(ctx context.Context, in *FinishSession) error {
	return s.ent.Session.Update().Where(
		session.IDEQ(*in.ID),
		session.UserIDEQ(*in.UserID)).
		SetDeleteAt(time.Now()).
		Exec(ctx)
}
