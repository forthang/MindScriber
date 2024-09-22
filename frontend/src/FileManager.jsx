import React, { useEffect, useState } from 'react';
import axios from 'axios';

const FileManager = ({ token }) => {
    const [files, setFiles] = useState([]);
    const [fileName, setFileName] = useState('');
    const [fileContent, setFileContent] = useState('');
    const [currentFile, setCurrentFile] = useState(null);

    useEffect(() => {
        fetchFiles();
    }, []);

    const fetchFiles = async () => {
        try {
            const response = await axios.get('http://localhost:8080/files', {
                headers: { Authorization: token },
            });
            setFiles(response.data);
        } catch (err) {
            console.error(err);
        }
    };

    const handleCreateFile = async () => {
        try {
            await axios.post('http://localhost:8080/files', { name: fileName, content: fileContent }, {
                headers: { Authorization: token },
            });
            fetchFiles(); // Refresh file list after creating a new file
            setFileName('');
            setFileContent('');
        } catch (err) {
            console.error(err);
        }
    };

    const handleSelectFile = (file) => {
        setCurrentFile(file);
        setFileName(file.name);
        setFileContent(file.content);
    };

    const handleUpdateFile = async () => {
        try {
            await axios.post('http://localhost:8080/files', { name: fileName, content: fileContent }, {
                headers: { Authorization: token },
            });
            fetchFiles(); // Refresh file list after updating the file
        } catch (err) {
            console.error(err);
        }
    };

    return (
        <div>
            <h2>File Manager</h2>
            <div>
                <input
                    type="text"
                    placeholder="File Name"
                    value={fileName}
                    onChange={(e) => setFileName(e.target.value)}
                />
                <textarea
                    placeholder="File Content"
                    value={fileContent}
                    onChange={(e) => setFileContent(e.target.value)}
                />
                <button onClick={currentFile ? handleUpdateFile : handleCreateFile}>
                    {currentFile ? 'Update File' : 'Create File'}
                </button>
            </div>

            <h3>Your Files</h3>
            <ul>
                {files.map((file) => (
                    <li key={file.name} onClick={() => handleSelectFile(file)}>
                        <strong>{file.name}</strong>: {file.content}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default FileManager;
