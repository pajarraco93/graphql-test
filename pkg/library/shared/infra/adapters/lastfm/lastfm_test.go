package lastfm

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/pajarraco93/graphql-test/pkg/library/domain/entities"
)

func TestGetGroupInfo(t *testing.T) {
	Convey("Given we want to get the group info", t, func() {
		Convey("When the lastFM returns a successfull response", func(c C) {
			expectedResponse := "LastFM Stuff & Info"

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				c.So(r.Method, ShouldEqual, http.MethodGet)
				c.So(r.Header.Get("user-agent"), ShouldEqual, UserAgent)
				c.So(r.Header.Get("Content-Type"), ShouldEqual, "application/json")

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")

				_, err := w.Write([]byte(expectedResponse))
				c.So(err, ShouldBeNil)
			}))

			Convey("Then", func() {
				api := NewLastFMAPI("test", ts.URL)

				info, err := api.GetGroupInfo(entities.Group{})
				Convey("Info should returned and no error", func() {
					So(err, ShouldBeNil)
					So(info, ShouldEqual, expectedResponse)
				})
			})
		})

		Convey("When the lastFM returns a bad response", func(c C) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				c.So(r.Method, ShouldEqual, http.MethodGet)
				c.So(r.Header.Get("user-agent"), ShouldEqual, UserAgent)
				c.So(r.Header.Get("Content-Type"), ShouldEqual, "application/json")

				w.WriteHeader(http.StatusGatewayTimeout)
				w.Header().Set("Content-Type", "application/json")
			}))

			Convey("Then", func() {
				api := NewLastFMAPI("test", ts.URL)

				info, err := api.GetGroupInfo(entities.Group{})
				Convey("Info should returned and no error", func() {
					So(err, ShouldNotBeNil)
					So(info, ShouldEqual, "")
				})
			})
		})
	})
}
