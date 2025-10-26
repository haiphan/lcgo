package problems

type Bank struct {
	balance []int64
}

func BankConstructor(balance []int64) Bank {
	return Bank{balance: balance}
}

func (bank *Bank) ValidateAcc(acc int) bool {
	return acc <= len(bank.balance)
}

func (bank *Bank) HasEnough(acc int, m int64) bool {
	return m <= bank.balance[acc-1]
}

func (bank *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if bank.ValidateAcc(account1) && bank.ValidateAcc(account2) && bank.HasEnough(account1, money) {
		bank.balance[account1-1] -= money
		bank.balance[account2-1] += money
		return true
	}
	return false
}

func (bank *Bank) Deposit(account int, money int64) bool {
	if bank.ValidateAcc(account) {
		bank.balance[account-1] += money
		return true
	}
	return false
}

func (bank *Bank) Withdraw(account int, money int64) bool {
	if bank.ValidateAcc(account) && bank.HasEnough(account, money) {
		bank.balance[account-1] -= money
		return true
	}
	return false
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
