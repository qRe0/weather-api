# Weather Tacker (Go + Echo) by qRe0

## Used technologies
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Echo](https://img.shields.io/badge/Echo-4AE1FF?style=for-the-badge)

## Overview

## Project structure
```shell
weather-api
├── cmd
│   └── main.go
├── internal
│   ├── api
│   │   ├── api-data-struct
│   │   │   └── apiConfigData.go
│   │   └── api-key-loader
│   │       └── loadApiKey.go
│   ├── middleware
│   │   ├── server-response
│   │   │   └── responseHandler.go
│   │   └── weather-query
│   │       └── query.go
│   └── weather
│       └── weather-struct
│           └── weatherData.go
├── pkg
│   └── temperature-converter
│       └── kelvin-to-celsius
│           └── kelvinToCelsius.go
└── README.md
```

## Additional notes
