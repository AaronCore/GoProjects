package main

// 责任链模式
func main() {
	medical := &medical{}

	//Set next for cashier department
	cashier := &cashier{}
	cashier.setNext(medical)
	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(cashier)
	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
