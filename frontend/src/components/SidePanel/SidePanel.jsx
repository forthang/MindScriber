import React, { useState } from 'react';
import './SidePanel.scss';

const SidePanel = () => {
    const [files, setFiles] = useState(['File 1', 'File 2', 'File 3']); 

    const addFile = () => {
        const newFileNumber = files.length + 1; 
        setFiles([...files, `File ${newFileNumber}`]); 
    };

    return (
        <div className="side-panel">
            <h2>Files</h2>
            <button className="add-file-button" onClick={addFile}>Add File</button>
            {files.map((file, index) => (
                <div key={index} className="file">{file}</div>
            ))}
        </div>
    );
};

export default SidePanel;
