package main
import(
	"fmt"
	"math"
)
type Abser interface{
	Abs() float64
}
type Uu struct{
	name string
	age int64
}

func (pu* Uu) getAge() string{
	return pu.name
}

func main(){
	var a Abser
	//f := MyFloat(10)
	v := Vertex{3,4}
	p := Uu{"maiev", 20}
	fmt.Println(p.getAge());

	/*
	a = f
	fmt.Println(a.Abs());

	a = &v
	fmt.Println(a.Abs());
	*/

	a = v
	describe(a);
	//fmt.Println(a.Abs());
}
type  MyFloat float64
func (f MyFloat) Abs() float64{
	if  f < 0{
		return float64(-f)
	}

	return float64(f)
}
type Vertex struct{
	X,Y float64
}

func (v Vertex) Abs()float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func describe(i Abser) {
	fmt.Printf("(%v, %T)\n", i, i)
}
