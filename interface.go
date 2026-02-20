package main

import "fmt"

type payment struct {
	gateway stripe
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)

	// stripePaymentGW := stripe{}
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("Making payment using stripe: ",amount)
}

func interfaceExample() {
	stripePaymentGW := stripe{}
	newPayment := payment{
		gateway: stripePaymentGW,
	}
	newPayment.makePayment(100)
}