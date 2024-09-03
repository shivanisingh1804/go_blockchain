// package main

// import (
// 	"crypto/md5"
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// type Block struct {
// 	Pos       int
// 	Data      CarPurchase
// 	TimeStamp string
// 	Hash      string
// 	PrevHash  string
// }
// type CarPurchase struct {
// 	AddharNo     string `json:"addhar_no"`
// 	Owner        string `json:"owner"`
// 	PurchaseDate string `json:"purchase_date"`
// 	IsGenesis    bool   `json:"is_genesis"`
// }
// type Car struct {
// 	ID         string `json:"id"`
// 	Model      string `json:"model"`
// 	Company    string `json:"company"`
// 	LaunchDate string `json:"launch_date"`
// 	ChessisNo  string `json:"chessis_no"`
// }
// type Blochain struct {
// 	blocks []*Block
// }

// var BlockChain *Blochain

// func (b *Block) generateHash() {
// 	bytes, _ := json.Marshal(b.Data)

// 	data := string(b.Pos) + b.TimeStamp + string(bytes) + b.PrevHash

// 	hash := sha256.New()
// 	hash.Write([]byte(data))
// 	b.Hash = hex.EncodeToString(hash.Sum(nil))

// }

// func CreateBlock(prevBlock *Block, checkoutitem CarPurchase) *Block {
// 	block := &Block{}
// 	block.TimeStamp = time.Now().String()
// 	block.Pos = prevBlock.Hash
// 	block.generateHash()
// 	return block

// }

// func (bc *Blochain) AddBlock(data CarPurchase) {
// 	prevBlock := bc.blocks[len(bc.blocks)-1]

// 	block := CreateBlock(prevBlock, data)
// 	if validBlock(block, prevBlock) {
// 		bc.blocks = append(bc.blocks, block)
// 	}

// }

// func validBlock(block, prevBlock *Block) bool {
// 	if prevBlock.Hash != block.PrevHash {
// 		return false
// 	}
// 	if !block.validateHash(block.Hash) {
// 		return false
// 	}
// 	if prevBlock.Pos+1 != block.Pos {
// 		return false
// 	}
// 	return true

// }

// func (b *Block) validateHash(hash string) bool {
// 	b.generateHash()
// 	if b.Hash != hash {
// 		return false
// 	}
// 	return true
// }

// func writeBlock(w http.ResponseWriter, r *http.Request) {
// 	var checkoutitem CarPurchase

// 	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("Could not create:%v", err)
// 		w.Write([]byte("counld not create new block"))
// 		return
// 	}
// 	BlockChain.AddBlock(checkoutitem)
// }

// func newCar(w http.ResponseWriter, r *http.Request) {
// 	var car Car

// 	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("Could not create:%v", err)
// 		w.Write([]byte("counld not create new car"))
// 		return
// 	}
// 	h := md5.New()
// 	io.WriteString(h, car.LaunchDate+car.ChessisNo)
// 	car.ID = fmt.Sprintf("%x", h.Sum(nil))

// 	resp, err := json.MarshalIndent(car, "", " ")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("could not marshal payload: %v", err)
// 		w.Write([]byte("could not save car data"))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resp)

// }

// func GenesisBlock() *Block {
// 	return CreateBlock(&Block{}, CarPurchase{IsGenesis: true})
// }

// func NewBlockchain() *Blochain {
// 	return &Blochain{[]*Block{GenesisBlock()}}

// }

// func getBlockchain(w http.ResponseWriter, r *http.Request) {
// 	jbytes, err := json.MarshalIndent(BlockChain.blocks, "", " ")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}
// 	io.WriteString(w, string(jbytes))
// }

// func main() {
// 	BlockChain = NewBlockchain()
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", getBlockchain).Methods("GET")
// 	r.HandleFunc("/", writeBlock).Methods("POST")
// 	r.HandleFunc("/new", newCar).Methods("POST")
// 	go func() {
// 		for _, block := range BlockChain.blocks {
// 			fmt.Printf("Prev.hash:%x\n", block.PrevHash)
// 			bytes, _ := json.MarshalIndent(block.Data, "", " ")
// 			fmt.Printf("Data:%v\n", string(bytes))
// 			fmt.Printf("Hash%x\n", block.Hash)
// 			fmt.Println()
// 		}
// 	}()

//		log.Println("Listening port 3000")
//		log.Fatal(http.ListenAndServe(":3000", r))
//	}
package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos       int
	Data      CarPurchase
	TimeStamp string
	Hash      string
	PrevHash  string
}

type CarPurchase struct {
	AddharNo     string `json:"addhar_no"`
	Owner        string `json:"owner"`
	PurchaseDate string `json:"purchase_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

type Car struct {
	ID         string `json:"id"`
	Model      string `json:"model"`
	Company    string `json:"company"`
	LaunchDate string `json:"launch_date"`
	ChessisNo  string `json:"chessis_no"`
}

type Blockchain struct {
	blocks []*Block
}

var BlockChain *Blockchain

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)
	data := string(b.Pos) + b.TimeStamp + string(bytes) + b.PrevHash

	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func CreateBlock(prevBlock *Block, checkoutitem CarPurchase) *Block {
	block := &Block{}
	block.TimeStamp = time.Now().String()
	block.Pos = prevBlock.Pos + 1
	block.PrevHash = prevBlock.Hash
	block.Data = checkoutitem
	block.generateHash()
	return block
}

func (bc *Blockchain) AddBlock(data CarPurchase) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := CreateBlock(prevBlock, data)
	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

func validBlock(block, prevBlock *Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	}
	if !block.validateHash(block.Hash) {
		return false
	}
	if prevBlock.Pos+1 != block.Pos {
		return false
	}
	return true
}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	if b.Hash != hash {
		return false
	}
	return true
}

func writeBlock(w http.ResponseWriter, r *http.Request) {
	var checkoutitem CarPurchase
	if err := json.NewDecoder(r.Body).Decode(&checkoutitem); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not create: %v", err)
		w.Write([]byte("Could not create new block"))
		return
	}
	BlockChain.AddBlock(checkoutitem)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Block added successfully"))
}

func newCar(w http.ResponseWriter, r *http.Request) {
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not create: %v", err)
		w.Write([]byte("Could not create new car"))
		return
	}
	h := md5.New()
	io.WriteString(h, car.LaunchDate+car.ChessisNo)
	car.ID = fmt.Sprintf("%x", h.Sum(nil))

	resp, err := json.MarshalIndent(car, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not marshal payload: %v", err)
		w.Write([]byte("Could not save car data"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, CarPurchase{IsGenesis: true})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	var carPurchases []CarPurchase
	for _, block := range BlockChain.blocks {
		carPurchases = append(carPurchases, block.Data)
	}

	jbytes, err := json.MarshalIndent(carPurchases, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	io.WriteString(w, string(jbytes))
}

func main() {
	BlockChain = NewBlockchain()
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newCar).Methods("POST")

	// Logging middleware
	r.Use(loggingMiddleware)

	go func() {
		for _, block := range BlockChain.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
	}()

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

// Logging middleware to log request details
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
