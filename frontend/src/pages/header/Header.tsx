import React, { FC } from 'react';

export type Props = {};

const Header: FC<Props> = () => (
    <header>
        <nav className="bg-white border-gray-200 px-4 lg:px-6 py-2.5">
            <div className="flex items-center">
                <img src="../../../../header-logo.png" className="mr-3 h-6 sm:h-9" alt="Weatherxz Logo" />
                <span className="self-center text-xl font-semibold whitespace-nowrap">Weatherxz</span>
            </div>
        </nav>
    </header>
);

export default Header;
