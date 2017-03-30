package main
import "fmt"
import (
	"time"
	"errors"
)
type Person  struct{
	Name string
	Age int
}

type  MyError struct{
	when  time.Time
	start string
}

func (e *MyError)Error()string{
	return fmt.Sprintf("at %v, %s", e.when, e.start);
}
func run() error{
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func aa(x int)(int, error){
	if x == 0{
		return 0,errors.New("meets a error");
	} else{
		return x,nil
	}

}

func (p Person) String()string{
	return fmt.Sprintf("%v (%v years", p.Name, p.Age);
}

func main(){
	x, err := aa(0);

	if err !=nil{
		fmt.Println(err);
	}

	fmt.Println(x);

	a := Person{"uu", 28};
	fmt.Println(a);
	var i interface{}
	//describe(i)


	var b = [4]int{1,2,3,4};
	for i:=0;i<4;i++ {
		fmt.Println(b[i]);	
	}

	i = 42
	//describe(i)

	i = "hello"
	//describe(i)

	i, ok := i.(string)
	fmt.Println(ok, i);

	switch v:= i.(type) {
		case int:
			fmt.Println("int %v", v);
		case string:
			fmt.Println("string", v);
		default:
			fmt.Println("error..%T", v);

	}
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}



