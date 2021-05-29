import React from 'react';
import './Home.css';
import Particles from 'react-particles-js';
import para from './particles.json'
import logo from '../../Assets/logo2.png'
import stable from '../../Assets/stable.png'

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

                    </div>
                </div>
                <div className="foot-note">
                    <div className="foot-note-wrapper">
                        <div className="left-footer">
                            <h1>Powered by: </h1>
                            <img className='stable-logo' src={stable} alt='stable.png'/>
                        </div>
                        <div className="right-fotter">

                        </div>   
                    </div>

                </div>
            </section>
        </>
    );
}