package common

type Account struct {
	Account string           `json:"account"`
	Passwd  string           `json:"Passwd"`
	Level   uint32           `json:"Level"`
	Score   uint64           `json:"Score"`
	Props   map[string]int64 `json:"Props"`
}
