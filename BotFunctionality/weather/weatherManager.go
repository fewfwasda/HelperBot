package weather

import (
	weather "HelperBot/Data/emojiMap"
	texts "HelperBot/Data/textsUI"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

const storagePath = "C:\\Users\\user\\Desktop\\MyProject\\Go\\HelperBot\\UserDataStorage\\userCities.json"

var (
	userCities = make(map[int64]string)
	mu         sync.Mutex
)

type Weather struct {
	City    string `json:"name"`
	Message string `json:"message"`
	Main    struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

func GetWeather(city string) (emoji string, desc string, temp float64, err error) {
	openWeatherAPIKey := os.Getenv("OPENWEATHER_API_KEY")
	if openWeatherAPIKey == "" {
		log.Fatalf("API is not set")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&lang=ru&appid=%s", city, openWeatherAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", "", 0, err
	}
	defer resp.Body.Close()

	var result Weather
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", "", 0, err
	}
	if result.Message == "city not found" {
		return "", "", 0, errors.New(texts.ErrCouldNotFindCity)
	}

	main := result.Weather[0].Main
	temp = result.Main.Temp
	desc = result.Weather[0].Description

	emoji = weather.EmojiMap[main]

	return emoji, desc, temp, nil
}
func ReverseGeocode(lat, lon float64) (string, error) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatalf(texts.ErrAPIKey)
	}

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%.6f&lon=%.6f&appid=%s&units=metric&lang=ru", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result Weather

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.City, nil
}
