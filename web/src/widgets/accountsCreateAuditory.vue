<template>
  <div>
    <div class="flex flex-col md:flex-row w-full md:w-auto items-stretch gap-1 p-2 rounded-lg self-start bg-blue-300">
      <div class="flex flex-row flex-wrap gap-1">
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
          
          <option class="text-lg" :value="-1" disabled>Выберите кампус</option>
          <option v-for="item in getCampusList" :key="item.id" class="text-lg" :value="item.id">{{ item.name }}</option>
        </select>
        
        <select 
          v-model="auditoryType"
          class="min-w-20 w-full md:w-auto max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option class="text-lg" :value="-1" disabled>Выберите тип аудитории</option>
          <option v-for="(item, index) in getAuditoryTypeList" :key="index" class="text-lg" :value="index">{{ item }}</option>
        </select>

        <select 
          v-model="auditoryProfile"
          class="min-w-20 w-full md:w-auto max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option class="text-lg" :value="-1" disabled>Выберите профиль аудитории</option>
          <option v-for="(item, index) in getAuditoryProfileList" :key="index" class="text-lg" :value="index">{{ item }}</option>
        </select>
      </div>
      <div class="flex flex-row flex-shrink-0 gap-1 items-center">
        <div @click="addAuditory" class="rounded-lg p-1 cursor-pointer btn">
          <img class="w-8 h-8" src="../assets/icons/icon-plus.svg"/>
        </div>
        <div v-if="auditoryList.length !== 0" @click="sendAuditories" class="px-2 h-10 rounded-lg cursor-pointer btn">
          <p class="text-white text-xl">Отправить</p>
        </div>
      </div>
    </div>

    <div v-if="auditoryList.length !== 0" class="flex flex-col gap-y-2 p-2 rounded-lg max-h-[400px] scrollable bg-white">
      <div v-for="item in auditoryList" :key="item.name" class="flex flex-row gap-x-2 px-2 py-1 items-center rounded-lg bg-color-light">
        
        <div class="flex flex-row flex-grow justify-start items-center">
          <p class="text-lg">{{item.name}}</p>
        </div>

        <div class="flex flex-row flex-wrap gap-2 px-2">

          <p class="font-semibold text-lg cursor-default">
            Кампус: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">{{ getCampusName(item.campus_id) }}</span>
          </p>

          <p class="font-semibold text-lg cursor-default">
            Вместимость: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">{{ item.capacity }} чел</span>
          </p>

          <p class="font-semibold text-lg cursor-default">
            Тип: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">{{ item.type }}</span>
          </p>

          <p class="font-semibold text-lg cursor-default">
            Профиль: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">{{ item.profile }}</span>
          </p>

        </div>
        <div @click="deleteAuditory(item)" class="w-9 h-9 flex flex-row flex-shrink-0 justify-center items-center rounded-lg cursor-pointer btn">
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
import { StatusCodes, type IAPI_Audience_Create, AUDITORY_TYPE_LIST, AUDITORY_PROFILE_LIST } from '../helpers/constants';
import { API_Audience_Create } from '@/api/api';

export default {
  data(){
    return {
      auditoryList: [] as IAPI_Audience_Create[],

      auditoryCampusID: -1,
      auditoryCapacity: 0,
      auditoryName: '',
      auditoryProfile: -1,
      auditoryType: -1,
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore, useUniversityStore),

    getAuditoryProfileList(){
      return AUDITORY_PROFILE_LIST;
    },
    getAuditoryTypeList(){
      return AUDITORY_TYPE_LIST;
    },
    getCampusList(){
      return this.universityStore.campusList;
    }
  },
  methods: {
    addAuditory(){
      if(
        this.auditoryCampusID !== -1 &&
        this.auditoryCapacity > 0 && 
        this.auditoryName !== '' && 
        this.auditoryProfile !== -1 && 
        this.auditoryType !== -1
      ){
        this.auditoryList.push({
          campus_id: this.auditoryCampusID, 
          capacity: this.auditoryCapacity, 
          name: this.auditoryName, 
          profile: AUDITORY_PROFILE_LIST[this.auditoryProfile], 
          type: AUDITORY_TYPE_LIST[this.auditoryType]
        });

        this.auditoryCapacity = 0;
        this.auditoryName = '';
        this.auditoryProfile = -1;
        this.auditoryType = -1;

        return;
      }

      if(this.auditoryCampusID === -1){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Выберите кампус!');
      }

      if(this.auditoryCapacity < 1){
        if(this.auditoryCapacity === -1) this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Укажите вместимость аудитории!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Некорректная вместимость аудитории!');
      }

      if(this.auditoryName === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите название аудитории!');
      }

      if(this.auditoryProfile === -1){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Выберите профиль аудитории!');
      }

      if(this.auditoryType === -1){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Выберите тип аудитории!');
      }
    },
    deleteAuditory(itemDelete: IAPI_Audience_Create){
      this.auditoryList = this.auditoryList.filter(item => !(
        item.campus_id === itemDelete.campus_id && 
        item.capacity === itemDelete.capacity &&
        item.name === itemDelete.name &&
        item.profile === itemDelete.profile &&
        item.type === itemDelete.type
      ));
    },
    sendAuditories() {
      if(this.auditoryList.length === 0)return;
      for(let item of this.auditoryList){
        this.universityStore.auditoriesList.push({...item, id: this.universityStore.tmpuserID++});
      }
      //очищаем буфер аудиторий
      this.auditoryList = [];
      setTimeout(() => {
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Аудитории сохранены!');
      }, 460);
      // const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);
      // API_Audience_Create(this.auditoryList)
      // .then(resolve => {
      //   this.statusWindowStore.deteleStatusWindow(stID);
      //   this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Аудитории сохранены!');
      //   //Добавить в universityStore
      // })
      // .catch(error => {
      //   this.statusWindowStore.deteleStatusWindow(stID);
      //   this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при сохранении аудиторий!');
      // });
    },
    getCampusName(campusID: number): string{
      for(let el of this.getCampusList) {
        if(el.id === campusID) return el.name;
      }
      return '';
    }
  }
};
</script>