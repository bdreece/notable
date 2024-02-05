import { defineConfig } from 'vite';
import autoprefixer from 'autoprefixer';
import cssnano from 'cssnano';
import dts from 'vite-plugin-dts';

export default defineConfig({
    plugins: [dts()],
    build: {
        assetsDir: '',
        lib: {
            entry: 'src/index.ts',
            name: 'Notable',
            fileName: 'index',
        }
    },
    css: {
        postcss: {
            plugins: [
                autoprefixer(),
                cssnano(),
            ]
        },
        preprocessorOptions: {
            scss: {
                file: 'src/styles/index.scss',
                includePaths: ['src/styles', 'node_modules'],
            },
        },
    },
});
