import './App.css';
import { WALLET_SUCCESS } from './Redux/Constants/walletConstants'
import { useDispatch, useSelector } from 'react-redux'
import Home from './Pages/Home/Home.js';
import Dashboard from './Pages/Upload/Dashboard'
import Nav from './Nav/Nav.js';
import { web3, portis } from './services/web3'
import { ethers } from 'ethers'
import {BrowserRouter as Router, Switch,Route} from 'react-router-dom';

function App() {
  const dispatch = useDispatch()
  const wallet = useSelector(state => state.walletList)
  const { wallets } = wallet

  const connect = () => {
    web3.eth.getAccounts((error, accounts) => {
      dispatch({type: WALLET_SUCCESS, payload:accounts})
    })
    .then(accounts => {
      const REGISTERED_USER = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("REGISTERED_USER"))
      const provider = new ethers.providers.Web3Provider(portis.provider);
      const address = "0xE31F4778593d284752120F0A6ad8f11F59bbb1bA"
      const signer = provider.getSigner()
      const abi = require('./ABI/Cicada').abi
      const contract = new ethers.Contract(address, abi, signer)

      contract
        .hasRole(REGISTERED_USER, accounts[0])
        .then(status => {
          if (!status) {
            web3.eth.sendTransaction({
              from: wallets[0],
              to: address,
              data: new ethers.utils.Interface(abi).encodeFunctionData('registerUser')         
            })
            .then(tx => console.log('You have been registered!'))
          } else {
            console.log('Already registered!')
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
          <Route path='/dashboard' exact component={Dashboard}/>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
