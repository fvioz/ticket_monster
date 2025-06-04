import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";
import path from "path";

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd());

  return {
    define: {
      API_DOMAIN: JSON.stringify(env.API_DOMAIN || ""),
      WS_DOMAIN: JSON.stringify(env.WS_DOMAIN || ""),
    },
    plugins: [react()],
    preview: {
      host: env.VITE_HOST || "0.0.0.0",
      port: parseInt(env.VITE_PORT || "3000"),
    },
    server: {
      host: env.VITE_HOST || "0.0.0.0",
      port: parseInt(env.VITE_PORT || "3000"),
    },
    build: {
      outDir: "dist",
    },
    resolve: {
      alias: {
        "@assets": path.resolve(__dirname, "./src/assets"),
        "@components": path.resolve(__dirname, "./src/components"),
        "@hooks": path.resolve(__dirname, "./src/hooks"),
        "@pages": path.resolve(__dirname, "./src/pages"),
        "@types": path.resolve(__dirname, "./src/types"),
        "@utils": path.resolve(__dirname, "./src/utils"),
      },
    },
  };
});
