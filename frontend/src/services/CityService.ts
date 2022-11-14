import axios, { AxiosResponse } from 'axios';
import { ENDPOINTS } from '../configs/constants';
import { CityDTO } from '../DTOs/CityDTO';

const getAll = (): Promise<AxiosResponse<CityDTO[]>> => axios.get(ENDPOINTS.CITY);

const getByName = (name: string): Promise<AxiosResponse<CityDTO>> => axios.get(`${ENDPOINTS.CITY}/${name}`);

export default {
    getAll,
    getByName,
};
