import { createRouter, createWebHistory } from "vue-router";


export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: 'WelcomePage',
      path: '/',
      component: () => import('@/pages/WelcomePage.vue'),
      meta: { requiresAuth: false },
    },
    {
      name: 'LoginPage',
      path: '/login',
      component: () => import('@/pages/LoginPage.vue'),
      meta: { requiresAuth: false },
    },
    {
      name: 'LoginSetPasswordPage',
      path: '/setPassword',
      component: () => import('@/pages/LoginSetPasswordPage.vue'),
      meta: { requiresAuth: false },
    },
    {
      name: 'ProfilePage',
      path: '/profile',
      component: () => import('@/pages/ProfilePage.vue'),
      meta: { requiresAuth: true },
    },
    {
      name: 'ProrectorPage',
      path: '/prorector',
      component: () => import('@/pages/MainPageProrector.vue'),
      meta: { requiresAuth: true },
    },
    {
      name: 'Schedule',
      path: '/schedule',
      component: () => import('@/pages/SchedulePage.vue'),
      meta: { requiresAuth: false },
    },
    {
      name: 'Performace',
      path: '/performance',
      component: () => import('@/pages/PerformancePage.vue'),
      meta: { requiresAuth: true },
    },
  ],
});