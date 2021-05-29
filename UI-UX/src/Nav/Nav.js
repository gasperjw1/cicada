import React,{useEffect} from 'react';
import './Nav.css';

export default function Nav(props){
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
            <div className="logo">
                Cicada-Blockchain
            </div>

            <div className="nav-item">
                <div className="underline"/>
                Pricing
            </div>
            <div className="nav-item">
                <div className="underline"/>
                Upload
               
            </div>
            <button className="connect-portis-button" onClick={props.connect_wall} >
                Connect Portis
            </button>
        </nav>
    );
}