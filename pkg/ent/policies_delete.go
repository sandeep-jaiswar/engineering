// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sandeep-jaiswar/engineering/pkg/ent/policies"
	"github.com/sandeep-jaiswar/engineering/pkg/ent/predicate"
)

// PoliciesDelete is the builder for deleting a Policies entity.
type PoliciesDelete struct {
	config
	hooks    []Hook
	mutation *PoliciesMutation
}

// Where appends a list predicates to the PoliciesDelete builder.
func (pd *PoliciesDelete) Where(ps ...predicate.Policies) *PoliciesDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PoliciesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PoliciesDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PoliciesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(policies.Table, sqlgraph.NewFieldSpec(policies.FieldID, field.TypeInt))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// PoliciesDeleteOne is the builder for deleting a single Policies entity.
type PoliciesDeleteOne struct {
	pd *PoliciesDelete
}

// Where appends a list predicates to the PoliciesDelete builder.
func (pdo *PoliciesDeleteOne) Where(ps ...predicate.Policies) *PoliciesDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *PoliciesDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{policies.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PoliciesDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}
