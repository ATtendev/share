package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/ATtendev/share/store/db/ent/mixins"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").MaxLen(100),
		field.String("email").MaxLen(100),
		field.String("avatar_url").Optional(),
		field.String("password_hash"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.TimeMixin{},
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
