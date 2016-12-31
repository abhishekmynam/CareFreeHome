package ConfigurationRepository

import (
	"fmt"
	"os"
	"time"
)

func LogError(message string){
	var today = time.Now().Format("01-02-2006")
	var ErrorPath = "C:/Users/abhis/OneDrive/Documents/GitHub/CareFreeHomeLogs/CFHErrorLog" +today+".txt"
	createFile(ErrorPath)
	writeFile(ErrorPath, message)
}

func LogProcess(message string){
	var LogProcessActive = "False"
	if(LogProcessActive =="True") {
		var today = time.Now().Format("01-02-2006")
		var ProcessLogPath = "C:/Users/abhis/OneDrive/Documents/GitHub/CareFreeHomeLogs/CFHProcessLog" + today + ".txt"
		createFile(ProcessLogPath)
		writeFile(ProcessLogPath, message)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

}
func createFile(path string) {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err)
		defer file.Close()
	}
}

func writeFile(path string, message string) {
	var file, err = os.OpenFile(path, os.O_APPEND, 0644)
	checkError(err)
	var message1 = "\n " + message
	_, err = file.WriteString(message1)
	checkError(err)
	err = file.Sync()
	checkError(err)
}