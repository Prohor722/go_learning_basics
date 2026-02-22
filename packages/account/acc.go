package account

var balance float32 = 0.00

func checkAmount(amount float32) bool{
	if amount<0 {
		print("Please Enter correct amount!!")
		return false
	}

	return true
}

func CheckBalance() float32{
	return balance
}

func AddAmoubt(amount float32){
	if(checkAmount(amount)){
		balance += amount
	}
}

func SubAmount(amount float32){
	if(checkAmount(amount)){
		balance-= amount
	}
}


