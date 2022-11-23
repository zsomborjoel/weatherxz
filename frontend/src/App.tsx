import React, { createElement } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Header from './pages/header/Header';
import routes from './configs/routes';
import NotFoundPage from './pages/notfound/NotFoundPage';
import TodayPage from './pages/today/TodayPage';

function App(): any {
    return (
        <div>
            <Router>
                <Header />
                <Routes>
                    <Route path="/home" element={<TodayPage />} />
                    {routes.map(({ path, element }) => (
                        <Route key={path} path={path} element={<div>{createElement(element)}</div>} />
                    ))}
                    <Route path="*" element={<NotFoundPage />} />
                </Routes>
            </Router>
        </div>
    );
}

export default App;
