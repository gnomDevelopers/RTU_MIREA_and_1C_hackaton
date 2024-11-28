<template>
  <div v-if="!isLoginPage" class="flex flex-row justify-between items-stretch mb:h-24 uus:h-16 shrink-0 w-full px-5 header-shadow border-b-2 z-10 bg-color-light">
    <div class="h-full grid content-center md:ml-16">
      <div @click="$router.push({name: 'MainPage'})" class="logo-hr">EventVUZ+</div>
    </div>
    <!--  это для норм размера и вкладки для студентоффф-->
    <div v-if="isBigScreen && isProfilePage" class="flex flex-row justify-center items-center gap-4 mt-2">
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-purple-100 active:bg-purple-200" @click="$router.push({name: 'EventsPage'})">Мероприятия</button>
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-purple-100 active:bg-purple-200" @click="$router.push({name: 'TeamPage'})">Команда</button>
    </div>
    <div v-if="isBigScreen && isEventPage" class="flex flex-row justify-center items-center gap-4 mt-2">
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-ev-color rounded-lg cursor-pointer transition-bg hover:bg-purple-100 active:bg-purple-200" @click="$router.push({name: 'EventsPage'})">Мероприятия</button>
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-purple-100 active:bg-purple-200" @click="$router.push({name: 'TeamPage'})">Команда</button>
    </div>
    <div v-if="isBigScreen && isTeamPage" class="flex flex-row justify-center items-center gap-4 mt-2">
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-b-gray-600 rounded-lg cursor-pointer transition-bg hover:bg-purple-100 active:bg-purple-200" @click="$router.push({name: 'EventsPage'})">Мероприятия</button>
      <button class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-ev-color rounded-lg cursor-pointer transition-bg hover:bg-purple-100 active:bg-purple-200" @click="$router.push({name: 'TeamPage'})">Команда</button>
    </div>


    <div class="flex flex-row gap-x-2 justify-center items-center md:mr-16">

      <img  class="w-8 mb:w-10 h-8 mb:h-10" src="../assets/icons/icon-profile.svg"/>


      <div @click="$router.push({name: 'ProfilePage'})" class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl uus:text-base text-color border-2 border-hr-color rounded-lg cursor-pointer transition-bg hover:bg-purple-100 active:bg-purple-200">
        Личный кабинет
      </div>
    </div>
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
    isEventPage(){
      return this.$route.fullPath === '/events';
    },
    isTeamPage(){
      return this.$route.fullPath === '/team';
    },

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

