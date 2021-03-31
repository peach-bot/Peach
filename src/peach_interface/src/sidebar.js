import React from 'react';
import { ProSidebar, SidebarHeader, SidebarContent, Menu, MenuItem } from 'react-pro-sidebar';
import 'react-pro-sidebar/dist/css/styles.css';
import { FaTachometerAlt, FaList } from 'react-icons/fa';
import { Link } from 'react-router-dom'
import logo from './img/logo.ico'
import './css/sidebar.css';

function Sidebar() {
    return(
        <ProSidebar>
            <SidebarHeader>
                <div>
                    <img src={logo} alt="logo"></img>
                    <Link to={'/'} class="sidebar-title">Peach</Link>
                </div>
            </SidebarHeader>
            <SidebarContent>
            <Menu iconShape="circle">
                <MenuItem icon={<FaTachometerAlt />}><Link to={'/dashboard/overview'}>Overview</Link></MenuItem>
                <MenuItem icon={<FaList />}><Link to={'/dashboard/settings'}>Settings</Link></MenuItem>
            </Menu>
            </SidebarContent>
        </ProSidebar>
    )
}


export default Sidebar;