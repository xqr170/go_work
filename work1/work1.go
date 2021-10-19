package main

import "fmt"

type Data struct {
	Name string
	Num  int
}

func main() {
	chA, chB, chC, chD := make(chan Data), make(chan Data), make(chan Data), make(chan Data)
	go func() {
		for i := 1; i < 41; {
			var data0 = Data{Name: "张三", Num: i}
			chA <- data0
			i += 4
		}
	}()
	go func() {
		for i := 2; i < 41; {
			var data0 = Data{Name: "李四", Num: i}
			chB <- data0
			i += 4
		}
	}()
	go func() {
		for i := 3; i < 41; {
			var data0 = Data{Name: "王五", Num: i}
			chC <- data0
			i += 4
		}
	}()
	go func() {
		for i := 4; i < 41; {
			var data0 = Data{Name: "赵六", Num: i}
			chD <- data0
			i += 4
		}
	}()

	flag := false
	for !flag {
		dataA := <-chA
		fmt.Println("" + dataA.Name + "")
		dataB := <-chB
		fmt.Println("" + dataB.Name + "")
		dataC := <-chC
		fmt.Println("" + dataC.Name + "")
		dataD := <-chD
		fmt.Println("" + dataD.Name + "")
	}
}
