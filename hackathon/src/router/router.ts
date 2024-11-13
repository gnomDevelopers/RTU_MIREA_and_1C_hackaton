import { createRouter, createWebHistory } from "vue-router";


export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: 'WelcomePage',
      path: '/',
      component: () => import('../pages/WelcomePage.vue'),
    },
    {
      name: 'LoginPage',
      path: '/login',
      component: () => import('../pages/LoginPage.vue'),
    },
    {
      name: 'LoginSetPasswordPage',
      path: '/setPassword',
      component: () => import('../pages/LoginSetPasswordPage.vue'),
    },
  ],
});