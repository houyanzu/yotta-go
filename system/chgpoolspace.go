package system

import (
	eos "github.com/houyanzu/yotta-go"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `eosio.system` contract.
func ChgPoolSpace(name eos.Name, isIncrease bool, deltaSpace uint64, account eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("hddpool12345"),
		Name:    ActN("chgpoolspace"),
		Authorization: []eos.PermissionLevel{
			{Actor: account, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(ChangePollSpace{
			Name:       name,
			IsIncrease: isIncrease,
			DeltaSpace: deltaSpace,
		}),
	}
}

// RegProducer represents the `eosio.system::regproducer` action
type ChangePollSpace struct {
	Name       eos.Name `json:"name"`
	IsIncrease bool     `json:"is_increase"`
	DeltaSpace uint64   `json:"delta_space"`
}
