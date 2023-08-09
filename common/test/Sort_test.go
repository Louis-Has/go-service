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

		go Dog()
		go Cat()
		go Fish()
	}
	fish <- struct{}{}

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

func countStair(step int) int {
	if step == 0 || step == 1 {
		return 1
	} else if step == 1 {
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
