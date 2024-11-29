<template>
  <div class="scrollable md:p-14 uus:p-5">
    <p class="md:text-3xl uus:text-xl">Кандидаты, которым уже был выслан отклик</p>
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mt-6">
      <div v-for="item in candidateResponses" class="border-4 border-hr-color rounded-xl p-6">
        <div class="flex flex-row">
          <img class="md:max-w-28 md:max-h-28 uus:max-w-24 uus:max-h-24 mr-4" src="../assets/icons/icon-profile.svg">
          <div class="flex flex-col">
            <p class="font-bold md:text-2xl uus:text-xl">{{item.last_name + " " + item.first_name + " " + item.father_name}}</p>
            <p class="md:text-xl uus:text-base">{{item.speciality}}</p>
            <div class="flex flex-col">
              <button class="sml-btn mb-2 mt-2">Посмотреть</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import axios from 'axios';
import { API_HRResponses } from '@/api/api';
import type { CandidateResponse } from '@/helpers/constants';
import { mapStores } from 'pinia';
import { useUserInfoStore } from '@/stores/userInfoStore';

export default {

  data() {
    return {
      candidateResponses: [] as CandidateResponse[]
    }
  },
  computed: {...mapStores(useUserInfoStore)},
  mounted() {
    this.fetchStudents(); // Вызов функции для получения студентов при монтировании компонента
  },
  methods: {
    async fetchStudents() {
      try {
        const response = await API_HRResponses(this.userInfoStore.userID!); // Вызов API
        this.candidateResponses = response; // Сохранение данных студентов в состоянии компонента
      } catch (error) {
        console.error('Ошибка при получении студентов:', error);
      }
    }
  }
}

</script>