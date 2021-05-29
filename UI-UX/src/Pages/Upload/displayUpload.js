import React from 'react'
import './Upload.css'

const Display = (props) => {
    return (
        <div className="display-files-container">
            {
                props.selectedFile.map((fileHandle,index)=>{
                    return(
                        <div key={`${fileHandle}${index}`} className="display-files" >
                            <p>{fileHandle.name}</p>
                            <p>SIZE</p>
                            <p>{fileHandle.lastModifiedDate.toString()}</p>
                            <div className="X-container"  onClick={()=>props.removeFile(index)}/>
                    </div>
                    );
                })
            }
            
        </div>
    )
}

export default Display
