import './App.css';
import { WALLET_SUCCESS } from './Redux/Constants/walletConstants'
import { useDispatch } from 'react-redux'
import Home from './Pages/Home/Home.js';
import Upload from './Pages/Upload/Upload'
import Nav from './Nav/Nav.js';
import { web3, portis } from './services/web3'
import { ethers } from 'ethers'
import {BrowserRouter as Router, Switch,Route} from 'react-router-dom';

function App() {
  const dispatch = useDispatch()
  const connect = () => {
    web3.eth.getAccounts((error, accounts) => {
      dispatch({type: WALLET_SUCCESS, payload:accounts})
    })
    .then(accounts => {
      web3
        .listAccounts()
        .then(wallet => {
          if (wallet.length) {
            const provider = new ethers.providers.Web3Provider(window.ethereum);
            const REGISTERED_USER = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("REGISTERED_USER"))
            const signer = provider.getSigner()
            const address = ""
            const abi = ""
            const contract = new ethers.Contract(address, abi, signer)

            contract.hasRole(REGISTERED_USER, wallet[0])
          }
        })
    })
  }
// hello {^V^}
  return (
    <Router>
      <div className="App">
        <Nav connect_wall={connect} />
        <Switch>
          <Route path='/' exact component={Home}/>
          <Route path='/upload' exact component={Upload}/>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
