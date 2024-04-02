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
  const [isSuccess, setIsSuccess] = useState(false); // State for success message
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const token = await signup(email, password, username, image);
      localStorage.setItem("token", token);
      setIsSuccess(true);
      setTimeout(() => {
        navigate("/posts");
      }, 1000);

    } catch (err) {
      console.error(err);
      navigate("/signup");
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
        <div className="left-half">
          <div className="signup-container">
            <form className="signup-form" onSubmit={handleSubmit}>
              <h1>Sign Up</h1>
              <h2>Create Your Account</h2>
              <input
                  id="email"
                  type="text"
                  value={email}
                  onChange={handleEmailChange}
                  placeholder="Email"
              />
              <input
                  id="password"
                  type="password"
                  value={password}
                  onChange={handlePasswordChange}
                  placeholder="Password"
              />
              <input 
                id="username"
                type="text"
                value={username}
                onChange={handleUsernameChange}
                placeholder="Username"
                />
              <input
                id="image"
                type="file" 
                onChange={handleImageChange}
                />
              <input
                  role="submit-button"
                  id="submit"
                  type="submit"
                  value="Submit"
              />
            </form>
          </div>
          {isSuccess && <p>Account created successfully!</p>}
        </div>
        <div className="right-half">
          <h3>Login Here</h3>
          <p>If you already have an account, login here:</p>
          <Link className="login" to="/login">Login</Link>
        </div>
      </div>
  );
};
