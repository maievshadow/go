package main
import (
	"fmt"
	"math"
)
func add(x int ,y int) int{
	return x + y
}

func swap(x int ,y int) (int, int){
	return y , x
}
func split(sum int )(y int, x int){
	x = sum - 1
	y = sum * 10
	return
}

type Vertex struct{
	X, Y float64
}
type  MyFloat float64
func  (v Vertex) Abs() float64{
	return math.Sqrt(v.X * 	v.X + v.Y * v.Y)
}

func Abs2(v Vertex) float64{
	return math.Sqrt(v.X * 	v.X + v.Y * v.Y)
}
func (v *Vertex) Scale(f float64) {
	v.X = f* v.X
	v.Y = f* v.Y
	fmt.Println(v);
}

func Scale2(v *Vertex, f float64) {
	v.X = f* v.X
	v.Y = f* v.Y
	fmt.Println(v);
}

func main(){
	v := Vertex{3,4}
	Scale2(&v, 10);
	fmt.Println(Abs2(v));
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)



	s := []int{2, 3, 5, 7, 11, 13}
	s2 := []string{"uu", "zz", "qq"};

	printSlice(s);
	printSlice2(s2);
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

