package functions

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	res,err := Auth("visit","17omqnu4s40cfj7zt38kcuk100cn2178")
	fmt.Println("res.Code=",res.Code)
	fmt.Println("err = ",err)
	fmt.Println(res.Message,err)
}
