package iteration

import "fmt"

func Repeat(character string, count int) string{
	var repeated string 
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	fmt.Println(Repeat("abc", 7))
}