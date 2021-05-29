import React, { useState } from 'react'
import axios from 'axios'
import './Upload.css'

const Upload = () => {

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
                const res = await axios.post('/upload', formData)
                console.log('its sent')
            }
        } catch (error) {
            
        }
    }
    
    return (
        <div className='Container'>
            <div className='Upload'>
                <form onSubmit={submitUploads}>
                    <input type='file' onChange={onFileChange}/>
                    <input type='submit' value='Upload'/>
                </form>
            </div>
            
        </div>
    )
}

export default Upload
