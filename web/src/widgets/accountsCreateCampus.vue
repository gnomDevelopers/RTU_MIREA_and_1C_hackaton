<template>
  <div>
    <div class="flex flex-col md:flex-row w-full md:w-auto items-stretch gap-1 p-2 rounded-lg self-start bg-blue-300">
      <div class="flex flex-row flex-wrap gap-1">
        <input 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="text" 
          v-model="campusName"
          placeholder="Название кампуса">
        <input 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="text" 
          v-model="campusAddress"
          placeholder="Адрес кампуса">
      </div>
      <div class="flex flex-row flex-shrink-0 gap-1 items-center">
        <div @click="addCampus" class="rounded-lg p-1 cursor-pointer btn">
          <img class="w-8 h-8" src="../assets/icons/icon-plus.svg"/>
        </div>
        <div v-if="campusList.length !== 0" @click="sendCampuses" class="px-2 h-10 rounded-lg cursor-pointer btn">
          <p class="text-white text-xl">Отправить</p>
        </div>
      </div>
    </div>

    <div v-if="campusList.length !== 0" class="flex flex-col gap-y-2 p-2 rounded-lg max-h-[400px] scrollable bg-white">
      <div v-for="item in campusList" :key="item.name" class="flex flex-row gap-x-2 px-2 py-1 items-center rounded-lg bg-color-light">
        <div class="flex flex-row flex-wrap gap-2 px-2">

          <p class="font-semibold text-lg cursor-default">
            Название: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">{{ item.name }}</span>
          </p>

          <p class="font-semibold text-lg cursor-default">
            Адрес: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">{{ item.address }}</span>
          </p>

        </div>
        <div @click="deleteCampus(item)" class="w-9 h-9 flex flex-row flex-shrink-0 justify-center items-center rounded-lg cursor-pointer btn">
          <img class="w-6 h-7" src="../assets/icons/icon-delete.svg"/>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUniversityStore } from '@/stores/universityStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { StatusCodes, type IAPI_Campus_Create } from '../helpers/constants';
import { API_Campus_Create } from '@/api/api';

export default {
  data(){
    return {
      campusList: [] as IAPI_Campus_Create[],

      campusName: '',
      campusAddress: '',
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore, useUniversityStore, useUserInfoStore),

  },
  methods: {
    addCampus(){
      if(this.campusName !== '' && this.campusAddress !== ''){
        this.campusList.push({
          name: this.campusName,
          address: this.campusAddress,
          university: this.userInfoStore.university!,
        });
        return;
      }

      if(this.campusName === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите название кампуса!');
      }
      if(this.campusAddress === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите адрес кампуса!');
      }
    },
    deleteCampus(campusDelete: IAPI_Campus_Create){
      for(let i = 0; i < this.campusList.length; i++){
        if(this.campusList[i].name === campusDelete.name){
          this.campusList.splice(i, 1);
          return;
        }
      }
    },
    sendCampuses(){
      if(this.campusList.length === 0)return;
      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);

      API_Campus_Create(this.campusList)
      .then(async (response: any) => {
        //заново подгружаем все кампусы
        await this.universityStore.loadCampus(this.userInfoStore.university!);

        //очищаем буфер кампусов
        this.campusList = [];

        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Кампусы успешно созданы!');
      })
      .catch(error => {
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при создании кампусов!');
      })
    }
  }
};
</script>