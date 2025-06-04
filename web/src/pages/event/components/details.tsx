import { useNavigate } from "react-router-dom";

import { formatAvailability, getAvailabilityClass } from "@utils/availability";
import { Plan } from "@types/api/v1/plans/plan";

import "./details.css";

interface EventDetailsProps {
  plan: Plan;
}

export const EventDetails: React.FC<EventDetailsProps> = ({ plan }) => {
  const navigate = useNavigate();

  const { basePlan, zones } = plan;

  const onClickBuy = () => {
    void navigate(`/event/${basePlan.id}/checkout`);
  };

  return (
    <div className="event-details">
      <section className="event-info-section">
        <h2>About This Event</h2>
        <div className="event-description">
          <p>{basePlan.title || "No description available for this event."}</p>
        </div>

        <div className="event-metadata">
          <div className="metadata-item">
            <span className="metadata-label">Date:</span>
            <span>{new Date(plan.planStartDate).toLocaleDateString()}</span>
          </div>
          <div className="metadata-item">
            <span className="metadata-label">Time:</span>
            <span>
              {new Date(plan.planStartDate).toLocaleTimeString([], {
                hour: "2-digit",
                minute: "2-digit",
              })}
            </span>
          </div>
        </div>
      </section>

      <section className="ticket-zones-section">
        <h2>Available Ticket Zones</h2>
        <p className="zone-instructions">
          Select a zone below to purchase tickets
        </p>

        {zones && zones.length > 0 ? (
          <div className="zones-list">
            {zones.map((zone) => (
              <div key={zone.id} className="zone-card">
                <div className="zone-header">
                  <h3 className="zone-name">{zone.name}</h3>
                  <div
                    className={`zone-availability ${getAvailabilityClass(zone.capacity)}`}
                  >
                    {formatAvailability(zone.capacity)}
                  </div>
                </div>
                <div className="zone-details">
                  <div className="zone-price">
                    ${zone.price.toFixed(2)}
                    <span className="price-note">per ticket</span>
                  </div>
                  <button
                    type="button"
                    onClick={onClickBuy}
                    className="btn-buy-tickets"
                  >
                    Buy Tickets
                  </button>
                </div>
              </div>
            ))}
          </div>
        ) : (
          <div className="no-zones-message">
            No ticket zones are currently available for this event.
          </div>
        )}
      </section>
    </div>
  );
};
