package main

import "fmt"

func main() {

	//shorthand variable
	fees := 998.86

	//constant value rewrite = no
	const gstcode = 18

	//it's safe to use float cause one .00 calculation and system fucked totally,
	// later i'll use the calculation via paise and whenever needed to display the amount will calculate and show.
	gstamount := fees * float64(gstcode) / 100
	total := fees + gstamount

	fmt.Println("Freight Charge:", fees)
	fmt.Println("GST:", gstcode, "%")
	fmt.Println("GST Amount:", gstamount)
	fmt.Println("Total:", total)

	//used struct make payment feature
	PaymentGW := Googlepay{}
	newPayment := Payment{
		gateway: PaymentGW,
	}
	newPayment.makePayment(float32(total))

	//21 jan 2026, learned if & else and here is the try
	// == exact value
	// also i can declare a value directly in if construct
	paymentstatus := 1

	if paymentstatus == 1 {
		fmt.Println("Payment received")
		fmt.Println("Order Booked")

	} else {
		fmt.Println("Payment pending")
		fmt.Println("Order Booked, Payment pending")
	}

	OrderBooked := paymentstatus + gstcode

	//learned switch
	switch OrderBooked {
	default:
		fmt.Println("AWB Not Generated")

	case 19:
		fmt.Println("AWB Generated")
	}

	//learned maps & range results name, address, email-id, mobile-number
	m := make(map[string]string)

	m["od"] = "Order details"
	m2 := map[string]string{"Name:": "John Doe"}

	for k, v := range m2 {
		fmt.Println(m["od"])
		fmt.Println(k, v)
	}

	m3 := map[string]string{"Address:": "12/3, 345 street, City, State, Country"}

	for k, v := range m3 {
		fmt.Println(k, v)
	}

	m4 := map[string]int{"Zipcode:": 100001}
	for k, v := range m4 {
		fmt.Println(k, v)
	}

	m5 := map[string]string{"item:": "item name"}
	for k, v := range m5 {
		fmt.Println(k, v)
	}

	m6 := map[string]int{"item value:": 4000}
	for k, v := range m6 {
		fmt.Println(k, v)
	}

	m7 := map[string]int{"HSN code:": 123456}
	for k, v := range m7 {
		fmt.Println(k, v)
	}

}

// Learned to make struct and make it work with interfaces
//type razorpay struct{}

//func (r razorpay) pay(amount float32) {
//	fmt.Println("making payment using RazorPay", amount)}

type Googlepay struct{}

func (G Googlepay) pay(amount float32) {
	fmt.Println("make payment via Google Pay", amount)
}

type Payment struct {
	gateway paymenter
}

func (p Payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type paymenter interface {
	pay(amount float32)
}
