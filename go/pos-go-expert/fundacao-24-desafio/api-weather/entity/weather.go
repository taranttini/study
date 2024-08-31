package entity

type Weather struct {
	Location WeatherLocation `json:"location"`
	Current  WeatherCurrent  `json:"current"`
}

type WeatherLocation struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzId           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type WeatherCurrent struct {
	TempCelcius    float64 `json:"temp_c"`
	TempFahrenheit float64 `json:"temp_f"`
}
