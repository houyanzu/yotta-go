package system

import (
	eos "github.com/houyanzu/yotta-go"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `eosio.system` contract.
func NewSeqProducer(producer eos.AccountName, shadow eos.AccountName, seq uint16, level uint8, authActor eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("seqproducer"),
		Authorization: []eos.PermissionLevel{
			{Actor: authActor, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(SeqProducer{
			Producer: producer,
			Shadow:   shadow,
			Seq:      seq,
			Level:    level,
		}),
	}
}

// RegProducer represents the `eosio.system::regproducer` action
type SeqProducer struct {
	Producer eos.AccountName `json:"producer"`
	Shadow   eos.AccountName `json:"shadow"`
	Seq      uint16          `json:"seq"`
	Level    uint8           `json:"level"`
}
