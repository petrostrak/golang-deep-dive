package myinterface

// mkdir ~/go/src/myInterface
// cp myInterface.go ~/go/src/myInterface
// go install myInterface
type Shape interface {
	Area() float64
	Perimeter() float64
}
