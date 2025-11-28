/*
Задача 2: Обработка данных с помощью каналов

Описание: Создайте программу, в которой одна горутина генерирует числа от 1 до 10 и отправляет их через канал.
Другая горутина читает числа из канала и выводит их на экран. После отправки всех чисел закройте канал и завершите программу.

Требования:

    Используйте каналы для передачи данных.
    Горутина-генератор должна отправлять числа и закрывать канал после этого.
    Горутина-потребитель должна читать из канала и выводить числа, пока канал не закрыт.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func NumberGenerator(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(c)
	var r int
	for range 10 {
		r = int(rand.Int31n(10))
		time.Sleep(500*time.Millisecond)
		c <- r
	}
}
func PrintNumber(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	wg.Add(2)
	go NumberGenerator(c, &wg)
	go PrintNumber(c, &wg)
	wg.Wait()	
}
