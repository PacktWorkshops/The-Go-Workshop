package main
	
import "fmt"
import "time"

func main(){
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	length := end.Sub(start)
	fmt.Println("The execution too exactly",length.Seconds(),"seconds!")
}