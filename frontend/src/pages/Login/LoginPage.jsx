import { useState } from "react";
import {Link, useNavigate} from "react-router-dom";
import "./LoginPage.css";
import { login } from "../../services/authentication";

export const LoginPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const token = await login(email, password);
      localStorage.setItem("token", token);
      navigate("/posts");
    } catch (err) {
      navigate("/login");
      setErrorMessage("User does not exist or invalid credentials.");
    }
  };

  const handleEmailChange = (event) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  return (
      <div className="container">
          <div className="login-container">
            <h1> Welcome back to acebook!</h1>
            <form className="login-form" onSubmit={handleSubmit}>
              <h1>Login to your account</h1>
              <h2>new here? <Link to="/signup" style={{ color: 'royalblue' }}>signup</Link></h2>
              {errorMessage && <p style={{ color: "red" }}>{errorMessage}</p>}
                <input
                  id="email"
                  type="text"
                  value={email}
                  onChange={handleEmailChange}
                  placeholder="Email 📩"
                />
                <input
                  id="password"
                  type="password"
                  value={password}
                  onChange={handlePasswordChange}
                  placeholder="Password 🔒"
                />
                <input role="submit-button" id="submit" type="submit" value="Login 🚀" />
            </form>
          </div>
      </div>
  );
};
