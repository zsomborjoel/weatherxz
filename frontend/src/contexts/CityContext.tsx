import React, { createContext, FC, useState, Dispatch, SetStateAction } from 'react';

interface CityContextInterface {
    cityIdState: [cityId: number, setCityId: Dispatch<SetStateAction<number>>];
}

const CityContext = createContext<CityContextInterface>({
    cityIdState: [1, () => {}],
});

const CityContextProvider: FC<any> = ({ children }) => {
    const [cityId, setCityId] = useState<number>(1);

    return (
        <CityContext.Provider
            // eslint-disable-next-line react/jsx-no-constructed-context-values
            value={{
                cityIdState: [cityId, setCityId],
            }}
        >
            {children}
        </CityContext.Provider>
    );
};

export { CityContext, CityContextProvider };
