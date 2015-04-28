package core_types

import (
	"github.com/tendermint/tendermint/account"
	sm "github.com/tendermint/tendermint/state"
	"github.com/tendermint/tendermint/types"
)

type ResponseGenPrivAccount struct {
	PrivAccount *account.PrivAccount
}

type ResponseGetAccount struct {
	Account *account.Account
}

type ResponseGetStorage struct {
	Key   []byte
	Value []byte
}

type ResponseCall struct {
	Return  []byte
	GasUsed uint64
	// TODO ...
}

type ResponseListAccounts struct {
	BlockHeight uint
	Accounts    []*account.Account
}

type StorageItem struct {
	Key   []byte
	Value []byte
}

type ResponseDumpStorage struct {
	StorageRoot  []byte
	StorageItems []StorageItem
}

type ResponseBlockchainInfo struct {
	LastHeight uint
	BlockMetas []*types.BlockMeta
}

type ResponseGetBlock struct {
	BlockMeta *types.BlockMeta
	Block     *types.Block
}

type ResponseBroadcastTx struct {
	Receipt Receipt
}

type ResponseListUnconfirmedTxs struct {
	Txs []types.Tx
}

type Receipt struct {
	TxHash          []byte
	CreatesContract uint8
	ContractAddr    []byte
}

type ResponseStatus struct {
	GenesisHash       []byte
	Network           string
	LatestBlockHash   []byte
	LatestBlockHeight uint
	LatestBlockTime   int64 // nano
}

type ResponseNetInfo struct {
	Moniker   string
	Network   string
	Listening bool
	Listeners []string
	Peers     []Peer
}

type Peer struct {
	types.NodeInfo
	IsOutbound bool
}

type ResponseSignTx struct {
	Tx types.Tx
}

type ResponseListValidators struct {
	BlockHeight         uint
	BondedValidators    []*sm.Validator
	UnbondingValidators []*sm.Validator
}

type ResponseDumpConsensusState struct {
	RoundState      string
	PeerRoundStates []string
}