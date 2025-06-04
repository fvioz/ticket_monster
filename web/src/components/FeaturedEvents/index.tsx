import React from "react";
import { Link } from "react-router-dom";

import { EventCard } from "@components/EventCard";
import { Loading } from "@components/Loading";
import { LoadingError } from "@components/LoadingError";
import { Plan } from "@types/api/v1/plans/plan";
import { useAPI } from "@hooks/useAPI";

import "./featuredEvents.css";

const MAX_EVENTS = 3;

export const FeaturedEvents: React.FC = () => {
  const url =
    "/v1/events/plans?starts_at=2006-01-02T15:04:05&ends_at=2026-01-02T15:04:05";
  const { data, loading, error } = useAPI<Plan[]>("plans", url);

  const plans = data ?? [];

  const upcomingEvents = plans
    .sort(
      (a, b) =>
        new Date(a.planStartDate).getTime() -
        new Date(b.planStartDate).getTime(),
    )
    .slice(0, MAX_EVENTS);

  return (
    <section className="content-section featured-events">
      <div className="content-container">
        <div className="section-header">
          <h2>Featured Events</h2>
          <Link to="/events" className="view-all-link">
            View All Events
          </Link>
        </div>

        {loading ? (
          <Loading />
        ) : error ? (
          <LoadingError error={error} />
        ) : (
          <div className="events-grid">
            {upcomingEvents.map((plan) => (
              <EventCard key={plan.id} plan={plan} />
            ))}
          </div>
        )}
      </div>
    </section>
  );
};
