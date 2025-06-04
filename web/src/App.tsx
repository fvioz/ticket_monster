import { Routes, Route } from "react-router-dom";

import { Header } from "@components/Header";
import { Footer } from "@components/Footer";

import { NotFound } from "@pages/notFound";

export const App: React.FC = () => {
  const routeWrapper = (Component: React.FC) => {
    // eslint-disable-next-line react/display-name
    return () => (
      <div className="route-wrapper">
        <Component />
      </div>
    );
  };

  // Wrapping components with routeWrapper to maintain consistent layout
  const NotFoundWrapper = routeWrapper(NotFound);

  return (
    <div className="app">
      <Header />
      <Routes>
        <Route path="*" element={<NotFoundWrapper />} />
      </Routes>
      <Footer />
    </div>
  );
};
