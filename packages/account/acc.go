package account

type UserAccount struct{
	Name string
	Age int
	balance float32
}

func (u UserAccount) CreateAccount(){
	if(u.balance < 0){
		u.balance = 0
	}
}

func checkAmount(amount float32) bool{
	if amount<0 {
		print("Please Enter correct amount!!")
		return false
	}

	return true
}

func (u UserAccount) CheckBalance() float32{
	return u.balance
}

func (u UserAccount) AddAmoubt(amount float32){
	if(checkAmount(amount)){
		u.balance += amount
	}
}

func (u UserAccount) SubAmount(amount float32){
	if(checkAmount(amount)){
		u.balance-= amount
	}
}


