package daily_temperatures

import (
	"fmt"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
