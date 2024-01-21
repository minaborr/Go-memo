package main

import "fmt"

//создаем тип 
type intFunc func(int64)

func main() {
	// анонимная функция которая подходит под сигнатуру типу intFunc
	// и она выводит преобразованное значение
	cor := func(i int64){
		fmt.Println(i*10)
	}

	fff(2, cor)
}

func fff(num int64, i intFunc){
	i(num)
}
