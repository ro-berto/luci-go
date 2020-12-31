package build

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/memlogger"
)

func TestLogableNoop(t *testing.T) {
	Convey(`Loggable Noop`, t, func() {
		Convey(`nop`, func() {
			var nop nopStream
			n, err := nop.Write([]byte("hey"))
			So(n, ShouldEqual, 3)
			So(err, ShouldBeNil)

			So(nop.Close(), ShouldBeNil)
		})

		Convey(`nopDatagram`, func() {
			var nop nopDatagramStream
			So(nop.WriteDatagram([]byte("hey")), ShouldBeNil)
			So(nop.Close(), ShouldBeNil)
		})

		Convey(`loggingWriter`, func() {
			ctx := memlogger.Use(context.Background())
			logs := logging.Get(ctx).(*memlogger.MemLogger)
			lw := makeLoggingWriter(ctx, "some log")

			Convey(`single line`, func() {
				n, err := lw.Write([]byte("hello world\n"))
				So(n, ShouldEqual, 12)
				So(err, ShouldBeNil)

				So(logs, memlogger.ShouldHaveLog, logging.Info, "hello world", logging.Fields{
					"build.logname": "some log",
				})

				So(lw.Close(), ShouldBeNil)
				So(logs.Messages(), ShouldHaveLength, 1)
			})

			Convey(`multi line`, func() {
				n, err := lw.Write([]byte("hello world\ncool\nbeans\n"))
				So(n, ShouldEqual, 23)
				So(err, ShouldBeNil)

				So(logs, memlogger.ShouldHaveLog, logging.Info, "hello world")
				So(logs, memlogger.ShouldHaveLog, logging.Info, "cool")
				So(logs, memlogger.ShouldHaveLog, logging.Info, "beans")

				So(lw.Close(), ShouldBeNil)
				So(logs.Messages(), ShouldHaveLength, 3)
			})

			Convey(`partial line`, func() {
				n, err := lw.Write([]byte("hello worl"))
				So(n, ShouldEqual, 10)
				So(err, ShouldBeNil)

				So(logs.Messages(), ShouldHaveLength, 0)

				n, err = lw.Write([]byte("d\ncool\n\n\n"))
				So(n, ShouldEqual, 9)
				So(err, ShouldBeNil)

				So(logs, memlogger.ShouldHaveLog, logging.Info, "hello world")
				So(logs, memlogger.ShouldHaveLog, logging.Info, "cool")
				So(logs, memlogger.ShouldHaveLog, logging.Info, "")

				So(lw.Close(), ShouldBeNil)
				So(logs.Messages(), ShouldHaveLength, 4)
			})

			Convey(`partial flush`, func() {
				n, err := lw.Write([]byte("hello worl"))
				So(n, ShouldEqual, 10)
				So(err, ShouldBeNil)

				So(lw.Close(), ShouldBeNil)
				So(logs.Messages(), ShouldHaveLength, 1)

				So(logs, memlogger.ShouldHaveLog, logging.Info, "hello worl")
			})
		})

		Convey(`loggingWriter - debug`, func() {
			ctx := memlogger.Use(context.Background())
			ctx = logging.SetLevel(ctx, logging.Debug)
			logs := logging.Get(ctx).(*memlogger.MemLogger)
			lw := makeLoggingWriter(ctx, "$some log")

			n, err := lw.Write([]byte("hello world\n"))
			So(n, ShouldEqual, 12)
			So(err, ShouldBeNil)

			So(logs, memlogger.ShouldHaveLog, logging.Debug, "hello world", logging.Fields{
				"build.logname": "$some log",
			})

			So(lw.Close(), ShouldBeNil)
			So(logs.Messages(), ShouldHaveLength, 1)
		})

		Convey(`loggingWriter - debug skip`, func() {
			ctx := memlogger.Use(context.Background())
			ctx = logging.SetLevel(ctx, logging.Info) // ignore debug
			logs := logging.Get(ctx).(*memlogger.MemLogger)
			lw := makeLoggingWriter(ctx, "$some log")

			n, err := lw.Write([]byte("hello world\n"))
			So(n, ShouldEqual, 12)
			So(err, ShouldBeNil)

			So(lw.Close(), ShouldBeNil)
			So(logs.Messages(), ShouldHaveLength, 0)
		})
	})
}
