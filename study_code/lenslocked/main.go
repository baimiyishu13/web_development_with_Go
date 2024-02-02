// 定义代码所属的包
package main

//
import (
	"fmt"
	"net/http"
)

// 基本的 web 应用程序
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 设置 Content-Type 类型
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// w.Header().Set("Content-Type", "/text/plain")
	fmt.Fprint(w, "<h1>Welcometo my great site! </h1>")
}

// 另一个 界面
func cantactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Cantact Page</h1> \n<p>To get in touch: \n<a href=\"https://github.com/baimiyishu13\">github</a>\n</p>\n<p> ")
}

// 构建路由器
// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprint(w, r.URL.Path)
// 	// fmt.Fprint(w, r.URL.RawPath)
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 		return
// 	case "/cantact":
// 		cantactHandler(w, r)
// 		return
// 	default:
// 		// TODO: 找不到的空网页
// 		// http.NotFound(w, r)
// 		// 可以http 包的 常量中查询到
// 		http.Error(w, "page not found", http.StatusNotFound)
// 		// http.Error(w, http.StatusText(404), http.StatusNotFound)
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "page not found")
// 	}
// 	// fmt.Fprint(w, r.URL.Path)
// }

// 新的路由
// type Server struct {
// 	DB *sql.DB
// }

// func (server *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
// 	// s.DB
// }

func fqaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> FQA </h1>")
}

type Router struct {
	name string
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.name = "page not found 3"
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
		return
	case "/cantact":
		cantactHandler(w, r)
		return
	case "/fqa":
		fqaHandler(w, r)
		return
	default:
		http.Error(w, router.name, http.StatusNotFound)
	}
}

/*
http.Hander - 这是一个接口 使用 ServerHttp 处理程序
http.HandleFunc - 接收函数 HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
http.HandlerFunc - type HandlerFunc func(ResponseWriter, *Request)
*/

func main() {
	var router Router
	// http.HandleFunc("/", pathHandler)
	// http.HandlerFunc() // 只是调用了底层的基本函数
	fmt.Println("Starting the server on 3000 ...")
	http.ListenAndServe(":3000", router)
}
