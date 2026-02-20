package main

import "fmt"

type payment struct {
	gateway paymenter
}

type paymenter interface{
	pay(amount float32)
	refund(amount float32)
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

type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("Making payment using paypal:",amount)
}

func (p paypal) refund(amount float32){
	fmt.Println("Refund done from paypal amount:",amount)
}

func interfaceExample() {
	// stripePaymentGW := stripe{}
	// razorPaymentGW := razorpay{}
	paypalPaymentGW := paypal{}

	newPayment := payment{
		// gateway: stripePaymentGW,
		// gateway: razorPaymentGW,
		gateway: paypalPaymentGW,
	}
	newPayment.makePayment(100)
	newPayment.gateway.refund(100)
}