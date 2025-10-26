package problems

import (
	"testing"
)

type BankOp struct {
	cmd      string
	account1 int
	account2 int
	money    int64
	want     bool
}

func TestBank(t *testing.T) {
	tests := []struct {
		name    string
		balance []int64
		ops     []BankOp
	}{
		{
			name:    "Example 1",
			balance: []int64{10, 100, 20, 50, 30},
			ops: []BankOp{
				{cmd: "withdraw", account1: 3, money: 10, want: true},
				{cmd: "transfer", account1: 5, account2: 1, money: 20, want: true},
				{cmd: "deposit", account1: 5, money: 20, want: true},
				{cmd: "transfer", account1: 3, account2: 4, money: 15, want: false},
				{cmd: "withdraw", account1: 10, money: 50, want: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := BankConstructor(tt.balance)
			for _, op := range tt.ops {
				var got bool
				switch op.cmd {
				case "transfer":
					got = bank.Transfer(op.account1, op.account2, op.money)
				case "withdraw":
					got = bank.Withdraw(op.account1, op.money)
				case "deposit":
					got = bank.Deposit(op.account1, op.money)
				}
				if got != op.want {
					t.Errorf("%s() = %v, want %v", op.cmd, got, op.want)
				}
			}
		})
	}
}
