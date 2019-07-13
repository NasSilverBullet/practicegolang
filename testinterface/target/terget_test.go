package target

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/practicegolang/testinterface/user"
)

type testUser struct {
	user.User // これでinterfaceを満たせる
	age       int
}

// 差し替えたいメソッドだけ定義する（オーバーライド）
func (u *testUser) Age() int {
	return u.age
}
func TestIsTestTarget(t *testing.T) {

	testCases := []struct {
		age      int
		expected bool
	}{
		{24, false},
		{25, false},
		{26, true},
		{30, true},
	}

	for _, tt := range testCases {
		u := testUser{age: tt.age}
		reflect.DeepEqual(tt.expected, IsTestTarget(&u))
	}
}
