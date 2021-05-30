import React from 'react';
import './Home.css';
import Particles from 'react-particles-js';
import para from './particles.json'
import logo from '../../Assets/logo2.png'
import stable from '../../Assets/stable.png'
import portis from '../../Assets/portis.png'
import golang from '../../Assets/golang.png'
import storj from '../../Assets/storj.png'

export default function Home(){
    return(
        <>
            <section className="home-container">
                <div className="home-landing">
                    <Particles 
                            params={para} 
                            className="background-part"
                    />
                    <div className="text-wrapper">
                        <div className="text-holder">
                            <span>Decentralized Cloud storage</span> <br/>
                            <span>for developers through</span> <br/>
                            <span>Portis</span>
                        </div>
                        <div className="button-wrapper">
                            Upload Files
                        </div>
                    </div>
                    <img src={logo} alt='logo2.png' className='logo-wallpaper'/> 
                </div>
                <div className="info-wrapper">
                    <div className="info-container">
                        <div className='storj-wrapper'> 
                            <img className='storj' src={storj} alt={storj.png}/>
                            <h3 className='info-text'>With Storj DCS (Decentralized Cloud Storage) files aren’t stored in centralized data centers— instead, they're encrypted, split into pieces, and distributed on a global cloud network.</h3>
                        </div>
                        <div className="portis-wrapper">
                            <img className='portis' src={portis} alt={portis.png}/>
                            <h3 className='info-text'>Your Dapp (Decentralized Application) communicates with the Portis SDK using the standard web3.js methods, meaning it will work automatically with your existing code.</h3>
                        </div>
                        <div className="golang-wrapper">
                            <img className='golang' src={golang} alt={golang.png}/>
                            <h3 className='info-text'>Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.</h3>
                        </div>
                    </div>
                </div>
                <div className="foot-note">
                    <div className="foot-note-wrapper">
                        <div className="left-footer">
                            <div className="left-footer-content">
                                <h1>Powered by: </h1>
                                <img className='stable-logo' src={stable} alt='stable.png'/>
                            </div>
                        </div>
                        <div className="right-footer">
                            <h2 className="footer-text">Contact Us</h2>
                            <h2 className="footer-text">About Us</h2>
                            <h2 className="footer-text">Join Us</h2>
                            <h2 className="footer-text">Accessebility</h2>
                        </div>   
                    </div>
                </div>
            </section>
        </>
    );
}