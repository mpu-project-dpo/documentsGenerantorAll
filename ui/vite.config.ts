import path from 'node:path'
import vue from '@vitejs/plugin-vue'
import { defineConfig, loadEnv } from 'vite'
// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  // eslint-disable-next-line node/prefer-global/process
  const env = loadEnv(mode, process.cwd())
  console.log(env)
  return {
    plugins: [vue()],
    resolve:
    {
      alias: {
        '@/*':
        path.resolve(__dirname, './src/*'),
        '@':
        path.resolve(__dirname, './src'),
      }
      ,
    },
    define: {
      VITE_API_BASE_URL: JSON.stringify(env.VITE_API_BASE_URL),
      VITE_SEND_FORM_PATH: JSON.stringify(env.VITE_SEND_FORM_PATH),
      VITE_GET_DPO_PATH: JSON.stringify(env.VITE_API_BASE_URL),
      VITE_CONF_POLICY_LINK: JSON.stringify(env.VITE_API_BASE_URL),
    }
    ,
  }
})
