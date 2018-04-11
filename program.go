package main

func main() {
	/*If we list all the natural numbers below 10 that are multiples of 3 and 5, we get 3, 5, 6, and 9. The sum of these multiples is 23.
	Find the sum of all the multiples of 3 and 5 below 1000.
	*/
	println("This application will list the sum of the natural numbers below 1000 which are a multiple of 3 and 5.")

	var i = 0
	var total = 0
	for i = 1; i < 1000; i++{
		if i % 3 == 0 || i % 5 == 0 {
			total = total + i

		}
	}
	println(total)
}
