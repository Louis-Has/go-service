package Test

import (
	"fmt"
	"go-service/Common/utils"
	"testing"
)

func Test(t *testing.T) {
	months := []int{-3, -2, -1, 0, 1, 2}
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
		fmt.Printf("getMonthDays ==>%v月%v天\n",
			int(utils.GetCurrentTime().Month())+v, utils.GetMonthDays(v),
		)
	}
}
