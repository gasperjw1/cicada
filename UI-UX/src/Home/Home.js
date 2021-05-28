import React from 'react';
import './Home.css';

export default function Home(){
    return(
        <>
            <section className="home-container">
                <div className="home-landing">
                    <div className="text-wrapper">
                        <div className="text-holder">
                            Decentralized Cloud storage <br/>
                            for developers through <br/>
                            Portis
                        </div>
                        <div className="button-wrapper">
                            Upload Files
                        </div>
                    </div>
                </div>
            </section>
        </>
    );
}