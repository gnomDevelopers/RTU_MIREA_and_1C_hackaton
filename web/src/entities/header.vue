<template>
  <div v-if="!isWelcomePage" class="flex flex-row justify-between items-stretch h-24 shrink-0 w-full px-5 header-shadow z-10 bg-color-light">
    <div class="h-full grid content-center md:ml-16">
      <div @click="$router.push({name: 'MainPage'})" class="header-logo">VUZ+</div>
    </div>
    <div v-if="checkUserLogin" class="flex flex-row gap-x-2 justify-center items-center md:mr-16">

      <img class="w-8 mb:w-10 h-8 mb:h-10" src="../assets/icons/icon-profile.svg"/>

      <p @click="$router.push({name: 'ProfilePage'})" class="text-2xl cursor-pointer ">{{ getNameFatherName }}</p>
      
      <!-- <div v-else class="box-shadow flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl text-color border-2 border-color rounded-lg cursor-pointer transition-bg hover:bg-cyan-50 active:bg-cyan-100">Войти</div> -->

    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useUserInfoStore } from '@/stores/userInfoStore';

export default {
  computed: {
    ...mapStores(useUserInfoStore),

    checkUserLogin(){
      return this.userInfoStore.authorized;
    },

    getNameFatherName(){
      return this.userInfoStore.first_name + ' ' + this.userInfoStore.father_name;
    },

    isWelcomePage(){
      return this.$route.fullPath === '/';
    }
  }

};
</script>
