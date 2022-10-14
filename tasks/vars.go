package tasks

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

const (
	UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"
)

var (
	Jar, _         = cookiejar.New(nil)
	CourseDocument = InitTasks()
	Client         = http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
		Jar:     Jar,
		Timeout: time.Second * 20,
	}
	Courses        []*Course = make([]*Course, 0)
	Username       string
	Password       string
	Authorized     = false

	LinksClient = http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar:     Jar,
		Timeout: time.Second * 20,
	}
)
