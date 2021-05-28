import Portis from "@portis/web3";
import Web3 from "web3";

export const portis = new Portis(
  "d009ec94-ee7d-40c3-a782-bc25048ba6e4",
  "ropsten",
  { gasRelay: true }
);
export const web3 = new Web3(portis.provider);

