package Test

import (
	"fmt"
	"github.com/TwiN/go-color"
	"log"
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
	log.Printf(color.InRed("ready to log"))
	t.Fail()
}

func TestScan(t *testing.T) {
	nowTime := time.Now()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fmt.Printf("this is %v\n;format is %v\n", nowTime, nowTime.Format("2006-01-02 15:04:05.000 Mon Jan"))
	log.Printf("this is log%v\n", nowTime)
}
