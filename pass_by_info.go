package main

import "fmt"




func main(){
	var  x,y,z,result int
	fmt.Scan(&x,&y,&z)
	result = hitung(&x,&y,z)
	fmt.Println(x,y,z)
	fmt.Println(result)

}


func hitung(b1 *int, b2 *int,b3 int ) int {
	*b1 = *b1 + 1
	b3 = b3 + 1
	return *b1 + *b2
}