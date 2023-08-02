package web3

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client 連線到節點的兩種物件
type Client struct {
	RpcClient *rpc.Client
	EthClient *ethclient.Client
}

// Key 公鑰及私鑰
type Key struct {
	PublicKey  common.Address
	PrivateKey *ecdsa.PrivateKey
}

type Response struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

// Connect 連線到以太節點
func Connect(host string) (*Client, error) {
	ctx, err := rpc.Dial(host)
	if err != nil {
		return nil, err
	}
	conn := ethclient.NewClient(ctx)

	return &Client{RpcClient: ctx, EthClient: conn}, nil
}

//GetBlockNumber 取得現在最新區塊號
func (ec *Client) GetBlockNumber(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := ec.RpcClient.CallContext(ctx, &result, "eth_blockNumber")
	return (*big.Int)(&result), err
}

// GetBlockTxs 取得該區塊中的所有tx
func (ec *Client) GetBlockTxs(blockHeight *big.Int) (types.Transactions, error) {
	time.Sleep(time.Second * 5)
	block, err := ec.EthClient.BlockByNumber(context.TODO(), blockHeight)
	if err != nil {
		log.Println("get Block error msg: ", err)
		return nil, err
	}
	txs := block.Transactions()
	return txs, nil
}

//IsTxPendding 檢查Tx是否正在上鏈
func (ec *Client) IsTxPendding(txHash common.Hash) (bool, error) {
	_, pedding, err := ec.EthClient.TransactionByHash(context.TODO(), txHash)
	if err != nil {
		if err.Error() == "not found" {
			return true, nil
		}
		return true, err
	}
	return pedding, nil
}

//SendRawTransaction 發送交易並回傳txHash
func (ec *Client) SendRawTransaction(tx *types.Transaction) (common.Hash, error) {
	var txHash common.Hash
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return txHash, err
	}

	err = ec.RpcClient.CallContext(
		context.TODO(),
		&txHash,
		"eth_sendRawTransaction",
		common.Bytes2Hex(data),
	)
	return txHash, err
}

// GetNonce 取得該 address 的 nonce
func (ec *Client) GetNonce(address common.Address) (uint64, error) {

	res, resErr := ec.EthClient.PendingNonceAt(context.TODO(), address)
	if resErr != nil {
		return 0, resErr
	}
	return res, nil
}

// GetFilterLogs 取得合約所有資料
func (ec *Client) GetFilterLogs(filter ethereum.FilterQuery) (logs []types.Log, err error) {
	logs, err = ec.EthClient.FilterLogs(context.TODO(), filter)
	return
}

// TxByHash 取得Hash後的Tx
func (ec *Client) TxByHash(hash common.Hash) (tx *types.Transaction, err error) {
	tx, _, err = ec.EthClient.TransactionByHash(context.TODO(), hash)
	return
}

// SubscribeNewBlock 訂閱新的區塊
func (ec *Client) SubscribeNewBlock() (ethereum.Subscription, chan *types.Header) {
	var head = make(chan *types.Header)
	sub, err := ec.EthClient.SubscribeNewHead(context.TODO(), head)
	if err != nil {
		log.Fatal(err)
	}
	return sub, head
}

// GetKey 取得公鑰及私鑰
func GetKey(privateKey string) (*Key, error) {
	var pvkBuffer *ecdsa.PrivateKey
	pvkBuffer, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	publicKey := crypto.PubkeyToAddress(pvkBuffer.PublicKey)

	return &Key{PrivateKey: pvkBuffer, PublicKey: publicKey}, nil
}

// MakeTxOpts 將交易資料包成byte
func MakeTxOpts(from common.Address, nonce *big.Int, value *big.Int, gasPrice *big.Int, gasLimit uint64, privKey *ecdsa.PrivateKey, chainID int64) *bind.TransactOpts {
	txOpts := &bind.TransactOpts{
		From:  from,
		Nonce: nonce,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			Singerbase := types.NewEIP155Signer(big.NewInt(chainID))
			signedTx, err := types.SignTx(tx, Singerbase, privKey)
			if err != nil {
				return nil, err
			}
			return signedTx, nil
		},
		Value:    value,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
	}
	return txOpts
}

//CallTxOpts 合約的 call function
func CallTxOpts(from common.Address) *bind.CallOpts {
	callOpts := &bind.CallOpts{
		From: from,
	}
	return callOpts
}
