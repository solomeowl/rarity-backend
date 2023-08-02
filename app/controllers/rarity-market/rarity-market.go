package raritymarket

import (
	"bytes"
	"io/ioutil"
	"math/big"
	"net/http"
	"rarity-backend/app/models"
	raritymarket "rarity-backend/app/models/rarity-market"
	"rarity-backend/utils/e"
	"rarity-backend/utils/ethereum"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetAllSummoners(c *gin.Context) (int, int, int, interface{}) {
	statusQ := c.Query("status")
	pageQ := c.Query("page")
	sizeQ := c.Query("size")
	addr := c.Query("address")
	status := -1
	var page, size int
	if statusQ != "" {
		status, _ = strconv.Atoi(c.Query("status"))
	}
	if pageQ == "" || sizeQ == "" {
		page = -1
		size = -1
	} else {
		page, _ = strconv.Atoi(c.Query("page"))
		size, _ = strconv.Atoi(c.Query("size"))
	}

	client := ethereum.GetClient()
	abiJson, err := ioutil.ReadFile("./contract/rarity-market/abi.json")
	if err != nil {
		return http.StatusOK, e.SERVER_ERROR, 0, err.Error()
	}
	contractAbi, err := abi.JSON(bytes.NewReader(abiJson))
	if err != nil {
		return http.StatusOK, e.SERVER_ERROR, 0, err.Error()
	}
	contract := raritymarket.BoundContract(common.HexToAddress(models.RARITY_MARKET_ADDR), contractAbi, client)
	contractResp, err := contract.GetAllSummoners()
	if err != nil {
		return http.StatusOK, e.SERVER_ERROR, 0, err.Error()
	}
	value := reflect.ValueOf(contractResp[0])
	var tmpList, list []raritymarket.RarityMarket
	for i := 0; i < value.Len(); i++ {
		tmp := value.Index(i).Interface().(struct {
			ListId  *big.Int       `json:"listId"`
			TokenID *big.Int       `json:"tokenID"`
			Owner   common.Address `json:"owner"`
			Buyer   common.Address `json:"buyer"`
			Price   *big.Int       `json:"price"`
			Payout  *big.Int       `json:"payout"`
			Status  uint8          `json:"status"`
		})
		tmpRarity := raritymarket.RarityMarket{
			ListId:  tmp.ListId.String(),
			TokenID: tmp.TokenID,
			Owner:   tmp.Owner,
			Buyer:   tmp.Buyer,
			Price:   tmp.Price,
			Payout:  tmp.Payout,
			Status:  tmp.Status,
		}
		if status != -1 {
			if status == 0 && tmp.Status == 0 {
				tmpList = append(tmpList, tmpRarity)
			}
			if status == 1 && tmp.Status == 1 {
				tmpList = append(tmpList, tmpRarity)
			}
			if status == 2 && tmp.Status == 2 {
				tmpList = append(tmpList, tmpRarity)
			}
		} else {
			tmpList = append(tmpList, tmpRarity)
		}
	}

	if addr != "" {
		for _, data := range tmpList {
			if common.HexToAddress(addr) == data.Owner {
				list = append(list, data)
			}
		}
	} else {
		list = tmpList
	}

	total := len(list)
	var min, max int
	if page == -1 || size == -1 {
		min = 0
		max = total
	} else {
		min = page * size
		max = size * (page + 1)
		if max >= total {
			return http.StatusOK, e.PARAMETER_ERROR, 0, "out of range"
		}
	}
	return http.StatusOK, e.SUCCESS, total, list[min:max]
}
