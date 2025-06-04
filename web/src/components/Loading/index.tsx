import "./loading.css";

export const Loading: React.FC = () => {
  return (
    <div className="loading-container">
      <div className="loading-spinner"></div>
      <p>Loading ...</p>
    </div>
  );
};
