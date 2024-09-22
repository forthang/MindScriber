import React, { useState } from 'react';
import './SidePanel.scss'
import FileInput from '../../Filemanagment/CreateFile';

const SidePanel = ({ files, onCreate, onSelect }) => {
    const handleCreateFile = () => {
        const newFile = { name: `New File ${files.length + 1}`, content: '' };
        onCreate(newFile);
    };

    return (
        <div className="side-panel">
            <h3>Files</h3>
            <button onClick={handleCreateFile}>Create New File</button>
            <ul>
                {files.map((file, index) => (
                    <li key={index} onClick={() => onSelect(file)}>
                        {file.name}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default SidePanel;

