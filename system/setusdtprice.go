package system

import (
	"github.com/houyanzu/yotta-go"
)

func SetUSDTPrice(price uint64, accType uint8, authActor eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("hddpool12345"),
		Name:    ActN("setusdprice"),
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
type USDTPrice struct {
	Price   uint64 `json:"price"`
	AccType uint8  `json:"acc_type"`
}
