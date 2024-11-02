import type { RouteRecordRaw } from 'vue-router'
import MainPage from '@/components/pages/MainPage.vue'

export const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: MainPage,
    children: [
      { path: '', name: 'education', component: () => import('@/components/organisms/EducationForm.vue') },
      { path: 'contacts', name: 'contacts', component: () => import('@/components/organisms/ContactsForm.vue') },
      { path: 'docs', name: 'docs', component: () => import('@/components/organisms/DocsForm.vue') },
    ],
  },
]
