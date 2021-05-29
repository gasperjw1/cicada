import React from 'react';
import './Home.css';
import Particles from 'react-particles-js';
import para from './particles.json'

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
                </div>
                <div className="info-wrapper">
                    <div className="info-container">

                    </div>
                </div>
            </section>
        </>
    );
}