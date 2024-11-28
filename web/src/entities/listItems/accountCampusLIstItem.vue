<template>
  <div class="flex flex-col gap-y-2 p-2 rounded-lg bg-color-light">
    <div class="flex flex-row gap-x-2 items-stretch">
      <div class="flex flex-row flex-grow flex-wrap gap-x-2 items-center">
        <p class="text-lg font-semibold">
          Название: 
          <span class="text-base font-normal">{{ data.name }}</span>
        </p>
        <p class="text-lg font-semibold">
          Адрес: 
          <span class="text-base font-normal">{{ data.address }} чел.</span>
        </p>
      </div>

      <div class="flex flex-row flex-shrink-0 gap-x-2">
        <div @click="showSettings = !showSettings" class="flex lfex-row justify-center items-center w-9 h-9 rounded-lg cursor-pointer bg-sky-600">
          <img class="w-7 h-7" src="../../assets/icons/icon-correct.svg"/>
        </div>
        <div @click="deleteCampus" class="flex lfex-row justify-center items-center w-9 h-9 rounded-lg cursor-pointer bg-red-500">
          <img class="w-6 h-6" src="../../assets/icons/icon-delete.svg"/>
        </div>
      </div>
    </div>

    <div v-if="showSettings" class="flex flex-col md:flex-row w-full items-stretch gap-1">
      <div class="flex flex-row flex-grow flex-wrap gap-1">
        <input 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="text" 
          v-model="campusName"
          placeholder="Название кампуса">
        
        <input 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="number" 
          min="1"
          v-model="campusAddress"
          placeholder="Адрес кампуса">
        
      </div>
      <div v-if="!isNoChanges" class="flex flex-row flex-shrink-0 gap-1 items-center">
        <div @click="updateCampus" class="rounded-lg p-1 cursor-pointer btn">
          Сохранить
        </div>
      </div>
    </div>
  </div>
  
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUniversityStore } from '@/stores/universityStore';
import { type IAPI_Audience_Update, type IAPI_Campus_Update, StatusCodes } from '@/helpers/constants';
import { type PropType } from 'vue';
import { API_Campus_Delete, API_Campus_Update } from '@/api/api';

export default {
  props: {
    data: {
      type: Object as PropType<IAPI_Campus_Update>,
      required: true,
    }
  },
  data() {
    return {
      showSettings: false,

      campusName: this.data.name,
      campusAddress: this.data.address,
    }
  },
  computed: {
    ...mapStores(useUniversityStore, useStatusWindowStore),

    isNoChanges(){
      return this.campusName === this.data.name && this.campusAddress === this.data.address;
    },
    getCampusList(){
      return this.universityStore.campusList;
    }
  },
  methods: {
    updateCampus(){
      if(this.isNoChanges) return;

      if(this.campusName !== '' && this.campusAddress !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Обновляем кампус...', -1);
        
        const body:IAPI_Campus_Update = {
          name: this.campusName,
          address: this.campusAddress,
          id: this.data.id,
          university: this.data.university,
        }

        API_Campus_Update(body)
        .then(response => {
          for(let i = 0; i < this.universityStore.campusList.length; i++){
            if(this.universityStore.campusList[i].id === this.data.id){
              this.universityStore.campusList[i].name = this.campusName;
              this.universityStore.campusList[i].address = this.campusAddress;
              break;
            }
          }

          this.statusWindowStore.deteleStatusWindow(stID);
          this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Кампус обновлен!');
        })
        .catch(error => {
          this.statusWindowStore.deteleStatusWindow(stID);
          this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при обновлении кампуса!');
        });
      }

      //ошибка вместимости аудитории
      if(this.campusName === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите название кампуса!');
      }
      //ошибка названия аудитории
      if(this.campusAddress === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите адрес кампуса!');
      }
    },
    deleteCampus() {
      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Удаляем аудиторию...', -1);
      API_Campus_Delete(this.data.id)
      .then((response: any) => {
        for(let i = 0; i < this.universityStore.campusList.length; i++){
          if(this.universityStore.campusList[i].id === this.data.id){
            this.universityStore.campusList.splice(i, 1);
            break;
          }
        }
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Кампус успешно удален!');
      })
      .catch(error => {
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при удалении кампуса!');
      });
    }
  }
};
</script>