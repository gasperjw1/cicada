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
                <form enctype="multipart/form-data" onSubmit={submitUploads}>
                    <input name='myFile' type='file' onChange={onFileChange}/>
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
                
            
        </div>
    )
}

export default Upload
