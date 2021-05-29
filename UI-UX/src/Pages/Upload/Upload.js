import React, { useState } from 'react'
import axios from 'axios'
import { useSelector} from 'react-redux'
import user from '../../Assets/user.png'
import upload from '../../Assets/upload.png'
import upload2 from '../../Assets/upload2.png'

import './Upload.css'

const Upload = () => {
    const userWallet = useSelector(state=>state.walletList)

    const [selectedFile, uploadFile] = useState([])
    const [fileName, setFileName] = useState('Choose File')

    const onFileChange = (event) => {
        uploadFile([...selectedFile, event.target.files[0]])
        setFileName(event.target.files[0])
        console.log(selectedFile)
    }


    const submitUploads = async (event) => {
        event.preventDefault()

        try {
            var i
            for (i=0; i < selectedFile.length; i++)
            {
                const formData = new FormData()
                formData.append('fileName', selectedFile[i])
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

    const removeFile = (idx) => {

        uploadFile( selectedFile.filter((file,index)=>{
            console.log(index)

            return index != idx;
        }));

    }
    
    return (
        <div className='Container'>

            <div className='Upload'>
                <form onSubmit={submitUploads}>
                    <input name='myFile' 
                           type='file' 
                           onChange={onFileChange}
                           />
                    <input type='submit' value='Upload'/>
                </form>
                <button onClick={()=>{console.log(selectedFile)}}>selected file</button>
            </div>
            <div className="display-files-container">
                {
                    selectedFile.map((fileHandle,index)=>{
                        return(
                            <div key={`${fileHandle}${index}`} className="display-files" >
                                <p>{fileHandle.name}</p>
                                <p>{fileHandle.lastModifiedDate.toString()}</p>
                                <div className="X-container"  onClick={()=>removeFile(index)}/>
                            </div>
                        );
                    })
                }
            </div>
            <div className='upload-container'>
                <div className='upload-wrapper'>
                    <div className="top-container">  
                    <div className='user-info-container'>
                        <img src={user} alt='user.png' className="user-wrapper"/>
                        <div className="user-wallet-wrapper">0x916c5F7ddcce69F0…f1e2a4f</div>
                        {console.log('Wallets:', userWallet.wallets.accounts)}
                    </div>
                    <div className="upload-file-wrapper">
                        <form className="form-wrapper">
                            <input 
                            className="upload-input"
                            name='myFile' 
                            type='file' 
                            id='file'
                            onChange={onFileChange}
                            />
                            <label className="upload-file" for="file"> <img className="upload-img"src={upload2} alt="upload.png"/> Choose a file</label>
                            <input className="upload-button"type='submit' value='Upload'></input>
                        </form>

                    </div>
                    </div>
                    <div className="break"/>
                    <div className="static-text">
                        <h1 className="static-h1">Name</h1> <h1 className="static-h1">Size</h1> <h1 className="static-h1">Upload Date</h1>
                    </div>
                    <div className="break"/>
                    <button onClick={()=>{console.log(selectedFile)}}>selected file</button>
                </div>
                </div>
            </div>
    )
}

export default Upload
