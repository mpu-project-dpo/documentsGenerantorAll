import { createRouter, createWebHistory } from 'vue-router'
import { routes } from './routes'

export const router = createRouter({
  history: createWebHistory(),
  routes,
})

let first = true

router.beforeEach(async (_, from) => {
  if (!from.name && first) {
    first = false
    await router.push({ name: 'education' })
  }
})
