package utils

import "fmt"

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	flag    bool
	details string
}

//编写工厂模块的构造方法
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("\n-----------------------当前收支明细记录-----------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细...")
	}

	fmt.Println()
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额: ")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明: ")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t             %v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额: ")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足...")
		fmt.Println("")
	} else {
		this.balance -= this.money
		fmt.Println("本次支出说明: ")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n收入\t%v\t             %v      \t%v", this.balance, this.money, this.note)
		this.flag = true
	}
}

func (this *FamilyAccount) quit() {
	fmt.Println("确定退出系统? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("输入错误，请重新输入...")
	}
	if choice == "y" {
		this.loop = false
	}
}

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("-----------------------家庭收支记账软件-----------------------")
		fmt.Println("                        1 收支明细")
		fmt.Println("                        2 登记收入")
		fmt.Println("                        3 登记支出")
		fmt.Println("                        4 退出软件")
		fmt.Println("请选择(1-4): ")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.quit()
		default:
			fmt.Println("请输入正确的选项...")
		}

		if !this.loop {
			break
		}
	}
}
