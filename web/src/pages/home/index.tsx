import { FeaturedEvents } from "@components/FeaturedEvents";
import { Newsletter } from "@components/Newsletter";

import "./home.css";

export const Home: React.FC = () => {
  return (
    <div className="home-page">
      <FeaturedEvents />
      <Newsletter />
    </div>
  );
};
