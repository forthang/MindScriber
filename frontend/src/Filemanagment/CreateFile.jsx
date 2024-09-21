import React, { useState } from 'react';
import MessageInput from '../components/CreateMessage';


const FileInput = ({ currentFile, onSave }) => {
    const [content, setContent] = useState(currentFile ? currentFile.content : '');

    const handleChange = (event) => {
        setContent(event.target.value);
    };

    const handleSave = () => {
        onSave(content);
    };

    return (
        <div>
            <h2>{currentFile ? currentFile.name : 'No file selected'}</h2>
            <MessageInput/>
            <button onClick={handleSave}>Save</button>
        </div>
    );
};

export default FileInput;

