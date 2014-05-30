
# fetcher.go

Marcus Kazmierczak, [@mkaz](https://twitter.com/mkaz)

A client library to help make HTTP requests a little easier in Go.

In general basic requests aren't too bad using the `net/http` library. A little extra boilerplate just to get the results, but workable.

However, when working and testing a REST API which required setting headers, uploading images and creating more complex requests, the basic `net/http` package becomes a bit challenging and verbose.


## Install

```
$ go get github.com/mkaz/fetcher
```


## Usage

### GET Example

Here's a basic example using fetcher and GET request

```go
import "github.com/mkaz/fetcher"

f := fetcher.NewFetcher()
result, err := f.Fetch("http://en.gravatar.com/mkaz.json", "GET")
if err != nil {
    fmt.Println("Error:". err)
}

fmt.Println(result)
```

### POST Example

Example using fetcher to POST params to a form

```go
f := fetcher.NewFetcher()
f.Params["foo"] = "bar"
f.Params["baz"] = "foz"
f.Fetch("/post-form", "POST")
```

### File Upload Example

Example using fetcher to upload files, set parameters and header variable

```go
f := fetcher.NewFetcher()
f.Params["foo"] = "bar"
f.Params["baz"] = "foz"
f.Files["filedata"] = "/home/mkaz/tmp/upload.jpg"
f.Header["X-Auth"] = "my-secret-token"
f.Fetch("/upload-file", "POST")
```

## License

This software is licensed under the MIT License.

