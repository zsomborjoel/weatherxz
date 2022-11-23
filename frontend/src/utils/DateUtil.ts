const getTodaysDate = (): string => {
    const today = new Date();
    return `${today.getFullYear()} ${today.toLocaleString('default', { month: 'long' })} ${today.getDate()}`;
};

const DateUtil = {
    getTodaysDate,
};

export default DateUtil;
