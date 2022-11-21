import { SERVER_URL } from './env';

export const ENDPOINTS = {
    CITY: `${SERVER_URL}/api/cities`,
    WEATHER: `${SERVER_URL}/api/weathers`,
};

export const ROUTES = {
    ROOT: '/',
    HOME: '/home',
};

export const QUERIES = {
    GET_ALL_CITY: 'getAllCity',
    GET_CITY_BY_NAME: 'getCityByName',
    GET_WEATHER_FORECAST: 'getWeatherForecast',
    GET_CURRENT_WEATHER: 'getCurrentWeather',
};
