/* eslint-disable jsx-a11y/label-has-associated-control */
import { FC, useContext, useEffect, useState } from 'react';
import LoadingIndicator from '../../components/LoadingIndicator';
import WeatherTypeImg from '../../components/WeatherTypeImg';
import { CityContext } from '../../contexts/CityContext';
import { City } from '../../models/City';
import { useGetAllCity } from '../../queries/CityQuery';
import { useGetTodayWeatherForCity } from '../../queries/WeatherQuery';
import DateUtil from '../../utils/DateUtil';

export type Props = {};

const CityDetails: FC<Props> = () => {
    const [curretCity, setCurrentCity] = useState<City | undefined>();
    const {
        cityIdState: [cityId, setCityId],
    } = useContext(CityContext);

    const cities = useGetAllCity();
    const { data: todaysWeather, refetch: todaysWeatherRefetch } = useGetTodayWeatherForCity(cityId);

    useEffect(() => {
        const foundCity = cities?.filter((city) => city.id === cityId);
        setCurrentCity(foundCity?.[0]);
    }, [cityId, cities]);

    useEffect(() => {
        todaysWeatherRefetch();
    }, [cityId]);

    const findSelectedCityId = (name: string): number | undefined => cities?.filter((c) => c.name === name)[0].id;

    if (cities === undefined && todaysWeather === undefined) {
        return <LoadingIndicator />;
    }

    return (
        <div className="flex justify-center bg-slate-700">
            <div className="m-5 w-full">
                <div className="h-1/6 text-center">
                    <label htmlFor="cities" className="block mb-2 text-sm font-medium text-white">
                        Select an option
                    </label>
                    <select
                        id="cities"
                        className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder:text-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                        onChange={(e) => setCityId(findSelectedCityId(e.target.value)!)}
                    >
                        <option selected>Choose a country</option>
                        {cities?.map((c) => (
                            <option value={c.name}>{c.name}</option>
                        ))}
                    </select>
                </div>
                <p className="h-1/6 flex justify-center">
                    <WeatherTypeImg value={todaysWeather?.weatherConditionMain} />
                </p>
                <div className="h-1/6 text-center flex justify-center">
                    <p className="font-extrabold text-center text-9xl text-white">{todaysWeather?.temp}</p>
                    <p className="font-light text-center text-6xl text-gray-500">°c</p>
                </div>
                <p className="font-extrabold h-1/6 text-center text-2xl text-gray-500">
                    {todaysWeather?.weatherConditionDesc}
                </p>
                <p className="h-1/6 text-center text-gray-500">{`Today - ${DateUtil.getTodaysDate()}`}</p>
                <p className="h-1/6 text-center text-gray-500">{`${curretCity?.name} / Country`}</p>
            </div>
        </div>
    );
};

export default CityDetails;
