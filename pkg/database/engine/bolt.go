package engine

import (
	"go.etcd.io/bbolt"

	"github.com/iotaledger/inx-api-core-v0/pkg/database/bolt"
)

// NewBoltDB creates a new bolt DB instance.
func NewBoltDB(path string) (*bbolt.DB, error) {
	opts := &bbolt.Options{
		NoSync: true,
	}

	return bolt.CreateDB(path, opts)
}
