package main

import "fmt"

type student struct {
	rollNo int
	avg    float64
	name   string
	result bool
}

// int, string, float, bool without pointer
func allBasicTypes(rollNo int, avg float64, name string, result bool) {
	rollNo = 10
	avg = 60.50
	name = "Naveen"
	result = true
}

// int, string, float, bool with pointer
func allBasicPointerTypes(rollNo *int, avg *float64, name *string, result *bool) {
	*rollNo = 10
	*avg = 60.50
	*name = "Naveen"
	*result = true
}

// Struct without pointer
func structTypes(stud student) {
	stud.rollNo = 50
	stud.name = "Sathihi Vavileti"
	stud.avg = 99.00
	stud.result = true
}

// Struct with pointer
func structPointerTypes(stud *student) {
	stud.rollNo = 50
	stud.name = "Sathihi Vavileti"
	stud.avg = 99.00
	stud.result = true
}

// slice doesnt require pointer
func slicetype(a []int) {
	for i := range a {
		a[i]++
	}
}

// map  doesnt require pointer
func maptype(mp map[string]int) {
	mp["sahi"] = 100
	mp["navi"] = 200
}

func main() {
	rollNo := 1
	avg := 40.20
	name := "Nimmala"
	result := false

	// int, string, float, bool without pointer
	allBasicTypes(rollNo, avg, name, result)
	fmt.Println(rollNo, avg, name, result) //1 40.2 Nimmala false

	// int, string, float, bool with pointer
	allBasicPointerTypes(&rollNo, &avg, &name, &result)
	fmt.Println(rollNo, avg, name, result) //10 60.5 Naveen true

	sahithi := student{
		rollNo: 100,
		avg:    30.30,
		name:   "sahithi",
		result: false,
	}
	// struct without pointers
	structTypes(sahithi)
	fmt.Println(sahithi) // output : {100 30.3 sahithi false}

	// struct with pointers
	structPointerTypes(&sahithi)
	fmt.Println(sahithi) // output: {50 99 Sathihi Vavileti true}

	// struct without pointers --> slices doent required pointers,  y defalul it it act as pass by pointer to the function
	sahi := []int{1, 2, 3, 4}
	slicetype(sahi)
	fmt.Println(sahi) // output: [2 3 4 5]

	namess := map[string]int{
		"sahi": 1,
		"navi": 2,
	}
	maptype(namess)
	fmt.Println(namess) //Output: map[navi:200 sahi:100]

}
