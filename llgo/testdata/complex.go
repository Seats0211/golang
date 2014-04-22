package main

func main() {
	var f32 float32 = 1
	var f64 float64 = 1
	c64 := complex(f32, f32+1)
	println(c64)
	println(-c64)
	println(c64 == c64)
	c128 := complex(f64, f64+1)
	println(c128)
	println(-c128)
	println(c128 == c128)
}
