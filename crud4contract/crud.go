package crud4contract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	todo "web3-golang/gen"
)

var path = "/Users/jiaxi/project/learning/web3/projects/web3-go/web3-golang/wallet/UTC--2025-01-06T15-03-07.577303000Z--71d682da0658de7e38da835c44b2d5e0ceca42c0"
var localUrl = "http://127.0.0.1:8545"

func ActionWithContract() {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	// 2 初始化client
	client, err := ethclient.Dial(localUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 3 根据client和账户获取address
	addr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		log.Fatal(err)
	}

	// 4 设置交易的汽油费
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 5 获取交易网络信息
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cAddr := "0x71d682da0658de7e38da835c44b2d5e0ceca42c0"
	address := common.HexToAddress(cAddr)
	t, err := todo.NewTodo(address, client)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice
	tx.Nonce = big.NewInt(int64(nonce))

	// 增
	tran, err := t.Add(tx, "First Task")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tran.Hash())

	// 查
	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	tasks, err := t.List(&bind.CallOpts{
		From: add,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)

	// 改
	tran, err = t.Update(tx, big.NewInt(0), "update task content")
	if err != nil {
		log.Fatal(err)
	}

	// 删
	tran, err = t.Remove(tx, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}
}
