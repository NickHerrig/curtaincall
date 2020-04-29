package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "net/http/cookiejar"
    "net/http/httptest"
    "testing"
)

func newTestApplication(t *testing.T) *application {
    return &application{
        errorLog: log.New(ioutil.Discard, "", 0),
        infoLog:  log.New(ioutil.Discard, "", 0),
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
