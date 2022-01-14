package ormtable

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/kvstore"
)

// ReadBackend defines the type used for read-only ORM operations.
type ReadBackend interface {
	// CommitmentStoreReader returns the reader for the commitment store.
	CommitmentStoreReader() kvstore.Reader

	// IndexStoreReader returns the reader for the index store.
	IndexStoreReader() kvstore.Reader

	private()
}

// Backend defines the type used for read-write ORM operations.
// Unlike ReadBackend, write access to the underlying kv-store
// is hidden so that this can be fully encapsulated by the ORM.
type Backend interface {
	ReadBackend

	// CommitmentStore returns the merklized commitment store.
	CommitmentStore() kvstore.Store

	// IndexStore returns the index store if a separate one exists,
	// otherwise it the commitment store.
	IndexStore() kvstore.Store

	// Hooks returns a Hooks instance or nil.
	Hooks() Hooks
}

// ReadBackendOptions defines options for creating a ReadBackend.
// Read context can optionally define two stores - a commitment store
// that is backed by a merkle tree and an index store that isn't.
// If the index store is not defined, the commitment store will be
// used for all operations.
type ReadBackendOptions struct {

	// CommitmentStoreReader is a reader for the commitment store.
	CommitmentStoreReader kvstore.Reader

	// IndexStoreReader is an optional reader for the index store.
	// If it is nil the CommitmentStoreReader will be used.
	IndexStoreReader kvstore.Reader
}

type readBackend struct {
	commitmentReader kvstore.Reader
	indexReader      kvstore.Reader
}

func (r readBackend) CommitmentStoreReader() kvstore.Reader {
	return r.commitmentReader
}

func (r readBackend) IndexStoreReader() kvstore.Reader {
	return r.indexReader
}

func (readBackend) private() {}

// NewReadBackend creates a new ReadBackend.
func NewReadBackend(options ReadBackendOptions) ReadBackend {
	indexReader := options.IndexStoreReader
	if indexReader == nil {
		indexReader = options.CommitmentStoreReader
	}
	return &readBackend{
		commitmentReader: options.CommitmentStoreReader,
		indexReader:      indexReader,
	}
}

type backend struct {
	commitmentStore kvstore.Store
	indexStore      kvstore.Store
	hooks           Hooks
}

func (backend) private() {}

func (c backend) CommitmentStoreReader() kvstore.Reader {
	return c.commitmentStore
}

func (c backend) IndexStoreReader() kvstore.Reader {
	return c.indexStore
}

func (c backend) CommitmentStore() kvstore.Store {
	return c.commitmentStore
}

func (c backend) IndexStore() kvstore.Store {
	return c.indexStore
}

func (c backend) Hooks() Hooks {
	return c.hooks
}

// BackendOptions defines options for creating a Backend.
// Context can optionally define two stores - a commitment store
// that is backed by a merkle tree and an index store that isn't.
// If the index store is not defined, the commitment store will be
// used for all operations.
type BackendOptions struct {

	// CommitmentStore is the commitment store.
	CommitmentStore kvstore.Store

	// IndexStore is the optional index store.
	// If it is nil the CommitmentStore will be used.
	IndexStore kvstore.Store

	// Hooks are optional hooks into ORM insert, update and delete operations.
	Hooks Hooks
}

// NewBackend creates a new Backend.
func NewBackend(options BackendOptions) Backend {
	indexStore := options.IndexStore
	if indexStore == nil {
		indexStore = options.CommitmentStore
	}
	return &backend{
		commitmentStore: options.CommitmentStore,
		indexStore:      indexStore,
		hooks:           options.Hooks,
	}
}

// WrapContextDefault performs the default wrapping of a backend in a context.
// This should be used primarily for testing purposes and production code
// should use some other framework specific wrapping (for instance using
// "store keys").
func WrapContextDefault(backend ReadBackend) context.Context {
	return context.WithValue(context.Background(), defaultContextKey, backend)
}

type contextKeyType string

var defaultContextKey = contextKeyType("backend")

func getBackendDefault(ctx context.Context) (Backend, error) {
	return ctx.Value(defaultContextKey).(Backend), nil
}

func getReadBackendDefault(ctx context.Context) (ReadBackend, error) {
	return ctx.Value(defaultContextKey).(ReadBackend), nil
}
