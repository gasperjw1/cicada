import React,{useState,useEffect} from 'react'
import {useSelector} from 'react-redux'
import './Pricing.css'
import renew from '../../Assets/renew.png'
import storage from '../../Assets/storage.png'
import { web3, portis } from '../../services/web3'
import { ethers } from 'ethers'

const Pricing = () => {
    const [storageSpace, setStorageSpace] = useState(1);
    const setStorage = (bytes) =>{
        if(bytes === -1){
            setStorageSpace(1);
            return;
        }

        if(bytes >= 0 ){
            setStorageSpace(bytes);
            return;
        }
        
    }

    const wallet = useSelector(state => state.walletList)
    const { wallets } = wallet
    const provider = new ethers.providers.Web3Provider(portis.provider);
    const address = "0xE31F4778593d284752120F0A6ad8f11F59bbb1bA"
    const signer = provider.getSigner()
    const abi = require('../../ABI/Cicada').abi
    const contract = new ethers.Contract(address, abi, signer)

    const renewMembership = async () => {
        const monthlyFee = await contract.monthlyFee();
        web3.eth.sendTransaction({
            from: wallets[0],
            to: address,
            value: monthlyFee._hex,
            data: new ethers.utils.Interface(abi).encodeFunctionData('renewSubscription')
        })
        .then(tx => console.log(tx))
    }

    const addStorage = async () => {
        const rate = await contract.memoryRate()
        console.log(rate * storageSpace)
        web3.eth.sendTransaction({
            from: wallets[0],
            to: address,
            value: rate * storageSpace,
            data: new ethers.utils.Interface(abi).encodeFunctionData('payForStorage', [storageSpace])
        })
        .then(tx => console.log(tx))
    }

    return (
        <div className='pricing-container'>
            <div className='pricing-wrapper'>
                <div className='text-info'>
                    <h1 className='info-text'>
                        Store files and maintain their integrity through blockchain technology
                    </h1>
                    <h2 className='second-text'>Using Storj and their ethereum nodes, your files will always maintain security through a reasonable price. Interacting with our Dapp will require you to have an Portis account, where those credentials will be used to verify your files and payment systems.</h2>
                </div>
                <div className='pricing-buttons'>
                    <div className='storage-wrapper'>
                        <div className='storage-container'>
                            <h3 className='storage-tag'>
                                Add more storage to your membership
                            </h3>
                            <img src={storage} alt="storage.png" className='renew-img'/>
                            <div className="input-storage-container">
                                <input type='text' className='input-storage' value={storageSpace} onChange={(e)=>{
                                    let GB = e.target.value;
                                    if(GB === "00" || GB === ''){
                                        setStorage(-1);
                                    } else {
                                        setStorage(parseInt(GB));
                                    }
                                }}/>
                                <p className="after-input">GB</p>
                            </div>
                            <div className='purchase-button' onClick={addStorage}>Purchase</div>
                        </div>
                    </div>
                    <div className='renew-wrapper'>
                        <div className='container-buttons'>
                            <h3 className='renew-tag'>
                                Renew your membership!
                            </h3>
                            <img  className='renew-img'src={renew} alt="renew.png"/>
                            <div className='renew-button' onClick={renewMembership}>
                            Renew Membership!
                            </div>

                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Pricing
