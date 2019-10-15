package main
	
import "fmt"
import "time"

func main(){
	Current := time.Now()
	NYtime, _ := time.LoadLocation("America/New_York")
	LA, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println("The local current time is:",Current.Format(time.ANSIC))
	fmt.Println("The time in New York is: ",Current.In(NYtime).Format(time.ANSIC))
	fmt.Println("The time in Los Angeles is: ",Current.In(LA).Format(time.ANSIC))
}