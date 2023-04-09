package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TokenBalanceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument.")
		return
	}

	arg := os.Args[1]
	// Define the URL
	url := "https://api-testnet.polygonscan.com/api?module=account&action=tokenbalance&contractaddress=0x16581f93797e33fd2b1a3497822adf1762ee36e2&address=" + arg + "&tag=latest&apikey=MWUUNIZK5Y9WZ2MUJJHKTU36PE73MT8ITB"

	// Send a GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to call API:", err)
		return
	}
	defer response.Body.Close()

	// Decode the JSON response
	var tokenBalanceResponse TokenBalanceResponse
	err = json.NewDecoder(response.Body).Decode(&tokenBalanceResponse)
	if err != nil {
		fmt.Println("Failed to decode JSON response:", err)
		return
	}

	// Check the response status
	if tokenBalanceResponse.Status != "1" {
		fmt.Println("API returned an error:", tokenBalanceResponse.Message)
		return
	}

	// Print the token balance
	fmt.Println("Token balance:", tokenBalanceResponse.Result)
}

// This is the code i tried to get the metadata from the polygon mumbai testnet.

// package main

// import (
// 	"fmt"
// 	"log"
// 	"math/big"
// 	"os"
// 	"strings"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// const (
// 	contractAddress = "0x16581f93797e33fd2b1a3497822adf1762ee36e2"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		log.Fatal("Usage: go run main.go <wallet address>")
// 	}
// 	key, _ := crypto.GenerateKey()
// 	auth := bind.NewKeyedTransactor(key)

// 	sim := backends.NewSimulatedBackend(core.GenesisAccount{Address: auth.From, Balance: big.NewInt(10000000000)})

// 	walletAddress := os.Args[1]
// 	client, err := ethclient.Dial("https://rpc-mumbai.matic.today")
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
// 	}
// 	fmt.Println(walletAddress)

// 	// balance, err := client.BalanceAt(context.Background(), common.HexToAddress(contractAddress), nil)
// 	// fmt.Println(balance)

// 	store, err := NewStorage(common.HexToAddress(contractAddress), client)

// 	// fmt.Println(store.Retrieve(nil))
// 	// totalTokens, err := store.Retrieve(&bind.CallOpts{})

// 	// fmt.Println(totalTokens)
// 	auth, err := bind.NewStorageTransactor(strings.NewReader(key), "strong_password")
// 	if err != nil {
// 		log.Fatalf("Failed to create authorized transactor: %v", err)
// 	}
// 	// Call the store() function
// 	tx, err := store.Store(auth, big.NewInt(420))
// 	if err != nil {
// 		log.Fatalf("Failed to update value: %v", err)
// 	}
// 	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
// 	// contractAbi, err := abi.JSON(strings.NewReader(abiJson))
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to parse ABI: %v", err)
// 	// }

// 	// contractAddress := common.HexToAddress(contractAddress)
// 	// contract, err := NewContract(contractAddress, client)
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to instantiate contract: %v", err)
// 	// }

// 	// totalTokens, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(walletAddress))
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to get balance: %v", err)
// 	// }

// 	// fmt.Printf("Total tokens for address %s: %d\n", walletAddress, totalTokens)

// 	// metadata, err := contract.TokenURI(&bind.CallOpts{}, big.NewInt(0))
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to get token metadata: %v", err)
// 	// }

// 	// if metadata != "" {
// 	// 	fmt.Printf("Metadata for first token:\n%s\n", metadata)
// 	// } else {
// 	// 	fmt.Println("No tokens in wallet")
// 	// }
// }

// // NewStorage creates a new instance of Storage, bound to a specific deployed contract.
// // func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
// // 	contract, err := bindStorage(address, backend, backend, backend)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
// // }

// const abiJson = `[
//     {
//         "constant": true,
//         "inputs": [
//             {
//                 "name": "owner",
//                 "type": "address"
//             }
//         ],
//         "name": "balanceOf",
//         "outputs": [
//             {
//                 "name": "",
//                 "type": "uint256"
//             }
//         ],
//         "payable": false,
//         "stateMutability": "view",
//         "type": "function"
//     },
//     {
//         "constant": true,
//         "inputs": [
//             {
//                 "name": "_tokenId",
//                 "type": "uint256"
//             }
//         ],
//         "name": "tokenURI",
//         "outputs": [
//             {
//                 "name": "",
//                 "type": "string"
//             }
//         ],
//         "payable": false,
//         "stateMutability": "view",
//         "type": "function"
//     }
// ]`
