/*
Package gostub is used for stubbing variables in tests,
and resetting the original value once the test has been run.
This can be used to stub static variables as well as static functions.
https://github.com/prashantv/gostub
*/
package stub

import (
	"fmt"
	"testing"
	"time"

	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetConfig(t *testing.T) {
	// For simple cases where you are only setting up simple stubs,
	// you can condense the setup and cleanup into a single line.
	// The Stub call must be passed a pointer to the variable that
	// should be stubbed, and a value which can be assigned to the variable.
	defer gostub.Stub(&configFile, "/tmp/test.config").Reset()
	data, err := GetConfig()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
}

func TestGetDate(t *testing.T) {
	Convey("Given date Jun 1, 2015", t, func() {
		now := time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
		// You should keep the return argument from the Stub call
		// if you need to change stubs or add more stubs during test execution.
		stubs := gostub.Stub(&timeNow, func() time.Time {
			return now
		})
		defer stubs.Reset()
		Convey("The day should be 1", func() {
			So(GetDate(), ShouldEqual, 1)
		})
		Convey("When a day past", func() {
			// If you are stubbing a function to return a constant value like
			// in the above test, you can use StubFunc instead.
			stubs.StubFunc(&timeNow, now.AddDate(0, 0, 1))
			Convey("The day should be 2", func() {
				So(GetDate(), ShouldEqual, 2)
			})
		})
	})
	Convey("Stub should be reset", t, func() {
		So(GetDate(), ShouldEqual, time.Now().Day())
	})
}
