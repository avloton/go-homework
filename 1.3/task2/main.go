package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Введите первое число: ")
	reader0 := bufio.NewReader(os.Stdin)
	rawInput0, err := reader0.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Введите второе число: ")
	reader1 := bufio.NewReader(os.Stdin)
	rawInput1, err := reader1.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Введите третье число: ")
	reader2 := bufio.NewReader(os.Stdin)
	rawInput2, err := reader2.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Введите четвертое число: ")
	reader3 := bufio.NewReader(os.Stdin)
	rawInput3, err := reader3.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Введите пятое число: ")
	reader4 := bufio.NewReader(os.Stdin)
	rawInput4, err := reader4.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input0 := strings.TrimSpace(rawInput0)
	input1 := strings.TrimSpace(rawInput1)
	input2 := strings.TrimSpace(rawInput2)
	input3 := strings.TrimSpace(rawInput3)
	input4 := strings.TrimSpace(rawInput4)

	fmt.Printf("%v %v %v %v %v \n", input0, input1, input2, input3, input4)

	a, err := strconv.Atoi(input0)
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := strconv.Atoi(input1)
	if err != nil {
		fmt.Println(err)
		return
	}
	c, err := strconv.Atoi(input2)
	if err != nil {
		fmt.Println(err)
		return
	}

	d, err := strconv.Atoi(input3)
	if err != nil {
		fmt.Println(err)
		return
	}

	e, err := strconv.Atoi(input4)
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Print("Отсортированные элементы: ")
	fmt.Println(sort(a,b,c,d,e))

	fmt.Print("Самое большое число: ")
	fmt.Println(max(a,b,c,d,e))

	fmt.Print("Самое маленькое число: ")
	fmt.Println(min(a,b,c,d,e))

	fmt.Print("Среднее арифметическое: ")
	fmt.Println(avg(a,b,c,d,e))


}

func sort(a int, b int, c int, d int, e int) (int, int, int, int, int) {
	 for i := 0; i < 5; i++ {
        if a < b { a, b = b, a }
        if b < c { b, c = c, b }
        if c < d { c, d = d, c }
        if d < e { d, e = e, d }
    }
	return a, b, c ,d, e
}

func max(a int, b int, c int, d int, e int) int {
	maximum, _, _, _, _ := sort(a, b, c, d, e)
	return maximum
}

func min(a int, b int, c int, d int, e int) int {
	_, _, _, _, minimum := sort(a, b, c, d, e)
	return minimum
}

func avg(a int, b int, c int, d int, e int) float32 {
	return (float32(a)+float32(b)+float32(c)+float32(d)+float32(e))/5.0
}