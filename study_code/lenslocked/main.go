// 定义代码所属的包
package main

//
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

// 解析 - 执行 gohtml文件
func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "The wars an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "The wars an error Executing the template", http.StatusInternalServerError)
		return
	}
}

// 基本的 web 应用程序
func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "./templates/home.gohtml")
}

// 另一个 界面
func cantactHandler(w http.ResponseWriter, r *http.Request) {
	// executeTemplate(w, "./templates/contact.gohtml")
	tplpath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplpath)
}

// FQA 界面
func fqaHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "./templates/fqa.gohtml")
}

// Gin

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/cantact", cantactHandler)
	r.Get("/fqa", fqaHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("404"))
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on 3000 ...")
	http.ListenAndServe(":3000", r)
}
