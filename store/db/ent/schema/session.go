package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/ATtendev/share/store/db/ent/mixins"
	"github.com/google/uuid"
)

type Session struct {
	ent.Schema
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	T int64   `json:"t"`
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).
			Comment("User's UUID").
			Annotations(entsql.WithComments(true)),
		field.String("description").Optional().Default("good trip"),
		field.String("title").Default("trip " + uuid.NewString()),
		field.JSON("position", []Point{}).Optional(),
		field.Bool("is_finished").Default(false),
	}
}

func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.TimeMixin{},
	}
}

func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Unique().Required().
			Field("user_id"),
	}
}

func (Session) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}
