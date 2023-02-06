package database

import (
	"github.com/iotaledger/hive.go/core/generics/lo"
	"github.com/iotaledger/inx-api-core-v0/pkg/hornet"
)

func (db *Database) WasAddressSpentFrom(address hornet.Hash) bool {
	return lo.PanicOnErr(db.spentAddressesStore.Has(address[:49]))
}
