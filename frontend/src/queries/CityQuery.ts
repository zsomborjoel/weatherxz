import { useQuery, UseQueryResult } from 'react-query';
import { QUERIES } from '../configs/constants';
import CityMapper from '../mappers/CityMapper';
import { City } from '../models/City';
import CityService from '../services/CityService';

export const useGetAllCity = (): City[] | undefined =>
    useQuery(QUERIES.GET_ALL_CITY, async () => CityService.getAll().then((res) => CityMapper.dtoToModelArray(res.data)))
        .data;

export const useGetCityByName = (name: string): City | undefined =>
    useQuery(QUERIES.GET_CITY_BY_NAME, async () =>
        CityService.getByName(name).then((res) => CityMapper.dtoToModel(res.data))
    ).data;
