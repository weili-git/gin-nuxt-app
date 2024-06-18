// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,
  devtools: { enabled: true },
  modules: [
    '@element-plus/nuxt',
    '@pinia/nuxt',
    '@pinia-plugin-persistedstate/nuxt',
  ],
  elementPlus: { /** Options */ },
  runtimeConfig: {
    public:{ // expose to client-side
      apiBase: process.env.BASE_URL || "http://127.0.0.1:8086", // localhost not recommend
      apiKey: process.env.API_KEY, // save your api in .env
    }
  },
  devServer: { // nuxt 3
    port: 3001,
    host: '127.0.0.1',
  },

})
