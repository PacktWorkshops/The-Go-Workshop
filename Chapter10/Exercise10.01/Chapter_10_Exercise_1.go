package main 

import "time"
import "fmt"

func whatstheclock() string {
	return time.Now().Format(time.ANSIC)
}

func main(){
	fmt.Println(whatstheclock())
}