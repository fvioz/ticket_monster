import { Routes, Route } from "react-router-dom";

import { Header } from "@components/Header";
import { Footer } from "@components/Footer";

import { Home } from "@pages/home";
import { Event } from "@pages/event";
import { Checkout } from "@pages/checkout";
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
  const HomeWrapper = routeWrapper(Home);
  const EventWrapper = routeWrapper(Event);
  const CheckoutWrapper = routeWrapper(Checkout);
  const NotFoundWrapper = routeWrapper(NotFound);

  return (
    <div className="app">
      <Header />
      <Routes>
        <Route path="/" element={<HomeWrapper />} />
        <Route path="/event/:id" element={<EventWrapper />} />
        <Route path="/event/:id/checkout" element={<CheckoutWrapper />} />
        <Route path="*" element={<NotFoundWrapper />} />
      </Routes>
      <Footer />
    </div>
  );
};
