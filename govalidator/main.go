package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

//URL
func f01() {
	type User struct {
		FirstName string
		LastName  string
	}
	str := govalidator.ToString(&User{"John", "Juan"})
	fmt.Println(str)
	fmt.Println(govalidator.IsURL(`www.baidu.com`))
	fmt.Println(strings.Index("fdjhfgujhrthfdssssdfgh", "s"))
}

//Each for slices
func f02() {
	data := []interface{}{1, 2, 3, 4, 5}
	var fn govalidator.Iterator = func(value interface{}, index int) {
		println(value.(int))
	}
	govalidator.Each(data, fn)
}

//Map for slices
func f03() {
	data := []interface{}{1, 2, 3, 4, 5}
	var fn govalidator.ResultIterator = func(value interface{}, index int) interface{} {
		return value.(int) * 3
	}
	fmt.Println(govalidator.Map(data, fn))
}

//Filter, Count for slices
func f04() {
	data := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,30}
	var fn govalidator.ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	fmt.Println(govalidator.Filter(data, fn)) // result = []interface{}{2, 4, 6, 8, 10}
	fmt.Println(govalidator.Count(data, fn))  // result = 5
}

func main() {
	//f01()
	//f02()
	//f03()
	f04()
}
