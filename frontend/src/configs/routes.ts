import TodayPage from '../pages/today/TodayPage';
import { ROUTES } from './constants';

export default [
    { path: ROUTES.ROOT, element: TodayPage, name: 'Home' },
    { path: ROUTES.HOME, element: TodayPage, name: 'Home' },
];
