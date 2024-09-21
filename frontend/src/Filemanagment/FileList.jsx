import React from 'react';

const FileList = ({ files, onSelect, onCreate }) => {
    return (
        <div>
            <h3>Files</h3>
            <button onClick={onCreate}>Create New File</button>
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

export default FileList;
