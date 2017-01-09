//package ServiceRepository
package main

import ("io"
	"net/http"
	"ConfigurationRepository"
)

func cfhService(res http.ResponseWriter, req *http.Request){
	io.WriteString(res,"Care Free Home Central Service running")
}

func main(){
	http.HandleFunc("/", cfhService)
	http.ListenAndServe(":8080",nil)
	http.HandleFunc("/saveuser", saveUser)
}

func saveUser(res http.ResponseWriter, req *http.Request){
	ConfigurationRepository.SaveUser(User)
}