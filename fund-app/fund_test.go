package funding

import "testing"

func BenchMarkFund(b *testing.B)  {

	// Add as many dollars as we have iterations this run

	fund := NewFund(b.N)
	
}
