import { FC, useContext, useEffect, useState } from 'react';
import LoadingIndicator from '../../components/LoadingIndicator';
import { CityContext } from '../../contexts/CityContext';
import { City } from '../../models/City';
import { useGetAllCity } from '../../queries/CityQuery';
import { useGetTodayWeatherForCity } from '../../queries/WeatherQuery';
import DateUtil from '../../utils/DateUtil';

export type Props = {};

const CityDetails: FC<Props> = () => {
    const [curretCity, setCurrentCity] = useState<City | undefined>();
    const {
        cityIdState: [cityId],
    } = useContext(CityContext);

    const cities = useGetAllCity();
    const todaysWeather = useGetTodayWeatherForCity(cityId);

    useEffect(() => {
        const foundCity = cities?.filter((city) => city.id === cityId);
        setCurrentCity(foundCity?.[0]);
    }, [cityId, cities]);

    if (cities === undefined && todaysWeather === undefined) {
        return <LoadingIndicator />;
    }

    return (
        <div className="flex justify-center bg-slate-700">
            <div className="m-5 w-full">
                <p className="h-1/6 text-center">Test</p>
                <p className="h-1/6 text-center">Test</p>
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
