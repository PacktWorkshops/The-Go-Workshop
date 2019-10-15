package payroll
import "fmt"
type Payer interface {
	Pay() (string, float64)
}
type Employee struct {
	Id        int
	FirstName string
	LastName  string
}
func PayDetails(p Payer) {
	fullName, yearPay := p.Pay()
	fmt.Printf("%s got paid %.2f for the year\n", fullName, yearPay)
}
