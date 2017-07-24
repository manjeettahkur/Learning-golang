package fund

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int

}


// A regular function returning a pointer to a fund

func NewFund(initialBalance int) *Fund {
	// We can return a pointer to a new struct without worrying ab
	//whether it's on the stack or heap: Go figures that out for
	return &Fund{
		balance:initialBalance,
	}
}

//Method start with a *receiver*, in this case a fund pointer
func (f *Fund) Balance() int {
	return  f.balance
}

func (f *Fund) Withdraw(amount int){
	f.balance -= amount
}