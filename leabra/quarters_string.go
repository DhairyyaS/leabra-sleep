// Code generated by "stringer -type=Quarters"; DO NOT EDIT.

package leabra

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Q1-0]
	_ = x[Q2-1]
	_ = x[Q3-2]
	_ = x[Q4-3]
	_ = x[QuartersN-4]
}

const _Quarters_name = "Q1Q2Q3Q4QuartersN"

var _Quarters_index = [...]uint8{0, 2, 4, 6, 8, 17}

func (i Quarters) String() string {
	if i < 0 || i >= Quarters(len(_Quarters_index)-1) {
		return "Quarters(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Quarters_name[_Quarters_index[i]:_Quarters_index[i+1]]
}

func (i *Quarters) FromString(s string) error {
	for j := 0; j < len(_Quarters_index)-1; j++ {
		if s == _Quarters_name[_Quarters_index[j]:_Quarters_index[j+1]] {
			*i = Quarters(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Quarters")
}
