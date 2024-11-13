<template>
  <div class="flex flex-row justify-between items-stretch h-24 shrink-0 w-full px-5 header-shadow z-10 bg-color-light">
    <div class="h-full grid content-center md:ml-16">
      <div class="font-semibold text-3xl mb:text-4xl md:text-5xl cursor-pointer">VUZ+</div>
    </div>
    <div class="flex flex-row gap-x-2 justify-center items-center md:mr-16">

      <svg class="w-8 mb:w-10 h-8 mb:h-10" viewBox="0 0 42 42" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M32.7683 39.4383C31.0816 39.9367 29.0883 40.1667 26.75 40.1667H15.25C12.9116 40.1667 10.9183 39.9367 9.23165 39.4383M32.7683 39.4383C32.3466 34.455 27.2291 30.5258 21 30.5258C14.7708 30.5258 9.65331 34.455 9.23165 39.4383M32.7683 39.4383C37.9816 37.9625 40.1666 33.995 40.1666 26.75V15.25C40.1666 5.66666 36.3333 1.83333 26.75 1.83333H15.25C5.66665 1.83333 1.83331 5.66666 1.83331 15.25V26.75C1.83331 33.995 4.01831 37.9625 9.23165 39.4383M21 25.1591C17.205 25.1591 14.1383 22.0734 14.1383 18.2784C14.1383 14.4834 17.205 11.4167 21 11.4167C24.795 11.4167 27.8617 14.4834 27.8617 18.2784C27.8617 22.0734 24.795 25.1591 21 25.1591Z" stroke="#063C73" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
      </svg>

      <p v-if="checkUserLogin" class=" underline cursor-pointer text-sky-600">{{ userInfoStore.first_name }} and {{ userInfoStore.father_name }}</p>
      
      <div v-else class="flex justify-center items-center px-3 mb:px-4 text-xl mb:text-2xl text-color border-2 border-color rounded-lg cursor-pointer transition-bg hover:bg-cyan-50 active:bg-cyan-100">Войти</div>

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
      return this.userInfoStore.authenticated;
    }
  }

};
</script>
