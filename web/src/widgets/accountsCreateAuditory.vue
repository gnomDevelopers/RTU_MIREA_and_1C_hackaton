<template>
  <div>
    <div class="flex flex-col md:flex-row w-full md:w-auto items-stretch gap-1 p-2 rounded-lg self-start bg-blue-300">
      <div class="flex flex-row flex-wrap gap-1">
        <input 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="text" 
          placeholder="Название аудитории">
        
        <input 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="number" 
          placeholder="Вместимость аудитории">

        <select 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option class="text-lg" :value="-1" disabled>Выберите кампус</option>
          <!-- <option v-for="item in campusList" :key="item.id" class="text-lg" :value="item.id">{{ item }}</option> -->
        </select>
        
        <select 
          class="min-w-20 w-full md:w-52 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option class="text-lg" :value="-1" disabled>Выберите тип аудитории</option>
          <option v-for="(item, index) in auditoryTypesList" :key="index" class="text-lg" :value="index">{{ item }}</option>
        </select>

        <select 
          class="min-w-20 w-full md:w-52 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option class="text-lg" :value="-1" disabled>Выберите профиль аудитории</option>
          <option v-for="(item, index) in auditoryProfilesList" :key="index" class="text-lg" :value="index">{{ item }}</option>
        </select>
      </div>
      <div class="flex flex-row flex-shrink-0 gap-1 items-center">
        <div class="rounded-lg p-1 cursor-pointer btn">
          <img class="w-8 h-8" src="../assets/icons/icon-plus.svg"/>
        </div>
        <div class="px-2 h-10 rounded-lg cursor-pointer btn">
          <p class="text-white text-xl">Отправить</p>
        </div>
      </div>
    </div>

    <div class="flex flex-col gap-y-2 p-2 rounded-lg max-h-[400px] scrollable bg-white">
      <div class="flex flex-row gap-x-2 px-2 py-1 items-center rounded-lg bg-color-light">
        
        <div class="flex flex-row flex-grow justify-start items-center">
          <p class="text-lg">Аудитория 1</p>
        </div>

        <div class="flex flex-row flex-wrap gap-2 px-2">

          <p class="font-semibold text-lg cursor-default">
            Кампус: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">МП-1</span>
          </p>

          <p class="font-semibold text-lg cursor-default">
            Вместимость: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">150 чел</span>
          </p>

          <p class="font-semibold text-lg cursor-default">
            Тип: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">Лекторий</span>
          </p>

          <p class="font-semibold text-lg cursor-default">
            Профиль: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">Компьютерный класс</span>
          </p>

        </div>
        <div class="w-9 h-9 flex flex-row flex-shrink-0 justify-center items-center rounded-lg cursor-pointer btn">
          <img class="w-6 h-7" src="../assets/icons/icon-delete.svg"/>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { StatusCodes, type IAPI_Audience_Create } from '../helpers/constants';
export default {
  data(){
    return {
      campusList: [] as IAPI_Audience_Create[],

      auditoryTypesList: ['Лекторий', 'Практика', 'Спортивный зал', 'Лаборантская', 'Склад'],
      auditoryProfilesList: ['Обычная', 'Компьютерный класс', 'Химический класс', 'Физический класс', 'Биологический класс', 'Географический класс', 'Литературный класс', 'Исторический класс'],
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore),
  },
  mounted() {
    //load all campuses
  },
  methods() {
    
  }
};
</script>