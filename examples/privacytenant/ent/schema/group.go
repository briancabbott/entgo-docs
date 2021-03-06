// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/briancabbott/entgo"
	"github.com/briancabbott/entgo/examples/privacytenant/ent/privacy"
	"github.com/briancabbott/entgo/examples/privacytenant/rule"
	"github.com/briancabbott/entgo/schema/edge"
	"github.com/briancabbott/entgo/schema/field"
)

// User holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Mixin of the User schema.
func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the User.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("Unknown"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("groups"),
	}
}

// Policy defines the privacy policy of the Group.
func (Group) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Limit DenyMismatchedTenants only for
			// Create operations
			privacy.OnMutationOperation(
				rule.DenyMismatchedTenants(),
				ent.OpCreate,
			),
			// Limit the FilterTenantRule only for
			// UpdateOne and DeleteOne operations.
			privacy.OnMutationOperation(
				rule.FilterTenantRule(),
				ent.OpUpdateOne|ent.OpDeleteOne,
			),
		},
	}
}
