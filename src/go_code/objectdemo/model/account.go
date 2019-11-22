package model

import "fmt"

type account struct {
	accountNo string
	pwd       string
	balance   float64
}

//工厂模式的函数-构造函数
func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度并不对...")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码的长度并不对...")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额数目不对...")
		return nil
	}
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}

func (account *account) Deposite(money float64, pwd string) {
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
}
