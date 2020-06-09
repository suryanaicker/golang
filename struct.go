package main
import(
	"fmt"
	"math"
)
type vertex struct{
	X,Y float64
}
type myint int
var x myint=87

func main(){
v:=vertex{4.5,3.6}
fmt.Println(v)
v.X=8.8
fmt.Println(v)


p:=&v// pointer to struct
fmt.Println(p)
fmt.Println(*p)
fmt.Println(p.X)//no need to explicitly derefence the pointer
fmt.Println((*p).X)

pv:=new(vertex)// we can also create a pointer using new functionpv.X=6.5
pv.Y=3.0
pv.X=4// pv is a pointer to the function
fmt.Println(pv)
fmt.Println("the calculate value is ",pv.calc())
fmt.Println(v)
fmt.Println("the calculate value is ",v.calc())


fmt.Println(pv)
fmt.Println("the calculate value is ",pv.add())
fmt.Println(v)
fmt.Println("the calculate value is ",v.add())
fmt.Println("the calculate value is ",p.add())
//i:=45
//pi:=&i
//fmt.Println(pi.calc1())

y := &x
fmt.Println("passing address")
fmt.Println( y.calc1() )
fmt.Println("passing value")
fmt.Println( x.calc1() )
fmt.Println("passing address")
fmt.Println( y.calc2() )
fmt.Println("passing value")
fmt.Println( x.calc2() )

}

func (v vertex)calc()float64{
	return math.Sqrt((v.X*v.X)+(v.Y*v.Y))
}

func (p *myint) calc1() myint{
	fmt.Println("here receiver is of pointer type")
	return *p
	
}

func (p myint) calc2() myint{
	fmt.Println("here receiver is of value type")
	return p
}



func (v vertex)add()float64{
	return v.X+v.Y
}

//doubt here the receiver of pointer type is taking both T and *T

/*func (v *int)calc1()int{		reciever can be of struct type or new datatype defines in the same package
	return ( *v+*v)         but cannot be of built in data type or new data type defined in other packages			
}




************************************************
there fore if the reciever of the method can be a struct or a new datatype 
in both the cases the receiver can be either pointer or value
but we can pass value to the pointer type and pointer to the value type
it works fine for both
*/
