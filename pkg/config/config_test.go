package config

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestGetContracts(t *testing.T) {
	var (
		bnbChain = &Contracts{
			Exchange:         common.HexToAddress("0xF99F5367ce708c66F0860B77B4331301A5597c86"),
			FeeModule:        common.Address{}, // Not used by Probable Markets
			NegRiskExchange:  common.Address{}, // Not used by Probable Markets
			NegRiskFeeModule: common.Address{}, // Not used by Probable Markets
			NegRiskAdapter:   common.Address{}, // Not used by Probable Markets
			Collateral:       common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
			Conditional:      common.HexToAddress("0x364d05055614B506e2b9A287E4ac34167204cA83"),
		}
	)

	c, err := GetContracts(56)
	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(c.Exchange[:], bnbChain.Exchange[:]))
	assert.True(t, bytes.Equal(c.FeeModule[:], bnbChain.FeeModule[:]))
	assert.True(t, bytes.Equal(c.NegRiskExchange[:], bnbChain.NegRiskExchange[:]))
	assert.True(t, bytes.Equal(c.NegRiskFeeModule[:], bnbChain.NegRiskFeeModule[:]))
	assert.True(t, bytes.Equal(c.NegRiskAdapter[:], bnbChain.NegRiskAdapter[:]))
	assert.True(t, bytes.Equal(c.Collateral[:], bnbChain.Collateral[:]))
	assert.True(t, bytes.Equal(c.Conditional[:], bnbChain.Conditional[:]))

	// Invalid chain ID
	c, err = GetContracts(100000)
	assert.Nil(t, c)
	assert.NotNil(t, err)
}
