package main
import "fmt"


func BilanganGanjil(inputValue int){

	if (inputValue % 2) == 0 {
		fmt.Printf(" %d adalah bilangan Genap \n", inputValue)
	} else {
		fmt.Printf(" %d adalah bilangan Ganjil \n", inputValue)
	}
}

func main() {

	value := 7
	BilanganGanjil(value)
}

// written by darma14