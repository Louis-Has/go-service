package Test

import (
	"fmt"
	"go-service/Common/utils"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	months := []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf(
		"currentTime ==>%v\n"+
			"addTime ==>%v\n"+
			"firstDay ==>%v\n"+
			"lastDay ==>%v\n"+
			"currentMonthDays ==>4月%v天\n",
		utils.GetCurrentTime(),
		utils.GetAddTime(),
		utils.GetFirstDay(),
		utils.GetLastDay(),
		utils.CurrentMonthDays(),
	)
	for _, v := range months {
		currentMonth, months := utils.GetMonthDays(v)
		fmt.Printf("getMonthDays ==>%v月%v天\n",
			currentMonth, months,
		)
	}
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	c := <-timer.C
	fmt.Printf("ddd%v\n", c)
}

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			fmt.Printf("now is writing%v\n", count)
			yield <- count
			count++
		}
	}()
	return yield
}

var resume chan int

func generateInteger() int {
	return <-resume
}

func TestResume(t *testing.T) {
	resume = integers()
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
}
