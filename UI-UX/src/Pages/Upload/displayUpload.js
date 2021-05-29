import React from 'react'
import './Dashboard.css'

const Display = (props) => {
    return (
        <div className="display-files-container">
            {
                props.selectedFile.map((fileHandle,index)=>{
                    console.log( fileHandle.type);
                    return(
                        <div key={`${fileHandle}${index}`} className="display-files" >
                            <p>{fileHandle.name}</p>
                            <p>{fileHandle.size} bytes</p>
                            {/* <p>{fileHandle.type}</p> */}
                            <p>Upload date:{Date()}</p>
                            <div className="X-container"  onClick={()=>props.removeFile(index)}/>
                    </div>
                    );
                })
            }
            
        </div>
    )
}

export default Display
