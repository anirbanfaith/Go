package main

import "fmt"

func main() {

	//shorthand variable
	fees := 798.86

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

}
