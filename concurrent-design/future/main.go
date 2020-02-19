package main

type Matrix struct{} // todo change to actual matrix
// futures used internally
type futureMatrix chan Matrix

// API remains the same
func Inverse (a Matrix) Matrix {
	return <-InverseAsync(promise(a))
}

func Product (a Matrix, b Matrix) Matrix {
	return <-ProductAsync(promise(a), promise(b))
}

// expose async version of the API
func InverseAsync (a futureMatrix) futureMatrix {
	c := make (futureMatrix);
	go func () { c <- inverse(<-a) } ();
	return c
}

func ProductAsync (a futureMatrix) futureMatrix {
	c := make (futureMatrix)
	go func () { c <- product(<-a) } ()
	return c
}

// actual implementation is the same as before
func product (a Matrix, b Matrix) Matrix {
	return struct{}{} //todo
}

func inverse (a Matrix) Matrix {
	return struct{}{} //todo
}



// utility fxn: create a futureMatrix from a given matrix
func promise (a Matrix) futureMatrix {
	future := make (futureMatrix, 1)
	future <- a
	return future
}


func InverseProduct (a Matrix, b Matrix) {
	a_inv := Inverse(a)
	b_inv := Inverse(b)
	return Product(a_inv, b_inv)
}

func main() {

}
