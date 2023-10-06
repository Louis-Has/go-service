package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"math"
	"sync"
	"testing"
)

// todo 有三个函数，分别打印"cat", "fish","dog"要求每一个函数都用一个goroutine，按照顺序打印100次。

var dog = make(chan struct{})
var cat = make(chan struct{})
var fish = make(chan struct{})

func Dog() {
	defer wg.Done()
	<-fish
	fmt.Println("dog")
	dog <- struct{}{}
}

func Cat() {
	defer wg.Done()
	<-dog
	fmt.Println("cat")
	cat <- struct{}{}
}

func Fish() {
	defer wg.Done()
	<-cat
	fmt.Println("fish")
	fish <- struct{}{}
}

var wg sync.WaitGroup

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go Cat()
		go Fish()
		go Dog()
	}
	dog <- struct{}{}

	wg.Wait()
}

// todo 批量生成几十万或者上百万的兑换码
// 用雪花算法实现
// 再解析打印出三段部分
func generateSonyCodes(step int) ([]string, error) {
	codes := make([]string, step)
	flake := sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return 12345, nil
		},
	})

	for i := range codes {
		code, err := flake.NextID()
		if err != nil {
			return nil, err
		}
		timestamp := code >> 24
		machineID := (code >> 16) & 0xFF
		sequence := code & 0xFFFF
		codes[i] = fmt.Sprintf("code: %v | ox: %x | timestamp: %x | machineID: %x | sequence:%d \n",
			code, code, timestamp, machineID, sequence)

		//time.Sleep(1 * time.Millisecond)
	}
	return codes, nil
}

func TestSonyCodes(t *testing.T) {
	result, err := generateSonyCodes(32)

	if err != nil {
		fmt.Printf("code gener err!")
	}

	for _, s := range result {
		fmt.Printf(s)
	}
}

// todo evaluate 方法递归计算表达式结果

func evaluate(nums []float64, total float64, expression string, solutions *[]string, target float64) {
	if len(nums) == 0 {
		if total == target {
			*solutions = append(*solutions, expression)
		}
		return
	}

	for i, num := range nums {
		restNums := append([]float64{}, nums[0:i]...)
		restNums = append(restNums, nums[i+1:]...)

		if expression == "" {
			evaluate(restNums, num, fmt.Sprintf("%g", num), solutions, target)
		} else {
			evaluate(restNums, total+num, fmt.Sprintf("(%s+%g)", expression, num), solutions, target)
			evaluate(restNums, total-num, fmt.Sprintf("(%s-%g)", expression, num), solutions, target)
			evaluate(restNums, total*num, fmt.Sprintf("(%s*%g)", expression, num), solutions, target)
			if math.Abs(num) != 0 {
				evaluate(restNums, total/num, fmt.Sprintf("(%s/%g)", expression, num), solutions, target)
			}
		}
	}
}

func findSolutions(nums []float64, target float64) []string {
	var solutions []string
	evaluate(nums, 0, "", &solutions, target)
	return solutions
}

func TestEvaluate(t *testing.T) {
	nums := []float64{1, 2, 3, 4}
	setTarget := float64(20)

	found := findSolutions(nums, setTarget)

	if len(found) > 0 {
		fmt.Printf("可以通过加减乘除得到期望值 %v \n", setTarget)
		for _, s := range found {
			fmt.Printf("exec: %v \n", s)
		}
	} else {
		fmt.Printf("无法通过加减乘除得到期望值 \n")
	}

}

// todo 台阶问题
// 台阶问题，假如对于上台阶，可以一次上一阶，也可以一次上两阶，写一个方法，实现输入台阶数，输出可以有多少种上法。

func countStair(step int) int {
	if step == 0 || step == 1 {
		return 1
	} else if step == 2 {
		return 2
	} else {
		a := 0
		b := 0
		sum := 1

		for i := 0; i < step; i++ {
			a, b = b, sum
			sum = a + b
		}
		return sum
	}
}

func TestCountStair(t *testing.T) {
	steps := 5
	if steps < 0 {
		fmt.Print("注意输入有意义的台阶数目")
	}

	ways := countStair(steps)
	fmt.Printf("对于 %d 个台阶，有 %d 种上法\n", steps, ways)
}

func TestChan(t *testing.T) { // mihoyo

	poolMax := 3
	pool := make(chan int, poolMax)

	for i := 0; i < 100; i++ {
		pool <- i
		go func() {
			u := <-pool
			fmt.Printf("this is %v\n", u)
		}()
	}

	close(pool)
}

func TestCount(t *testing.T) { // con slice
	buffer := make(chan int)
	var sli []int

	// cus
	go func() {
		defer close(buffer)
		for i := range buffer {
			sli = append(sli, i)
		}
	}()

	//pro
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			buffer <- 1
		}(i)
	}

	wg.Wait()
	fmt.Printf("this is len: %v\n", len(sli))
}

func TestCont(t *testing.T) { // slice con error
	sli := make([]int, 0)

	for i := 0; i < 1000; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			sli = append(sli, i)
		}()
	}

	wg.Wait()
	fmt.Printf("this is len: %v\n", len(sli))
}

func TestArr(t *testing.T) { // map con panic
	arr := make(map[int]int)

	for i := 0; i < 1000; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			arr[i] = i
		}()
	}

	wg.Wait()
	fmt.Printf("this is len: %v\n", len(arr))
}

//  1. 给定无序整数数组，打印最长上升子序列，例如[10,9,2,5,3,7,20,18], 打印[2,5,7,20]

func TestTrueSort(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 20, 18}

	var res [][]int

	res = getLangSort(nums)
	maxI := 0

	for k, v := range res {
		fmt.Printf("this is %v nums %v :%v\n", k, len(v), v)
		if len(v) > len(res[maxI]) {
			maxI = k
		}
	}
	fmt.Printf("so the longest sort is %v\n", res[maxI])
}

func getLangSort(nums []int) [][]int {
	var res [][]int
	for i := 0; i < len(nums); i++ {
		var tic []int
		for u := i; u < len(nums); u++ {
			if len(tic) == 0 || tic[len(tic)-1] <= nums[u] {
				tic = append(tic, nums[u])
			}
		}

		res = append(res, tic)
	}
	return res
}
