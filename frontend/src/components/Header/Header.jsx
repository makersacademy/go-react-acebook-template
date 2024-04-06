import "./Header.css"
import capybara from "../../static/img/capybara-logo.png"
import { Link } from "react-router-dom";
// import {useState } from "react";

export default function Header({isLoggedIn}) {

  function removeToken() {
    localStorage.removeItem("token");
  }

  return (
    <div className="header">
      <div className="logo">
        <img src={capybara} alt="logo" />
        <Link to="/">Cappybook</Link>
      </div>
      <div className="nav-bar">
        {isLoggedIn ? (
          <>
            <Link onClick={removeToken} to="/">Log out</Link>
            <Link to="/posts">Feed</Link>
          </>
        ) : (
          <>
            <Link to="/signup">Sign up</Link>
            <Link to="/login">Log in</Link>
          </>
        )}
      </div>
    </div>
  );
}
