import { createRouter, createWebHistory } from 'vue-router';

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'LoginPage',
      component: () => import('../pages/LoginPage.vue'),
      meta: {requiresAuth: false},
    },
  ],
});
