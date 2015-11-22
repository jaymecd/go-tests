package dot

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSomething(t *testing.T) {
	Convey("Description of cool behavior", t, func() {
		So(1, ShouldEqual, 1)
	})
}
