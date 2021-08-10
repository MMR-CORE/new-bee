package config

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	// chain ID
	goerliChainID = int64(5)
	xdaiChainID   = int64(100)

	// start block
	goerliStartBlock = uint64(5208103)
	xdaiStartBlock   = uint64(17511766)

	// price oracle address
	goerliPriceOracleContractAddress = common.HexToAddress("0xD871739e523849C15773649C1b617Df1cDa302C1")
	xdaiPriceOracleContractAddress = common.HexToAddress("0xAB4Ff9942EA3b6EcbBAD84f3B0B1f2FCDB739674")

	// token address
	goerliTokenFBZZ = "0x5E4B0229565643fD9f5613aBf4B1498598033445"
	xdaiTokenXFBZZ = "0x36F8096d0308C786F56f7cFE90ff2F120367197b"

	// postage stamp
	goerliPostageStampContractAddress = common.HexToAddress("0x208689548f803150021142f21F2428551A060AA5")
	xdaiPostageStampContractAddress   = common.HexToAddress(" 0x8106747bFdBed1c45691000BaE42dFE5a1F5692f")

	// factory address
	goerliFactoryAddress             = common.HexToAddress("0xcD87A69Ebd39FB48E8Ca8428026c47d9F3b2BDa5")
	xdaiFactoryAddress  = common.HexToAddress("0x830b0711C3F052C1cEd24abfD27BFbB86AcB51Bf")
)

type ChainConfig struct {
	StartBlock         uint64
	LegacyFactories    []common.Address
	PostageStamp       common.Address
	CurrentFactory     common.Address
	PriceOracleAddress common.Address
}

func GetChainConfig(chainID int64) (*ChainConfig, bool) {
	var cfg ChainConfig
	switch chainID {
	case goerliChainID:
		cfg.PostageStamp = goerliPostageStampContractAddress
		cfg.StartBlock = goerliStartBlock
		cfg.CurrentFactory = goerliFactoryAddress
		cfg.LegacyFactories = []common.Address{}
		cfg.PriceOracleAddress = goerliPriceOracleContractAddress
		return &cfg, true
	case xdaiChainID:
		cfg.PostageStamp = xdaiPostageStampContractAddress
		cfg.StartBlock = xdaiStartBlock
		cfg.CurrentFactory = xdaiFactoryAddress
		cfg.LegacyFactories = []common.Address{}
		cfg.PriceOracleAddress = xdaiPriceOracleContractAddress
		return &cfg, true
	default:
		return &cfg, false
	}
}
