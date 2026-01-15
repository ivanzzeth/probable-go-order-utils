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

	_AMOY_CONTRACTS = &Contracts{
		Exchange:         common.HexToAddress("0xdFE02Eb6733538f8Ea35D585af8DE5958AD99E40"),
		FeeModule:        common.HexToAddress("0xE3f18aCc55091e2c48d883fc8C8413319d4Ab7b0"),
		NegRiskExchange:  common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
		NegRiskFeeModule: common.HexToAddress("0xB768891e3130F6dF18214Ac804d4DB76c2C37730"),
		NegRiskAdapter:   common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
		Collateral:       common.HexToAddress("0x9c4e1703476e875070ee25b56a58b008cfb8fa78"),
		Conditional:      common.HexToAddress("0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB"),
	}

	_MATIC_CONTRACTS = &Contracts{
		Exchange:         common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
		FeeModule:        common.HexToAddress("0xE3f18aCc55091e2c48d883fc8C8413319d4Ab7b0"),
		NegRiskExchange:  common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
		NegRiskFeeModule: common.HexToAddress("0xB768891e3130F6dF18214Ac804d4DB76c2C37730"),
		NegRiskAdapter:   common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
		Collateral:       common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
		Conditional:      common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
	}
)

func GetContracts(chainId int64) (*Contracts, error) {
	switch chainId {
	case 56:
		return _BNB_CHAIN_CONTRACTS, nil
	case 137:
		return _MATIC_CONTRACTS, nil
	case 80002:
		return _AMOY_CONTRACTS, nil
	default:
		return nil, fmt.Errorf("invalid chain id: %d, only BNB Chain (56), Polygon (137), and Amoy (80002) are supported", chainId)
	}
}
