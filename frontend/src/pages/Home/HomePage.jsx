import { Link } from "react-router-dom";

import "./HomePage.css";

export const HomePage = () => {
    return (
        <div className="home">
            <div className="welcome-message">
                <h1>Welcome to Cappybook!</h1>
            </div>
                <div className="button-container">
                    <Link className="login" to="/login">Log In</Link>
                    <Link className="signup" to="/signup">Sign Up</Link>
                </div>
        </div>
    );
};

