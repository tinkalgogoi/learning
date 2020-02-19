package main

type inter1 interface {
	Minter1()
	inter2
}

type inter2 interface {
	Minter2()
}

func main()  {
	var v1 inter1
	v1.Minter2()

	var v2 inter2
	v2 = v1 // v1 = v2 is not possible as v2 d'ont have M1inter1()
	v2.Minter2()

}
