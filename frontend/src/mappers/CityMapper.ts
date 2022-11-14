import { CityDTO } from '../DTOs/CityDTO';
import { City } from '../models/City';

const dtoToModel = (dto: CityDTO): City => ({
    id: dto.id,
    name: dto.name,
    coordLat: dto.coordLat,
    coordLon: dto.coordLon,
    population: dto.population,
    timeZone: dto.timeZone,
    sunrise: dto.sunrise,
    sunset: dto.sunset,
    countryName: dto.countryName,
});

const dtoToModelArray = (dtos: CityDTO[]): City[] => dtos.map(dtoToModel);

export default {
    dtoToModel,
    dtoToModelArray,
};
