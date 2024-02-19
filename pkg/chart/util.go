package chart

import (
	"fmt"
	"time"
)

func randomImageName() string {
	return fmt.Sprintf("%d.png", time.Now().Unix())
}
