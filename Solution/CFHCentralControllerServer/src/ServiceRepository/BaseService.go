//package ServiceRepository
package main

import ("io"
	"net/http"
)

func CFHService(res http.ResponseWriter, req *http.Request){
	io.WriteString(res,"Care Free Home Central Service running")
}

func main(){
	http.HandleFunc("/", CFHService)
	http.ListenAndServe(":8080",nil)
}