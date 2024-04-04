import { useState } from "react";

import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";
import "./SignupPage.css";
import { signup } from "../../services/authentication";

export const SignupPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [username, setUsername] = useState("");
  const [image, setImage] = useState("");
  const [errorMessage, setErrorMessage] = useState(""); // State for error message
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const token = await signup(email, password, username, image);
      localStorage.setItem("token", token);
      setTimeout(() => {
        navigate("/posts");
      }, 1000);
    } catch (err) {
      setErrorMessage(err.message); // Set error message state with the message from the backend
    }
  };


  const handleEmailChange = (event) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleUsernameChange = (event) => {
    setUsername(event.target.value)
  };

  const handleImageChange = (event) => {
    const file = event.target.files[0];
    setImage(file);
  };

  return (
      <div className="container">
          <div className="signup-container">
            <h1> Welcome to Acebook!</h1>
            <form className="signup-form" onSubmit={handleSubmit}>
              <h1>Create your account</h1>
              <h3> already have an account? <Link to="/login" style={{ color: 'royalblue' }}>login here</Link></h3>
              {/*<h2>Create your account</h2>*/}
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
              <input
                  id="username"
                  type="text"
                  value={username}
                  onChange={handleUsernameChange}
                  placeholder="Username 💡"
              />
              <label htmlFor="image" className="file-upload-label">
                Upload profile picture 🖼️
                <input
                    id="image"
                    type="file"
                    onChange={handleImageChange}
                    className="file-upload-input"
                />
                {image && (
                    <span className="file-name">[{image.name}]</span>
                    )}
              </label>
              <input
                  role="submit-button"
                  id="submit"
                  type="submit"
                  value="Lets Go! ✅ "
              />
            </form>
            <div className="error-container">
              {errorMessage && <p className="error-message">{errorMessage}</p>}
            </div>
          </div>
      </div>
  );
};
