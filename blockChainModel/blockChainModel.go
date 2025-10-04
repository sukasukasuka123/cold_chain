package blockChainModel

import "time"

// 交易
type Transaction struct {
	TxID string      `json:"tx_id"`
	Data interface{} `json:"data"`
	SenderPubKey string      `json:"sender_pub_key"` 
	Signature    string      `json:"signature"`  
}

// 区块头
type BlockHeader struct {
	Index     int       `json:"index"`
	Timestamp time.Time `json:"timestamp"`
	PrevHash  string    `json:"prev_hash"`
	Hash      string    `json:"hash"`
	Nonce     int       `json:"nonce"`
}

// 区块体
type BlockBody struct {
	Transactions []Transaction `json:"transactions"`
}

// 区块
type Block struct {
	Header BlockHeader `json:"header"`
	Body   BlockBody   `json:"body"`
}
