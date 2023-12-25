import React from "react";
import {useHandleSignInCallback} from '@logto/react';
import {Navigate} from 'react-router-dom';

const Callback: React.FC = () => {
    const {isLoading} = useHandleSignInCallback(() => {
        window.location.href = '/';
    });

    return isLoading ? <p>Redirecting...</p> : null;
};

export default Callback;