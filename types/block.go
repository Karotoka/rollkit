package types

type Header struct {
	// Block and App version
	Version Version
	// TODO this is redundant; understand if
	// anything like the ChainID in the Header is
	// required for IBC though.
	NamespaceID [8]byte

	Height uint64
	Time   uint64 // time in tai64 format

	// prev block info
	LastHeaderHash [32]byte

	// hashes of block data
	LastCommitHash [32]byte // commit from aggregator(s) from the last block
	DataHash       [32]byte // Block.Data root aka Transactions
	ConsensusHash  [32]byte // consensus params for current block
	AppHash        [32]byte // state after applying txs from the current block

	// root hash of all results from the txs from the previous block
	LastResultsHash [32]byte // TODO this is ABCI specific: do we really need it though?

	// TODO: do we need this to be included in the header?
	// the address can be derived from the pubkey which can be derived
	// from the signature when using secp256k.
	ProposerAddress []byte // original proposer of the block
}

// Version captures the consensus rules for processing a block in the blockchain,
// including all blockchain data structures and the rules of the application's
// state transition machine.
// This is equivalent to the tmversion.Consensus type in Tendermint.
type Version struct {
	Block uint32
	App   uint32
}

type Block struct {
	Header     Header
	Data       Data
	LastCommit *Commit
}

type Data struct {
	Txs                    Txs
	IntermediateStateRoots IntermediateStateRoots
	Evidence               EvidenceData
}

type EvidenceData struct {
	Evidence []Evidence
}

type Commit struct {
	Height     uint64
	HeaderHash [32]byte
	Signatures []Signature // most of the time this is a single signature
}

type Signature []byte

type IntermediateStateRoots struct {
	RawRootsList [][]byte
}
