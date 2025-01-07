package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math"
	"math/big"
)

// online test web: https://developer.metamask.io/key/bc339063a9b54795a3097d2e9b7d413f/active-endpoints
var httpUrl = "https://sepolia.infura.io/v3/bc339063a9b54795a3097d2e9b7d413f"
var localUrl = "http://127.0.0.1:8545" // commond line input: ganache-cli
var dir = "/Users/jiaxi/project/learning/web3/projects/web3-go/web3-golang/wallet"

func initClient() *ethclient.Client {
	fmt.Println("hello")
	client, err := ethclient.DialContext(context.Background(), localUrl)
	if err != nil {
		log.Fatalf("error to create a eth client:%v", err)
		return nil
	}
	defer client.Close()
	fmt.Println("get client success!")
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("error to get block:%v", err)
		return nil
	}
	fmt.Println(block.Number())
	// section 2: https://www.youtube.com/watch?v=nuivtRUaSw8&list=PLay9kDOVd_x7hbhssw4pTKZHzzc6OG0e_&index=2
	return client
}

func getBalance() {
	addr := "0xf88554657c1b821dB79B305A606d6fB31CF32df1"
	address := common.HexToAddress(addr)
	ethClient := initClient()

	balance, err := ethClient.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("error to get balance:%v", err)
	}
	fmt.Println("get balance is: ", balance) //100000000000000000000
	// convert big int to big float 1eth = 10^18 wei
	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	fmt.Println("get float balance is: ", floatBalance)
	value := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("the etc num is: ", value)
}

func ethKeys() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("error to generate key:%v", err)
	}

	pData := crypto.FromECDSA(pvk)
	fmt.Println("get pData: ", pData)  //[39 35 9 84 171 116 218 17 155 167 33 149 71 228 196 54 21 131 55 181 191 14 126 58 10 25 226 190 84 221 231 244]
	fmt.Println(hexutil.Encode(pData)) //0x10a5501763def2089674b32fd70082d97fcdc837b73a5f3b61436474682aa55b

	publicKey := pvk.PublicKey
	pubPdata := crypto.FromECDSAPub(&publicKey)
	fmt.Println(hexutil.Encode(pubPdata)) // 0x04fc0f2428d012385c9ca1e1e4ae56b9e8c2b2994b043da2c97bc1f7b6be2a9f0334b2c87b39e0bd02205c7f81402c40aa1d437448f56428c3814135e2b1fd5739

	address := crypto.PubkeyToAddress(publicKey).Hex()
	fmt.Println("the address is: ", address) // 0x1FAac65B4Fd225055972D41acc239A4813320A96

}

func wallet() {

	key := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	passwd := "passwd"
	account, err := key.NewAccount(passwd)
	if err != nil {
		log.Fatalf("error to generate wallet:%v", err)
	}
	fmt.Println("get the new account: ", account.Address)

}

func tryGetPrivateKeyFromWalletByPasswd() {
	walletPath := "/Users/jiaxi/project/learning/web3/projects/web3-go/web3-golang/wallet/UTC--2025-01-06T14-47-15.202426000Z--948e85a429c18ff8ed550df1a2424424bb989d5c"
	passwd := "passwdd"
	b, err := ioutil.ReadFile(walletPath)
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, passwd)
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println(hexutil.Encode(pData)) //0xc7447f7ebe3d0ce6d0e46fc6d34d573beb1ef8556d353f5000f336d0776bd22a

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("get publicKey: ", hexutil.Encode(pData)) // 0x04a2041fe747385c357e9d0cac5c35614eb9b58f980e05076ab0b3aa89b138c2afab8e6d4f5aaa1ee59bde51b94347b40bcbfaf1e8bc12675bf6df41f510d24692

	fmt.Println(crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()) //0x948e85a429c18fF8Ed550dF1A2424424bb989d5c
}

func test4ether() {
	//ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	//_, err := ks.NewAccount("password")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_, err = ks.NewAccount("password")
	//if err != nil {
	//	log.Fatal(err)
	//}
	addr1 := "0x2CD1b8cCe7b1F3c2e2514CfB26Ef4C02daE59143"
	addr2 := "71d682da0658de7e38da835c44b2d5e0ceca42c0"
	client, err := ethclient.Dial(localUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	a1 := common.HexToAddress(addr1)
	a2 := common.HexToAddress(addr2)
	balance1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}
	balance2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance1) // 100000000000000000000
	fmt.Println(balance2) // 0

	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}
	amount := big.NewInt(1000000000000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("aaaaaa")
	prvKeyStr := "af6f17ec2368f152a2cabe8423bc108fe8ef864c86d9285358e3662fb5a1b045"
	pvk, err := converPrvKey2Struct(prvKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), pvk)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx send ", tx.Hash().Hex()) //0xa32fab687701b4871ac621100bcd0f25c2225cfe9015ec48c7c6b1a7da1cb7e3
}

func converPrvKey2Struct(prvStr string) (*ecdsa.PrivateKey, error) {
	privateKeyHex := prvStr // 替换为你的私钥字符串（不带 "0x" 前缀）

	// 将十六进制字符串转换为字节数组
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode private key hex: %v", err)
		return nil, err
	}

	// 使用 go-ethereum 的 crypto 库将字节数组解析为 *ecdsa.PrivateKey
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to convert to ECDSA private key: %v", err)
		return nil, err
	}
	return privateKey, nil
}

func main() {
	// getBalance()
	// ethKeys()
	//wallet()
	//tryGetPrivateKeyFromWalletByPasswd()
	//test4ether()
	// https://www.youtube.com/watch?v=NFVakpTKvUg
}
