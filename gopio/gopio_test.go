package gopio

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var tempDir, _ = ioutil.TempDir("", "")
var pin = Pin{N22}

func TestMain(m *testing.M) {
	getSignalPath = func(pin *Pin) string {
		return tempDir
	}

	ControlPath = tempDir

	f, _ := os.Create(filepath.Join(tempDir, "value"))
	f.Write([]byte("0"))

	os.Exit(m.Run())

}

func TestPin(t *testing.T) {

	Convey("Given I have a Pin", t, func() {
		Convey("When I export it it", func() {
			pin.Export()

			Convey("Then the kernel id of the Pin should be written to the control path.", func() {
				b, _ := ioutil.ReadFile(filepath.Join(tempDir, "export"))
				val, _ := strconv.Atoi(string(b))

				So(val, ShouldEqual, 116)
			})
		})

		Convey("When I read its value", func() {
			val := pin.Read()

			Convey("Then the direction must be \"in\"", func() {
				b, _ := ioutil.ReadFile(filepath.Join(tempDir, "direction"))
				direction := string(b)
				So(direction, ShouldEqual, In)
			})

			Convey("And the value should be 0.", func() {
				So(val, ShouldEqual, 0)
			})
		})

		Convey("When I write value 1 to it", func() {
			pin.Write(1)
			Convey("Then the direction must be \"out\"", func() {
				b, _ := ioutil.ReadFile(filepath.Join(tempDir, "direction"))
				direction := string(b)
				So(direction, ShouldEqual, Out)
			})
			Convey("And and the Pin should read 1.", func() {
				b, _ := ioutil.ReadFile(filepath.Join(tempDir, "value"))

				val, _ := strconv.Atoi(string(b))
				So(val, ShouldEqual, 1)
			})
		})
	})
}

func BenchmarkReadPin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pin.Read()
	}
}

func BenchmarkWritePin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pin.Write(0)
	}
}
