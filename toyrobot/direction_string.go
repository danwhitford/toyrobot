// Code generated by "stringer -type=Direction"; DO NOT EDIT.

package toyrobot

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NORTH-0]
	_ = x[EAST-1]
	_ = x[SOUTH-2]
	_ = x[WEST-3]
}

const _Direction_name = "NORTHEASTSOUTHWEST"

var _Direction_index = [...]uint8{0, 5, 9, 14, 18}

func (i Direction) String() string {
	if i >= Direction(len(_Direction_index)-1) {
		return "Direction(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Direction_name[_Direction_index[i]:_Direction_index[i+1]]
}
