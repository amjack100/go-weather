package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiUrl string = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s"
const apiKey string = "01341bed01b00d329e3f312b510df4b2"
const cityName string = "Indianapolis"
const units string = "imperial"

func makeRequest() string {
	
	out := fmt.Sprintf(apiUrl, cityName, apiKey, units)
	resp, err := http.Get(out)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("%s", err)
	}
	return string(body)

}

func main() {

	text := makeRequest()
	bytes := []byte(text)

	var objmap map[string]interface{}
	

	json.Unmarshal(bytes, &objmap)

	MainValues, _ := objmap["main"]

	WeatherInfo := objmap["weather"].([]interface{})[0]

	fmt.Printf("Main description: %v\n", WeatherInfo.(map[string]interface{})["main"]) 
	fmt.Printf("Sub-description: %v\n", WeatherInfo.(map[string]interface{})["description"]) 
	fmt.Printf("Temperature: %v\n", MainValues.(map[string]interface{})["temp"])
	fmt.Printf("Feels like: %v\n", MainValues.(map[string]interface{})["feels_like"])

}
