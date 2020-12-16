package zcase

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

const length int = 10000

func testString() {
	var itmeItem = time.Now().Nanosecond()
	var o = test2()
	var endTime = time.Now().Nanosecond()
	var useTime = endTime - itmeItem
	fmt.Printf("useTime %d, str len : %d \n", useTime, len(o))
}

func test1() string {
	var i string = ""
	for a := 0; a < length; a++ {
		i = i + "sdf|"
	}
	return i
}

func test2() string {
	var bufferString bytes.Buffer
	for a := 0; a < length; a++ {
		bufferString.WriteString("sdf|")
	}
	return bufferString.String()
}

func test3() string {
	var bufferString strings.Builder
	for a := 0; a < length; a++ {
		bufferString.WriteString("sdf|")
	}
	return bufferString.String()
}
