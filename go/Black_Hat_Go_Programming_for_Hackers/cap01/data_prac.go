var x= "Hello Wold"
z:=int(42)



var s = make ([]string,0)
var m = make(map[string]string)
s = append(s,"some string")
m["some key"]="some value"


var count := int(42)
ptr := &count
fmt.Println(*ptr)
*ptr=100
fmt.Println(count)



type Person struct{

	Name string
	Age int
}
func(p *Person)SayHello(){
	fmt.Println("Hello,",p.Name)
}
func main(){
	var guy =new(Person)
	guy.Name="Dave"
	guy.SayHello()
}


type Friend interface{
	SayHello()
}

