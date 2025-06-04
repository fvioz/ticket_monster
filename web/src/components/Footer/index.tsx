import React from "react";
import { Link } from "react-router-dom";
import "./footer.css";

export const Footer: React.FC = () => {
  return (
    <footer className="footer">
      <div className="footer-container">
        <div className="footer-section">
          <h3>TickerMonster</h3>
          <p>
            Your premier destination for concert and event tickets. We connect
            fans to unforgettable experiences.
          </p>
          <div className="social-icons">
            <a href="https://facebook.com" aria-label="Facebook">
              <i className="fab fa-facebook"></i>
            </a>
            <a href="https://twitter.com" aria-label="Twitter">
              <i className="fab fa-twitter"></i>
            </a>
            <a href="https://instagram.com" aria-label="Instagram">
              <i className="fab fa-instagram"></i>
            </a>
          </div>
        </div>

        <div className="footer-section">
          <h3>Quick Links</h3>
          <ul className="footer-links">
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/events">Events</Link>
            </li>
            <li>
              <Link to="/venues">Venues</Link>
            </li>
            <li>
              <Link to="/artists">Artists</Link>
            </li>
          </ul>
        </div>

        <div className="footer-section">
          <h3>Support</h3>
          <ul className="footer-links">
            <li>
              <Link to="/faq">FAQ</Link>
            </li>
            <li>
              <Link to="/contact">Contact Us</Link>
            </li>
            <li>
              <Link to="/terms">Terms of Service</Link>
            </li>
            <li>
              <Link to="/privacy">Privacy Policy</Link>
            </li>
          </ul>
        </div>

        <div className="footer-section">
          <h3>Contact</h3>
          <address>
            <p>
              Fake Street 123
              <br />
              Springfield, USA
            </p>
            <p>
              Email: support@test.com
              <br />
              Phone: (555) 123-4567
            </p>
          </address>
        </div>
      </div>

      <div className="footer-bottom">
        <p>
          &copy; {new Date().getFullYear()} Concert Ticketing. All rights
          reserved.
        </p>
      </div>
    </footer>
  );
};
