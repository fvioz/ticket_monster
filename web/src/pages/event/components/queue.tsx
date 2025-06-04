import { useMemo } from "react";

import { getRandomImage } from "@utils/images";

import "./queue.css";

interface QueueParams {
  position: number;
  queueLimit: number;
  title: string;
}

export const Queue: React.FC<QueueParams> = ({
  position,
  queueLimit,
  title,
}: QueueParams) => {
  const image = useMemo(() => getRandomImage(), []);
  const calculateProgress = () => {
    if (position <= 0 || queueLimit <= 0) return 0;
    const progress = ((queueLimit - position) / queueLimit) * 100;
    return Math.min(Math.max(progress, 0), 100); // Ensure progress is between 0-100
  };

  const progress = calculateProgress();

  const estimatedTimeMinutes = Math.ceil(position * 0.5); // Rough estimate of 30 seconds per person

  const estimationText =
    estimatedTimeMinutes > 0
      ? `${estimatedTimeMinutes} minutes`
      : "Calculating...";

  const positionText =
    position === 0 ? (
      <span className="position-number position-calculating">Calculating</span>
    ) : (
      <span className="position-number">{position}</span>
    );

  return (
    <div className="queue-page">
      <div
        className="queue-header"
        style={{ backgroundImage: `url(${image})` }}
      >
        <h1>{title}</h1>
      </div>

      <div className="queue-container">
        <div className="queue-content">
          <div className="queue-info">
            <h2>You&apos;re in line!</h2>
            <p>Our system is currently experiencing high demand.</p>
            <p>Please wait while we prepare your access to purchase tickets.</p>

            <div className="queue-status">
              <div className="queue-position">
                {positionText}
                <span className="position-label">Your position</span>
              </div>

              <div className="queue-total">
                <span className="total-number">{queueLimit}</span>
                <span className="total-label">Total capacity</span>
              </div>
            </div>

            <div className="queue-progress-container">
              <div
                className="queue-progress-bar"
                style={{ width: `${progress == 0 ? "100" : progress}%` }}
              ></div>
            </div>

            <div className="queue-time">
              <p>
                Estimated wait time: <strong>{estimationText}</strong>
              </p>
              <p className="queue-tips">
                Please don&apos;t refresh this page or you may lose your place
                in line
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
