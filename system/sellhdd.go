package system

import (
	eos "github.com/houyanzu/yotta-go"
)

func NewSellHdd(account eos.AccountName, value uint64, memo string) *eos.Action {
	return &eos.Action{
		Account: AN("hddpool12345"),
		Name:    ActN("sellhdd"),
		Authorization: []eos.PermissionLevel{
			{Actor: account, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(SellHdd{
			User:   account,
			Amount: value,
			Memo:   memo,
		}),
	}
}

type SellHdd struct {
	User   eos.AccountName `json:"user"`
	Amount uint64          `json:"amount"`
	Memo   string          `json:"memo"`
}
