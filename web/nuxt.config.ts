// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  runtimeConfig: {
    supabaseRestUrl: process.env.SUPABASE_REST_URL,
    supabaseApiKey: process.env.SUPABASE_API_KEY,
  },
  app: {
    pageTransition: {
      name: "page",
      mode: "out-in",
    },
  },
  nitro: {
    preset: "static",
  },
  modules: ["@nuxthq/ui", "nuxt-icon"],
});
