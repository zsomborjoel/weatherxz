import { CityDTO } from '../DTOs/CityDTO';
import { WeatherDTO } from '../DTOs/WeatherDTO';
import { Weather } from '../models/Weather';

const dtoToModel = (dto: WeatherDTO): Weather => ({
    id: dto.id,
    dateTime: dto.dateTime,
    temp: dto.temp,
    feelsLike: dto.feelsLike,
    tempMin: dto.tempMin,
    tempMax: dto.tempMax,
    pressure: dto.pressure,
    preassureSeaLevel: dto.preassureSeaLevel,
    preassureGroundLevel: dto.preassureGroundLevel,
    humidity: dto.humidity,
    weatherConditionMain: dto.weatherConditionMain,
    weatherConditionDesc: dto.weatherConditionDesc,
    clouds: dto.clouds,
    windSpeed: dto.windSpeed,
    windDeg: dto.windDeg,
    windGust: dto.windGust,
    visibility: dto.visibility,
    probabilityOfPrecipitation: dto.probabilityOfPrecipitation,
    rainVolume: dto.rainVolume,
    snowVolume: dto.snowVolume,
    partOfDay: dto.partOfDay,
    cityId: dto.cityId,
});

const dtoToModelArray = (dtos: WeatherDTO[]): Weather[] => dtos.map(dtoToModel);

export default {
    dtoToModel,
    dtoToModelArray,
};
