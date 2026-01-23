package saudacao
import "fmt"   

// Diminutivo é privada
func sayHi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}

// Maisculo é pública 
func BoasVindas(name string){
	fmt.Println(sayHi(name))
}