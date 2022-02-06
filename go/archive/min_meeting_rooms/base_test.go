package min_meeting_rooms

import (
	"fmt"
	"testing"
)

func Test_minMeetingRooms(t *testing.T) {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(minMeetingRooms([][]int{{7, 10}, {2, 4}}))
	fmt.Println(minMeetingRooms([][]int{{13, 15}, {1, 13}}))
}
