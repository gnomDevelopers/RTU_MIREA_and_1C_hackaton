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

    {
      path: '/profile',
      name: 'ProfilePage',
      component: () => import('../pages/ProfilePage.vue'),
      meta: {requiresAuth: false},
    },
     {
       path: '/events',
       name: 'EventsPage',
       component: () => import('../pages/EventsPage.vue'),
       meta: {requiresAuth: false},
     },
    {
      path: '/team',
      name: 'TeamPage',
      component: () => import('../pages/TeamPage.vue'),
      meta: {requiresAuth: false},
    },


  ],
});
