package main


import(
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
)


type Product struct {
	Name string
}


func main(){

	http.HandleFunc("/createbox", handlerCreateBox)
	http.HandleFunc("/getname", handlerGetName)
	log.Fatal(http.ListenAndServe(":8080", nil))


	// url := "http://localhost:8081/getNameFromDatabase"
	// req, err := http.NewRequest("GET", url, nil)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// res, err := http.DefaultClient.Do(req)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(string(body))
}


func handlerGetName(w http.ResponseWriter, r *http.Request) {

	fmt.Println("handlerGetName was invoked")

	fmt.Println("Calling to backend : http://localhost:8081/getNameFromDatabase")
	url := "http://localhost:8081/getNameFromDatabase"	
	req, err := http.NewRequest(http.MethodGet, url, nil)	


	if err != nil {
		fmt.Println(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}


	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}


	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Println("Success")

		


}

func handlerCreateBox(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handlerCreateBox was invoked")

	fmt.Println("Calling to backend : http://localhost:8081/createBoxToDatabase")


	url := "http://localhost:8081/createBoxToDatabase"	
	req, err := http.NewRequest(http.MethodPost, url, nil)	


	if err != nil {
		fmt.Println(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}


	defer resp.Body.Close()

	fmt.Println("Success")



}