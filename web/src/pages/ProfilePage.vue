<template>
  <div class="flex flex-col w-full h-full bg-color-ultralight scrollable">
    <div class="flex flex-col justify-evenly items-center gap-y-8 pt-8 md:gap-y-0 md:flex-row md:items-stretch md:pt-16">
      <div class="flex flex-col gap-y-4 md:gap-y-16 items-stretch">
        <p class="text-center text-3xl">Личный Кабинет</p>
        <div class="flex flex-col items-center">
          <img class="w-32 h-32" src="../assets/icons/icon-profile.svg"/>
          <p class="text-2xl mt-10">{{ getNameSurname }}</p>
          <p class="text-lg text-sky-600">{{ getEmail }}</p>
        </div>
        <div class="flex flex-col gap-y-3 items-stretch py-3 px-6 bg-color-light rounded-lg profile-shadow">
          <div @click="$router.push({name: 'AchievementsPage'})" class="box-shadow py-2 px-12 text-center rounded-lg cursor-pointer text-sky-500 bg-white hover:underline header-shadow active:text-sky-300">Посмотреть мои достижения</div>
          <div class="box-shadow py-2 px-12 text-center rounded-lg cursor-pointer text-sky-500 bg-white hover:underline header-shadow active:text-sky-300">Сменить пароль</div>
          <div @click="LogOut" class="box-shadow py-2 px-12 text-center rounded-lg cursor-pointer text-red-700 bg-white hover:underline header-shadow active:text-red-400">Выйти из аккаунта</div>
        </div>
      </div>
      <div class="flex flex-col gap-y-4 md:gap-y-16 pb-10">
        <p class="text-center text-3xl">Личная информация</p>
        <div class="flex justify-center items-center p-10 aspect-square rounded-xl bg-color-light profile-shadow">
          <div class="p-5 rounded-lg aspect-square lg:w-[400px] bg-white header-shadow">
            <ul class="list-disc markers profile-list">
              <li>Вуз: <span>{{ getUniversity }}</span></li>
              <li>Должность: <span>{{ getRoleName }}</span></li>
              <li>ФИО: <span>{{ getFIO }}</span></li>
              <li>Факультет: <span>{{ getFaculty }}</span></li>
              <li>Кафедра: <span>{{ getDepartment }}</span></li>
              <li>Направление: <span>{{ getEducationalDirection }}</span></li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { useUniversityStore } from '@/stores/universityStore';
import { ROLES_NAME } from '@/helpers/constants';
export default{
  computed: {
    ...mapStores(useUserInfoStore, useUniversityStore),

    getUniversity(){
      return 'РТУ МИРЭА';
    },
    getRoleName(){
      return ROLES_NAME[this.userInfoStore.role];
    },
    getFIO(){
      return this.userInfoStore.last_name + ' ' + this.userInfoStore.first_name + ' ' + this.userInfoStore.father_name;
    },
    getNameSurname(){
      return this.userInfoStore.last_name + ' ' + this.userInfoStore.first_name;
    },
    getEmail(){
      return this.userInfoStore.email;
    },
    getFaculty(){
      if(this.userInfoStore.faculty_id === null) return '-';
      for(let item of this.universityStore.facultiesList){
        if(item.id === this.userInfoStore.faculty_id) return item.name;
      }
      return '';
    },
    getDepartment(){
      if(this.userInfoStore.department_id === null) return '-';
      for(let item of this.universityStore.deparmentsList){
        if(item.id === this.userInfoStore.department_id) return item.name;
      }
      return '';
    },
    getEducationalDirection(){
      if(this.userInfoStore.educationalDirection === null) return '-';
      return this.userInfoStore.educationalDirection;
    },

  },
  methods: {
    LogOut(){

    }
  }
};

</script>