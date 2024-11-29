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
    //   студенческие возможности
    {
      path: '/profile',
      name: 'ProfilePage',
      component: () => import('../pages/ProfilePage.vue'),
      meta: {requiresAuth: false},
    },
    {
      path: '/incoming',
      name: 'IncomingOffersPage',
      component: () => import('../pages/IncomingOffersPage.vue'),
      meta: {requiresAuth: false},
    },
    {
      path: '/myresponses',
      name: 'MyResponsesPage',
      component: () => import('../pages/MyResponsesPage.vue'),
      meta: {requiresAuth: false},
    },

  //     инфа для хров
    {
      path: '/candidatessearch',
      name: 'CandidatesSearchPage',
      component: () => import('../pages/CandidatesSearchPage.vue'),
      meta: {requiresAuth: false},
    },
    {
      path: '/candidatesresponses',
      name: 'CandidatesResponsesPage',
      component: () => import('../pages/CandidatesResponsesPage.vue'),
      meta: {requiresAuth: false},
    },
    {
      path: '/candidatesprofile',
      name: 'CandidatesProfilePage',
      component: () => import('../pages/CandidatesProfilePage.vue'),
      meta: {requiresAuth: false},
    },

  ],
});
