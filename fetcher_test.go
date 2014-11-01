package fetcher_test

// Tests for Fetcher
// Users net/http/httptest for server stubbing

import (

    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/mkaz/fetcher"
)

// This tests basic GET request
func TestGetBasic(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func( w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "hola mundo")
    }))
    defer ts.Close()

    f := fetcher.NewFetcher()
    result, err := f.Fetch(ts.URL, "GET")
    if err != nil {
        t.Errorf("Error: %v", err)
    }

    if result != "hola mundo" {
        t.Errorf("Unexpected result: %v", result)
    }
}

// This tests GET request with passing in a parameter
func TestGetParams(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(func( w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, r.FormValue("p"))
    }))
    defer ts.Close()


    f := fetcher.NewFetcher()
    f.Params.Add("p", "hello")
    result, err := f.Fetch(ts.URL, "GET")
    if err != nil {
        t.Errorf("Error: %v", err)
    }

    if result != "hello" {
        t.Errorf("Unexpected result: %v", result)
    }

}
