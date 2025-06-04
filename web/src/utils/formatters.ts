export const formatDate = (dateString: string | number): string => {
  const options: Intl.DateTimeFormatOptions = {
    year: "numeric",
    month: "long",
    day: "numeric",
  };
  return new Date(dateString).toLocaleDateString(undefined, options);
};

export const formatTitle = (title: string): string => {
  return title.charAt(0).toUpperCase() + title.slice(1);
};
