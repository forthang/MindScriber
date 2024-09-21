import React, { useState } from 'react';
import FileInput from './CreateFile';

const Folder = ({ folder, onFileSelect, onCreateFile }) => {
    const [selectedFile, setSelectedFile] = useState(null);

    const handleFileSelect = (file) => {
        setSelectedFile(file);
        onFileSelect(file);
    };

    const handleCreateFile = () => {
        const newFile = { name: `New File ${folder.files.length + 1}`, content: '' };
        onCreateFile(folder.name, newFile);
        setSelectedFile(newFile); // Select the newly created file
    };

    return (
        <div>
            <h3>{folder.name}</h3>
            <button onClick={handleCreateFile}>Create New File</button>
            <ul>
                {folder.files.map((file, index) => (
                    <li key={index} onClick={() => handleFileSelect(file)}>
                        {file.name}
                    </li>
                ))}
            </ul>
            {selectedFile && (
                <FileInput currentFile={selectedFile} onSave={(content) => {
                    const updatedFiles = folder.files.map(f =>
                        f.name === selectedFile.name ? { ...f, content } : f
                    );
                    onCreateFile(folder.name, updatedFiles);
                }} />
            )}
        </div>
    );
};

export default Folder;
