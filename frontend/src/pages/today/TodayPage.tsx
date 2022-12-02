import { FC } from 'react';
import { CityContextProvider } from '../../contexts/CityContext';
import CityDetails from './CityDetails';
import WeatherDetails from './WeatherDetails';

export type Props = {};

const TodayPage: FC<Props> = () => (
    <CityContextProvider>
        <div className="grid grid-cols-2 h-[calc(100vh-56px)] overflow-hidden">
            <CityDetails />
            <WeatherDetails />
        </div>
    </CityContextProvider>
);

export default TodayPage;
