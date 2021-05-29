import './App.css';
import { WALLET_SUCCESS } from './Redux/Constants/walletConstants'
import { useDispatch } from 'react-redux'
import Home from './Pages/Home/Home.js';
import Upload from './Pages/Upload/Upload'
import Nav from './Nav/Nav.js';
import { web3, portis } from './services/web3'
import {BrowserRouter as Router, Switch,Route} from 'react-router-dom';

function App() {
  const dispatch = useDispatch()
  const connect = () => {
    web3.eth.getAccounts((error, accounts) => {
      console.log(accounts);
      dispatch({type: WALLET_SUCCESS, payload:{accounts}})
    });
  }

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
