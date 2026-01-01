import React, { useEffect, useState } from 'react';
import ReactDOM from 'react-dom';
import { CheckCircle, AlertCircle, X } from 'lucide-react';

interface ToastProps {
    message: string;
    type: 'success' | 'error';
    onClose: () => void;
}

export const Toast: React.FC<ToastProps> = ({ message, type, onClose }) => {
    const [isExiting, setIsExiting] = useState(false);

    useEffect(() => {
        const timer = setTimeout(() => {
            handleClose();
        }, 3000);
        return () => clearTimeout(timer);
    }, []);

    const handleClose = () => {
        setIsExiting(true);
        setTimeout(onClose, 300); // Wait for animation
    };

    return ReactDOM.createPortal(
        <div
            className={isExiting ? 'toast-exit' : 'toast-enter'}
            style={{
                position: 'fixed',
                top: '1rem',
                left: '50%',
                transform: 'translateX(-50%)',
                backgroundColor: type === 'success' ? '#166534' : '#991b1b',
                color: '#ffffff',
                padding: '0.75rem 1rem',
                borderRadius: '8px',
                boxShadow: '0 10px 25px -3px rgba(0, 0, 0, 0.3), 0 4px 10px -2px rgba(0, 0, 0, 0.2)',
                zIndex: 9999,
                display: 'flex',
                alignItems: 'center',
                gap: '0.75rem',
                minWidth: '300px',
                border: type === 'success' ? '1px solid #22c55e' : '1px solid #ef4444',
            }}
        >
            {type === 'success' ? (
                <CheckCircle size={20} color="#22c55e" />
            ) : (
                <AlertCircle size={20} color="#fca5a5" />
            )}

            <span style={{ flex: 1, fontSize: '0.9rem', fontWeight: 500 }}>{message}</span>

            <button
                onClick={handleClose}
                style={{
                    background: 'transparent',
                    color: 'rgba(255, 255, 255, 0.7)',
                    padding: '0.25rem',
                    display: 'flex'
                }}
            >
                <X size={16} />
            </button>
        </div>,
        document.body
    );
};
