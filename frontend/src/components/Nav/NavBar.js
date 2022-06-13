import React from 'react'
import './navbar.css'
import {  Link } from "react-router-dom";
const NavBar = () => {
  return (
    <div>
   <nav className='navbar'>
    <Link className='instaclone-logo' to='/'>Products GO</Link>
    <div className='navbar-ul'>
        <ul className='navbar-content'>
          <li><Link to="/products">products</Link></li>
            <li><Link to='/login'>login</Link></li>
            <li><Link to='/register'>register</Link></li>
        </ul>
    </div>
   </nav>
    </div>
  )
}

export default NavBar