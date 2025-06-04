import { useParams } from "react-router-dom";

import { EventDetails } from "./components/details";
import { getRandomImage } from "@utils/images";
import { Loading } from "@components/Loading";
import { LoadingError } from "@components/LoadingError";
import { NotFound } from "@pages/notFound";
import { Plan } from "@types/api/v1/plans/plan";
import { useAPI } from "@hooks/useAPI";

import "./event.css";

export const Event: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  const url =
    "/v1/events/plans?starts_at=2006-01-02T15:04:05&ends_at=2026-01-02T15:04:05";

  const { data, loading, error } = useAPI<Plan[]>("plans", url);

  const plans = data ?? [];

  const plan = plans.find((plan) => String(plan.basePlan.id) === id);

  return (
    <div className="event-page">
      {loading ? (
        <Loading />
      ) : error ? (
        <LoadingError error={error} />
      ) : plan === undefined ? (
        <NotFound />
      ) : (
        <div className="event-container">
          <div
            className="event-header"
            style={{ backgroundImage: `url(${getRandomImage()})` }}
          >
            <h1>{plan.basePlan.title}</h1>
          </div>

          <div className="event-content">
            <div className="event-main">
              <EventDetails plan={plan} />
            </div>

            <div className="event-sidebar">
              <div className="ticket-purchase-card">
                <h2>Get Tickets</h2>
                <p className="ticket-info">Select from available zones below</p>
                <button type="button" className="btn-select-all">
                  Select Best Available
                </button>

                <div className="ticket-availability-indicator">
                  <div className="indicator-item">
                    <span className="dot high"></span>
                    <span>High Availability</span>
                  </div>
                  <div className="indicator-item">
                    <span className="dot medium"></span>
                    <span>Limited Availability</span>
                  </div>
                  <div className="indicator-item">
                    <span className="dot low"></span>
                    <span>Almost Sold Out</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
