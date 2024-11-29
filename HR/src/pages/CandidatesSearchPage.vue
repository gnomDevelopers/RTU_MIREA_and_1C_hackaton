<template>
  <div class="scrollable flex h-auto flex-row">
    <div class="scrollable md:w-4/6 h-auto border-r-2 border-gray-200">
      <div class="w-11/12 h-auto m-10 flex flex-col justify-center">
        <div class="grid grid-cols-1 gap-4 mt-6">
          <div v-for="item in filteredWorkUsers" class="border-4 border-hr-color rounded-xl p-6">
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
    </div>
    <div class="md:w-2/6 h-auto">
      <p style="margin-left: 30px; color: #8F0101;" class="md:text-3xl uus:text-xl md:mt-4">Фильтры</p>
      <div class="p-6">
        <div class="mb-4">
          <label for="spec" class="block text-gray-700 text-lg font-bold mb-2">Специальность:</label>
          <input
              type="text"
              id="spec"
              v-model="filterData.speciality"
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              placeholder="Специальность"
          />
        </div>
        <div class="mb-4">
          <label for="skl" class="block text-gray-700 text-lg font-bold mb-2">Навыки:</label>
          <input
              type="text"
              id="skl"
              v-model="filterData.skills"
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              placeholder="Навыки"
          />
        </div>
        <div class="mb-4">
          <label for="gpa" class="block text-gray-700 text-lg font-bold mb-2">GPA:</label>
          <input
              type="number"
              id="gpa"
              v-model="filterData.gpa"
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              placeholder="GPA"
          />
        </div>
      </div>
      <div class="flex justify-center w-full p-2">
        <button @click="filter" class=" cursor-pointer transition-colors py-2 px-5 text-lg rounded-xl font-semibold btn text-slate-100"> Применить фильтры </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import axios from 'axios';
import { API_AllWorkUsers } from '@/api/api';
import type { WorkUser } from '@/helpers/constants';

export default {

  data() {
    return {
      workUsers: [] as WorkUser[],
      filteredWorkUsers: [] as WorkUser [],
      filterData: {
        speciality: "",
        skills: "",
        gpa: 0,
      } // Массив для хранения студентов
    }
  },
  mounted() {
    this.fetchStudents(); // Вызов функции для получения студентов при монтировании компонента
  },
  methods: {
    async fetchStudents() {
      try {
        const response = await API_AllWorkUsers(); // Вызов API
        this.workUsers = response; // Сохранение данных студентов в состоянии компонента
        this.filteredWorkUsers = this.workUsers;
      } catch (error) {
        console.error('Ошибка при получении студентов:', error);
      }
    },
    filter() {
      this.filteredWorkUsers = this.workUsers.filter(user => {
        const matchesSpeciality = this.filterData.speciality != "" ? user.speciality.toLowerCase() === this.filterData.speciality.toLowerCase() : true;
        //const matchesSkills = user.skills.toLowerCase().includes(this.filterData.skills.toLowerCase());
        const matchesGPA = this.filterData.gpa ? user.gpa >= this.filterData.gpa : true;

        return matchesSpeciality && matchesGPA;
      });
    }
  }
}

</script>