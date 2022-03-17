package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

// 如何永久保存,将page的body保存在文本文件中
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) //返回一个错误，将错误交给调用者处理
	//第三个参数表示用户只拥有读写权限
}

// 将保存的文件读取回来
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil { //防止读失败
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// func handler(rw http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(rw, "hi there , i love %s", r.URL.Path[1:])
// }

func viewHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("viewhander:", r.URL.Path) //历史记录
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.NotFound(rw, r)
		return

		// p = &Page{Title: title}
		// t, _ := template.ParseFiles("view.html")
		// t.Execute(rw, p)

	}
	renderTemplate(rw, "view", p)
	// fmt.Fprintf(rw, "<h1>%s</h1><div>%s<div>", p.Title, p.Body) //将读到的内容以html形式写入rw中
}

// 以下是编辑页面
// 编辑页面需要编辑与保存
//编辑
func editHandler(rw http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(rw, "edit", p) //执行p到rw内
}

// 保存
func saveHandler(rw http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	fmt.Fprintf(rw, "success, page %s has been  stored", title)
}

//重构
func renderTemplate(rw http.ResponseWriter, templ string, p *Page) {
	t, _ := template.ParseFiles(templ + ".html")
	t.Execute(rw, p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
