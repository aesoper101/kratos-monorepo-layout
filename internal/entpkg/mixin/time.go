package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// CreatedAt adds created at time field.
type CreatedAt struct{ ent.Schema }

// Fields of the create time mixin.
func (CreatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}).
			Immutable(),
	}
}

// create time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*CreatedAt)(nil)

// UpdatedAt adds updated at time field.
type UpdatedAt struct{ ent.Schema }

// Fields of the update time mixin.
func (UpdatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("updated_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}).
			UpdateDefault(func() int64 {
				return time.Now().Unix()
			}),
	}
}

// update time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*UpdatedAt)(nil)

// Time composes create/update time mixin.
type Time struct{ ent.Schema }

// Fields of the time mixin.
func (Time) Fields() []ent.Field {
	return append(
		CreatedAt{}.Fields(),
		UpdatedAt{}.Fields()...,
	)
}

// time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*Time)(nil)
