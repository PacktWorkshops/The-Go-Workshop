package main
	
import "fmt"
import "time"
import "strconv"

func main(){
	date := time.Date(2019, 1, 31, 2, 49, 21, 324359102, time.UTC)

	day := strconv.Itoa(date.Day())

	month := strconv.Itoa(int(date.Month()))
	year := strconv.Itoa(date.Year())
	hour := strconv.Itoa(date.Hour())
	minute := strconv.Itoa(date.Minute())
	second := strconv.Itoa(date.Second())
	fmt.Println(hour + ":" + minute + ":" + second + " " + year + "/" + month + "/" + day) 
}
