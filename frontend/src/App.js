import React, { useState } from 'react';
import SidePanel from './components/SidePanel';
import FileInput from './components/FileInput';
import './App.css'; // Import global styles if needed
import Header from './components/Header/Header';


const App = () => {
    const [files, setFiles] = useState([]);
    const [currentFile, setCurrentFile] = useState(null);

    const handleCreateFile = (newFile) => {
        if (Array.isArray(newFile)) {
            setFiles(newFile); // Update files with new content
        } else {
            setFiles([...files, newFile]); // Add new file to the list
            setCurrentFile(newFile); // Select newly created file
        }
    };

    const handleSelectFile = (file) => {
        setCurrentFile(file); // Set selected file for editing
    };

    return (
        <div className="app" style={{ display: 'flex' }}>
          <Header/>
            <div style={{ flex: '0 0 300px' }}> {/* Fixed width for SidePanel */}
                <SidePanel files={files} onCreate={handleCreateFile} onSelect={handleSelectFile} />
            </div>
            <div style={{ flex: 1, padding: '20px' }}> {/* Flexible width for FileInput */}
                {currentFile ? (
                    <FileInput currentFile={currentFile} onSave={(content) => {
                        const updatedFiles = files.map(f =>
                            f.name === currentFile.name ? { ...f, content } : f
                        );
                        setFiles(updatedFiles);
                    }} />
                ) : (
                    <h2>Select a file to view or create files.</h2>
                )}
            </div>
        </div>
    );
};

export default App;
