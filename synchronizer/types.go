package main

import (
	"context"

	"github.com/ARR552/Hackaton_Circom/smc/hackaton"
	"github.com/ethereum/go-ethereum/common"
)

// EventOrder is the the type used to identify the events order
type EventOrder string

type DataEvent struct {
	Name        EventOrder
	Information string
	BlockNumber uint64
	Pos         int
}

// ClientEtherMan struct
type ClientEtherMan struct {
	EtherClient ethClienter
	hackaton    *hackaton.Hackaton
	SCAddresses []common.Address
}

type ClientSynchronizer struct {
	etherMan       *ClientEtherMan
	events         []DataEvent
	ctx            context.Context
	cancelCtx      context.CancelFunc
	genBlockNumber uint64
	cfg            Config
}
