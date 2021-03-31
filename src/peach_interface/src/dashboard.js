import React from 'react';
import Sidebar from './sidebar';
import Overview from './overview'
import { Route, Switch } from 'react-router-dom'
import './css/dashboard.css';


function Dashboard() {
  return (
    <div class="app">
      <Sidebar />
      <div class="main">
        <Switch>
          <Route path = "/dashboard/overview">
            <Overview />
          </Route>
        </Switch>
      </div>
    </div>
    )
};

export default Dashboard
