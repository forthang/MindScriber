import React, { useState } from 'react';
import axios from 'axios';
import './CreateMessage.scss'; 

const MessageInput = () => {
    const [message, setMessage] = useState('');
    const [messages, setMessages] = useState([]);

    const handleChange = (event) => {
        setMessage(event.target.value);
    };

    const handleKeyPress = async (event) => {
        if (event.key === 'Enter') {
            event.preventDefault();     
            if (!message) return; 

            try {
              
                const response = await axios.post('http://localhost:8080/messages', { message });
                setMessages([...messages, response.data.message]); 
                setMessage(''); 
            } catch (error) {
                console.error('Error saving message:', error);
            }
        }
    };

    return (
        <div className="message-input-container">
            <form className="message-input-form">
                <input 
                    type="text" 
                    value={message} 
                    onChange={handleChange} 
                    onKeyPress={handleKeyPress} 
                    placeholder="Type your message here..." 
                    className="message-input"
                />
            </form>
            <div className="messages">
                <h3>Saved Messages:</h3>
                <ul>
                    {messages.map((msg, index) => (
                        <li key={index}>{msg}</li>
                    ))}
                </ul>
            </div>
        </div>
    );
};

export default MessageInput;
