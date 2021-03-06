// RUN: llgo -o %t %s
// RUN: %t > %t1 2>&1
// RUN: go run %s > %t2 2>&1
// RUN: diff -u %t1 %t2

package main

func f1() (x int) {
	x = 123
	return
}

func f2() (x int) {
	return 456
}

func f3() (x, y int) {
	y, x = 2, 1
	return
}

func f4() (x, _ int) {
	x = 666
	return
}

func main() {
	x := f1()
	println(x)
	x = f2()
	println(x)

	var y int
	x, y = f3()
	println(x, y)

	x, y = f4()
	println(x, y)
}
