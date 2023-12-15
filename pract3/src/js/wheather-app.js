const WEATHER_API_KEY = '2954094ae71e46ef0b0bfc1905b15d7f';

var settings = {
    "url": `https://api.openweathermap.org/data/2.5/weather?q=Moscow&units=metric&appid=${WEATHER_API_KEY}`,
    "method": "GET",
    "timeout": 0,
};

$.ajax(settings).done(function (response) {
    console.log("Full response: ", response);
    var [main, wind] = [response.main, response.wind];
    $("#root").append(`
        <div>
            <p>Average temprature: ${main.temp}</p>
            <p>Minimum temprature: ${main.temp_min}</p>
            <p>Maximum temprature: ${main.temp_max}</p>
        </div>
        <p>Pressure: ${main.pressure}</p>
        <div>
            <p>Wind speed: ${wind.speed}</p>
            <p>Wind deg: ${wind.deg}</p>
            <p>Wind gust: ${wind.gust}</p>
        </div>
    `);
}).catch(() => {  // If city name is NOT valid
    $("#root").append(`<p>City not found</p>`);
});