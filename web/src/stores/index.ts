import { createContext, useContext } from 'react';

export const rootStoreContext = createContext({
  // ResultsStore: resultsStore
});

export const useStores = () => {
  const store = useContext(rootStoreContext);
  if (!store) {
    throw new Error('useStores must be used within a provider');
  }
  return store;
};
