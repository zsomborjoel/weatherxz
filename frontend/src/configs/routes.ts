import HomePage from '../pages/home/HomePage';
import { ROUTES } from './constants';

export default [
    { path: ROUTES.ROOT, element: HomePage, name: 'Home' },
    { path: ROUTES.HOME, element: HomePage, name: 'Home' },
];
