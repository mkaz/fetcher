package fetcher_test

// Tests for Fetcher
//
// Testing requires a web server to response to tests requests.
// If running Apache locally with PHP, create the followiing as
// echo.php in the localhost directory:
//
//      <?php
//
// 	    if ( isset( $_REQUEST['p'] ) ) {
//   		$p = strip_tags( $_REQUEST['p'] );
// 	    }
//
// 	    if ( empty( $p ) ) {
// 		   echo "No P";
// 	    } else {
// 		   echo $p;
// 	    }

import (
    "github.com/mkaz/fetcher"
    "testing"
)

// This tests basic GET request
func TestGetBasic(t *testing.T) {
    url := "http://localhost/echo.php"
    f := fetcher.NewFetcher()
    result, err := f.Fetch(url, "GET")
    if err != nil {
        t.Errorf("Error: %v", err)
    }

    if result != "No P" {
        t.Errorf("Unexpected result: %v", result)
    }
}

// This tests GET request with passing in a parameter
func TestGetParams(t *testing.T) {
    url := "http://localhost/echo.php"
    f := fetcher.NewFetcher()
    f.Params.Add("p", "hello")
    result, err := f.Fetch(url, "GET")
    if err != nil {
        t.Errorf("Error: %v", err)
    }

    if result != "hello" {
        t.Errorf("Unexpected result: %v", result)
    }

}
