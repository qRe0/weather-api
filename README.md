# Weather Tacker (Go + Echo) by qRe0

## Used technologies
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Echo](https://img.shields.io/badge/Echo-4AE1FF?style=for-the-badge)
![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)

## Overview
This project is an application to show the weather forecast for the specified city. The work is performed by querying the corresponding site using API key. The project implements a client part with a menu interface for entering the city and a weather window interface for providing information. 

## Project structure
```shellweather-api
├── client-part
│   ├── bg.jpeg
│   ├── index.html
│   └── menu.html
├── cmd
│   └── main.go
├── internal
│   ├── api
│   │   ├── api-data-struct
│   │   │   └── apiConfigData.go
│   │   └── api-key-loader
│   │       └── loadApiKey.go
│   ├── middleware
│   │   ├── crud-handler
│   │   │   └── handler.go
│   │   ├── server-response
│   │   │   └── responseHandler.go
│   │   └── weather-query
│   │       └── query.go
│   └── weather
│       └── weather-struct
│           └── weatherData.go
├── LICENSE
├── pkg
│   └── temperature-converter
│       └── kelvin-to-celsius
│           └── kelvinToCelsius.go
└── README.md
```

## Additional notes
* [OpenWeatherMap](https://openweathermap.org/)
* [Echo framework](https://echo.labstack.com/)
