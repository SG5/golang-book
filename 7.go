package main
import "fmt"

func main() {

	xs := []float32{3.14, 5, 4.55, 9.1}

	err, result := average(xs)
	fmt.Println("Error: ", err, ", resut:", result)

	fmt.Println(sum(1, 2, 3))

	closure := getClosure()
	fmt.Println(closure(5))

	fmt.Println(haltAndIsEven(1))
	fmt.Println(haltAndIsEven(2))

	fmt.Println("Maximum number is:", max(1, 5, 100, 10, 789, -1))

	oddGenerator := makeOddGenerator();
	fmt.Println(oddGenerator(), oddGenerator(), oddGenerator(), oddGenerator())
	oddGenerator = makeOddGenerator();
	fmt.Println(oddGenerator(), oddGenerator())

	fmt.Println(fib(11))
	fmt.Println(fib(15))
}

func average(xs []float32) (error, float32) {
	total := float32(.0)
	for _, i := range xs {
		total += i
	}
	return nil, total / float32(len(xs))
}

func sum (argv ...int) int {
	total := 0;
	for _, value := range argv {
		total += value
	}

	return total
}

func getClosure () func(int) int {
	return func(a int) int {
		return a*2;
	}
}

func haltAndIsEven (number int) (half int, isEven bool) {
	half = number / 2
	isEven = false
	if (0 == number%2) {
		isEven = true
	}
	return
}

func max (numbers ...int) (max int) {
	max = numbers[0]
	for _, number := range numbers {
		if (number > max) {
			max = number
		}
	}
	return
}

func makeOddGenerator () func() uint {
	var start uint = 1
	return func () uint {
		defer func () {
			start += 2
		}()
		return start
	}
}

func fib(n int) int {
	if 2 > n {return n}
	return fib(n-1) + fib(n-2)
}