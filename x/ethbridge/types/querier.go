package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/peggy/x/oracle"
)

// defines the params for the following queries:
// - 'custom/ethbridge/prophecies/'
type QueryEthProphecyParams struct {
	ChainID               int             `json:"chain_id"`
	BridgeContractAddress EthereumAddress `json:"bridge_contract_address"`
	Nonce                 int             `json:"nonce"`
	Symbol                string          `json:"symbol"`
	TokenContractAddress  EthereumAddress `json:"token_contract_address"`
	EthereumSender        EthereumAddress `json:"ethereum_sender"`
}

func NewQueryEthProphecyParams(chainID int, bridgeContractAddress EthereumAddress, nonce int, symbol string, tokenContractAddress EthereumAddress, ethereumSender EthereumAddress) QueryEthProphecyParams {
	return QueryEthProphecyParams{
		ChainID:               chainID,
		BridgeContractAddress: bridgeContractAddress,
		Nonce:                 nonce,
		Symbol:                symbol,
		TokenContractAddress:  tokenContractAddress,
		EthereumSender:        ethereumSender,
	}
}

// Query Result Payload for an eth prophecy query
type QueryEthProphecyResponse struct {
	ID     string           `json:"id"`
	Status oracle.Status    `json:"status"`
	Claims []EthBridgeClaim `json:"claims"`
}

func NewQueryEthProphecyResponse(id string, status oracle.Status, claims []EthBridgeClaim) QueryEthProphecyResponse {
	return QueryEthProphecyResponse{
		ID:     id,
		Status: status,
		Claims: claims,
	}
}

func (response QueryEthProphecyResponse) String() string {
	prophecyJSON, err := json.Marshal(response)
	if err != nil {
		return fmt.Sprintf("Error marshalling json: %v", err)
	}

	return string(prophecyJSON)
}
