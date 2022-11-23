import { useQuery } from 'react-query';
import { QUERIES } from '../configs/constants';
import WeatherMapper from '../mappers/WeatherMapper';
import { Weather } from '../models/Weather';
import WeatherService from '../services/WeatherService';

export const useGetWeatherForecastForCity = (cityId: number): Weather[] | undefined =>
    useQuery(QUERIES.GET_WEATHER_FORECAST, async () =>
        WeatherService.getWeatherForecast(cityId).then((res) => WeatherMapper.dtoToModelArray(res.data))
    ).data;

export const useGetTodayWeatherForCity = (cityId: number): Weather | undefined =>
    useQuery(QUERIES.GET_CURRENT_WEATHER, async () =>
        WeatherService.getCurrentWeather(cityId).then((res) => WeatherMapper.dtoToModel(res.data))
    ).data;
