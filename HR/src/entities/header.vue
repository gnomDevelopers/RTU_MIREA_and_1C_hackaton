<template>
  <div v-if="!isLoginPage" class="flex flex-row justify-between items-stretch mb:h-24 uus:h-16 shrink-0 w-full px-5 header-shadow border-b-2 z-10 bg-color-light">
    <div class="h-full grid content-center md:ml-16">
      <div @click="$router.push({name: 'MainPage'})" class="logo-hr">WorkVUZ+</div>
    </div>
    <!--  это для норм размера и вкладки для студентоффф-->
    <div v-if="isBigScreen && isProfilePage" class="flex flex-row justify-center items-center gap-4 mt-2">
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'MyResponsesPage'})">Мои отклики</button>
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'IncomingOffersPage'})">Входящие предложения</button>
    </div>
    <div v-if="isBigScreen && isMyResponsesPage" class="flex flex-row justify-center items-center gap-4 mt-2">
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'MyResponsesPage'})">Мои отклики</button>
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'IncomingOffersPage'})">Входящие предложения</button>
    </div>
    <div v-if="isBigScreen && isIncomingOffersPage" class="flex flex-row justify-center items-center gap-4 mt-2">
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'MyResponsesPage'})">Мои отклики</button>
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'IncomingOffersPage'})">Входящие предложения</button>
    </div>

    <!--  это для норм размера и вкладки для ХРов-->
    <div v-if="isBigScreen && isCandidatesSearchPage" class="flex flex-row justify-center items-center gap-4 mt-2">
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'CandidatesSearchPage'})">Поиск кандидатов</button>
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'CandidatesResponsesPage'})">Отклики кандидатов</button>
    </div>
    <div v-if="isBigScreen && isCandidatesResponsesPage" class="flex flex-row justify-center items-center gap-4 mt-2">
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'CandidatesSearchPage'})">Поиск кандидатов</button>
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'CandidatesResponsesPage'})">Отклики кандидатов</button>
    </div>

    <div class="flex flex-row gap-x-2 justify-center items-center md:mr-16">

      <img v-if="isProfilePage || isMyResponsesPage || isIncomingOffersPage" class="w-8 mb:w-10 h-8 mb:h-10" src="../assets/icons/icon-profile.svg"/>
      <img v-if="isCandidatesSearchPage || isCandidatesResponsesPage" class="w-8 mb:w-10 h-8 mb:h-10" src="../assets/icons/icon-company.svg"/>

      <div v-if="isProfilePage || isMyResponsesPage || isIncomingOffersPage" @click="$router.push({name: 'ProfilePage'})" class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300">
        Личный кабинет
      </div>
    </div>
  </div>
<!--  это для телефона и вкладки для студентоффф-->
  <div v-if="isSmallScreen && isProfilePage" class="flex flex-row justify-center gap-4 mt-2 pb-2 border-b-2">
    <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'MyResponsesPage'})">Мои отклики</button>
    <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'IncomingOffersPage'})">Входящие предложения</button>
  </div>
  <div v-if="isSmallScreen && isMyResponsesPage" class="flex flex-row justify-center gap-4 mt-2 pb-2 border-b-2">
    <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'MyResponsesPage'})">Мои отклики</button>
    <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'IncomingOffersPage'})">Входящие предложения</button>
  </div>
  <div v-if="isSmallScreen && isIncomingOffersPage" class="flex flex-row justify-center gap-4 mt-2 pb-2 border-b-2">
    <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'MyResponsesPage'})">Мои отклики</button>
    <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'IncomingOffersPage'})">Входящие предложения</button>
  </div>

  <!--  это телефона и вкладки для ХРов-->
  <div v-if="isSmallScreen && isCandidatesSearchPage" class="flex flex-row justify-center items-center gap-4 mt-2 pb-2 border-b-2">
    <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'CandidatesSearchPage'})">Поиск кандидатов</button>
    <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'CandidatesResponsesPage'})">Отклики кандидатов</button>
  </div>
  <div v-if="isSmallScreen && isCandidatesResponsesPage" class="flex flex-row justify-center items-center gap-4 mt-2 pb-2 border-b-2">
    <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'CandidatesSearchPage'})">Поиск кандидатов</button>
    <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-red-100 active:bg-red-300" @click="$router.push({name: 'CandidatesResponsesPage'})">Отклики кандидатов</button>
  </div>




</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useUserInfoStore } from '@/stores/userInfoStore';

export default {
  data() {
    return {
      isSmallScreen: false,
      isBigScreen: false,
    };
  },
  computed: {
    ...mapStores(useUserInfoStore),
    isLoginPage(){
      return this.$route.fullPath === '/';
    },
    isProfilePage(){
      return this.$route.fullPath === '/profile';
    },
    isMyResponsesPage(){
      return this.$route.fullPath === '/myresponses';
    },
    isIncomingOffersPage(){
      return this.$route.fullPath === '/incoming';
    },
    isCandidatesSearchPage() {
      return this.$route.fullPath === '/candidatessearch';
    },
    isCandidatesResponsesPage(){
      return this.$route.fullPath === '/candidatesresponses';
    }
  },
  methods: {
    updateScreenSize() {
      this.isSmallScreen = window.innerWidth <= 768;
      this.isBigScreen = window.innerWidth > 768;
    },
  },
  mounted() {
    this.updateScreenSize();
    window.addEventListener('resize', this.updateScreenSize);
  },
  beforeUnmount() {
    window.removeEventListener('resize', this.updateScreenSize);
  },

};
</script>

