package main

import "fmt"

func main(){
	var a, b float32
	var ch string
	fmt.Scan(&a, &b, &ch)
	if(ch == "+"){
		fmt.Println(a + b)
	}else if(ch == "-"){
		fmt.Println(a - b)
	}else if(ch == "*"){
		fmt.Println(a * b)
	}else if(ch == "/"){
		if(b != 0){
			fmt.Println(a / b)
		}else{
			fmt.Println("Делить на ноль нельзя!")
		}
	}else if(ch == "%"){
		if(b != 0){
			fmt.Println(int(a) % int(b))
		}else{
			fmt.Println("Делить на ноль нельзя!")
		}
	}else{
		fmt.Println("Такой операции пока нет")
	}
}