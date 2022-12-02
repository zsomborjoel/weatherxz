import { FC } from 'react';
import { WeatherType } from '../enums/WeatherType';

type Props = {
    value?: WeatherType;
};

const WeatherTypeImg: FC<Props> = ({ value }): any => {
    switch (value) {
        case WeatherType.Clouds:
            return <img src="../../../weather/clear.png" className="mr-3 h-6 sm:h-9" alt="Cloudy Weather" />;
        case WeatherType.Rain:
            return <img src="../../../weather/cloud.png" className="mr-3 h-6 sm:h-9" alt="Rainy Weather" />;
        case WeatherType.Snow:
            return <img src="../../../weather/snow.png" className="mr-3 h-6 sm:h-9" alt="Snowy Weather" />;
        default:
            return <img src="../../../weather/clear.png" className="mr-3 h-6 sm:h-9" alt="Clear Weather" />;
    }
};

export default WeatherTypeImg;
