package main

import (
	"context"
	"go-service/common/log"
	"go.uber.org/zap"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				//fmt.Println("ready to write", n)
				n = n + 7
			}
		}
	}()
	return dst
}

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel() // 当我们取完需要的整数后调用cancel
	//
	//for n := range gen(ctx) {
	//	fmt.Println(n)
	//	if n >= 55 {
	//		break
	//	}
	//}
	for i := 1; i < 25; i++ {
		switch i % 4 {
		case 0:
			logger.logger.Debug("this is Debug", zap.Int("ik", i))
		case 1:
			logger.Info("this is Info", zap.Int("ik", i))
		case 2:
			logger.Warn("this is Warn", zap.Int("ik", i))
		case 3:
			logger.Error("this is Error", zap.Int("ik", i))

		}
	}
}
