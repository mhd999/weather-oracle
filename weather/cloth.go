package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/mhd999/weather-oracle/proto/go/services/weather"
)

type Weather struct {
	Main    map[string]float64       `json:"main"`
	Weather []map[string]interface{} `json:"weather"`
}

type Status struct {
	Temp   int
	Status string // Rain - Snow - Sun
}

type Recomendation struct {
	Temp               int
	RecommendedClothes string
}

func getClothesRecommendation(city string) weather.GetWeatherReply {
	c := make(chan Status, 2)

	go fetchOpenWeather(c, city)
	go getRandomWeather(c)

	status := make([]Status, 0)
	for weather := range c {
		time.Sleep(3 * time.Second)
		status = append(status, Status{
			Status: weather.Status,
			Temp:   weather.Temp,
		})
		if len(c) == 0 {
			close(c)
		}
	}
	likelyness := getLikelyness(status)

	switch likelyness.Status {
	case "Rain":
		return weather.GetWeatherReply{
			Tempreture:         int32(likelyness.Temp),
			ClothRecomendation: "Rain coat",
			Status:             likelyness.Status,
		}
	case "Snow":
		return weather.GetWeatherReply{
			Tempreture:         int32(likelyness.Temp),
			ClothRecomendation: "Biggest jacket you have - gloves - proper shoes",
			Status:             likelyness.Status,
		}
	case "Sun":
		return weather.GetWeatherReply{
			Tempreture:         int32(likelyness.Temp),
			ClothRecomendation: "Shorts and, shade",
			Status:             likelyness.Status,
		}
	}

	return weather.GetWeatherReply{
		Tempreture:         int32(likelyness.Temp),
		ClothRecomendation: "Check the Tempreture",
		Status:             likelyness.Status,
	}
}

func fetchOpenWeather(c chan Status, city string) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, os.Getenv("API_ID"))

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println(err)
	}

	status := Status{
		Temp:   int(math.Round(weather.Main["temp"] - 273.15)), // convert Kelvin into Celsius is C = K - 273.15
		Status: fmt.Sprint(weather.Weather[0]["main"]),
	}

	if err != nil {
		fmt.Println(err)
	} else {
		c <- status
	}
}

func getRandomWeather(c chan Status) {
	status := Status{
		Temp:   3,
		Status: "Rain",
	}

	c <- status
}

func getLikelyness(statuses []Status) Status {
	var temps []int
	var weatherState []string
	for _, status := range statuses {
		temps = append(temps, status.Temp)
		weatherState = append(weatherState, status.Status)
	}
	status := Status{
		Temp:   getAverageTemp(temps),
		Status: getLikelyWeather(weatherState),
	}

	return status

}

func getAverageTemp(temps []int) int {
	total := 0
	for _, temp := range temps {
		total = total + temp
	}
	average := total / len(temps)

	return average
}

// source: https://www.tutorialspoint.com/write-a-golang-program-to-find-duplicate-elements-in-a-given-array
func getLikelyWeather(weatherStatuses []string) string {
	visited := make(map[string]bool)
	for i := 0; i < len(weatherStatuses); i++ {
		if visited[weatherStatuses[i]] == true {
			return weatherStatuses[i]
		} else {
			visited[weatherStatuses[i]] = true
		}
	}
	return "Unknown"

}
