package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ARR552/Hackaton_Circom/smc/hackaton"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.vocdoni.io/dvote/log"
)

var (
	dataEventSignatureHash = crypto.Keccak256Hash([]byte("Data(string)"))
)

type ethClienter interface {
	ethereum.ChainReader
	ethereum.LogFilterer
	ethereum.TransactionReader
}

// NewEtherman creates a new etherman
func NewEtherman(cfg Config) (*ClientEtherMan, error) {
	// TODO: PoEAddr can be got from bridge smc. Son only bridge smc is required
	// Connect to ethereum node
	ethClient, err := ethclient.Dial(cfg.URL)
	if err != nil {
		log.Errorf("error connecting to %s: %+v", cfg.URL, err)
		return nil, err
	}
	// Create smc clients
	hackaton, err := hackaton.NewHackaton(cfg.HackatonAddr, ethClient)
	if err != nil {
		return nil, err
	}
	var scAddresses []common.Address
	scAddresses = append(scAddresses, cfg.HackatonAddr)

	return &ClientEtherMan{EtherClient: ethClient, hackaton: hackaton, SCAddresses: scAddresses}, nil
}

func (etherMan *ClientEtherMan) getEventsByBlockRange(ctx context.Context, fromBlock uint64, toBlock *uint64) ([]DataEvent, error) {
	// First filter query
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		Addresses: etherMan.SCAddresses,
	}
	if toBlock != nil {
		query.ToBlock = new(big.Int).SetUint64(*toBlock)
	}
	dataEvents, err := etherMan.readEvents(ctx, query)
	if err != nil {
		return nil, err
	}
	return dataEvents, nil
}

func (etherMan *ClientEtherMan) readEvents(ctx context.Context, query ethereum.FilterQuery) ([]DataEvent, error) {
	logs, err := etherMan.EtherClient.FilterLogs(ctx, query)
	if err != nil {
		return []DataEvent{}, err
	}

	dataEvents := make([]DataEvent, 0)
	blockPos := make(map[common.Hash]int)

	for _, vLog := range logs {
		if vLog.Topics[0] == dataEventSignatureHash {
			log.Debug("Data event detected")
			data, err := etherMan.hackaton.ParseData(vLog)
			if err != nil {
				return []DataEvent{}, err
			}

			var dataEvent DataEvent
			dataEvent.Information = data.Information
			dataEvent.BlockNumber = vLog.BlockNumber

			if pos, exists := blockPos[vLog.BlockHash]; exists {
				dataEvent.Pos = pos
				blockPos[vLog.BlockHash] = pos + 1
			} else {
				dataEvent.Pos = 0
				blockPos[vLog.BlockHash] = 1
			}
			dataEvents = append(dataEvents, dataEvent)
		} else {
			log.Debug("Event not registered: ", vLog)
			return []DataEvent{}, fmt.Errorf("event not registered: %s", string(vLog.Data))
		}
	}
	return dataEvents, nil
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (etherMan *ClientEtherMan) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return etherMan.EtherClient.HeaderByNumber(ctx, number)
}

// Synchronizer interface
type Synchronizer interface {
	Sync() error
	Stop()
}

// NewSynchronizer creates and initializes an instance of Synchronizer
func NewSynchronizer(genBlockNumber uint64, cfg Config) (Synchronizer, error) {
	ctx, cancel := context.WithCancel(context.Background())
	etherMan, err := NewEtherman(cfg)
	if err != nil {
		return nil, err
	}
	return &ClientSynchronizer{
		etherMan:       etherMan,
		ctx:            ctx,
		cancelCtx:      cancel,
		genBlockNumber: genBlockNumber,
		cfg:            cfg,
	}, nil
}

// Sync function will read the last state synced and will continue from that point.
// Sync() will read blockchain events to detect bridge updates
func (s *ClientSynchronizer) Sync() error {
	go func() {
		// If there is no lastEthereumBlock means that sync from the beginning is necessary. If not, it continues from the retrieved ethereum block
		// Get the latest synced block. If there is no block on db, use genesis block
		log.Info("Sync started")
		waitDuration := time.Duration(0)
		lastSyncedBlockNumber := s.genBlockNumber
		var err error
		for {
			select {
			case <-s.ctx.Done():
				return
			case <-time.After(waitDuration):
				if lastSyncedBlockNumber, err = s.syncBlocks(lastSyncedBlockNumber); err != nil {
					if s.ctx.Err() != nil {
						continue
					}
				}
				if waitDuration != s.cfg.SyncInterval.Duration {
					// Check latest Ethereum Block
					header, err := s.etherMan.HeaderByNumber(s.ctx, nil)
					if err != nil {
						log.Warn("error getting latest block from Ethereum. Error: ", err)
						continue
					}
					lastKnownBlock := header.Number
					if lastSyncedBlockNumber == lastKnownBlock.Uint64() {
						waitDuration = s.cfg.SyncInterval.Duration
					}
				}
			}
		}
	}()
	return nil
}

// Stop function stops the synchronizer
func (s *ClientSynchronizer) Stop() {
	s.cancelCtx()
}

// This function syncs the node from a specific block to the latest
func (s *ClientSynchronizer) syncBlocks(lastBlockNumber uint64) (uint64, error) {
	// Call the blockchain to retrieve data
	var fromBlock, lastSyncedBlockNumber uint64
	if lastBlockNumber > 0 {
		fromBlock = lastBlockNumber + 1
	}

	header, err := s.etherMan.HeaderByNumber(s.ctx, nil)
	if err != nil {
		return 0, err
	}
	lastKnownBlock := header.Number

	for {
		toBlock := fromBlock + s.cfg.SyncChunkSize

		log.Debugf("Getting bridge info from block %d to block %d", fromBlock, toBlock)
		events, err := s.etherMan.getEventsByBlockRange(s.ctx, fromBlock, &toBlock)
		s.events = append(s.events, events...)
		if err != nil {
			return 0, err
		}
		lastSyncedBlockNumber = toBlock
		fromBlock = toBlock + 1

		if lastKnownBlock.Cmp(new(big.Int).SetUint64(fromBlock)) < 1 {
			break
		}
	}

	return lastSyncedBlockNumber, nil
}
