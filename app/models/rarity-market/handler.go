package raritymarket

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BoundContract(contractAddr common.Address, contractAbi abi.ABI, client *ethclient.Client) *RarityMarketContract {
	c := bind.NewBoundContract(contractAddr, contractAbi, client, client, client)
	return &RarityMarketContract{c}
}

func (r *RarityMarketContract) GetAllSummoners() (result []interface{}, err error) {
	err = r.contract.Call(nil, &result, "getAllSummoners")
	if err != nil {
		log.Println("GetAllSummoners err :", err)
	}
	return
}
