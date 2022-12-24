// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "time"

// EthDetailTX is an object representing the database table.
type EthDetailTX struct { // ID
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	// eth_tx table ID
	TXID int64 `boil:"tx_id" json:"tx_id" toml:"tx_id" yaml:"tx_id"`
	// UUID
	UUID string `boil:"uuid" json:"uuid" toml:"uuid" yaml:"uuid"`
	// current transaction type
	CurrentTXType int8 `boil:"current_tx_type" json:"current_tx_type" toml:"current_tx_type" yaml:"current_tx_type"`
	// sender account
	SenderAccount string `boil:"sender_account" json:"sender_account" toml:"sender_account" yaml:"sender_account"`
	// sender address
	SenderAddress string `boil:"sender_address" json:"sender_address" toml:"sender_address" yaml:"sender_address"`
	// receiver account
	ReceiverAccount string `boil:"receiver_account" json:"receiver_account" toml:"receiver_account" yaml:"receiver_account"`
	// receiver address
	ReceiverAddress string `boil:"receiver_address" json:"receiver_address" toml:"receiver_address" yaml:"receiver_address"`
	// amount of coin to receive
	Amount uint64 `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	// fee
	Fee uint64 `boil:"fee" json:"fee" toml:"fee" yaml:"fee"`
	// gas limit
	GasLimit uint32 `boil:"gas_limit" json:"gas_limit" toml:"gas_limit" yaml:"gas_limit"`
	// nonce
	Nonce uint64 `boil:"nonce" json:"nonce" toml:"nonce" yaml:"nonce"`
	// HEX string for unsigned transaction
	UnsignedHexTX string `boil:"unsigned_hex_tx" json:"unsigned_hex_tx" toml:"unsigned_hex_tx" yaml:"unsigned_hex_tx"`
	// HEX string for signed transaction
	SignedHexTX string `boil:"signed_hex_tx" json:"signed_hex_tx" toml:"signed_hex_tx" yaml:"signed_hex_tx"`
	// Hash for sent transaction
	SentHashTX string `boil:"sent_hash_tx" json:"sent_hash_tx" toml:"sent_hash_tx" yaml:"sent_hash_tx"`
	// updated date for unsigned transaction created
	UnsignedUpdatedAt *time.Time `boil:"unsigned_updated_at" json:"unsigned_updated_at,omitempty" toml:"unsigned_updated_at" yaml:"unsigned_updated_at,omitempty" database:"default:null"`
	// updated date for signed transaction sent
	SentUpdatedAt *time.Time `boil:"sent_updated_at" json:"sent_updated_at,omitempty" toml:"sent_updated_at" yaml:"sent_updated_at,omitempty" database:"default:null"`

	R *ethDetailTXR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ethDetailTXL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EthDetailTXColumns = struct {
	ID                string
	TXID              string
	UUID              string
	CurrentTXType     string
	SenderAccount     string
	SenderAddress     string
	ReceiverAccount   string
	ReceiverAddress   string
	Amount            string
	Fee               string
	GasLimit          string
	Nonce             string
	UnsignedHexTX     string
	SignedHexTX       string
	SentHashTX        string
	UnsignedUpdatedAt string
	SentUpdatedAt     string
}{
	ID:                "id",
	TXID:              "tx_id",
	UUID:              "uuid",
	CurrentTXType:     "current_tx_type",
	SenderAccount:     "sender_account",
	SenderAddress:     "sender_address",
	ReceiverAccount:   "receiver_account",
	ReceiverAddress:   "receiver_address",
	Amount:            "amount",
	Fee:               "fee",
	GasLimit:          "gas_limit",
	Nonce:             "nonce",
	UnsignedHexTX:     "unsigned_hex_tx",
	SignedHexTX:       "signed_hex_tx",
	SentHashTX:        "sent_hash_tx",
	UnsignedUpdatedAt: "unsigned_updated_at",
	SentUpdatedAt:     "sent_updated_at",
}

var EthDetailTXTableColumns = struct {
	ID                string
	TXID              string
	UUID              string
	CurrentTXType     string
	SenderAccount     string
	SenderAddress     string
	ReceiverAccount   string
	ReceiverAddress   string
	Amount            string
	Fee               string
	GasLimit          string
	Nonce             string
	UnsignedHexTX     string
	SignedHexTX       string
	SentHashTX        string
	UnsignedUpdatedAt string
	SentUpdatedAt     string
}{
	ID:                "eth_detail_tx.id",
	TXID:              "eth_detail_tx.tx_id",
	UUID:              "eth_detail_tx.uuid",
	CurrentTXType:     "eth_detail_tx.current_tx_type",
	SenderAccount:     "eth_detail_tx.sender_account",
	SenderAddress:     "eth_detail_tx.sender_address",
	ReceiverAccount:   "eth_detail_tx.receiver_account",
	ReceiverAddress:   "eth_detail_tx.receiver_address",
	Amount:            "eth_detail_tx.amount",
	Fee:               "eth_detail_tx.fee",
	GasLimit:          "eth_detail_tx.gas_limit",
	Nonce:             "eth_detail_tx.nonce",
	UnsignedHexTX:     "eth_detail_tx.unsigned_hex_tx",
	SignedHexTX:       "eth_detail_tx.signed_hex_tx",
	SentHashTX:        "eth_detail_tx.sent_hash_tx",
	UnsignedUpdatedAt: "eth_detail_tx.unsigned_updated_at",
	SentUpdatedAt:     "eth_detail_tx.sent_updated_at",
}


// ethDetailTXR is where relationships are stored.
type ethDetailTXR struct {
}

// NewStruct creates a new relationship struct
func (*ethDetailTXR) NewStruct() *ethDetailTXR {
	return &ethDetailTXR{}
}

// ethDetailTXL is where Load methods for each relationship are stored.
type ethDetailTXL struct{}