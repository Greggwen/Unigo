package route

import "net/http"

type Handle func(http.ResponseWriter, *http.Request, Params)

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (ps Params) GName(name string) string {
	for i := range ps {
		if ps[i].Key == name {
			return ps[i].Value
		}
	}
	return ""
}

// 实现 http.Handler接口的 ServeHttp 方法
type RouteRegister struct {
}

func (rh RouteRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/rose" {
		w.Write([]byte("Hello Rose, Go Http rose"))
		return
	}

	if r.URL.Path == "/swen" {
		w.Write([]byte("Hello Swen!Change you!"))
		return
	}

	w.Write([]byte("Hello World, This is the homepage! Welcome using Go!"))
}
