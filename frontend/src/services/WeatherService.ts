import axios, { AxiosResponse } from 'axios';
import { ENDPOINTS } from '../configs/constants';
import { CityDTO } from '../DTOs/CityDTO';
import { WeatherDTO } from '../DTOs/WeatherDTO';

const getWeatherForecast = (cityId: number): Promise<AxiosResponse<WeatherDTO[]>> =>
    axios.get(`${ENDPOINTS.WEATHER}/${cityId}`);

const getCurrentWeather = (cityId: number): Promise<AxiosResponse<WeatherDTO>> =>
    axios.get(`${ENDPOINTS.WEATHER}/today/${cityId}`);

export default {
    getWeatherForecast,
    getCurrentWeather,
};
