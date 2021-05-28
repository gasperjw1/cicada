import './App.css';
import Portis from '@portis/web3';
import Web3 from 'web3';
import { WALLET_SUCCESS } from './Redux/Constants/walletConstants'
import { useDispatch } from 'react-redux'
import Home from './Home/Home.js';
import Nav from './Nav/Nav.js';
import { web3, portis } from './services/web3'

function App() {
  const dispatch = useDispatch()
  const connect = () => {
    web3.eth.getAccounts((error, accounts) => {
      console.log(accounts);
      dispatch({type: WALLET_SUCCESS, payload:{accounts}})
    });
  }

  return (
    <div className="App">
      <Nav connect_wall={connect} />
      <Home/>
    </div>
  );
}

export default App;
