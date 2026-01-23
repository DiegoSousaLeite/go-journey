package main
import "fmt"

func main(){
	original := []int{10, 20, 30, 40, 50}
	subSlice := original[1:4] // Pega os elementos do índice 1 ao 3

	fmt.Printf("Original Slice: %v, Len: %d, Cap: %d\n", original, len(original), cap(original))
	fmt.Printf("Subslice: %v, Len: %d, Cap: %d\n", subSlice, len(subSlice), cap(subSlice))

	// Modificando o subslice
	subSlice[0] = 99

	fmt.Printf("After modification:\n")
	fmt.Printf("Original Slice: %v\n", original)
	fmt.Printf("Subslice: %v\n", subSlice)

}