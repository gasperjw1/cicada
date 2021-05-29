import React,{useEffect} from 'react';
import {Link} from 'react-router-dom';
import logo from '../Assets/logo.png'
import './Nav.css';

export default function Nav(props){
    const clear_underline = () =>{
        const underline = document.querySelector('.active-nav-item');
        underline?.classList.remove('active-nav-item');
    }

    useEffect(()=>{
        const nav_items = document.querySelectorAll('.nav-item');
        nav_items.forEach((nav_item,index)=>{
            nav_item.addEventListener('click',()=>{
               
                let underline = nav_item.childNodes[0];
                if(!underline.classList.contains('active-nav-item')){
                    let remove = document.querySelector('.active-nav-item');
                    remove?.classList.remove('active-nav-item');
                    console.log(underline.classList.contains('active-nav-item'));
                    underline.classList.add('active-nav-item');
                    console.log(underline);
                }
                
            })
        })
    },[])


    return(
        <nav>
            <Link to="/" target="_self">
            <div className="logo" onClick={clear_underline}>
                <img src={logo} className="logo-img" alt='logo.png'/>
                Cicada-Drop
            </div>
            </Link>
            <div className="nav-item pricing">
                <div className="underline"/>
                <Link to="/pricing" target="_self">Pricing</Link>
            </div>
            <div className="nav-item upload">
                <div className="underline"/>
                <Link to="/dashboard" target="_self">Dashboard</Link>
            </div>
            <button className="connect-portis-button" onClick={props.connect_wall} >
                Connect Portis
            </button>
        </nav>
    );
}