import React from "react";
import { Link } from "react-router-dom";
import "./header.css";

export const Header: React.FC = () => {
  return (
    <header className="header">
      <div className="header-container">
        <div className="logo">
          <Link to="/">
            <h1>Ticket Monster</h1>
          </Link>
        </div>

        <nav className="main-nav">
          <ul></ul>
        </nav>

        <div className="user-actions">
          <button type="button" className="btn-login">
            Log In
          </button>
          <button type="button" className="btn-signup">
            Sign Up
          </button>
        </div>
      </div>
    </header>
  );
};
