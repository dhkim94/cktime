package cktime

import (
	"testing"
	"time"
	"fmt"
)

func TestToString(t *testing.T) {
	format := "YYYYMMDD"

	ctm := NewCktime(time.Now(), format)
	str := ctm.ToString()
	if len(str) != 8 {
		t.Fatalf("fail time now to string format: %s\n", format)
	}
	fmt.Printf("time now to string. format: %s, string: %s\n", format, str)
}

func TestToString2(t *testing.T) {
	format := "YYYY-MM-DD hh:mm:ss"

	ctm := NewCktime2(format)
	ctm.SetTime(time.Now())

	str := ctm.ToString()
	if len(str) != 19 {
		t.Fatalf("fail time now to string format: %s\n", format)
	}
	fmt.Printf("time now to string. format: %s, string: %s\n", format, str)
}