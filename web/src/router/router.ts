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
      name: 'MainPage',
      path: '/main',
      component: () => import('@/pages/MainPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      name: 'AccountsPage',
      path: '/accounts',
      component: () => import('@/pages/AccountsPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      name: 'SchedulePage',
      path: '/schedule',
      component: () => import('@/pages/SchedulePage.vue'),
      meta: { requiresAuth: true },
    },
    {
      name: 'PerformancePage',
      path: '/performance',
      component: () => import('@/pages/PerformancePage.vue'),
      meta: { requiresAuth: true },
    },
    {
      name: 'GroupCorrectPage',
      path: '/group',
      component: () => import('@/pages/GroupCorrectPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      name: 'AttendancePage',
      path: '/attendance',
      component: () => import('@/pages/AttendancePage.vue'),
      meta: { requiresAuth: true },
    },
  ],
});