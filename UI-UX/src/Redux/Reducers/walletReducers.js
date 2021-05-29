import { WALLET_FAIL, WALLET_REQUEST, WALLET_SUCCESS }  from '../Constants/walletConstants'
import { REHYDRATE } from 'redux-persist'


export const walletReducer = (state= { wallets: [] }, action) => {
    switch (action.type){
        case WALLET_REQUEST:
            return { loading: true, wallets: [] }
        case WALLET_SUCCESS:
            return { loading: false, wallets: action.payload }
        case WALLET_FAIL:
            return { loading: false, error: action.payload }
        case REHYDRATE:
            return { loading: false, wallets: action.payload?.walletList?.wallets}
        default:
            return state    
    }
}   