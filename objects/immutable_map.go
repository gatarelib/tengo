package objects

import (
	"fmt"
	"strings"

	"github.com/d5/tengo/compiler/token"
)

// ImmutableMap represents an immutable map object.
type ImmutableMap struct {
	Value map[string]Object
}

// TypeName returns the name of the type.
func (o *ImmutableMap) TypeName() string {
	return "immutable-map"
}

func (o *ImmutableMap) String() string {
	var pairs []string
	for k, v := range o.Value {
		pairs = append(pairs, fmt.Sprintf("%s: %s", k, v.String()))
	}

	return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
}

// BinaryOp returns another object that is the result of
// a given binary operator and a right-hand side object.
func (o *ImmutableMap) BinaryOp(op token.Token, rhs Object) (Object, error) {
	return nil, ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *ImmutableMap) Copy() Object {
	c := make(map[string]Object)
	for k, v := range o.Value {
		c[k] = v.Copy()
	}

	return &ImmutableMap{Value: c}
}

// IsFalsy returns true if the value of the type is falsy.
func (o *ImmutableMap) IsFalsy() bool {
	return len(o.Value) == 0
}

// Get returns the value for the given key.
func (o *ImmutableMap) Get(key string) (Object, bool) {
	val, ok := o.Value[key]

	return val, ok
}

// Equals returns true if the value of the type
// is equal to the value of another object.
func (o *ImmutableMap) Equals(x Object) bool {
	t, ok := x.(*ImmutableMap)
	if !ok {
		return false
	}

	if len(o.Value) != len(t.Value) {
		return false
	}

	for k, v := range o.Value {
		tv := t.Value[k]
		if !v.Equals(tv) {
			return false
		}
	}

	return true
}
