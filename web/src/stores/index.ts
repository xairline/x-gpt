import {createContext, useContext} from 'react';
import {flightLogStore} from "./FlightLog";

export const rootStoreContext = createContext({
    FlightLogStore: flightLogStore,
});

export const useStores = () => {
    const store = useContext(rootStoreContext);
    if (!store) {
        throw new Error('useStores must be used within a provider');
    }
    return store;
};
