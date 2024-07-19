package main

import "strconv"

func main() {

	println("dff")
	println(maxSub(1236321))
}

func maxSub(x int) bool {
	s := strconv.Itoa(x)
	var midi int
	nu := len(s) % 2
	if nu == 0 {
		midi = len(s) / 2
		s = string([]byte(s[:midi]))
	}
	//midi := len(s) / 2
	for i := 1; i < midi+1; i++ {

		println(string(s[midi-i]))
		println(string(s[midi+i]))
		println(".......")
		if s[midi+i] != s[midi-i] {
			return false
		}
	}
	//for i := 1; i < midi+1; i++ {
	//
	//	println(string(s[midi-i]))
	//	println(string(s[midi+i]))
	//	println(".......")
	//	if s[midi+i] != s[midi-i] {
	//		return false
	//	}
	//}
	return true

}
