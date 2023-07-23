package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showWeatherApp() {

	// Getting weather API
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Bangalore&appid=08c6511dc9a92503249df0fdd7b48939")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Image will fit based on its original size
	img := canvas.NewImageFromFile("weathers.png")
	img.FillMode = canvas.ImageFillOriginal

	// To display weather details
	label1 := canvas.NewText("Weather Details", color.Opaque)
	label1.TextStyle = fyne.TextStyle{Bold: true}

	// Creating separate canvas.Text objects for each weather detail
	label2 := canvas.NewText(fmt.Sprintf("Country: %s", weather.Sys.Country), color.Black)
	label3 := canvas.NewText(fmt.Sprintf("Wind Speed: %.2f", weather.Wind.Speed), color.Black)
	label4 := canvas.NewText(fmt.Sprintf("Humidity: %d", weather.Main.Humidity), color.Black)
	label5 := canvas.NewText(fmt.Sprintf("Temperature: %.2f", weather.Main.Temp), color.Black)
	label6 := canvas.NewText(fmt.Sprintf("Pressure: %d", weather.Main.Pressure), color.Black)
	label7 := canvas.NewText(fmt.Sprintf("Latitude: %.2f", weather.Coord.Lat), color.Black)
	label8 := canvas.NewText(fmt.Sprintf("Longitude: %.2f", weather.Coord.Lon), color.Black)

	// Creating a VBox container to hold all the weather details
	weatherDetailsContainer := container.NewVBox(
		label1,
		img,
		label2,
		label3,
		label4,
		label5,
		label6,
		label7,
		label8,
	)

	w := myApp.NewWindow("Weather App")
	w.Resize(fyne.NewSize(500, 280))

	w.SetContent(weatherDetailsContainer)
	w.Show()
}

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
