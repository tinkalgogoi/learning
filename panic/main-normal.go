package main
import "fmt"

func f1(ch chan int) {
	defer func () {
		if r := recover(); r!= nil {
			fmt.Println("recovered in f1 from : ", r)
			ch <- 0
		}
	}()

	defer func() {
		fmt.Println("Deferred by f")
	}()
	f2()
	//ch <- 0
}

func f2() {
	defer func() {
		fmt.Println("Deferred by f2")
	}()
	f3()
}

func f3() {
	defer func() {
		fmt.Println("Deferred by f3")
	}()
	//defer func() {
	//	panic("2nd explosion!")
	//}()
	panic("explosion!")
}

func main() {
	ch := make(chan int)
	go f1(ch)
	<-ch
	fmt.Println("last main")
}