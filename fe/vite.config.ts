/// <reference types="vitest/config" />

import preact from "@preact/preset-vite";
import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [preact(), tailwindcss()],
  test: {
    environment: "jsdom",
    coverage: {
      provider: "v8",
      reporter: ["text", "lcovonly"],
      reportsDirectory: "coverage",
      clean: true,
      exclude: ["src/client/**/*.{ts,tsx}"],
    },
  },
});
