package raritycraftingone

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"rarity-backend/app/models"
	raritycraftingone "rarity-backend/app/models/rarity-crafting-one"
	"rarity-backend/utils/e"
	"rarity-backend/utils/ethereum"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetCraftingByOwner(c *gin.Context) (int, int, int, interface{}) {
	addr := c.Query("address")
	baseTypeQ := c.Query("base_type")
	baseTypeReq := -1
	if addr == "" {
		return http.StatusOK, e.PARAMETER_ERROR, 0, nil
	}
	if baseTypeQ != "" {
		baseTypeReq, _ = strconv.Atoi(baseTypeQ)
	}
	client := ethereum.GetClient()
	abiJson, err := ioutil.ReadFile("./contract/rarity-crafting-one/abi.json")
	if err != nil {
		return http.StatusOK, e.SERVER_ERROR, 0, err.Error()
	}
	contractAbi, err := abi.JSON(bytes.NewReader(abiJson))
	if err != nil {
		return http.StatusOK, e.SERVER_ERROR, 0, err.Error()
	}
	contract := raritycraftingone.BoundContract(common.HexToAddress(models.RARITY_CRAFTING_ONE_ADDR), contractAbi, client)
	balanceResp, err := contract.BalanceOf(common.HexToAddress(addr))
	if err != nil {
		return http.StatusOK, e.SERVER_ERROR, 0, err.Error()
	}
	value := reflect.ValueOf(balanceResp)
	balance := value.Index(0).Interface().(*big.Int)
	var resp []raritycraftingone.RarityCraftingOneItemResp
	for i := 0; i < int(balance.Int64()); i++ {
		tokenResp, err := contract.TokenOfOwnerByIndex(common.HexToAddress(addr), big.NewInt(int64(i)))
		if err != nil {
			return http.StatusOK, e.SERVER_ERROR, 0, err.Error()
		}
		valueToken := reflect.ValueOf(tokenResp)
		token := valueToken.Index(0).Interface().(*big.Int)

		itemResp, err := contract.Items(token)
		if err != nil {
			return http.StatusOK, e.SERVER_ERROR, 0, err.Error()
		}
		var baseType, itemType uint8
		valueItem := reflect.ValueOf(itemResp)
		log.Println(valueItem)
		for i := 0; i < valueItem.Len(); i++ {
			if i == 0 {
				baseType = valueItem.Index(i).Interface().(uint8)
			} else if i == 1 {
				log.Println(valueItem.Index(i))
				itemType = valueItem.Index(i).Interface().(uint8)
			}
		}
		tmp := raritycraftingone.RarityCraftingOneItemResp{
			ItemId:   token,
			BaseType: baseType,
			ItemType: itemType,
		}

		if baseTypeReq != -1 {
			if baseTypeReq == 1 && baseType == 1 {
				resp = append(resp, tmp)
			}
			if baseTypeReq == 2 && baseType == 2 {
				resp = append(resp, tmp)
			}
			if baseTypeReq == 3 && baseType == 3 {
				resp = append(resp, tmp)
			}
		} else {
			resp = append(resp, tmp)
		}

	}
	return http.StatusOK, e.SUCCESS, len(resp), resp
}
