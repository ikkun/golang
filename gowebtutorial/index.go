package main

//go get github.com/gorilla/mux
import (
	"fmt"
	"net/http"

	"html/template"

	"github.com/gorilla/mux"

	"io"
	"os"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	var templates = template.Must(template.ParseFiles("index.html"))
	router := mux.NewRouter()
	userDB := map[string]int{
		"kreangkrai": 37,
		"ikkun":      50,
		"annie":      40,
		"or":         17,
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Kreangkrai")
		// http.ServeFile(w, r, "index.html")
		myProduct := Product{"นมสด", 500}
		templates.ExecuteTemplate(w, "index.html", myProduct)
	})

	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Kreangkrai")
		http.ServeFile(w, r, "signup.html")
	})

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Kreangkrai")
		http.ServeFile(w, r, "login.html")

	})

	router.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Kreangkrai")
		http.ServeFile(w, r, "upload.html")

	})

	router.HandleFunc("/uploadHandle", uploadHandle)

	router.HandleFunc("/authen", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method", r.Method)
		r.ParseForm()
		fmt.Println("Username:", r.Form["username"])
		fmt.Println("Password:", r.Form["password"])
	})

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Kreangkrai")
	// })

	http.HandleFunc("/product", product)
	// http.HandleFunc("/user/", user)
	router.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		name := vars["name"]
		// name := r.URL.Path[len("/user/"):]
		age := userDB[name]
		fmt.Fprintf(w, "My Name is %s %d year old", name, age)

	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product Request")
}

func user(w http.ResponseWriter, r *http.Request) {
	userDB := map[string]int{
		"kreangkrai": 37,
		"ikkun":      50,
		"annie":      40,
		"or":         17,
	}
	vars := mux.Vars(r)
	name := vars["name"]
	// name := r.URL.Path[len("/user/"):]
	age := userDB[name]
	fmt.Fprintf(w, "My Name is %s %d year old", name, age)
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%v", handle.Header)
	f, err := os.OpenFile("./file/"+handle.Filename, os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintf(w, "Upload Complete")
}
