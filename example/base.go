package main

import (
	"fmt"
	"net/http"

	"github.com/mozillazg/request"
)

func get(a *request.Args) {
	resp, err := request.Get("http://httpbin.org/get", a)
	defer resp.Body.Close()
	if err == nil {
		fmt.Println(resp.Ok())
		fmt.Println(resp.Reason())
	}
	d, _ := resp.Json()
	fmt.Println(d.Get("url"))
	fmt.Println(d.Get("args"))
}

func head(a *request.Args) {
	resp, err := request.Head("http://httpbin.org/get", a)
	if err == nil {
		fmt.Println(resp.Ok())
		fmt.Println(resp.Reason())
	}
	defer resp.Body.Close()
}

func json(a *request.Args) {
	resp, err := request.Get("http://httpbin.org/get", a)
	if err != nil {
		return
	}

	d, err := resp.Json()
	if err != nil {
		return
	}
	fmt.Println(d.Get("headers").Get("User-Agent"))
	defer resp.Body.Close()
}

func gzip(a *request.Args) {
	resp, err := request.Get("http://httpbin.org/gzip", a)
	if err != nil {
		return
	}

	d, err := resp.Json()
	if err != nil {
		return
	}
	fmt.Println(d.Get("headers").Get("Accept-Encoding"))
	fmt.Println(resp.Header.Get("Content-Encoding"))
	s, err := resp.Text()
	fmt.Println(s)
	defer resp.Body.Close()
}

func post(a *request.Args) {
	resp, err := request.Post("http://httpbin.org/post", a)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(resp.Ok())
	d, err := resp.Json()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(d.Get("headers").Get("Content-Type"))
	fmt.Println(d.Get("form"))
	fmt.Println(d.Get("url"))
	fmt.Println(d.Get("args"))
}

func customHeaders(a *request.Args) {
	a.Headers = map[string]string{
		"Accept-Encoding": "gzip,deflate,sdch",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	}
	resp, err := request.Get("http://httpbin.org/get", a)
	defer resp.Body.Close()
	if err == nil {
		fmt.Println(resp.Ok())
	}
	d, err := resp.Json()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(d.Get("headers").Get("User-Agent"))
	fmt.Println(d.Get("headers").Get("Accept-Encoding"))
	fmt.Println(d.Get("headers").Get("Accept"))
}

func put(a *request.Args) {
	resp, _ := request.Put("http://httpbin.org/put", a)
	defer resp.Body.Close()

	fmt.Println(resp.Ok())
	d, _ := resp.Json()
	fmt.Println(d.Get("headers").Get("Content-Type"))
	fmt.Println(d.Get("form"))
}

func patch(a *request.Args) {
	resp, _ := request.Patch("http://httpbin.org/patch", a)
	defer resp.Body.Close()

	fmt.Println(resp.Ok())
	d, _ := resp.Json()
	fmt.Println(d.Get("headers").Get("Content-Type"))
	fmt.Println(d.Get("form"))
}

func deleteF(a *request.Args) {
	resp, _ := request.Delete("http://httpbin.org/delete", a)
	defer resp.Body.Close()

	fmt.Println(resp.Ok())
	d, _ := resp.Json()
	fmt.Println(d.Get("headers").Get("Content-Type"))
	fmt.Println(d.Get("form"))
}

func options(a *request.Args) {
	resp, _ := request.Options("http://httpbin.org/get", a)
	defer resp.Body.Close()

	fmt.Println(resp.Ok())
	fmt.Println("Allow:", resp.Header.Get("Allow"))
}

func getParams(a *request.Args) {
	resp, err := request.Get("http://httpbin.org/get?foobar=123", a)
	defer resp.Body.Close()
	if err == nil {
		fmt.Println(resp.Ok())
		fmt.Println(resp.Reason())
	}
	d, _ := resp.Json()
	fmt.Println(d.Get("url"))
	fmt.Println(d.Get("args"))
}

func cookies(a *request.Args) {
	resp, _ := request.Get("http://httpbin.org/cookies", a)
	defer resp.Body.Close()
	d, _ := resp.Json()
	fmt.Println(d.Get("cookies"))
	fmt.Println(resp.Cookies())
	url := resp.Request.URL
	fmt.Println(a.Client.Jar.Cookies(url))
}

func main() {
	c := &http.Client{}
	a := request.NewArgs(c)
	//
	// fmt.Println("=====GET: ")
	// get(a)
	// fmt.Println("=====HEAD: ")
	// head(a)
	// fmt.Println("=====JSON: ")
	// json(a)
	// fmt.Println("=====GZIP: ")
	// gzip(a)
	//
	// a.Data = map[string]string{
	// 	"key": "value",
	// 	"a":   "123",
	// }
	// fmt.Println("=====POST: ")
	// post(a)
	//
	// fmt.Println("=====Custom Headers: ")
	// customHeaders(a)
	//
	// a = request.NewArgs(c)
	// a.Data = map[string]string{
	// 	"key": "value",
	// 	"a":   "123",
	// }
	// fmt.Println("=====PUT: ")
	// put(a)
	// fmt.Println("=====PATCH: ")
	// patch(a)
	// fmt.Println("=====DELTE: ")
	// deleteF(a)
	// fmt.Println("=====OPTIONS: ")
	// options(a)
	//
	// a = request.NewArgs(c)
	// a.Params = map[string]string{
	// 	"a":   "abc",
	// 	"key": "value",
	// }
	// fmt.Println("=====Params: ")
	// get(a)
	// getParams(a)
	// post(a)

	fmt.Println("=====Cookies: ")
	a = request.NewArgs(c)
	cookies(a)
	a.Cookies = map[string]string{
		"name": "value",
		"foo":  "bar",
	}
	cookies(a)
}