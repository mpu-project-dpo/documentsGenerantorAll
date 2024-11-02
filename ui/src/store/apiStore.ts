import axios from 'axios'
import { defineStore } from 'pinia'

export const useApiStore = defineStore('api', {
  state: () => ({
    api: axios.create({
      baseURL: import.meta.env.VITE_API_BASE_URL, // или ваш базовый URL
    }),
  }),
})
