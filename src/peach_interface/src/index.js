import React from 'react';
import ReactDOM from 'react-dom';
import * as serviceWorker from './serviceWorker';
import { BrowserRouter as Router, Route } from 'react-router-dom'
import Dashboard from './dashboard';
import Background from './background'
import './css/dashboard.css';


ReactDOM.render(
  <React.StrictMode>
    <Router>
        <Route path = "/dashboard" component={ Dashboard } />
    </Router>
    <Background />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
