package main
import "fmt"
import "net/http"
import "./model"

func main(){
	// fs:=http.FileServer(http.Dir("public"))
	// http.Handle("/",fs)
	// http.ListenAndServe(":8080",nil)
	http.HandleFunc("/saludar",saludar)
	http.ListenAndServe(":8080",nil)

	

}

func saludar(w http.ResponseWriter,r *http.Request){
	persons := make(map[int]Person)
	fmt.Fprintf(w, persons)	
}