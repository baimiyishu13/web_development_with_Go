// 定义代码所属的包
package main

//
import (
	"fmt"
	"net/http"
	"os"
)

// 基本的 web 应用程序
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcometo my awesome site</h1>")
}
func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Fprintln(os.Stdout, "Hello World!")
	fmt.Println("Starting the server on 3000 ...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
