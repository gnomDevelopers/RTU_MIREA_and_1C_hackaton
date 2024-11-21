import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

import { VitePWA } from 'vite-plugin-pwa';


// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    VitePWA({
      registerType: "autoUpdate",
      outDir: 'dist',
      manifest: {
        name: 'VUZ+ App',
        short_name: 'VUZ+',
        description: 'Mobile App to VUZ+',
        theme_color: '#EBF5FF',
        background_color: '#EBF5FF', 
        display: 'standalone',      
        icons: [
          {
            src: 'assets/icon-app-png.png',
            sizes: '640x640',
            type: 'image/png',
          },
          {
            src: 'assets/icon-app.svg',
            sizes: 'any',
            type: 'image/svg+xml',
          },
        ],
        screenshots: [
          {
            src: 'assets/screenshot-wide-800-452.jpg',
            sizes: '800x452',
            form_factor: 'wide',
          },
          {
            src: 'assets/screenshot-narrow-505-807.jpg',
            sizes: '505x807',
            form_factor: 'narrow',
          },
        ],
      }
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
})
