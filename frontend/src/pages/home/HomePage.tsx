import { FC } from 'react';
import LoadingIndicator from '../../components/LoadingIndicator';
import { useGetTodayWeatherForCity } from '../../queries/WeatherQuery';

export type Props = {};

const HomePage: FC<Props> = () => {
    const { data } = useGetTodayWeatherForCity(1);

    const getDate = (): string => {
        const today = new Date();
        return `${today.getFullYear()} ${today.toLocaleString('default', { month: 'long' })} ${today.getDate()}`;
    };

    if (data === undefined) {
        return <LoadingIndicator />;
    }

    return (
        <div className="grid grid-cols-2 h-screen">
            <div className="flex justify-center bg-slate-700">
                <div className="m-5 w-full">
                    <p className="h-1/6 text-center">Test</p>
                    <p className="h-1/6 text-center">Test</p>
                    <div className="h-1/6 text-center flex justify-center">
                        <p className="font-extrabold text-center text-9xl text-white">{data.temp}</p>
                        <p className="font-light text-center text-6xl text-gray-500">°c</p>
                    </div>
                    <p className="font-extrabold h-1/6 text-center text-2xl text-gray-500">
                        {data.weatherConditionDesc}
                    </p>
                    <p className="h-1/6 text-center text-gray-500">{`Today - ${getDate()}`}</p>
                    <p className="h-1/6 text-center text-gray-500">Budapest / Hungary</p>
                </div>
            </div>
            <div className="flex h-screen bg-gray-400">
                <div className="m-auto">
                    <h2 className="justify-start m-5 font-bold text-2xl text-white flex">{`Today's Highlishgts`}</h2>
                    <div className="grid grid-cols-3 gap-3">
                        <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                            <h2 className="m-3 text-white justify-center flex">Pressure</h2>
                            <h1 className="text-white flex justify-center">{data.pressure}</h1>
                        </div>
                        <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                            <h2 className="m-3 text-white justify-center flex">Pop</h2>
                            <h1 className="text-white flex justify-center">{data.probabilityOfPrecipitation}</h1>
                        </div>
                        <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                            <h2 className="m-3 text-white justify-center flex">Humidity</h2>
                            <h1 className="text-white flex justify-center">{data.humidity}</h1>
                        </div>
                        <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                            <h2 className="m-3 text-white justify-center flex">Clouds</h2>
                            <h1 className="text-white flex justify-center">{data.clouds}</h1>
                        </div>
                        <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                            <h2 className="m-3 text-white justify-center flex">Wind</h2>
                            <h1 className="text-white flex justify-center">{data.windSpeed}</h1>
                        </div>
                        <div className="relative overflow-hidden bg-no-repeat shadow-lg rounded-lg bg-slate-700">
                            <h2 className="m-3 text-white justify-center flex">Visibility</h2>
                            <h1 className="text-white flex justify-center">{data.visibility}</h1>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default HomePage;
