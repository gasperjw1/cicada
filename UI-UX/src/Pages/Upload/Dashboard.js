import React, { useState,useEffect } from 'react'
import axios from 'axios'
import { useSelector} from 'react-redux'
import user from '../../Assets/user.png'
import uploadIcon from '../../Assets/uploadicon.png'
import myfilesIcon from '../../Assets/myfilesicon.png'
import upload2 from '../../Assets/upload2.png'
import Display from './displayUpload';
import Download_Icon from '../../Assets/download.png';

import './Dashboard.css'

const Dashboard = () => {
    const userWallet = useSelector(state=>state.walletList)

    const [selectedFile, uploadFile] = useState([])
    const [fileName, setFileName] = useState('Choose File')
    const [viewFiles, setViewFiles] = useState(false);

    const onFileChange = (event) => {
        uploadFile([...selectedFile, event.target.files[0]])
        setFileName(event.target.files[0])
        console.log(selectedFile)
    }



    useEffect(()=>{
        const setNav = document.querySelector('.upload');
        setNav.childNodes[0].classList.add('active-nav-item');
    },[])


    const submitUploads = async (event) => {
        event.preventDefault();

        try {
            var i
            for (i=0; i < selectedFile.length; i++)
            {   
                const formData = new FormData()
                formData.append('fileName' , selectedFile[i])
                formData.append('size' , selectedFile[i].size.toString())
                formData.append('type' , selectedFile[i].type.toString())
                const res = await axios.post('http://localhost:8080/upload', formData, {
                    headers: {
                        'Content-Type':'multipart/form-data'
                    }
                });
                console.log('its sent')
            }
            } catch (error) {
            
        }
    }

    const displayFiles = async () => {
        try {
            const res = await axios.get(`http://localhost:8080/display`)
            const data = await res;
            console.log(data);
        } catch (error){

        }
    }

    const getMyFile = async (event) =>{
        event.preventDefault();
        try {
            const res = await axios.get(`http://localhost:8080/download?uuid=${event}`)
            const data = await res;
            console.log(data);
        } catch (error){

        }
    }
    const removeFile = (idx) => {

        uploadFile( selectedFile.filter((file,index)=>{
            console.log(index)

            return index != idx;
        }));

    }

    const setView = () =>{
        setViewFiles(true);
    }
    const setViewNT = () =>{
        setViewFiles(false);
    }
    // if(!userWallet.wallets){
    //     return(
    //         <>
    //             <div ClassName="no-account-landing">
    //                 <p>Sorry, no account has been detected. If you don't have an account please register with PORTIS. If you are not subrcribed please consider doing so. Click here for more information PRICING</p>

    //             </div>
                    
    //         </>
    //     )
    // }
    return (
        <div className='Container'>
            <div className='upload-container'>

                <div className="sideNav">
                    
                        <div onClick={setViewNT}><img className='icon-upload' src={uploadIcon} alt='uploadicon.png'/> <span>Upload</span></div>
                        <div onClick={setView} ><img className='icon-file' src={myfilesIcon} alt='myfilesicon.png'/><span>My Files</span></div>
                 
                </div>

                {!viewFiles 
                    ?<div className='upload-wrapper'>
                        <div className="top-container">  
                            <div className='user-info-container'>
                                <img src={user} alt='user.png' className="user-wrapper"/>
                                {
                                    userWallet.wallets
                                    ?<div className="user-wallet-wrapper">{userWallet.wallets}</div>
                                    : null
                                }
                            </div>
                            <div className="upload-file-wrapper">
                                <form className="form-wrapper" onSubmit={submitUploads}>
                                    <input 
                                    className="upload-input"
                                    name='myFile' 
                                    type='file' 
                                    id='file'
                                    onChange={onFileChange}
                                    />
                                    <label className="upload-file" htmlFor="file"> <img className="upload-img"src={upload2} alt="upload.png"/> Choose a file</label>
                                    <input className="upload-button"type='submit' value='Upload'></input>
                                </form>
                            </div>
                        </div>
                        <div className="break"/>
                        <div className="static-text">
                            <h1 className="static-h1-one">Name</h1> <h1 className="static-h1-two">Size</h1> <h1 className="static-h1-three">Upload Date</h1>
                        </div>
                        <div className="break"/>

                        <div className="second-container">
                        <Display selectedFile={selectedFile} removeFile={removeFile} />
                        </div>
                    </div>
                        
                    :<div className="display-my-files">
                        <button onClick={displayFiles} >display</button>
                      
                        <div className="my-files-container">
                            <div className="my-file-item">
                                <div className="file-handle"> the one asdasd randy's adjhbajdkfhbakjjdadsadaskjdaskjdasjdhasakjshdnfkjbs Resume.pdf</div>
                                <div className="date-uploaded">date Uploaded</div>
                                <div className="file-size">file size</div>
                                <img className="download-icon" src={Download_Icon} alt="download.png" />
                            </div>
                            
                        </div>
                    </div>
                    
                }
                </div>
            </div>
    );
}

export default Dashboard
