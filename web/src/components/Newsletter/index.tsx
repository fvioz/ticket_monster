import "./newsletter.css";

export const Newsletter: React.FC = () => {
  return (
    <section className="content-section newsletter">
      <div className="content-container">
        <div className="newsletter-content">
          <h2>Never Miss an Event</h2>
          <p>
            Subscribe to our newsletter for exclusive updates and special offers
          </p>
          <div className="newsletter-form">
            <input
              type="email"
              placeholder="Enter your email address"
              className="newsletter-input"
            />
            <button type="button" className="newsletter-button">
              Subscribe
            </button>
          </div>
        </div>
      </div>
    </section>
  );
};
