/**
*  @file
*  @copyright defined in go-scdo/LICENSE
 */

package factory

import (
	"crypto/ecdsa"
	"fmt"
	"path/filepath"

	"github.com/scdoproject/go-scdo/common"
	"github.com/scdoproject/go-scdo/common/errors"
	"github.com/scdoproject/go-scdo/consensus"
	"github.com/scdoproject/go-scdo/consensus/bft"
	"github.com/scdoproject/go-scdo/consensus/bft/server"
	"github.com/scdoproject/go-scdo/consensus/ethash"
	"github.com/scdoproject/go-scdo/consensus/istanbul"
	"github.com/scdoproject/go-scdo/consensus/istanbul/backend"
	"github.com/scdoproject/go-scdo/consensus/pow"
	"github.com/scdoproject/go-scdo/consensus/spow"
	"github.com/scdoproject/go-scdo/database/leveldb"
)

// GetConsensusEngine get consensus engine according to miner algorithm name
// WARNING: engine may be a heavy instance. we should have as less as possible in our process.
func GetConsensusEngine(minerAlgorithm string, folder string) (consensus.Engine, error) {
	var minerEngine consensus.Engine
	if minerAlgorithm == common.EthashAlgorithm {
		minerEngine = ethash.New(ethash.GetDefaultConfig(), nil, false)
	} else if minerAlgorithm == common.Sha256Algorithm {
		minerEngine = pow.NewEngine(1)
	} else if minerAlgorithm == common.SpowAlgorithm {
		minerEngine = spow.NewSpowEngine(1, folder)
	} else {
		return nil, fmt.Errorf("unknown miner algorithm")
	}

	return minerEngine, nil
}

func GetBFTEngine(privateKey *ecdsa.PrivateKey, folder string) (consensus.Engine, error) {
	path := filepath.Join(folder, common.BFTDataFolder)
	db, err := leveldb.NewLevelDB(path)
	if err != nil {
		return nil, errors.NewStackedError(err, "create bft folder failed")
	}

	return backend.New(istanbul.DefaultConfig, privateKey, db), nil
}

func MustGetConsensusEngine(minerAlgorithm string) consensus.Engine {
	engine, err := GetConsensusEngine(minerAlgorithm, "temp")
	if err != nil {
		panic(err)
	}

	return engine
}

// subchain bft engine engine
// here need to input the privatekey
// TODO: not use privateKey as an input
func GetBFTSubchainEngine(privateKey *ecdsa.PrivateKey, folder string) (consensus.Engine, error) {
	path := filepath.Join(folder, common.BFTDataFolder)
	db, err := leveldb.NewLevelDB(path)
	if err != nil {
		return nil, errors.NewStackedError(err, "create bft folder failed")
	}

	return server.NewServer(bft.DefaultConfig, privateKey, db), nil
}
