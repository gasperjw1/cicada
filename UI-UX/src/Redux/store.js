import { createStore, combineReducers, applyMiddleware } from 'redux'
import thunk from 'redux-thunk'
import { composeWithDevTools } from 'redux-devtools-extension'
import { walletReducer } from './Reducers/walletReducers'
import storageSession from 'redux-persist/lib/storage/session'
import { persistReducer, persistStore } from 'redux-persist';
import autoMergeLevel2 from 'redux-persist/lib/stateReconciler/autoMergeLevel2';


const persistConfig = {
    key: 'root',
    storage: storageSession,
    whitelist: ['walletList'],
    stateReconciler: autoMergeLevel2
}

const reducer = combineReducers({
    walletList: walletReducer,
})

const initialState = {}

const middleware = [thunk]

const persistorRedux = persistReducer(persistConfig, reducer);

export const store = createStore(persistorRedux, initialState, composeWithDevTools(applyMiddleware(...middleware)))
export const persistor = persistStore(store);

export default {store, persistor}

