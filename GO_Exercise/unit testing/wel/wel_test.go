package wel

import (
	"testing"
)

type test struct {
	x        float64
	y        float64
	expected float64
}

func TestCal_power_of_x(t *testing.T) {
	testcases := []test{
		{2, 2, 4},
		{2, 3, 8},
		{2, 4, 16},
	}
	// testcases := test{
	// 	x 	  : 2,
	// 	y 	  : 2,
	// 	expected : 4,
	// }

	for _, tc := range testcases {
		actual := Calc_power_of_x(tc.x, tc.y)
		if actual != tc.expected {
			// t.Errorf("do_cal_power_of_x(%f, %f) = %f; expected %f", tc.x, tc.y, actual, tc.expected)
			t.Fail()
		}
	}
}
