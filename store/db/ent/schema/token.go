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

type Token struct {
	ent.Schema
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).
			Comment("User's UUID").
			Annotations(entsql.WithComments(true)),
		field.String("token").
			Comment("Token string").
			Annotations(entsql.WithComments(true)),
	}
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.TimeMixin{},
	}
}

func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Unique().Required().
			Field("user_id"),
	}
}

func (Token) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}
