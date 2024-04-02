import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";
import "./SignupPage.css";
import { signup } from "../../services/authentication";

export const SignupPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isSuccess, setIsSuccess] = useState(false); // State for success message
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      await signup(email, password);
   
      setIsSuccess(true); // Set isSuccess to true on successful sign-up
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
