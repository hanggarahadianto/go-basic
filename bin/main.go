package main1

import (
	"fmt"
	"time"
)

var pl = fmt.Println
var pf = fmt.Printf

func main1(){

	var x int = 5
	var y int = 27
	var total = x + y

	pl(total)


// ARRAY -----------------------------

	var arr [5]int
	arr[2] = 433
	arr[3] = 541
	pl(arr)

// CONDITION ---------------------------

x = 5
if x<3 {
	pl("aye")
} else if x== 5{
	pl("equal")
} else{
	pl("bitch")
}


// MAPPING ---------------------------


// LOOPING --------------------------

for i := 0; i <=5; i++ {
	pl(i)
}

// FUNCTION --------------------------


	result := sum(33,2)
	pl(result)

	personName := Person{name: "Jamal"}
	age := Person{age : 40}
	pf("Nama %v, Umur %v \n", personName, age)

	timer("push up")
	timer("sit up")

}

func sum(x int, y int) int {
	return x + y

}

// STRUCT ------------------

type Person struct {
	name string
	age int
}


func timer(todo string){
	for i := 0; i <5; i++ {
		pl(todo, i)
		time.Sleep(time.Second)
	}


}
