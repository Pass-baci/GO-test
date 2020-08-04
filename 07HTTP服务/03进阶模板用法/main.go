package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Inventory struct {
	SKU       string
	Name      string
	UnitPrice float64
	Quantity  int64
}

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	tmp1, err := template.New("test").Parse(`
	//{{$name := "baci"}}
	//{{$name := "abc"}}
	//{{$age := 18}}
	//{{$round2 := true}}
	//Name: {{$name}}
	//Age: {{$age}}
	//Round2: {{$round2}}
	//`)
	//	if err != nil {
	//		fmt.Fprintln(w,"Parse : %v",err)
	//		return
	//	}
	//	err = tmp1.Execute(w, nil)
	//	if err != nil {
	//		fmt.Fprintln(w,"%v",err)
	//		return
	//	}
	//})


//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		tmp1, err := template.New("test").Parse(`
//{{if .yIsZero}}
//除数不能为0
//{{else}}
//	{{.result}}
//{{end}}
//`)
//		if err != nil {
//			fmt.Fprintf(w, "Parse: %v", err)
//			return
//		}
//		x, _ := strconv.ParseInt(r.URL.Query().Get("x"), 10, 64)
//		y, _ := strconv.ParseInt(r.URL.Query().Get("y"), 10, 64)
//		yIsZero := y==0
//		result := 0.0
//		if !yIsZero {
//			result = float64(x) / float64(y)
//		}
//		err = tmp1.Execute(w,map[string]interface{}{
//			"yIsZero": yIsZero,
//			"result": result,
//		})
//		if err != nil {
//			fmt.Fprintf(w, "Execute: %v", err)
//			return
//		}
//	})

//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		tmp1, err := template.New("test").Parse(`
//{{if eq .name1 .name2}}
//名字相同
//{{else}}
//	名字不相同
//{{end}}
//{{if ne .name1 .name2}}
//名字不相同
//{{else}}
//	名字相同
//{{end}}
//{{if gt .age1 .age2}}
//{{.name1}}年龄大
//{{else}}
//	{{.name2}}年龄大
//{{end}}
//`)
//		if err != nil {
//			fmt.Fprintln(w,"Parse err :",err)
//			return
//		}
//		age1, _ := strconv.ParseInt(r.URL.Query().Get("age1"), 10, 64)
//		age2, _ := strconv.ParseInt(r.URL.Query().Get("age2"), 10, 64)
//		name1 := r.URL.Query().Get("name1")
//		name2 := r.URL.Query().Get("name2")
//
//		tmp1.Execute(w, map[string]interface{}{
//			"age1": age1,
//			"age2": age2,
//			"name1": name1,
//			"name2": name2,
//		})
//	})

//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		tmp1, err := template.New("test").Parse(`
//{{range $name, $val := .}}
//	{{$name}}. {{$val}}
//{{end}}
//`)
//		if err != nil {
//			fmt.Fprintln(w,"Parse err:",err)
//		}
//		err = tmp1.Execute(w,map[string]interface{}{
//			"Names": []string{
//				"baci",
//				"Tom",
//				"Tony",
//			},
//			"Numbers": []int{1, 3, 5, 7},
//		})
//		if err != nil {
//			fmt.Fprintln(w,"Execute err:",err)
//		}
//	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp1, err := template.New("test").Parse(`
Inventory
{{- with .Inventory}}
SKU:{{.SKU}}
Name:{{.Name}}
UnitPrice:{{.UnitPrice}}
Quantity:{{.Quantity}}
{{end}}
`)
		if err != nil {
			fmt.Fprintln(w,"Parse err:",err)
			return
		}
		err = tmp1.Execute(w,map[string]interface{}{
			"Inventory": Inventory{
				SKU:       "11000",
				Name:      "Phone",
				UnitPrice: 699.99,
				Quantity:  666,
			},
		})
		if err != nil {
			fmt.Fprintln(w,"Execute err:",err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
