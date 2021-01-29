/*
https://blog.codecentric.de/en/2017/08/gomock-tutorial/
First, we need to install the gomock package github.com/golang/mock/gomock as well as the
mockgen code generation tool github.com/golang/mock/mockgen. Technically, we could do without
the code generation tool, but then we’d have to write our mocks by hand, which is tedious and error-prone.

Usage of GoMock follows four basic steps:
1. Use mockgen to generate a mock for the interface you wish to mock.
2. In your test, create an instance of gomock.Controller and pass it to your mock object’s constructor to obtain a mock object.
3. Call EXPECT() on your mocks to set up their expectations and return values
4. Call Finish() on the mock controller to assert the mock’s expectations
*/
package mock

import (
	"testing"

	"github.com/archerwq/go-lab/concept/test/mock/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCheck(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockFoo := mocks.NewMockFoo(mockCtrl)
	mockFoo.EXPECT().Bar(2).Return(4).Times(1)
	mockFoo.EXPECT().Bar(5).Return(32).Times(1)
	mockFoo.EXPECT().Bar(10).Return(1024).Times(1)
	mockFoo.EXPECT().Bar(20).Return(1024 * 1024).Times(1)
	Convey("Given a Foo with Bar method returning 2 power the given arg", t, func() {
		Convey("call Bar with 2", func() {
			So(Check(mockFoo, 2), ShouldBeFalse)
		})
		Convey("call Bar with 5", func() {
			So(Check(mockFoo, 5), ShouldBeFalse)
		})
		Convey("call Bar with 10", func() {
			So(Check(mockFoo, 10), ShouldBeFalse)
		})
		Convey("call Bar with 20", func() {
			So(Check(mockFoo, 20), ShouldBeTrue)
		})
	})
}
