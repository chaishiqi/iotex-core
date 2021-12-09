// Copyright (c) 2019 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// usage: make minicluster

package main

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"math"
// 	"math/big"
// 	"math/rand"
// 	"os"
// 	"runtime"
// 	"sync"
// 	"time"

// 	"github.com/iotexproject/go-pkgs/cache/ttl"
// 	"github.com/iotexproject/go-pkgs/crypto"
// 	"github.com/iotexproject/iotex-proto/golang/iotexapi"
// 	"go.uber.org/zap"
// 	"google.golang.org/grpc"

// 	"github.com/iotexproject/iotex-core/blockchain"
// 	"github.com/iotexproject/iotex-core/config"
// 	"github.com/iotexproject/iotex-core/pkg/log"
// 	"github.com/iotexproject/iotex-core/pkg/probe"
// 	"github.com/iotexproject/iotex-core/pkg/unit"
// 	"github.com/iotexproject/iotex-core/pkg/util/fileutil"
// 	"github.com/iotexproject/iotex-core/server/itx"
// 	"github.com/iotexproject/iotex-core/state/factory"
// 	"github.com/iotexproject/iotex-core/testutil"
// 	"github.com/iotexproject/iotex-core/tools/executiontester/assetcontract"
// 	bc "github.com/iotexproject/iotex-core/tools/executiontester/blockchain"
// 	"github.com/iotexproject/iotex-core/tools/util"
// )

// const (
// 	numNodes  = 4
// 	numAdmins = 2
// )

// func main() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())

// 	// timeout indicates the duration of running nightly build in seconds. Default is 300
// 	var timeout int
// 	// aps indicates how many actions to be injected in one second. Default is 0
// 	var aps float64
// 	// smart contract deployment data. Default is "608060405234801561001057600080fd5b506102f5806100206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632885ad2c8114610066578063797d9fbd14610070578063cd5e3c5d14610091578063d0e30db0146100b8575b600080fd5b61006e6100c0565b005b61006e73ffffffffffffffffffffffffffffffffffffffff600435166100cb565b34801561009d57600080fd5b506100a6610159565b60408051918252519081900360200190f35b61006e610229565b6100c9336100cb565b565b60006100d5610159565b6040805182815290519192507fbae72e55df73720e0f671f4d20a331df0c0dc31092fda6c573f35ff7f37f283e919081900360200190a160405173ffffffffffffffffffffffffffffffffffffffff8316906305f5e100830280156108fc02916000818181858888f19350505050158015610154573d6000803e3d6000fd5b505050565b604080514460208083019190915260001943014082840152825180830384018152606090920192839052815160009360059361021a9360029391929182918401908083835b602083106101bd5780518252601f19909201916020918201910161019e565b51815160209384036101000a600019018019909216911617905260405191909301945091925050808303816000865af11580156101fe573d6000803e3d6000fd5b5050506040513d602081101561021357600080fd5b5051610261565b81151561022357fe5b06905090565b60408051348152905133917fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c919081900360200190a2565b600080805b60208110156102c25780600101602060ff160360080260020a848260208110151561028d57fe5b7f010000000000000000000000000000000000000000000000000000000000000091901a810204029190910190600101610266565b50929150505600a165627a7a72305820a426929891673b0a04d7163b60113d28e7d0f48ea667680ba48126c182b872c10029"
// 	var deployExecData string
// 	// smart contract interaction data. Default is "d0e30db0"
// 	var interactExecData string
// 	// switch of fp token smart contract test. Default is false
// 	var testFpToken bool

// 	flag.IntVar(&timeout, "timeout", 100, "duration of running nightly build")
// 	flag.Float64Var(&aps, "aps", 1500, "actions to be injected per second")
// 	flag.StringVar(&deployExecData, "deploy-data", "60806040526005600055600560015534801561001a57600080fd5b506102558061002a6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806358931c461461003b5780637f353d5514610045575b600080fd5b61004361004f565b005b61004d610097565b005b60006001905060005b6000548110156100935760028261006f9190610114565b915060028261007e91906100e3565b9150808061008b90610178565b915050610058565b5050565b60005b6001548110156100e057600281908060018154018082558091505060019003906000526020600020016000909190919091505580806100d890610178565b91505061009a565b50565b60006100ee8261016e565b91506100f98361016e565b925082610109576101086101f0565b5b828204905092915050565b600061011f8261016e565b915061012a8361016e565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610163576101626101c1565b5b828202905092915050565b6000819050919050565b60006101838261016e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156101b6576101b56101c1565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fdfea26469706673582212207db3d448bded0c08719540b308d7e6bab7d999edbac77e578522d01aad800e7064736f6c63430008070033",
// 		"smart contract deployment data")
// 	flag.StringVar(&interactExecData, "interact-data", "d0e30db0", "smart contract interaction data")
// 	flag.BoolVar(&testFpToken, "fp-token", false, "switch of fp token smart contract test")
// 	flag.Parse()

// 	// path of config file containing all the public/private key paris of addresses getting transfers
// 	// from Creator in genesis block
// 	injectorConfigPath := "./tools/minicluster/gentsfaddrs.yaml"

// 	chainAddrs, err := util.LoadAddresses(injectorConfigPath, uint32(1))
// 	if err != nil {
// 		log.L().Fatal("Failed to load addresses from config path", zap.Error(err))
// 	}
// 	admins := chainAddrs[len(chainAddrs)-numAdmins:]
// 	delegates := chainAddrs[:len(chainAddrs)-numAdmins]

// 	dbFilePaths := make([]string, 0)
// 	//a flag to indicate whether the DB files should be cleaned up upon completion of the minicluster.
// 	deleteDBFiles := false

// 	// Set mini-cluster configurations
// 	configs := make([]config.Config, numNodes)
// 	for i := 0; i < numNodes; i++ {
// 		chainDBPath := fmt.Sprintf("./chain%d.db", i+1)
// 		dbFilePaths = append(dbFilePaths, chainDBPath)
// 		trieDBPath := fmt.Sprintf("./trie%d.db", i+1)
// 		dbFilePaths = append(dbFilePaths, trieDBPath)
// 		indexDBPath := fmt.Sprintf("./index%d.db", i+1)
// 		dbFilePaths = append(dbFilePaths, indexDBPath)
// 		bloomfilterIndexDBPath := fmt.Sprintf("./bloomfilter.index%d.db", i+1)
// 		dbFilePaths = append(dbFilePaths, bloomfilterIndexDBPath)
// 		consensusDBPath := fmt.Sprintf("./consensus%d.db", i+1)
// 		dbFilePaths = append(dbFilePaths, consensusDBPath)
// 		systemLogDBPath := fmt.Sprintf("./systemlog%d.db", i+1)
// 		dbFilePaths = append(dbFilePaths, systemLogDBPath)
// 		candidateIndexDBPath := fmt.Sprintf("./candidate.index%d.db", i+1)
// 		dbFilePaths = append(dbFilePaths, candidateIndexDBPath)
// 		networkPort := config.Default.Network.Port + i
// 		apiPort := config.Default.API.Port + i
// 		web3APIPort := config.Default.API.Web3Port + i
// 		HTTPAdminPort := config.Default.System.HTTPAdminPort + i
// 		config := newConfig(chainAddrs[i].PriKey, networkPort, apiPort, web3APIPort, HTTPAdminPort)
// 		config.Chain.ChainDBPath = chainDBPath
// 		config.Chain.TrieDBPatchFile = ""
// 		config.Chain.TrieDBPath = trieDBPath
// 		config.Chain.IndexDBPath = indexDBPath
// 		config.Chain.BloomfilterIndexDBPath = bloomfilterIndexDBPath
// 		config.Chain.CandidateIndexDBPath = candidateIndexDBPath
// 		config.Consensus.RollDPoS.ConsensusDBPath = consensusDBPath
// 		config.System.SystemLogDBPath = systemLogDBPath
// 		if i == 0 {
// 			config.Network.BootstrapNodes = []string{}
// 			config.Network.MasterKey = "bootnode"
// 		}
// 		config.Genesis.AleutianBlockHeight = 1
// 		config.Genesis.PacificBlockHeight = 1
// 		configs[i] = config
// 	}

// 	// Create mini-cluster
// 	svrs := make([]*itx.Server, numNodes)
// 	for i := 0; i < numNodes; i++ {
// 		svr, err := itx.NewServer(configs[i])
// 		if err != nil {
// 			log.L().Fatal("Failed to create server.", zap.Error(err))
// 		}
// 		svrs[i] = svr
// 	}
// 	defer func() {
// 		if !deleteDBFiles {
// 			return
// 		}
// 		for _, dbFilePath := range dbFilePaths {
// 			if fileutil.FileExists(dbFilePath) && os.RemoveAll(dbFilePath) != nil {
// 				log.L().Error("Failed to delete db file")
// 			}
// 		}
// 	}()
// 	// Create a probe server
// 	probeSvr := probe.New(7788)

// 	// Start mini-cluster
// 	for i := 0; i < numNodes; i++ {
// 		ctx, cancel := context.WithCancel(context.Background())
// 		defer cancel()
// 		go itx.StartServer(ctx, svrs[i], probeSvr, configs[i])
// 	}

// 	// target address for grpc connection. Default is "127.0.0.1:14014"
// 	grpcAddr := "127.0.0.1:14014"

// 	grpcctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	conn, err := grpc.DialContext(grpcctx, grpcAddr, grpc.WithBlock(), grpc.WithInsecure())
// 	if err != nil {
// 		log.L().Error("Failed to connect to API server.")
// 	}
// 	defer conn.Close()

// 	client := iotexapi.NewAPIServiceClient(conn)

// 	counter, err := util.InitCounter(client, chainAddrs)
// 	if err != nil {
// 		log.L().Fatal("Failed to initialize nonce counter", zap.Error(err))
// 	}

// 	// Inject actions to first node
// 	if aps > 0 {
// 		// transfer gas limit. Default is 1000000
// 		transferGasLimit := 1000000
// 		// transfer gas price. Default is 10
// 		transferGasPrice := unit.Qev
// 		// transfer payload. Default is ""
// 		transferPayload := ""
// 		// vote gas limit. Default is 1000000
// 		voteGasLimit := 1000000
// 		// vote gas price. Default is 10
// 		voteGasPrice := unit.Qev
// 		// execution amount. Default is 0
// 		executionAmount := 0
// 		// execution gas limit. Default is 1200000
// 		executionGasLimit := 200000
// 		// execution gas price. Default is 10
// 		executionGasPrice := unit.Qev
// 		// maximum number of rpc retries. Default is 5
// 		retryNum := 5
// 		// sleeping period between two consecutive rpc retries in seconds. Default is 1
// 		retryInterval := 1
// 		// reset interval indicates the interval to reset nonce counter in seconds. Default is 60
// 		resetInterval := 60
// 		// fpTotal indicates the total amount value of a fp token
// 		fpTotal := int64(20000)
// 		// fpRisk indicates the risk amount value of a fp token
// 		fpRisk := int64(1000)

// 		d := time.Duration(timeout) * time.Second

// 		// First deploy a user specified smart contract which can be interacted by injected executions
// 		executor, nonce := util.CreateExecutionInjection(counter, delegates)
// 		contract, err := util.DeployContract(client, executor, nonce, executionGasLimit, executionGasPrice,
// 			deployExecData, retryNum, retryInterval)
// 		if err != nil {
// 			log.L().Fatal("Failed to deploy smart contract", zap.Error(err))
// 		}

// 		var fpToken bc.FpToken
// 		var fpContract string
// 		var debtor *util.AddressKey
// 		var creditor *util.AddressKey
// 		if testFpToken {
// 			// Deploy asset smart contracts
// 			ret, err := assetcontract.StartContracts(configs[0])
// 			if err != nil {
// 				log.L().Fatal("Failed to deploy asset contracts.", zap.Error(err))
// 			}
// 			fpToken = ret.FpToken
// 			// Randomly pick two accounts from delegate list as fp_token debtor and creditor
// 			first := rand.Intn(len(admins))
// 			second := first
// 			for second == first {
// 				second = rand.Intn(len(admins))
// 			}
// 			debtor = admins[first]
// 			creditor = admins[second]

// 			// Create fp token
// 			assetID := assetcontract.GenerateAssetID()
// 			open := time.Now().Unix()
// 			exp := open + 100000

// 			if _, err := fpToken.CreateToken(assetID, debtor.EncodedAddr, creditor.EncodedAddr, fpTotal, fpRisk, open,
// 				exp); err != nil {
// 				log.L().Fatal("Failed to create fp token", zap.Error(err))
// 			}

// 			fpContract, err = fpToken.TokenAddress(assetID)
// 			if err != nil {
// 				log.L().Fatal("Failed to get token contract address", zap.Error(err))
// 			}

// 			// Transfer full amount from debtor to creditor
// 			debtorPriKey := debtor.PriKey.HexString()
// 			if _, err := fpToken.Transfer(fpContract, debtor.EncodedAddr, debtorPriKey,
// 				creditor.EncodedAddr, fpTotal); err != nil {
// 				log.L().Fatal("Failed to transfer total amount from debtor to creditor", zap.Error(err))
// 			}

// 			// Transfer amount of risk from creditor to contract
// 			creditorPriKey := creditor.PriKey.HexString()
// 			if _, err := fpToken.RiskLock(fpContract, creditor.EncodedAddr, creditorPriKey,
// 				fpRisk); err != nil {
// 				log.L().Fatal("Failed to transfer amount of risk from creditor to contract", zap.Error(err))
// 			}
// 		}

// 		expectedBalancesMap := util.GetAllBalanceMap(client, chainAddrs)
// 		pendingActionMap, _ := ttl.NewCache(ttl.EvictOnErrorOption())

// 		log.L().Info("Start action injections.")

// 		wg := &sync.WaitGroup{}
// 		util.InjectByApsV2(wg, aps, counter, transferGasLimit, transferGasPrice, transferPayload, voteGasLimit,
// 			voteGasPrice, contract, executionAmount, executionGasLimit, executionGasPrice, interactExecData, fpToken,
// 			fpContract, debtor, creditor, client, admins, delegates, d, retryNum, retryInterval, resetInterval,
// 			expectedBalancesMap, svrs[0].ChainService(1), pendingActionMap)
// 		wg.Wait()

// 		err = testutil.WaitUntil(100*time.Millisecond, 60*time.Second, func() (bool, error) {
// 			empty, err := util.CheckPendingActionList(
// 				svrs[0].ChainService(1),
// 				pendingActionMap,
// 				expectedBalancesMap,
// 			)
// 			if err != nil {
// 				log.L().Error(err.Error())
// 				return false, err
// 			}
// 			return empty, nil
// 		})

// 		totalPendingActions := 0
// 		pendingActionMap.Range(func(selphash, vi interface{}) error {
// 			totalPendingActions++
// 			return nil
// 		})

// 		if err != nil {
// 			log.L().Error("Not all actions are settled")
// 		}

// 		chains := make([]blockchain.Blockchain, numNodes)
// 		sfs := make([]factory.Factory, numNodes)
// 		stateHeights := make([]uint64, numNodes)
// 		bcHeights := make([]uint64, numNodes)
// 		idealHeight := make([]uint64, numNodes)

// 		var netTimeout int
// 		var minTimeout int

// 		for i := 0; i < numNodes; i++ {
// 			chains[i] = svrs[i].ChainService(configs[i].Chain.ID).Blockchain()
// 			sfs[i] = svrs[i].ChainService(configs[i].Chain.ID).StateFactory()

// 			stateHeights[i], err = sfs[i].Height()
// 			if err != nil {
// 				log.S().Errorf("Node %d: Can not get State height", i)
// 			}
// 			bcHeights[i] = chains[i].TipHeight()
// 			minTimeout = int(configs[i].Consensus.RollDPoS.Delay/time.Second - configs[i].Genesis.BlockInterval/time.Second)
// 			netTimeout = 0
// 			if timeout > minTimeout {
// 				netTimeout = timeout - minTimeout
// 			}
// 			idealHeight[i] = uint64((time.Duration(netTimeout) * time.Second) / configs[i].Genesis.BlockInterval)

// 			log.S().Infof("Node#%d blockchain height: %d", i, bcHeights[i])
// 			log.S().Infof("Node#%d state      height: %d", i, stateHeights[i])
// 			log.S().Infof("Node#%d ideal      height: %d", i, idealHeight[i])

// 			if bcHeights[i] != stateHeights[i] {
// 				log.S().Errorf("Node#%d: State height does not match blockchain height", i)
// 			}
// 			if bcHeights[i] < idealHeight[i] {
// 				log.S().Errorf("blockchain in Node#%d is behind the expected height", i)
// 			}
// 		}

// 		for i := 0; i < numNodes; i++ {
// 			for j := i + 1; j < numNodes; j++ {
// 				if math.Abs(float64(bcHeights[i]-bcHeights[j])) > 1 {
// 					log.S().Errorf("blockchain in Node#%d and blockchain in Node#%d are not sync", i, j)
// 				} else {
// 					log.S().Infof("blockchain in Node#%d and blockchain in Node#%d are sync", i, j)
// 				}
// 			}
// 		}

// 		m := util.GetAllBalanceMap(client, chainAddrs)
// 		balanceCheckPass := true
// 		for k, v := range m {
// 			if len(expectedBalancesMap) != 0 && v.Cmp(expectedBalancesMap[k]) != 0 {
// 				balanceCheckPass = false
// 				log.S().Info("Balance mismatch on account ", k)
// 				log.S().Info("Real balance: ", v.String(), " Expected balance: ", expectedBalancesMap[k].String())

// 			}
// 		}
// 		if balanceCheckPass {
// 			log.S().Info("Balance Check PASS")
// 		} else {
// 			log.S().Fatal("Balance Mismatch")
// 		}

// 		log.S().Info("Total Transfer created: ", util.GetTotalTsfCreated())
// 		log.S().Info("Total Transfer inject through grpc: ", util.GetTotalTsfSentToAPI())
// 		log.S().Info("Total Transfer succeed: ", util.GetTotalTsfSucceeded())
// 		log.S().Info("Total Transfer failed: ", util.GetTotalTsfFailed())
// 		log.S().Info("Total pending actions: ", totalPendingActions)

// 		if testFpToken {
// 			// Check fp token asset balance
// 			debtorBalance, err := fpToken.ReadValue(fpContract, "70a08231", debtor.EncodedAddr)
// 			if err != nil {
// 				log.S().Error("Failed to get debtor's asset balance.", zap.Error(err))
// 			}
// 			log.L().Info("Debtor's asset balance: ", zap.Int64("balance", debtorBalance))

// 			creditorBalance, err := fpToken.ReadValue(fpContract, "70a08231", creditor.EncodedAddr)
// 			if err != nil {
// 				log.S().Error("Failed to get creditor's asset balance.", zap.Error(err))
// 			}
// 			log.L().Info("Creditor's asset balance: ", zap.Int64("balance", creditorBalance))

// 			if debtorBalance+creditorBalance != fpTotal-fpRisk {
// 				log.S().Error("Sum of asset balance is incorrect.")
// 				return
// 			}

// 			log.S().Info("Fp token transfer test pass!")
// 		}
// 		deleteDBFiles = true
// 	}
// }

// func newConfig(
// 	producerPriKey crypto.PrivateKey,
// 	networkPort,
// 	apiPort int,
// 	web3APIPort int,
// 	HTTPAdminPort int,
// ) config.Config {
// 	cfg := config.Default

// 	cfg.Plugins[config.GatewayPlugin] = true
// 	cfg.Chain.EnableAsyncIndexWrite = false

// 	cfg.System.HTTPAdminPort = HTTPAdminPort
// 	cfg.Network.Port = networkPort
// 	cfg.Network.BootstrapNodes = []string{"/ip4/127.0.0.1/tcp/4689/ipfs/12D3KooWJwW6pUpTkxPTMv84RPLPMQVEAjZ6fvJuX4oZrvW5DAGQ"}

// 	cfg.Chain.ID = 1
// 	cfg.Chain.CompressBlock = true
// 	cfg.Chain.ProducerPrivKey = producerPriKey.HexString()

// 	cfg.ActPool.MinGasPriceStr = big.NewInt(0).String()

// 	cfg.Consensus.Scheme = config.RollDPoSScheme
// 	cfg.Consensus.RollDPoS.FSM.UnmatchedEventInterval = 2400 * time.Millisecond
// 	cfg.Consensus.RollDPoS.FSM.AcceptBlockTTL = 1800 * time.Millisecond
// 	cfg.Consensus.RollDPoS.FSM.AcceptProposalEndorsementTTL = 1800 * time.Millisecond
// 	cfg.Consensus.RollDPoS.FSM.AcceptLockEndorsementTTL = 1800 * time.Millisecond
// 	cfg.Consensus.RollDPoS.FSM.CommitTTL = 600 * time.Millisecond
// 	cfg.Consensus.RollDPoS.FSM.EventChanSize = 100000
// 	cfg.Consensus.RollDPoS.ToleratedOvertime = 1200 * time.Millisecond
// 	cfg.Consensus.RollDPoS.Delay = 6 * time.Second

// 	cfg.API.Port = apiPort
// 	cfg.API.Web3Port = web3APIPort

// 	cfg.Genesis.BlockInterval = 6 * time.Second
// 	cfg.Genesis.Blockchain.NumSubEpochs = 2
// 	cfg.Genesis.Blockchain.NumDelegates = numNodes
// 	cfg.Genesis.Blockchain.TimeBasedRotation = true
// 	cfg.Genesis.Delegates = cfg.Genesis.Delegates[3 : numNodes+3]
// 	cfg.Genesis.EnableGravityChainVoting = false
// 	cfg.Genesis.PollMode = "lifeLong"
// 	// unlimit tx
// 	cfg.Genesis.BlockGasLimit *= 1
// 	cfg.Network.EnableRateLimit = false
// 	return cfg
// }