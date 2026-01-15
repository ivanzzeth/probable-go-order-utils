package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Contracts struct {
	Exchange         common.Address
	FeeModule        common.Address
	NegRiskExchange  common.Address
	NegRiskFeeModule common.Address
	NegRiskAdapter   common.Address
	Collateral       common.Address
	Conditional      common.Address
}

var (
	// BNB Chain Mainnet (Chain ID: 56)
	// Contract addresses from probable-go-contracts/config.go
	_BNB_CHAIN_CONTRACTS = &Contracts{
		Exchange:         common.HexToAddress("0xF99F5367ce708c66F0860B77B4331301A5597c86"), // CTF Exchange (from probable-go-contracts)
		FeeModule:        common.Address{},                                                  // Not used by Probable Markets
		NegRiskExchange:  common.Address{},                                                  // Not used by Probable Markets
		NegRiskFeeModule: common.Address{},                                                  // Not used by Probable Markets
		NegRiskAdapter:   common.Address{},                                                  // Not used by Probable Markets
		Collateral:       common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), // USDT (from probable-go-contracts)
		Conditional:      common.HexToAddress("0x364d05055614B506e2b9A287E4ac34167204cA83"), // ConditionalTokens (from probable-go-contracts)
	}
)

func GetContracts(chainId int64) (*Contracts, error) {
	switch chainId {
	case 56:
		return _BNB_CHAIN_CONTRACTS, nil
	default:
		return nil, fmt.Errorf("invalid chain id: %d, only BNB Chain (56) is supported", chainId)
	}
}
