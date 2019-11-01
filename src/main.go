package main
import "fmt"
//import "strconv"
import "reflect"
func main() {
	var b bool = true
	fmt.Println(sayHello("your"))
	fmt.Println(reflect.TypeOf(b))
}

func sayHello(str string) string{
	return "Hello " + str + " world"
}

func add(val int) int {
	var num int = 5
	var arr [5]string
	arr[0] = "xmx"
	//var str string
	//str = strconv.FormatBool(true)
	return val + 99 + num * 2// + strconv.ParseInt("456",0,0)
}