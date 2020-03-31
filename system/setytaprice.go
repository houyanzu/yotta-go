package system

import (
	"github.com/houyanzu/yotta-go"
)

func SetYTAPrice(price uint64, accType uint8, authActor eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("hddpool12345"),
		Name:    ActN("setytaprice"),
		Authorization: []eos.PermissionLevel{
			{Actor: authActor, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(YTAPrice{
			Price:   price,
			AccType: accType,
		}),
	}
}

// RegProducer represents the `eosio.system::regproducer` action
type YTAPrice struct {
	Price   uint64 `json:"price"`
	AccType uint8  `json:"acc_type"`
}
