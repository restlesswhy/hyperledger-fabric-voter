package models

type ThreadParams struct {
	ID          string   `json:"-" db:"thread_id"`
	Category    string   `json:"category"`
	Theme       string   `json:"theme"`
	Description string   `json:"description"`
	Options     []string `json:"options"`
}

type Thread struct {
	Category    string              `json:"category"`
	Theme       string              `json:"theme"`
	Description string              `json:"description"`
	Creator     string              `json:"creator"`
	Options     map[string][]string `json:"options"`
	WinOption   []string            `json:"win_option"`
	Status      string              `json:"status"`
}

type Vote struct {
	ThreadID string `json:"thread_id"`
	VoteID   string `json:"vote_id"`
	Option   string `json:"option"`
}

type AnonThread struct {
	Category    string              `json:"category"`
	Theme       string              `json:"theme"`
	Description string              `json:"description"`
	Creator     string              `json:"creator"`
	Votes       []string            `json:"votes"`
	Options     map[string][]string `json:"options"`
	WinOption   []string            `json:"win_option"`
	Status      string              `json:"status"`
}

type AnonVote struct {
	ThreadID   string `json:"thread_id"`
	TxID       string `json:"tx_id"`
	Option     string `json:"option"`
	PrivateKey string `json:"private_key"`
}

type EndAnonData struct {
	ThreadID string   `json:"thread_id"`
	Keys     []string `json:"keys"`
	VoteTxs  []string `json:"vote_txs"`
}