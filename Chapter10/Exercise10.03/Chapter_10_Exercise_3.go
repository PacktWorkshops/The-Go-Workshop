package main 
import "time"
import "fmt"
func timeDiff(timezone string) (string, string)  {
	Current := time.Now()
	RemoteZone, _ := time.LoadLocation(timezone)
	RemoteTime := Current.In(RemoteZone)
	fmt.Println("The current time is: ",Current.Format(time.ANSIC))
	fmt.Println("The timezone:",timezone,"time is:",RemoteTime)
	return Current.Format(time.ANSIC), RemoteTime.Format(time.ANSIC)
}
func main(){
	fmt.Println(timeDiff("America/Los_Angeles"))
}
