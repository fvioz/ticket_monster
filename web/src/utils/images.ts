const IMAGES = [
  "https://images.unsplash.com/photo-1506748686214-e9df14d4d9d0?auto=format&fit=crop&w=400&q=60",
  "https://images.unsplash.com/photo-1521747116042-5a810fda9664?auto=format&fit=crop&w=400&q=60",
  "https://images.unsplash.com/photo-1498050108023-c5249f4df085?auto=format&fit=crop&w=400&q=60",
];

export const getRandomImage = () => {
  return IMAGES[Math.floor(Math.random() * IMAGES.length)];
};
