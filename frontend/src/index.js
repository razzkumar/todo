import React from 'react';
import { createRoot } from 'react-dom/client';

import "./asset/style/reset.css"
import "./asset/style/style.css"

import App from './App';

const container = document.getElementById('todo');
const root = createRoot(container); // createRoot(container!) if you use TypeScript
root.render(<App />);
