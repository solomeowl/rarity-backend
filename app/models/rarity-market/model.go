package raritymarket

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type RarityMarketContract struct {
	contract *bind.BoundContract
}

type RarityMarket struct {
	ListId  *big.Int       `json:"listId"`
	TokenID *big.Int       `json:"tokenID"`
	Owner   common.Address `json:"owner"`
	Buyer   common.Address `json:"buyer"`
	Price   *big.Int       `json:"price"`
	Payout  *big.Int       `json:"payout"`
	Status  uint8          `json:"status"`
}
