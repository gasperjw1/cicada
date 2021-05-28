import './App.css';
import Portis from '@portis/web3';
import Web3 from 'web3';

const portis = new Portis('d009ec94-ee7d-40c3-a782-bc25048ba6e4', 'mainnet');
const web3 = new Web3(portis.provider);

web3.eth.getAccounts((error, accounts) => {
  console.log(accounts);
});

function App() {
  return (
    <div className="App">
    </div>
  );
}

export default App;
