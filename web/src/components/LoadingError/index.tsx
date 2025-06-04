import "./loadingError.css";

interface LoadingErrorProps {
  error: string | null;
}

export const LoadingError: React.FC<LoadingErrorProps> = ({ error }) => {
  return (
    <div className="error-container">
      <div className="error-icon">⚠️</div>
      <p>Error loading the event</p>
      {error && <p>{error}</p>}
    </div>
  );
};
