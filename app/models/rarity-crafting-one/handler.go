package raritycraftingone

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BoundContract(contractAddr common.Address, contractAbi abi.ABI, client *ethclient.Client) *RarityCraftingOneContract {
	c := bind.NewBoundContract(contractAddr, contractAbi, client, client, client)
	return &RarityCraftingOneContract{c}
}

func (r *RarityCraftingOneContract) BalanceOf(addr common.Address) (result []interface{}, err error) {
	err = r.contract.Call(nil, &result, "balanceOf", addr)
	if err != nil {
		log.Println("BalanceOf err :", err)
	}
	return
}

func (r *RarityCraftingOneContract) TokenOfOwnerByIndex(addr common.Address, index *big.Int) (result []interface{}, err error) {
	err = r.contract.Call(nil, &result, "tokenOfOwnerByIndex", addr, index)
	if err != nil {
		log.Println("TokenOfOwnerByIndex err :", err)
	}
	return
}

func (r *RarityCraftingOneContract) Items(input *big.Int) (result []interface{}, err error) {
	err = r.contract.Call(nil, &result, "items", input)
	if err != nil {
		log.Println("Items err :", err)
	}
	return
}
