import vue from '@vitejs/plugin-vue';
import path from 'path';
import { defineConfig } from 'vite';

export default defineConfig({
    plugins: [vue()],
    resolve: {
        alias: {
            '@': path.resolve(__dirname, 'src'),
        }
    },
    server: {
        port: 3000,
        proxy: {
            '/v1': {
                target: 'http://localhost:8000',
                changeOrigin: true,
            }
        }
    }
});