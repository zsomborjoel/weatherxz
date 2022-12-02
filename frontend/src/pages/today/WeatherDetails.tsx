import { FC, useContext } from 'react';
import LoadingIndicator from '../../components/LoadingIndicator';
import { CityContext } from '../../contexts/CityContext';
import { useGetTodayWeatherForCity } from '../../queries/WeatherQuery';

export type Props = {};

const WeatherDetails: FC<Props> = () => {
    const {
        cityIdState: [cityId],
    } = useContext(CityContext);

    const todaysWeather = useGetTodayWeatherForCity(cityId);

    if (todaysWeather === undefined) {
        return <LoadingIndicator />;
    }

    return (
        <div className="flex h-screen bg-gray-400">
            <div className="m-auto">
                <h2 className="justify-center m-5 font-bold text-2xl text-white flex">{`Today's Highlights`}</h2>
                <div className="grid grid-cols-3 gap-6">
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700 md:w-32 md:h-32 lg:h-40 lg:w-40 xl:w-72 xl:h-72">
                        <h2 className="text-white justify-center flex md:m-6 xl:m-14">Pressure</h2>
                        <h1 className="text-white flex justify-center">{`${todaysWeather?.pressure} hPa`}</h1>
                    </div>
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700 md:w-32 md:h-32 lg:h-40 lg:w-40 xl:w-72 xl:h-72">
                        <h2 className="m-3 text-white justify-center flex md:m-6 xl:m-14">Pop</h2>
                        <h1 className="text-white flex justify-center">
                            {`${todaysWeather?.probabilityOfPrecipitation ?? 0 * 100} %`}
                        </h1>
                    </div>
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700 md:w-32 md:h-32 lg:h-40 lg:w-40 xl:w-72 xl:h-72">
                        <h2 className="m-3 text-white justify-center flex md:m-6 xl:m-14">Humidity</h2>
                        <h1 className="text-white flex justify-center">{`${todaysWeather?.humidity} %`}</h1>
                    </div>
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700 md:w-32 md:h-32 lg:h-40 lg:w-40 xl:w-72 xl:h-72">
                        <h2 className="m-3 text-white justify-center flex md:m-6 xl:m-14">Clouds</h2>
                        <h1 className="text-white flex justify-center">{`${todaysWeather?.clouds} %`}</h1>
                    </div>
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700 md:w-32 md:h-32 lg:h-40 lg:w-40 xl:w-72 xl:h-72">
                        <h2 className="m-3 text-white justify-center flex md:m-6 xl:m-14">Wind</h2>
                        <h1 className="text-white flex justify-center">{`${todaysWeather?.windSpeed} m/sec`}</h1>
                    </div>
                    <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700 md:w-32 md:h-32 lg:h-40 lg:w-40 xl:w-72 xl:h-72">
                        <h2 className="m-3 text-white justify-center flex md:m-6 xl:m-14">Visibility</h2>
                        <h1 className="text-white flex justify-center">{`${todaysWeather?.visibility} m`}</h1>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default WeatherDetails;
