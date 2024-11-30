<template>
  <div class="scrollable md:p-14 uus:p-5">
    <p class="md:text-3xl uus:text-xl">Входящие предложения</p>
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mt-6">
      <div class="border-4 border-hr-color rounded-xl p-6">
        <div class="flex flex-row">
          <img class="md:max-w-28 md:max-h-28 uus:max-w-24 uus:max-h-24 mr-4" src="../assets/icons/icon-company.svg">
          <div class="flex flex-col">
            <p class="font-bold md:text-2xl uus:text-xl">1C</p>
            <p class="md:text-xl uus:text-base">Младший разработчик</p>

            <div v-if="!agree1 && !deny1" class="flex flex-col">
              <button class="sml-btn mb-2 mt-2" @click="isAgreed1">Принять</button>
              <button class="sml-btn" @click="isDenied1">Отказаться</button>
            </div>
            <div v-if="agree1" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы приняли предложение!</p>
            </div>
            <div v-if="deny1" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы отказались от предложения</p>
            </div>

          </div>
        </div>
      </div>
      <div class="border-4 border-hr-color rounded-xl p-6">
        <div class="flex flex-row">
          <img class="md:max-w-28 md:max-h-28 uus:max-w-24 uus:max-h-24 mr-4" src="../assets/icons/icon-company.svg">
          <div class="flex flex-col">
            <p class="font-bold md:text-2xl uus:text-xl">Яндекс</p>
            <p class="md:text-xl uus:text-base">Разработчик C++</p>
            <div v-if="!agree2 && !deny2" class="flex flex-col">
              <button class="sml-btn mb-2 mt-2" @click="isAgreed2">Принять</button>
              <button class="sml-btn" @click="isDenied2">Отказаться</button>
            </div>
            <div v-if="agree2" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы приняли предложение!</p>
            </div>
            <div v-if="deny2" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы отказались от предложения</p>
            </div>
          </div>
        </div>
      </div>
      <div class="border-4 border-hr-color rounded-xl p-6">
        <div class="flex flex-row">
          <img class="md:max-w-28 md:max-h-28 uus:max-w-24 uus:max-h-24 mr-4" src="../assets/icons/icon-company.svg">
          <div class="flex flex-col">
            <p class="font-bold md:text-2xl uus:text-xl">Вкусно - и точка</p>
            <p class="md:text-xl uus:text-base">Менеджер отдела</p>
            <div v-if="!agree3 && !deny3" class="flex flex-col">
              <button class="sml-btn mb-2 mt-2" @click="isAgreed3">Принять</button>
              <button class="sml-btn" @click="isDenied3">Отказаться</button>
            </div>
            <div v-if="agree3" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы приняли предложение!</p>
            </div>
            <div v-if="deny3" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы отказались от предложения</p>
            </div>
          </div>
        </div>
      </div>
      <div class="border-4 border-hr-color rounded-xl p-6">
        <div class="flex flex-row">
          <img class="md:max-w-28 md:max-h-28 uus:max-w-24 uus:max-h-24 mr-4" src="../assets/icons/icon-company.svg">
          <div class="flex flex-col">
            <p class="font-bold md:text-2xl uus:text-xl">ВТБ Банк</p>
            <p class="md:text-xl uus:text-base">Go-разработчик</p>
            <div v-if="!agree4 && !deny4" class="flex flex-col">
              <button class="sml-btn mb-2 mt-2" @click="isAgreed4">Принять</button>
              <button class="sml-btn" @click="isDenied4">Отказаться</button>
            </div>
            <div v-if="agree4" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы приняли предложение!</p>
            </div>
            <div v-if="deny4" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы отказались от предложения</p>
            </div>
          </div>
        </div>
      </div>
      <div class="border-4 border-hr-color rounded-xl p-6">
        <div class="flex flex-row">
          <img class="md:max-w-28 md:max-h-28 uus:max-w-24 uus:max-h-24 mr-4" src="../assets/icons/icon-company.svg">
          <div class="flex flex-col">
            <p class="font-bold md:text-2xl uus:text-xl">АльфаБанк</p>
            <p class="md:text-xl uus:text-base">Backend-разработчик</p>
            <div v-if="!agree5 && !deny5" class="flex flex-col">
              <button class="sml-btn mb-2 mt-2" @click="isAgreed5">Принять</button>
              <button class="sml-btn" @click="isDenied5">Отказаться</button>
            </div>
            <div v-if="agree5" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы приняли предложение!</p>
            </div>
            <div v-if="deny5" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы отказались от предложения</p>
            </div>
          </div>
        </div>
      </div>
      <div class="border-4 border-hr-color rounded-xl p-6">
        <div class="flex flex-row">
          <img class="md:max-w-28 md:max-h-28 uus:max-w-24 uus:max-h-24 mr-4" src="../assets/icons/icon-company.svg">
          <div class="flex flex-col">
            <p class="font-bold md:text-2xl uus:text-xl">ГК Астра</p>
            <p class="md:text-xl uus:text-base">Стажёр в IT-отдел</p>
            <div v-if="!agree6 && !deny6" class="flex flex-col">
              <button class="sml-btn mb-2 mt-2" @click="isAgreed6">Принять</button>
              <button class="sml-btn" @click="isDenied6">Отказаться</button>
            </div>
            <div v-if="agree6" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы приняли предложение!</p>
            </div>
            <div v-if="deny6" class="flex flex-col">
              <p class="mt-4 text-xl" style="color:#8F0101">Вы отказались от предложения</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import {useUserFormStore} from "@/stores/userFormStore";

export default {
  data(){
    return {
      agree1: false,
      deny1: false,
      agree2: false,
      deny2: false,
      agree3: false,
      deny3: false,
      agree4: false,
      deny4: false,
      agree5: false,
      deny5: false,
      agree6: false,
      deny6: false,
    }
  },
  methods: {
    isAgreed1(){
      this.agree1 = true;
    },
    isDenied1(){
      this.deny1 = true;
    },
    isAgreed2(){
      this.agree2 = true;
    },
    isDenied2(){
      this.deny2 = true;
    },
    isAgreed3(){
      this.agree3 = true;
    },
    isDenied3(){
      this.deny3 = true;
    },
    isAgreed4(){
      this.agree4 = true;
    },
    isDenied4(){
      this.deny4 = true;
    },
    isAgreed5(){
      this.agree5 = true;
    },
    isDenied5(){
      this.deny5 = true;
    },
    isAgreed6(){
      this.agree6 = true;
    },
    isDenied6(){
      this.deny6 = true;
    },
  }
};
</script>