let city = "";
let temperature = 1.0;
let weatherDescription = "";

temperature = temperature.toFixed(1);

document.getElementById("city").innerText = decodeURIComponent(city);
document.getElementById("temperature").innerText = temperature;
document.getElementById("weatherDescription").innerText = weatherDescription;

function adjustFontSize() {
    let cityName = document.getElementById("city");
    let windowWidth = window.innerWidth;

    let baseFontSizeCity = 3;
    let baseFontSizeTemperature = 2.5;

    let vwUnit = windowWidth / 100;

    let cityFontSize = baseFontSizeCity + vwUnit;

    let temperatureFontSize = baseFontSizeTemperature + vwUnit;

    cityName.style.fontSize = "calc(" + cityFontSize + "px + 1vmin)";
    document.getElementById("temperature").style.fontSize = "calc(" + temperatureFontSize + "px + 1vmin)";
}

adjustFontSize();
window.addEventListener('resize', adjustFontSize);
