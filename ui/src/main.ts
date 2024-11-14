import VMask from '@ssibrahimbas/v-mask'
import { createPinia } from 'pinia'
import { createApp } from 'vue'
import App from './App.vue'
import { router } from './router'
import './styles/style.css'
import './styles/fonts.css'

createApp(App)
  .use(router)
  .use(createPinia())
  .use(VMask)
  .mount('#app')
