export default defineAppConfig({
  title: "Binocular | Advanced Subdomain Finder",
  ui: {
    primary: "purple",
    gray: "neutral",
  },
  binocularRestApi: process.env.BINOCULAR_REST_API || "http://127.0.0.1:8080/binocular",
});
