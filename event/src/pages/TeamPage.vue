<template>
  <div class="scrollable flex h-auto flex-col">
    <div class="flex justify-center w-full">
      <p class="text-3xl m-8">Моя команда</p>
    </div>
    <div>
      <button class="cursor-pointer transition-colors py-2 px-5 text-xl rounded-md h-16 border-2 border-ev-color btn ml-14" style="background-color: #F3F0FF"><img class="items-start" src="../assets/icons/icon-team.svg" style="padding-right: 20px">Здесь могла быть ваша реклама</button>
    </div>
    <div class="flex flex-row w-full justify-center mt-8">
      <div class="m-6 border-2 rounded-md w-1/5 border-ev-color p-4 flex flex-col justify-center items-center" style="background-color: #F3F0FF">
        <img src="../assets/icons/icon-profile.svg" class="max-w-[100px] max-h-[100px]">
        <p class="text-xl text-center mt-2">Постников Даниил Сергеевич</p>
        <p class="text-lg text-center mt-2 text-gray-400">postnikov_d_s</p>
        <p class="text-xl text-center mt-4" style="color:#6a0793">ID участника: 7</p>
      </div>
      <div class="m-6 border-2 rounded-md w-1/5 border-ev-color p-4 flex flex-col justify-center items-center" style="background-color: #F3F0FF">
        <img src="../assets/icons/icon-profile.svg" class="max-w-[100px] max-h-[100px]">
        <p class="text-xl text-center mt-2">Жаров Степан Дмитриевич</p>
        <p class="text-lg text-center mt-2 text-gray-400">zharov_s_d</p>
        <p class="text-xl text-center mt-4" style="color:#6a0793">ID участника: 5</p>
      </div>
      <div class="m-6 border-2 rounded-md w-1/5 border-ev-color p-4 flex flex-col justify-center items-center" style="background-color: #F3F0FF">
        <img src="../assets/icons/icon-profile.svg" class="max-w-[100px] max-h-[100px]">
        <p class="text-xl text-center mt-2">Игаев Даниил Владимирович</p>
        <p class="text-lg text-center mt-2 text-gray-400">igayev_d_v</p>
        <p class="text-xl text-center mt-4" style="color:#6a0793">ID участника: 2</p>
      </div>
      <div class="m-6 border-2 rounded-md w-1/5 border-ev-color p-4 flex flex-col justify-center items-center" style="background-color: #F3F0FF">
        <img src="../assets/icons/icon-profile.svg" class="max-w-[100px] max-h-[100px]">
        <p class="text-xl text-center mt-2">Малчевский Максим Андреевич</p>
        <p class="text-lg text-center mt-2 text-gray-400">malchevsky_m_a</p>
        <p class="text-xl text-center mt-4" style="color:#6a0793">ID участника: 3</p>
      </div>
      <div class="m-6 border-2 rounded-md w-1/5 border-ev-color p-4 flex flex-col justify-center items-center" style="background-color: #F3F0FF">
        <img src="../assets/icons/icon-profile.svg" class="max-w-[100px] max-h-[100px]">
        <p class="text-xl text-center mt-2">Ермолина Софья Алексеевна</p>
        <p class="text-lg text-center mt-2 text-gray-400">yermolina_s_a</p>
        <p class="text-xl text-center mt-4" style="color:#6a0793">ID участника: 4</p>
      </div>
    </div>
    <div class="w-full">
      <button class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-12 font-semibold btn m-10 float-end" style="background-color: #D9D9D9">Покинуть команду</button><br><br>
    </div>

  </div>

</template>

<script lang="ts">
import { useUserFormStore } from '@/stores/userFormStore';

export default {
  setup() {
    const userFormStore = useUserFormStore();
    return {
      formData: userFormStore.formData,
    };
  },
  computed: {
    getDownloadLink() {
      if (this.formData.filesList.length > 0) {
        return URL.createObjectURL(this.formData.filesList[0]);
      }
      return '';
    },
  },
  methods: {
    hideProfile() {
      if (!this.formData.isClosed) {
        this.formData.isClosed = true;
      }
      else this.formData.isClosed = false;
    }
  }
};
</script>