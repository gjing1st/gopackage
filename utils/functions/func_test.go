package functions

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	res, err := Auth("visit", "2178")
	fmt.Println("res.Code=", res.Code)
	fmt.Println("err = ", err)
	fmt.Println(res.Message, err)
}

func TestGetLocationFromQQ(t *testing.T) {
	l, err := GetLocationFromQQ()
	fmt.Println(l)
	fmt.Println(err)
	hexStr := "fee9ecaadafeee72d2eb66a0bd344cdd"
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		// handle error
		fmt.Println("err=", err)
	}
	fmt.Println(data)
}

func TestRetry(t *testing.T) {
	a := make(chan struct{})
	done := make(chan bool)
	go func() {
		_ = Retry(time.Second, a, A)
		done <- true
	}()
	select {
	case <-time.After(time.Second * 3):
		a <- struct{}{}
	case <-done:
		break
	}
}

func A() error {
	fmt.Println(111111)
	return errors.New("111")
}
