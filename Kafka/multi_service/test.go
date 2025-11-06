package main

import "fmt"


func modify(s *string) {
	*s = "world"

	
}

func changeNum(b *int){
	*b = 5
}


func main() {
	a := "hello"
	b := 1
	fmt.Println(a)
	fmt.Println(b)

	modify(&a)
	changeNum(&b)
	fmt.Println(a)
	fmt.Println(b)

}








/* 


//user service
go run main.go user_service.go


//res_service
go run crud.go dto.


//kafka

go run dto.go main.go process_event.go producer.go


*/