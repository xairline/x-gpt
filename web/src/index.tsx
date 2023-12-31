import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {BrowserRouter} from "react-router-dom";
import {LogtoConfig, LogtoProvider, UserScope} from '@logto/react';

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);
const config: LogtoConfig = {
    endpoint: 'https://auth.xairline.org',
    appId: 'sebl8fjkssu2oaw2s8gsv',
    scopes: [UserScope.Email],
};

const domain = window.location.origin + "/api";
root.render(
    <BrowserRouter>
        <LogtoProvider config={config}>
            <App/>
        </LogtoProvider>
    </BrowserRouter>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
