// Copyright 2016 The go-ethereum Authors
// (original work)
// Copyright 2024 The Erigon Authors
// (modifications)
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package params

import (
    "embed"
    "encoding/json"
    "fmt"
    "math/big"
	"path"
    "github.com/erigontech/erigon-lib/chain"
    "github.com/erigontech/erigon-lib/chain/networkname"
    libcommon "github.com/erigontech/erigon-lib/common"
    "github.com/erigontech/erigon/polygon/bor/borcfg"
	"github.com/erigontech/erigon-lib/common/paths"
)

//go:embed chainspecs
var chainspecs embed.FS

func readChainSpec(filename string) *chain.Config {
        f, err := chainspecs.Open(filename)
        if err != nil {
                panic(fmt.Sprintf("Could not open chainspec for %s: %v", filename, err))
        }
        defer f.Close()
        decoder := json.NewDecoder(f)
        spec := &chain.Config{}
        err = decoder.Decode(&spec)
        if err != nil {
                panic(fmt.Sprintf("Could not parse chainspec for %s: %v", filename, err))
        }
        if spec.BorJSON != nil {
                borConfig := &borcfg.BorConfig{}
                err = json.Unmarshal(spec.BorJSON, borConfig)
                if err != nil {
                        panic(fmt.Sprintf("Could not parse 'bor' chainspec for %s: %v", filename, err))
                }
                spec.Bor = borConfig
        }
        return spec
}

// Genesis hashes to enforce below configs on.
var (
     MainnetGenesisHash    = libcommon.HexToHash("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")
        HoleskyGenesisHash    = libcommon.HexToHash("0xb5f7f912443c940f21fd611f12828d75b534364ed9e95ca4e307729a4661bde4")
        SepoliaGenesisHash    = libcommon.HexToHash("0x25a5cc106eea7138acab33231d7160d69cb777ee0c2c553fcddf5138993e6dd9")
        HoodiGenesisHash      = libcommon.HexToHash("0xbbe312868b376a3001692a646dd2d7d1e4406380dfd86b98aa8a34d1557c971b")
        AmoyGenesisHash       = libcommon.HexToHash("0x7202b2b53c5a0836e773e319d18922cc756dd67432f9a1f65352b61f4406c697")
        BorMainnetGenesisHash = libcommon.HexToHash("0xa9c28ce2141b56c474f1dc504bee9b01eb1bd7d1a507580d5519d4437a97de1b")
        BorDevnetGenesisHash  = libcommon.HexToHash("0x5a06b25b0c6530708ea0b98a3409290e39dce6be7f558493aeb6e4b99a172a87")
        GnosisGenesisHash     = libcommon.HexToHash("0x4f1dd23188aab3a76b463e4af801b52b1248ef073c648cbdc4c9333d3da79756")
        ChiadoGenesisHash     = libcommon.HexToHash("0xada44fd8d2ecab8b08f256af07ad3e777f17fb434f8f8e678b312f576212ba9a")
        TestGenesisHash       = libcommon.HexToHash("0x6116de25352c93149542e950162c7305f207bbc17b0eb725136b78c80aed79cc")
    Olym3TestnetS3GenesisHash = libcommon.HexToHash("0xdbabbed39334fd59b612c675a555a9fbb13f9b1f1f56c15c4831603f09c9318b")
)

var (
    	GnosisGenesisStateRoot = libcommon.HexToHash("0x40cf4430ecaa733787d1a65154a3b9efb560c95d9e324a23b97f0609b539133b")
        ChiadoGenesisStateRoot = libcommon.HexToHash("0x9ec3eaf4e6188dfbdd6ade76eaa88289b57c63c9a2cde8d35291d5a29e143d31")
        TestGenesisStateRoot   = libcommon.HexToHash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
)

var (
    MainnetChainConfig       = readChainSpec("chainspecs/mainnet.json")
    HoleskyChainConfig       = readChainSpec("chainspecs/holesky.json")
    SepoliaChainConfig       = readChainSpec("chainspecs/sepolia.json")
    HoodiChainConfig         = readChainSpec("chainspecs/hoodi.json")
    AmoyChainConfig          = readChainSpec("chainspecs/amoy.json")
    BorMainnetChainConfig    = readChainSpec("chainspecs/bor-mainnet.json")
    BorDevnetChainConfig     = readChainSpec("chainspecs/bor-devnet.json")
    GnosisChainConfig        = readChainSpec("chainspecs/gnosis.json")
    ChiadoChainConfig        = readChainSpec("chainspecs/chiado.json")
    Olym3TestnetS3ChainConfig = readChainSpec("chainspecs/olym3-testnet-s3.json")
    // AllProtocolChanges contains every protocol change (EIPs) introduced
        // and accepted by the Ethereum core developers into the main net protocol.
        AllProtocolChanges = &chain.Config{
                ChainID:                       big.NewInt(1337),
                Consensus:                     chain.EtHashConsensus,
                HomesteadBlock:                big.NewInt(0),
                TangerineWhistleBlock:         big.NewInt(0),
                SpuriousDragonBlock:           big.NewInt(0),
                ByzantiumBlock:                big.NewInt(0),
                ConstantinopleBlock:           big.NewInt(0),
                PetersburgBlock:               big.NewInt(0),
                IstanbulBlock:                 big.NewInt(0),
                MuirGlacierBlock:              big.NewInt(0),
                BerlinBlock:                   big.NewInt(0),
                LondonBlock:                   big.NewInt(0),
                ArrowGlacierBlock:             big.NewInt(0),
                GrayGlacierBlock:              big.NewInt(0),
                TerminalTotalDifficulty:       big.NewInt(0),
                TerminalTotalDifficultyPassed: true,
                ShanghaiTime:                  big.NewInt(0),
                CancunTime:                    big.NewInt(0),
                PragueTime:                    big.NewInt(0),
                Ethash:                        new(chain.EthashConfig),
        }
        // AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
        // and accepted by the Ethereum core developers into the Clique consensus.
        AllCliqueProtocolChanges = &chain.Config{
                ChainID:               big.NewInt(1337),
                Consensus:             chain.CliqueConsensus,
                HomesteadBlock:        big.NewInt(0),
                TangerineWhistleBlock: big.NewInt(0),
                SpuriousDragonBlock:   big.NewInt(0),
                ByzantiumBlock:        big.NewInt(0),
                ConstantinopleBlock:   big.NewInt(0),
                PetersburgBlock:       big.NewInt(0),
                IstanbulBlock:         big.NewInt(0),
                MuirGlacierBlock:      big.NewInt(0),
                BerlinBlock:           big.NewInt(0),
                LondonBlock:           big.NewInt(0),
                Clique:                &chain.CliqueConfig{Period: 0, Epoch: 30000},
        }
        CliqueSnapshot = NewSnapshotConfig(10, 1024, 16384, true, "")
        TestChainConfig = &chain.Config{
                ChainID:               big.NewInt(1337),
                Consensus:             chain.EtHashConsensus,
                HomesteadBlock:        big.NewInt(0),
                TangerineWhistleBlock: big.NewInt(0),
                SpuriousDragonBlock:   big.NewInt(0),
                ByzantiumBlock:        big.NewInt(0),
                ConstantinopleBlock:   big.NewInt(0),
                PetersburgBlock:       big.NewInt(0),
                IstanbulBlock:         big.NewInt(0),
                MuirGlacierBlock:      big.NewInt(0),
                BerlinBlock:           big.NewInt(0),
                Ethash:                new(chain.EthashConfig),
        }
        TestChainAuraConfig = &chain.Config{
                ChainID:               big.NewInt(1),
                Consensus:             chain.AuRaConsensus,
                HomesteadBlock:        big.NewInt(0),
                TangerineWhistleBlock: big.NewInt(0),
                SpuriousDragonBlock:   big.NewInt(0),
                ByzantiumBlock:        big.NewInt(0),
                ConstantinopleBlock:   big.NewInt(0),
                PetersburgBlock:       big.NewInt(0),
                IstanbulBlock:         big.NewInt(0),
                MuirGlacierBlock:      big.NewInt(0),
                BerlinBlock:           big.NewInt(0),
                LondonBlock:           big.NewInt(0),
                Aura:                  &chain.AuRaConfig{},
        }
        TestRules = TestChainConfig.Rules(0, 0)
)

func ChainConfigByChainName(chain string) *chain.Config {
    switch chain {
    case networkname.Mainnet:
        return MainnetChainConfig
    case networkname.Holesky:
        return HoleskyChainConfig
    case networkname.Sepolia:
        return SepoliaChainConfig
    case networkname.Hoodi:
        return HoodiChainConfig
    case networkname.Amoy:
        return AmoyChainConfig
    case networkname.BorMainnet:
        return BorMainnetChainConfig
    case networkname.BorDevnet:
        return BorDevnetChainConfig
    case networkname.Gnosis:
        return GnosisChainConfig
    case networkname.Chiado:
        return ChiadoChainConfig
    case networkname.Test:
        return TestChainConfig
    case networkname.Olym3TestnetS3:
        return Olym3TestnetS3ChainConfig
    default:
        return nil
    }
}
type ConsensusSnapshotConfig struct {
        CheckpointInterval uint64 // Number of blocks after which to save the vote snapshot to the database
        InmemorySnapshots  int    // Number of recent vote snapshots to keep in memory
        InmemorySignatures int    // Number of recent block signatures to keep in memory
        DBPath             string
        InMemory           bool
}
const cliquePath = "clique"
func NewSnapshotConfig(checkpointInterval uint64, inmemorySnapshots int, inmemorySignatures int, inmemory bool, dbPath string) *ConsensusSnapshotConfig {
        if len(dbPath) == 0 {
                dbPath = paths.DefaultDataDir()
        }
        return &ConsensusSnapshotConfig{
                checkpointInterval,
                inmemorySnapshots,
                inmemorySignatures,
                path.Join(dbPath, cliquePath),
                inmemory,
        }
}

func ChainConfigByGenesisHash(genesisHash libcommon.Hash) *chain.Config {
    switch {
    case genesisHash == MainnetGenesisHash:
        return MainnetChainConfig
    case genesisHash == HoleskyGenesisHash:
        return HoleskyChainConfig
    case genesisHash == SepoliaGenesisHash:
        return SepoliaChainConfig
    case genesisHash == HoodiGenesisHash:
        return HoodiChainConfig
    case genesisHash == Olym3TestnetS3GenesisHash:
        return Olym3TestnetS3ChainConfig
    default:
        return nil
    }
}


func GenesisHashByChainName(chain string) *libcommon.Hash {
    switch chain {
    case networkname.Mainnet:
        return &MainnetGenesisHash
    case networkname.Holesky:
        return &HoleskyGenesisHash
    case networkname.Sepolia:
        return &SepoliaGenesisHash
    case networkname.Hoodi:
        return &HoodiGenesisHash
    case networkname.Amoy:
        return &AmoyGenesisHash
    case networkname.BorMainnet:
        return &BorMainnetGenesisHash
    case networkname.BorDevnet:
        return &BorDevnetGenesisHash
    case networkname.Gnosis:
        return &GnosisGenesisHash
    case networkname.Chiado:
        return &ChiadoGenesisHash
    case networkname.Test:
        return &TestGenesisHash
    case networkname.Olym3TestnetS3:
        return &Olym3TestnetS3GenesisHash
    default:
        return nil
    }
}
func NetworkIDByChainName(chain string) uint64 {
        config := ChainConfigByChainName(chain)
        if config == nil {
                return 0
        }
        return config.ChainID.Uint64()
}
func IsChainPoS(chainConfig *chain.Config, currentTDProvider func() *big.Int) bool {
        return isChainIDPoS(chainConfig.ChainID) || hasChainPassedTerminalTD(chainConfig, currentTDProvider)
}
func isChainIDPoS(chainID *big.Int) bool {
        ids := []*big.Int{
                MainnetChainConfig.ChainID,
                HoleskyChainConfig.ChainID,
                SepoliaChainConfig.ChainID,
                HoodiChainConfig.ChainID,
                GnosisChainConfig.ChainID,
                ChiadoChainConfig.ChainID,
		Olym3TestnetS3ChainConfig.ChainID,
        }
        for _, id := range ids {
                if id.Cmp(chainID) == 0 {
                        return true
                }
        }
        return false
}
func hasChainPassedTerminalTD(chainConfig *chain.Config, currentTDProvider func() *big.Int) bool {
        if chainConfig.TerminalTotalDifficultyPassed {
                return true
        }
        terminalTD := chainConfig.TerminalTotalDifficulty
        if terminalTD == nil {
                return false
        }
        currentTD := currentTDProvider()
        return (currentTD != nil) && (terminalTD.Cmp(currentTD) <= 0)
}








