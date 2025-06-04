export const getAvailabilityClass = (availability?: number): string => {
  if (!availability) return "unknown";
  if (availability > 0.7) return "high";
  if (availability > 0.3) return "medium";
  return "low";
};

export const formatAvailability = (availability?: number): string => {
  if (!availability) return "Unknown";
  if (availability > 0.7) return "High Availability";
  if (availability > 0.3) return "Limited Availability";
  return "Almost Sold Out";
};
