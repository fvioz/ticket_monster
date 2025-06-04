import React from "react";
import { Link } from "react-router-dom";

import { Plan } from "@types/api/v1/plans/plan";
import { formatDate, formatTitle } from "@utils/formatters";
import { getRandomImage } from "@utils/images";

import "./eventCard.css";

interface EventCardProps {
  plan: Plan;
}

export const EventCard: React.FC<EventCardProps> = ({ plan }) => {
  const minPrice = plan.zones.sort((a, b) => a.price - b.price)[0]?.price;

  return (
    <div className="event-card">
      <Link to={`/event/${plan.basePlan.id}`} className="event-card-link">
        <div className="event-image">
          <img src={getRandomImage()} alt={formatTitle(plan.basePlan.title)} />
          <span className="featured-badge">Featured</span>
        </div>
        <div className="event-details">
          <h3 className="event-title">{plan.basePlan.title}</h3>
          <div className="event-metadata">
            <span className="event-date">{formatDate(plan.planStartDate)}</span>
            <span className="event-venue">Barcelona</span>
          </div>
          {minPrice && <div className="event-price">From ${minPrice}</div>}
        </div>
      </Link>
    </div>
  );
};
