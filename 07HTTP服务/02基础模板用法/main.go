package main

import (
	"github.com/golang/go/src/pkg/fmt"
	"github.com/golang/go/src/pkg/html/template"
	"github.com/golang/go/src/pkg/strconv"
	"log"
	"net/http"
)

//type Inventory struct {
//	SKU       string
//	Name      string
//	UnitPrice float64
//	Quantity  int64
//}

//func (i *Inventory)Subtotal() float64 {
//	return i.UnitPrice * float64(i.Quantity)
//}

func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	parse, err := template.New("123").Parse("Inventory SKU: {{.SKU}} Name: {{.Name}} UnitPrice: {{.UnitPrice}} Quantity: {{.Quantity}} Subtotal:{{.Subtotal}}")
	//	if err != nil {
	//		fmt.Fprintln(writer,"%v",err)
	//		return
	//	}
	//	inventory := &Inventory{
	//		SKU:  request.URL.Query().Get("sku"),
	//		Name: request.URL.Query().Get("name"),
	//	}
	//	inventory.UnitPrice, _ = strconv.ParseFloat(request.URL.Query().Get("unitPrice"), 64)
	//	inventory.Quantity, _ = strconv.ParseInt(request.URL.Query().Get("quantity"), 10, 64)
	//	err = parse.Execute(writer,inventory)
	//	if err != nil {
	//		fmt.Fprintf(writer, "Execute: %v", err)
	//		return
	//	}
	//})
	//log.Println("Starting HTTP server...")
	//log.Fatal(http.ListenAndServe("localhost:4000", nil))

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte(request.URL.Query().Get("Val")))
	//})
	//log.Println("Starting HTTP server...")
	//log.Fatal(http.ListenAndServe("localhost:4000", nil))

	//使用map作为一个临时的根对象
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp1, err := template.New("test").Parse("Inventory SKU: {{.SKU}} Name: {{.Name}} UnitPrice: {{.UnitPrice}} Quantity: {{.Quantity}}")
		if err != nil {
			fmt.Fprintln(w,"Parse: %v", err)
			return
		}
		sku := r.URL.Query().Get("sku")
		name := r.URL.Query().Get("name")
		unitPrice, _ := strconv.ParseFloat(r.URL.Query().Get("unitPrice"), 64)
		quantity, _ := strconv.ParseInt(r.URL.Query().Get("quantity"),10,64)
		err = tmp1.Execute(w, map[string]interface{}{
			"SKU":       sku,
			"Name":      name,
			"UnitPrice": unitPrice,
			"Quantity":  quantity,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
