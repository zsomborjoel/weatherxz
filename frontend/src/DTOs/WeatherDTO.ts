export interface WeatherDTO {
    id: number;
    dateTime: number;
    temp: number;
    feelsLike: number;
    tempMin: number;
    tempMax: number;
    pressure: number;
    preassureSeaLevel: number;
    preassureGroundLevel: number;
    humidity: number;
    weatherConditionMain: string;
    weatherConditionDesc: string;
    clouds: number;
    windSpeed: number;
    windDeg: number;
    windGust: number;
    visibility: number;
    probabilityOfPrecipitation: number;
    rainVolume: number;
    snowVolume: number;
    partOfDay: string;
    cityId: number;
}
