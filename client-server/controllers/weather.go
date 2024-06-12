package controllers

import (
	"client-server/helpers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// {
//     "err": 0,
//     "errmsg": null,
//     "data": {
//         "weather": {
//             "timestamp": 1716660004,
//             "cloud": 0.193,
//             "wind_speed": 2,
//             "wind_direction": 4,
//             "wind_gustiness": 0.19,
//             "rain": 2,
//             "rain_intensity": 0.319,
//             "fog": 0.008,
//             "temp": 21,
//             "pressure": 766,
//             "date": "2024-05-25",
//             "time": "2024-05-25 21:00:04"
//         },
//         "date": "2024-05-25",
//         "time": "10:21:40",
//         "acceleration": 7,
//         "season": 3
//     }
// }

type GetWeatherResponseDataWeather struct {
	Timestamp     int64   `json:"timestamp"`
	Cloud         float64 `json:"cloud"`
	WindSpeed     int     `json:"wind_speed"`
	WindDirection int     `json:"wind_direction"`
	WindGustiness float64 `json:"wind_gustiness"`
	Rain          int     `json:"rain"`
	RainIntensity float64 `json:"rain_intensity"`
	Fog           float64 `json:"fog"`
	Temp          int     `json:"temp"`
	Pressure      int     `json:"pressure"`
	Date          string  `json:"date"`
	Time          string  `json:"time"`
}

type GetWeatherResponseData struct {
	Weather      GetWeatherResponseDataWeather `json:"weather"`
	Date         string                        `json:"date"`
	Time         string                        `json:"time"`
	Acceleration int                           `json:"acceleration"`
	Season       int                           `json:"season"`
}

func GetWeather(c *gin.Context) {
	currentTime := time.Now()

	helpers.JSONResponse(c, http.StatusOK, "", GetWeatherResponseData{
		Weather: GetWeatherResponseDataWeather{
			Timestamp:     currentTime.Unix(),
			Cloud:         0.193,
			WindSpeed:     2,
			WindDirection: 4,
			WindGustiness: 0.19,
			Rain:          2,
			RainIntensity: 0.319,
			Fog:           0.008,
			Temp:          21,
			Pressure:      766,
			Date:          currentTime.Format("2006-01-02"),
			Time:          currentTime.Format("2006-01-02 15:04:05"),
		},
		Date:         currentTime.Format("2006-01-02"),
		Time:         currentTime.Format("15:04:05"),
		Acceleration: 7,
		Season:       3,
	},
	)
}
