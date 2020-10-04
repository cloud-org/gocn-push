package server

import (
	"fmt"
	"testing"
	"time"
)

func Test_GetNewsContent(t *testing.T) {
	publishTime := time.Now().Add(time.Hour * -24)
	_, contents := GetNewsContent(publishTime)
	fmt.Println(contents)
}
