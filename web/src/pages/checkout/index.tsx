import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import { formatDate } from "@utils/formatters";

import "./checkout.css";

export const Checkout: React.FC = () => {
  const [orderDetails] = useState({
    orderNumber: `ORD-${Math.floor(Math.random() * 1000000)}`,
    purchaseDate: formatDate(Date.now()),
  });

  const [countdown, setCountdown] = useState(30);

  const navigate = useNavigate();

  useEffect(() => {
    const timer = setTimeout(() => {
      if (countdown > 0) {
        setCountdown(countdown - 1);
      }
      if (countdown === 0) {
        void navigate(`/`);
      }
    }, 1000);

    return () => clearTimeout(timer);
  }, [countdown]);

  return (
    <div className="checkout-page">
      <div className="checkout-container">
        <div className="checkout-header">
          <div className="success-icon">✓</div>
          <h1>Thank You for Your Purchase!</h1>
          <p className="confirmation-message">
            Your tickets have been confirmed and sent to your email
          </p>
        </div>

        <div className="order-summary">
          <h2>Order Summary</h2>

          <div className="order-info">
            <div className="order-info-item">
              <span className="info-label">Order Number:</span>
              <span className="info-value">{orderDetails.orderNumber}</span>
            </div>
            <div className="order-info-item">
              <span className="info-label">Purchase Date:</span>
              <span className="info-value">{orderDetails.purchaseDate}</span>
            </div>
          </div>
        </div>

        <div className="redirect-message">
          <p>You will be redirected to tthe home page {countdown} seconds...</p>
        </div>
      </div>
    </div>
  );
};
