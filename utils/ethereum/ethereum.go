package ethereum

import (
	"fmt"
	"rarity-backend/app/models"
	"rarity-backend/utils/web3"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetClient() *ethclient.Client {
	client, err := ethclient.Dial(models.FTM_RPC_URL)
	if err != nil {
		fmt.Println("get client err :", err)
	}
	return client
}
func CallFunc() *bind.CallOpts {
	addr := common.HexToAddress(models.ADMIN_ADDR)
	txOpt := web3.CallTxOpts(addr)
	return txOpt
}
