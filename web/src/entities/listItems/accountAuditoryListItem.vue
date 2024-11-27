<template>
  <div class="flex flex-col gap-y-2 p-2 rounded-lg bg-color-light">
    <div class="flex flex-row gap-x-2 items-stretch">
      <div class="flex flex-row flex-grow flex-wrap gap-x-2 items-center">
        <p class="text-lg font-semibold">
          Название: 
          <span class="text-base font-normal">{{ data.name }}</span>
        </p>
        <p class="text-lg font-semibold">
          Вместимость: 
          <span class="text-base font-normal">{{ data.capacity }} чел.</span>
        </p>
        <p class="text-lg font-semibold">
          Кампус: 
          <span class="text-base font-normal">{{ getCampusName(data.campus_id) }}</span>
        </p>
        <p class="text-lg font-semibold">
          Тип: 
          <span class="text-base font-normal">{{ data.type }}</span>
        </p>
        <p class="text-lg font-semibold">
          Профиль: 
          <span class="text-base font-normal">{{ data.profile }}</span>
        </p>
      </div>

      <div class="flex flex-row flex-shrink-0 gap-x-2">
        <div @click="showSettings = !showSettings" class="flex lfex-row justify-center items-center w-9 h-9 rounded-lg cursor-pointer bg-sky-600">
          <img class="w-7 h-7" src="../../assets/icons/icon-correct.svg"/>
        </div>
        <div @click="deleteAuditory" class="flex lfex-row justify-center items-center w-9 h-9 rounded-lg cursor-pointer bg-red-500">
          <img class="w-6 h-6" src="../../assets/icons/icon-delete.svg"/>
        </div>
      </div>
    </div>

    <div v-if="showSettings" class="flex flex-col md:flex-row w-full items-stretch gap-1">
      <div class="flex flex-row flex-grow flex-wrap gap-1">
        <input 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="text" 
          v-model="auditoryName"
          placeholder="Название аудитории">
        
        <input 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="number" 
          min="1"
          v-model="auditoryCapacity"
          placeholder="Вместимость аудитории">

        <select 
          v-model="auditoryCampusID"
          class="min-w-20 w-full md:w-auto max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option v-for="item in getCampusList" :key="item.id" class="text-lg" :value="item.id">{{ item.name }}</option>
        </select>
        
        <select 
          v-model="auditoryType"
          class="min-w-20 w-full md:w-auto max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option v-for="(item, index) in getAuditoryTypeList" :key="index" class="text-lg" :value="index">{{ item }}</option>
        </select>

        <select 
          v-model="auditoryProfile"
          class="min-w-20 w-full md:w-auto max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option v-for="(item, index) in getAuditoryProfileList" :key="index" class="text-lg" :value="index">{{ item }}</option>
        </select>
      </div>
      <div v-if="!isNoChanges" class="flex flex-row flex-shrink-0 gap-1 items-center">
        <div @click="updateAuditory" class="rounded-lg p-1 cursor-pointer btn">
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
import { type IAPI_Audience_Update, AUDITORY_PROFILE_LIST, AUDITORY_TYPE_LIST, StatusCodes } from '@/helpers/constants';
import { type PropType } from 'vue';
import { API_Audience_Delete, API_Audience_Update } from '@/api/api';

export default {
  props: {
    data: {
      type: Object as PropType<IAPI_Audience_Update>,
      required: true,
    }
  },
  data() {
    return {
      showSettings: false,

      auditoryCampusID: this.data.campus_id,
      auditoryCapacity: this.data.capacity,
      auditoryName: this.data.name,
      auditoryProfile: AUDITORY_PROFILE_LIST.indexOf(this.data.profile),
      auditoryType: AUDITORY_TYPE_LIST.indexOf(this.data.type),
    }
  },
  computed: {
    ...mapStores(useUniversityStore, useStatusWindowStore),

    getAuditoryProfileList(){
      return AUDITORY_PROFILE_LIST;
    },
    getAuditoryTypeList(){
      return AUDITORY_TYPE_LIST;
    },
    isNoChanges(){
      return this.auditoryCampusID === this.data.campus_id && 
        this.auditoryCapacity === this.data.capacity && 
        this.auditoryName === this.data.name && 
        this.auditoryProfile === AUDITORY_PROFILE_LIST.indexOf(this.data.profile) && 
        this.auditoryType ===  AUDITORY_TYPE_LIST.indexOf(this.data.type);
    },
    getCampusList(){
      return this.universityStore.campusList;
    }
  },
  methods: {
    getCampusName(campusID: number): string {
      for(let campus of this.universityStore.campusList){
        if(campus.id === campusID) return campus.name;
      }
      return '';
    },
    updateAuditory(){
      if(this.isNoChanges) return;

      if(this.auditoryCapacity >= 1 && this.auditoryName !== ''){
        // const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Обновляем информацию об аудитории...', -1);
        const body:IAPI_Audience_Update = {
          campus_id: this.auditoryCampusID,
          capacity: this.auditoryCapacity,
          name: this.auditoryName,
          profile: AUDITORY_PROFILE_LIST[this.auditoryProfile],
          type: AUDITORY_TYPE_LIST[this.auditoryType],
          id: this.data.id
        }
        for(let i = 0; i < this.universityStore.auditoriesList.length; i++){
          if(this.universityStore.auditoriesList[i].id === this.data.id){
            this.universityStore.auditoriesList[i].campus_id = this.auditoryCampusID;
            this.universityStore.auditoriesList[i].capacity = this.auditoryCapacity;
            this.universityStore.auditoriesList[i].name = this.auditoryName;
            this.universityStore.auditoriesList[i].profile = AUDITORY_PROFILE_LIST[this.auditoryProfile];
            this.universityStore.auditoriesList[i].type = AUDITORY_TYPE_LIST[this.auditoryType];
            break;
          }
        }
        setTimeout(() => {
          this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Информация об аудитории обновлена!');
        }, 300)
        // API_Audience_Update(body)
        // .then(response => {
        //   this.statusWindowStore.deteleStatusWindow(stID);
        //   this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Информация об аудитории обновлена!');
        // })
        // .catch(error => {
        //   this.statusWindowStore.deteleStatusWindow(stID);
        //   this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при обновлении аудиториии!');
        // })
      }

      //ошибка вместимости аудитории
      if(this.auditoryCapacity < 1){
        if(this.auditoryCapacity === -1) this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Укажите вместимость аудитории!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Некорректная вместимость аудитории!');
      }
      //ошибка названия аудитории
      if(this.auditoryName === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите название аудитории!');
      }
    },
    deleteAuditory() {
      for(let i = 0; i < this.universityStore.auditoriesList.length; i++){
          if(this.universityStore.auditoriesList[i].id === this.data.id){
            this.universityStore.auditoriesList.splice(i, 1);
            break;
          }
        }
      setTimeout(() => {
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Аудитория успешно удалена!');
      }, 300);
      // const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Удаляем аудиторию...', -1);
      // API_Audience_Delete(this.data.id)
      // .then(response => {
      //   this.statusWindowStore.deteleStatusWindow(stID);
      //   this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Аудитория успешно удалена!');
      // })
      // .catch(error => {
      //   this.statusWindowStore.deteleStatusWindow(stID);
      //   this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при удалении аудиториии!');
      // })
    }
  }
};
</script>