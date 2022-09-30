package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-res", "/make-reservation", "GET", []postData{}, http.StatusOK},

	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "start", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "start", value: "2020-01-02"},
	}, http.StatusOK},
	{"make reservation post", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "me@here.com"},
		{key: "phone", value: "555 555 5555"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes) //test server
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL + e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}

	}
}

// Step 1 check test - if error 400 -mux.Use(NoSurf) should be commented in setup_test.go
// xxx@xxx:~/go/src/Web_Part1/bookings/internal/handlers(master)$ go test -v
// === RUN   TestHandlers
//     handlers_test.go:74: for post-search-availability expected 200 but got 400
//     handlers_test.go:74: for post-search-availability-json expected 200 but got 400
//     handlers_test.go:74: for make reservation post expected 200 but got 400
// --- FAIL: TestHandlers (0.03s)
// FAIL
// exit status 1
// FAIL    github.com/RaminCH/bookings/internal/handlers   0.045s

//Step2  PASSED
// xxx@xxx:~/go/src/Web_Part1/bookings/internal/handlers(master)$ go test -v
// === RUN   TestHandlers
// 2022/09/30 20:34:20 cannot get item for session
// --- PASS: TestHandlers (0.04s)
// PASS
// ok      github.com/RaminCH/bookings/internal/handlers   0.050s


//!!!!!!!!!!!!!!
// In order to run main.go
// xxx@xxx:~/go/src/Web_Part1/bookings(master)$ go run cmd/web/*.go
// go run: cannot run *_test.go files (cmd/web/main_test.go)
// xxx@xxx:~/go/src/Web_Part1/bookings(master)$ go run cmd/web/main.go /cmd/web/middleware.go /cmd/web/routes.go
// stat /cmd/web/middleware.go: no such file or directory
// xxx@xxx:~/go/src/Web_Part1/bookings(master)$ go run cmd/web/main.go cmd/web/middleware.go cmd/web/routes.go
// Staring application on port :8080
// ^Csignal: interrupt