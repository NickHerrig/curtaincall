package main

import (
    "html"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/cookiejar"
    "net/http/httptest"
    "net/url"
    "regexp"
    "testing"
    "time"

    "curtaincall.tech/pkg/models/mock"

    "github.com/golangcollege/sessions"
)

var csrfTokenRX = regexp.MustCompile(`<input type='hidden' name='csrf_token' value='(.+)'>`)

func extractCSRFToken(t *testing.T, body []byte) string {
    matches := csrfTokenRX.FindSubmatch(body)
    if len(matches) < 2 {
        t.Fatal("No csrf token found in body")
    }

    return html.UnescapeString(string(matches[1]))
}

func newTestApplication(t *testing.T) *application {
    templateCache, err := newTemplateCache("./../../ui/html/")
    if err != nil{
        t.Fatal(err)
    }

    session := sessions.New([]byte("akdjiekdjfldjfhuejdkiofadsalfjckj"))
    session.Lifetime = 12 * time.Hour
    session.Secure = true

    return &application{
        errorLog:      log.New(ioutil.Discard, "", 0),
        infoLog:       log.New(ioutil.Discard, "", 0),
        session:       session,
        theaters:      &mock.TheaterModel{},
        templateCache: templateCache,
        users:         &mock.UserModel{},
    }
}

type testServer struct {
    *httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
    ts := httptest.NewTLSServer(h)

    jar, err := cookiejar.New(nil)
    if err != nil {
        t.Fatal(err)
    }

    ts.Client().Jar = jar

    ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse
    }

    return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
    rs, err := ts.Client().Get(ts.URL + urlPath)
    if err != nil {
        t.Fatal(err)
    }

    defer rs.Body.Close()
    body, err := ioutil.ReadAll(rs.Body)
    if err != nil {
        t.Fatal(err)
    }
    return rs.StatusCode, rs.Header, body
}

func (ts *testServer) postForm(t *testing.T, urlPath string, form url.Values) (int, http.Header, []byte) {
    rs, err := ts.Client().PostForm(ts.URL + urlPath, form)
    if err != nil {
        t.Fatal(err)
    }

    defer rs.Body.Close()
    body, err := ioutil.ReadAll(rs.Body)
    if err != nil {
        t.Fatal(err)
    }
    return rs.StatusCode, rs.Header, body
}
