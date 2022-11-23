import { FC, useState, useEffect, useContext } from 'react';
import LoadingIndicator from '../../components/LoadingIndicator';
import { CityContext, CityContextProvider } from '../../contexts/CityContext';
import { Weather } from '../../models/Weather';
import { useGetTodayWeatherForCity } from '../../queries/WeatherQuery';
import CityDetails from './CityDetails';

export type Props = {};

const TodayPage: FC<Props> = () => {
    const {
        cityIdState: [cityId],
    } = useContext(CityContext);

    const todaysWeather = useGetTodayWeatherForCity(cityId);

    if (todaysWeather === undefined) {
        return <LoadingIndicator />;
    }

    return (
        <CityContextProvider>
            <div className="grid grid-cols-2 h-screen">
                <CityDetails />
                <div className="flex h-screen bg-gray-400">
                    <div className="m-auto">
                        <h2 className="justify-center m-5 font-bold text-2xl text-white flex">{`Today's Highlights`}</h2>
                        <div className="grid grid-cols-3 gap-3">
                            <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                                <h2 className="m-3 text-white justify-center flex">Pressure</h2>
                                <h1 className="text-white flex justify-center">{todaysWeather.pressure}</h1>
                            </div>
                            <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                                <h2 className="m-3 text-white justify-center flex">Pop</h2>
                                <h1 className="text-white flex justify-center">
                                    {todaysWeather.probabilityOfPrecipitation}
                                </h1>
                            </div>
                            <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                                <h2 className="m-3 text-white justify-center flex">Humidity</h2>
                                <h1 className="text-white flex justify-center">{todaysWeather.humidity}</h1>
                            </div>
                            <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                                <h2 className="m-3 text-white justify-center flex">Clouds</h2>
                                <h1 className="text-white flex justify-center">{todaysWeather.clouds}</h1>
                            </div>
                            <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                                <h2 className="m-3 text-white justify-center flex">Wind</h2>
                                <h1 className="text-white flex justify-center">{todaysWeather.windSpeed}</h1>
                            </div>
                            <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                                <h2 className="m-3 text-white justify-center flex">Visibility</h2>
                                <h1 className="text-white flex justify-center">{todaysWeather.visibility}</h1>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </CityContextProvider>
    );
};

export default TodayPage;
