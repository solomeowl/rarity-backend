package raritycraftingone

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type RarityCraftingOneContract struct {
	contract *bind.BoundContract
}

type RarityCraftingOneItemResp struct {
	ItemId   *big.Int `json:"item_id"`
	BaseType uint8    `json:"base_type"`
	ItemType uint8    `json:"item_type"`
}
