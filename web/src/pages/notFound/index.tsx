import React from "react";
import { Link } from "react-router-dom";
import "./notFound.css";

export const NotFound: React.FC = () => {
  return (
    <div className="not-found-page">
      <div className="not-found-container">
        <div className="not-found-content">
          <div className="error-code">404</div>
          <h1>Page Not Found</h1>
          <p>
            The page you&apos;re looking for doesn&apos;t exist or has been
            moved.
          </p>

          <div className="not-found-actions">
            <Link to="/" className="btn-primary">
              Back to Home
            </Link>
          </div>

          <div className="not-found-suggestions">
            <h2>You might be interested in:</h2>
            <ul>
              <li>
                <Link to="/events">Upcoming Events</Link>
              </li>
              <li>
                <Link to="/faq">Frequently Asked Questions</Link>
              </li>
              <li>
                <Link to="/contact">Contact Support</Link>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
};
