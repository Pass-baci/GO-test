package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"html/template"
	"log"
	"net/http"
)

func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		tmp1, err := template.New("test").Parse(`
//{{$name1 := "baci"}}
//name1: {{$name1}}
//{{- with true}}
//	{{- $name1 := "Tom"}}
//	{{- $name2 := "Tony"}}
//	name2：{{$name2}}
//{{- end}}
//name1 after with: {{$name1}}
//`)
//		if err != nil {
//			fmt.Fprintln(w,"parse err:",err)
//			return
//		}
//		err = tmp1.Execute(w, nil)
//		if err != nil {
//			fmt.Fprintln(w,"Execute err:",err)
//			return
//		}
//	})
//	log.Println("Starting Http Server...")
//	log.Fatal(http.ListenAndServe(":4000",nil))

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	tmp1 := template.New("test").Funcs(template.FuncMap{
	//		"Add": func(a, b int) int {
	//			fmt.Println(a,b)
	//			return a + b
	//		},
	//	})
	//	_, err := tmp1.Parse("result: {{Add 1 2 | Add 2 | Add 2}}")
	//	if err != nil {
	//		fmt.Fprintln(w,"Parse err:",err)
	//		return
	//	}
	//	err = tmp1.Execute(w, nil)
	//	if err != nil {
	//		fmt.Fprintln(w,"Execute err:",err)
	//		return
	//	}
	//})
	//log.Println("Starting Http Server...")
	//log.Fatal(http.ListenAndServe(":4000",nil))

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	tmp1 := template.New("test").Funcs(template.FuncMap{
	//		"join": strings.Join,
	//	})
//		_, err := tmp1.Parse(`
//{{define "list"}}
//	{{join . ","}}
//{{end}}
//Names: {{template "list" .names }}
//`)
//		if err != nil {
//			fmt.Fprintf(w, "Parse: %v", err)
//			return
//		}
//		err = tmp1.Execute(w, map[string]interface{}{
//			"names": []string{"Alice", "Bob", "Cindy", "David"},
//		})
//		if err != nil {
//			fmt.Fprintf(w, "Execute: %v", err)
//			return
//		}
//	})
//	log.Println("Starting HTTP server...")
//	log.Fatal(http.ListenAndServe("localhost:4000", nil))

	//tmp1, err := template.ParseFiles("./07HTTP服务/04高级模板用法/template_local.tmpl")
	//if err != nil {
	//	log.Fatalf("Parse: %v", err)
	//}
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	// 调用模板对象的渲染方法
	//	err = tmp1.ExecuteTemplate(w,"template_local.tmpl", map[string]interface{}{
	//		"names": []string{"Alice", "Bob", "Cindy", "David"},
	//	})
	//	if err != nil {
	//		fmt.Fprintf(w, "Execute: %v", err)
	//		return
	//	}
	//})
	//
	//log.Println("Starting HTTP server...")
	//log.Fatal(http.ListenAndServe("localhost:4000", nil))

//		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//			// 创建模板对象并添加自定义模板函数
//			tmpl := template.New("test").Funcs(template.FuncMap{
//				"safe": func(s string) template.HTML {
//					return template.HTML(s)
//				},
//			})
//
//			// 解析模板内容
//			_, err := tmpl.Parse(`
//<html>
//<body>
//	<p>{{safe .content}}</p>
//</boyd>
//</html>
//`)
//			if err != nil {
//				fmt.Fprintf(w, "Parse: %v", err)
//				return
//			}
//
//			// 调用模板对象的渲染方法
//			err = tmpl.Execute(w, map[string]interface{}{
//				"content": "<b>Hello world!</b>",
//			})
//			if err != nil {
//				fmt.Fprintf(w, "Execute: %v", err)
//				return
//			}
//		})
//
//		log.Println("Starting HTTP server...")
//		log.Fatal(http.ListenAndServe("localhost:4000", nil))

	p := bluemonday.UGCPolicy()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建模板对象并添加自定义模板函数
		tmpl := template.New("test").Funcs(template.FuncMap{
			"sanitize": func(s string) template.HTML {
				return template.HTML(p.Sanitize(s))
			},
		})

		// 解析模板内容
		_, err := tmpl.Parse(`
<html>
<body>
	<p>{{.content | sanitize}}</p>
</boyd>
</html>
`)
		tmpl1, err := template.ParseFiles("./07HTTP服务/04高级模板用法/template_local.tmpl")
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		// 调用模板对象的渲染方法
		err = tmpl1.Execute(w, map[string]interface{}{
			"content": `<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
