import React, { useState } from 'react'
import './Upload.css'

const Upload = () => {

    const [selectedFile, uploadFile] = useState('')

    const onFileChange = (event) => {
        uploadFile(event.target.files[0])
        console.log(selectedFile)
    }
    
    return (
        <div className='Container'>
            <div className='Upload'>
                <input type='file' onChange={onFileChange}/>

                <button>Upload</button>
            </div>
            
        </div>
    )
}

export default Upload
