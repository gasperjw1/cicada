import React from 'react';
import './Nav.css';

export default function Nav(props){
    return(
        <nav>
            <div className="logo">
                Cicada-Blockchain
            </div>
            <button className="connect-portis-button" onClick={props.connect_wall} >
                connect Portis
            </button>
        </nav>
    );
}