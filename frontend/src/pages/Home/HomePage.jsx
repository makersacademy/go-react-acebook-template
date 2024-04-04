import { Link } from "react-router-dom";

import "./HomePage.css";

export const HomePage = () => {
  return (
    <div className="home">
      <div className="left-half">
        <h1>Welcome to Acebook!</h1>
        <h2>Login to Your Account</h2>
        <Link className="login" to="/login">Log In</Link>
      </div>
      <div className="right-half">
        <h3>New Here?</h3>
        <p>Sign up and start sharing moments with your friends today! Join our community and explore endless possibilities together.</p>
        <Link className="signup" to="/signup">Sign Up</Link>
      </div>
    </div>
  );
};

