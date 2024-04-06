import Header from '../Header/Header'
import Footer from '../Footer/Footer'
import { Outlet } from 'react-router-dom';
import "./Layout.css"
import React from 'react';
import { useState } from 'react';
import { useLocation } from 'react-router-dom';

export const Layout = ({ children }) => {
  const [isLoggedIn, setIsLoggedIn] = useState(!!localStorage.getItem('token'));
  const location = useLocation();

  React.useEffect(() => {
    setIsLoggedIn(!!localStorage.getItem('token'));
  }, [location]);
  return (
    <>
    <div className='layout-container'>
      <Header isLoggedIn={isLoggedIn}/>
      {children}
      <Outlet/>
      <Footer />
    </div>
    
    </>
  )
}
