import { FC } from 'react';
import { WeatherType } from '../enums/WeatherType';

type Props = {
    value?: string;
};

const WeatherTypeImg: FC<Props> = ({ value }): any => {
    switch (value) {
        case WeatherType.Clouds:
            return <img src="../../../weather/cloud.png" className="w-32 h-32" alt="Cloudy Weather" />;
        case WeatherType.Rain:
            return <img src="../../../weather/rain.png" className="w-32 h-32" alt="Rainy Weather" />;
        case WeatherType.Snow:
            return <img src="../../../weather/snow.png" className="w-32 h-32" alt="Snowy Weather" />;
        default:
            return <img src="../../../weather/clear.png" className="w-32 h-32" alt="Clear Weather" />;
    }
};

export default WeatherTypeImg;
