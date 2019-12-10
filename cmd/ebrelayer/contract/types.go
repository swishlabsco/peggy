package contract

import "errors"

// Type : enum containing supported contract names
type Type int

const (
	// BridgeRegistry : registers deployed addresses of the other contracts
	BridgeRegistry Type = iota + 1
	// Valset : manages the validator set and signature verification
	Valset
	// Oracle : enables validators to make OracleClaims and processes ProphecyClaims
	Oracle
	// CosmosBridge : enables validators to make ProphecyClaims
	CosmosBridge
	// BridgeBank : manages protocol assets on both Ethereum and Cosmos
	BridgeBank
)

// TypeToString : returns the string associated with a Type
var TypeToString = [...]string{"BridgeRegistry", "Valset", "Oracle", "CosmosBridge", "BridgeBank"}

// StringToType : returns the Type associated with a string
var StringToType = map[string]Type{
	"BridgeRegistry": BridgeRegistry,
	"Valset":         Valset,
	"Oracle":         Oracle,
	"CosmosBridge":   CosmosBridge,
	"BridgeBank":     BridgeBank,
}

// String : returns the Type as a string
func (contract Type) String() string {
	return TypeToString[contract-1]
}

// IsValid : checks if the string is a valid Type
func (contract Type) IsValid() error {
	switch contract {
	case BridgeRegistry, Valset, Oracle, CosmosBridge, BridgeBank:
		return nil
	}
	return errors.New("invalid contract type")
}
