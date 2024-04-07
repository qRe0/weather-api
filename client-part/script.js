var city = "";
var temperature = 1.0;
var weatherDescription = "";

temperature = temperature.toFixed(1);

document.getElementById("city").innerText = decodeURIComponent(city);
document.getElementById("temperature").innerText = temperature;
document.getElementById("weatherDescription").innerText = weatherDescription;

function adjustFontSize() {
    var cityName = document.getElementById("city");
    var windowWidth = window.innerWidth;

    var baseFontSizeCity = 3;
    var baseFontSizeTemperature = 2.5;

    var vwUnit = windowWidth / 100;

    var cityFontSize = baseFontSizeCity + vwUnit;

    var temperatureFontSize = baseFontSizeTemperature + vwUnit;

    cityName.style.fontSize = "calc(" + cityFontSize + "px + 1vmin)";
    document.getElementById("temperature").style.fontSize = "calc(" + temperatureFontSize + "px + 1vmin)";
}

adjustFontSize();
window.addEventListener('resize', adjustFontSize);
