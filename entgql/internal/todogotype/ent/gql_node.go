// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entgql/internal/todogotype/ent/billproduct"
	"entgo.io/contrib/entgql/internal/todogotype/ent/category"
	"entgo.io/contrib/entgql/internal/todogotype/ent/friendship"
	"entgo.io/contrib/entgql/internal/todogotype/ent/group"
	"entgo.io/contrib/entgql/internal/todogotype/ent/pet"
	"entgo.io/contrib/entgql/internal/todogotype/ent/schema/bigintgql"
	"entgo.io/contrib/entgql/internal/todogotype/ent/schema/uintgql"
	"entgo.io/contrib/entgql/internal/todogotype/ent/todo"
	"entgo.io/contrib/entgql/internal/todogotype/ent/user"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

// IsNode implements the Node interface check for GQLGen.
func (n *BillProduct) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Category) IsNode() {}

func (c Category) marshalID() string {
	var buf bytes.Buffer
	c.ID.MarshalGQL(&buf)
	return buf.String()
}

// IsNode implements the Node interface check for GQLGen.
func (n *Friendship) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Group) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Pet) IsNode() {}

func (pe Pet) marshalID() string {
	var buf bytes.Buffer
	pe.ID.MarshalGQL(&buf)
	return buf.String()
}

// IsNode implements the Node interface check for GQLGen.
func (n *Todo) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *User) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, string) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
// The value should be a table name in the database.
func WithFixedNodeType(tableName string) NodeOption {
	return WithNodeType(func(context.Context, string) (string, error) {
		return tableName, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, string) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id string) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id string, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id string) (Noder, error) {
	switch table {
	case billproduct.Table:
		query := c.BillProduct.Query().
			Where(billproduct.ID(id))
		query, err := query.CollectFields(ctx, "BillProduct")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case category.Table:
		var uid bigintgql.BigInt
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Category.Query().
			Where(category.ID(uid))
		query, err := query.CollectFields(ctx, "Category")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case friendship.Table:
		query := c.Friendship.Query().
			Where(friendship.ID(id))
		query, err := query.CollectFields(ctx, "Friendship")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case group.Table:
		query := c.Group.Query().
			Where(group.ID(id))
		query, err := query.CollectFields(ctx, "Group")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case pet.Table:
		var uid uintgql.Uint64
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Pet.Query().
			Where(pet.ID(uid))
		query, err := query.CollectFields(ctx, "Pet")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case todo.Table:
		query := c.Todo.Query().
			Where(todo.ID(id))
		query, err := query.CollectFields(ctx, "Todo")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []string, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]string)
	id2idx := make(map[string][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []string) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[string][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case billproduct.Table:
		query := c.BillProduct.Query().
			Where(billproduct.IDIn(ids...))
		query, err := query.CollectFields(ctx, "BillProduct")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case category.Table:
		uids := make([]bigintgql.BigInt, len(ids))
		for i, id := range ids {
			if err := uids[i].UnmarshalGQL(id); err != nil {
				return nil, err
			}
		}
		query := c.Category.Query().
			Where(category.IDIn(uids...))
		query, err := query.CollectFields(ctx, "Category")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.marshalID()] {
				*noder = node
			}
		}
	case friendship.Table:
		query := c.Friendship.Query().
			Where(friendship.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Friendship")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case group.Table:
		query := c.Group.Query().
			Where(group.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Group")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case pet.Table:
		uids := make([]uintgql.Uint64, len(ids))
		for i, id := range ids {
			if err := uids[i].UnmarshalGQL(id); err != nil {
				return nil, err
			}
		}
		query := c.Pet.Query().
			Where(pet.IDIn(uids...))
		query, err := query.CollectFields(ctx, "Pet")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.marshalID()] {
				*noder = node
			}
		}
	case todo.Table:
		query := c.Todo.Query().
			Where(todo.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Todo")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
