// Code generated by entc, DO NOT EDIT.

package codegen

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/alexgtn/go-linkshort/tools/ent/codegen/link"
	"github.com/alexgtn/go-linkshort/tools/ent/codegen/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeLink = "Link"
)

// LinkMutation represents an operation that mutates the Link nodes in the graph.
type LinkMutation struct {
	config
	op                Op
	typ               string
	id                *int
	short_path        *string
	long_uri          *string
	accessed_times    *int
	addaccessed_times *int
	created_at        *time.Time
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*Link, error)
	predicates        []predicate.Link
}

var _ ent.Mutation = (*LinkMutation)(nil)

// linkOption allows management of the mutation configuration using functional options.
type linkOption func(*LinkMutation)

// newLinkMutation creates new mutation for the Link entity.
func newLinkMutation(c config, op Op, opts ...linkOption) *LinkMutation {
	m := &LinkMutation{
		config:        c,
		op:            op,
		typ:           TypeLink,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLinkID sets the ID field of the mutation.
func withLinkID(id int) linkOption {
	return func(m *LinkMutation) {
		var (
			err   error
			once  sync.Once
			value *Link
		)
		m.oldValue = func(ctx context.Context) (*Link, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Link.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLink sets the old Link of the mutation.
func withLink(node *Link) linkOption {
	return func(m *LinkMutation) {
		m.oldValue = func(context.Context) (*Link, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LinkMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LinkMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("codegen: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LinkMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *LinkMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Link.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetShortPath sets the "short_path" field.
func (m *LinkMutation) SetShortPath(s string) {
	m.short_path = &s
}

// ShortPath returns the value of the "short_path" field in the mutation.
func (m *LinkMutation) ShortPath() (r string, exists bool) {
	v := m.short_path
	if v == nil {
		return
	}
	return *v, true
}

// OldShortPath returns the old "short_path" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldShortPath(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldShortPath is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldShortPath requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldShortPath: %w", err)
	}
	return oldValue.ShortPath, nil
}

// ClearShortPath clears the value of the "short_path" field.
func (m *LinkMutation) ClearShortPath() {
	m.short_path = nil
	m.clearedFields[link.FieldShortPath] = struct{}{}
}

// ShortPathCleared returns if the "short_path" field was cleared in this mutation.
func (m *LinkMutation) ShortPathCleared() bool {
	_, ok := m.clearedFields[link.FieldShortPath]
	return ok
}

// ResetShortPath resets all changes to the "short_path" field.
func (m *LinkMutation) ResetShortPath() {
	m.short_path = nil
	delete(m.clearedFields, link.FieldShortPath)
}

// SetLongURI sets the "long_uri" field.
func (m *LinkMutation) SetLongURI(s string) {
	m.long_uri = &s
}

// LongURI returns the value of the "long_uri" field in the mutation.
func (m *LinkMutation) LongURI() (r string, exists bool) {
	v := m.long_uri
	if v == nil {
		return
	}
	return *v, true
}

// OldLongURI returns the old "long_uri" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldLongURI(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLongURI is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLongURI requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLongURI: %w", err)
	}
	return oldValue.LongURI, nil
}

// ResetLongURI resets all changes to the "long_uri" field.
func (m *LinkMutation) ResetLongURI() {
	m.long_uri = nil
}

// SetAccessedTimes sets the "accessed_times" field.
func (m *LinkMutation) SetAccessedTimes(i int) {
	m.accessed_times = &i
	m.addaccessed_times = nil
}

// AccessedTimes returns the value of the "accessed_times" field in the mutation.
func (m *LinkMutation) AccessedTimes() (r int, exists bool) {
	v := m.accessed_times
	if v == nil {
		return
	}
	return *v, true
}

// OldAccessedTimes returns the old "accessed_times" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldAccessedTimes(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccessedTimes is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccessedTimes requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccessedTimes: %w", err)
	}
	return oldValue.AccessedTimes, nil
}

// AddAccessedTimes adds i to the "accessed_times" field.
func (m *LinkMutation) AddAccessedTimes(i int) {
	if m.addaccessed_times != nil {
		*m.addaccessed_times += i
	} else {
		m.addaccessed_times = &i
	}
}

// AddedAccessedTimes returns the value that was added to the "accessed_times" field in this mutation.
func (m *LinkMutation) AddedAccessedTimes() (r int, exists bool) {
	v := m.addaccessed_times
	if v == nil {
		return
	}
	return *v, true
}

// ResetAccessedTimes resets all changes to the "accessed_times" field.
func (m *LinkMutation) ResetAccessedTimes() {
	m.accessed_times = nil
	m.addaccessed_times = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *LinkMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *LinkMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *LinkMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the LinkMutation builder.
func (m *LinkMutation) Where(ps ...predicate.Link) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *LinkMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Link).
func (m *LinkMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LinkMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.short_path != nil {
		fields = append(fields, link.FieldShortPath)
	}
	if m.long_uri != nil {
		fields = append(fields, link.FieldLongURI)
	}
	if m.accessed_times != nil {
		fields = append(fields, link.FieldAccessedTimes)
	}
	if m.created_at != nil {
		fields = append(fields, link.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LinkMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case link.FieldShortPath:
		return m.ShortPath()
	case link.FieldLongURI:
		return m.LongURI()
	case link.FieldAccessedTimes:
		return m.AccessedTimes()
	case link.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LinkMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case link.FieldShortPath:
		return m.OldShortPath(ctx)
	case link.FieldLongURI:
		return m.OldLongURI(ctx)
	case link.FieldAccessedTimes:
		return m.OldAccessedTimes(ctx)
	case link.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Link field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LinkMutation) SetField(name string, value ent.Value) error {
	switch name {
	case link.FieldShortPath:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetShortPath(v)
		return nil
	case link.FieldLongURI:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLongURI(v)
		return nil
	case link.FieldAccessedTimes:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccessedTimes(v)
		return nil
	case link.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Link field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LinkMutation) AddedFields() []string {
	var fields []string
	if m.addaccessed_times != nil {
		fields = append(fields, link.FieldAccessedTimes)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LinkMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case link.FieldAccessedTimes:
		return m.AddedAccessedTimes()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LinkMutation) AddField(name string, value ent.Value) error {
	switch name {
	case link.FieldAccessedTimes:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAccessedTimes(v)
		return nil
	}
	return fmt.Errorf("unknown Link numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LinkMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(link.FieldShortPath) {
		fields = append(fields, link.FieldShortPath)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LinkMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LinkMutation) ClearField(name string) error {
	switch name {
	case link.FieldShortPath:
		m.ClearShortPath()
		return nil
	}
	return fmt.Errorf("unknown Link nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LinkMutation) ResetField(name string) error {
	switch name {
	case link.FieldShortPath:
		m.ResetShortPath()
		return nil
	case link.FieldLongURI:
		m.ResetLongURI()
		return nil
	case link.FieldAccessedTimes:
		m.ResetAccessedTimes()
		return nil
	case link.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Link field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LinkMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LinkMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LinkMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LinkMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LinkMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LinkMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LinkMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Link unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LinkMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Link edge %s", name)
}
