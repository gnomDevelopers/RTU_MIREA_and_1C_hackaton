<template>
  <div class="flex flex-row justify-start items-stretch min-h-14 mb:min-w-[500px] w-full mb:w-auto">
    <p class="text-xl font-medium w-14 flex flex-row flex-shrink-0 justify-center items-center rounded-l-xl cursor-default bg-color-bold text-white">{{ index }}</p>
    <div @click="selectClass" class="flex flex-row flex-grow items-stretch bg-white rounded-r-xl px-2 cursor-pointer">
      
      <div class="flex flex-col flex-grow items-start gap-y-1">
        <p class="flex flex-row flex-wrap gap-x-1 text-lg">
          <span class="font-bold">{{ data.time_start }}-{{ data.time_end }}</span>
          {{ data.name }}
        </p>
        <div v-if="showAddMySchedule" class="flex flex-row gap-x-1 w-full">
          <input class="flex-grow text-lg px-2 outline-none header-shadow rounded" type="text" placeholder="Введите ваше расписание" v-model="mySchedule">
          <svg @click="closeMySchedule" class="w-6 h-6 rotate-45" viewBox="0 0 32 32" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M16 2V30M2 16H30" stroke="#e64545" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </div>
        <p v-else class="flex flex-row flex-wrap gap-x-1 text-base">
          <span class="font-bold">{{ data.auditory }}</span>
          {{ data.group_names.join(', ') }}{{ mySheduleText }}
        </p>
      </div>
      <div class="flex flex-col items-center justify-center gap-y-2 cursor-pointer px-2" v-if="data.name === '' && canAddFaculties">
        <svg @click="addMySchedule" class="w-6 h-6" viewBox="0 0 32 32" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M16 2V30M2 16H30" stroke="#063C73" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </div>

    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { StatusCodes, type IScheduleItem } from '@/helpers/constants';
import type { PropType } from 'vue';
export default{
  props: {
    index: {
      type: Number,
      required: true,
    },
    data: {
      type: Object as PropType<IScheduleItem>,
      required: true,
    },
    canAddFaculties:{
      type:Boolean,
      required: false,
      default: false,
    }
  },
  data() {
    return {
      showAddMySchedule: false,
      mySchedule: '',
      mySheduleText: '',
    }
  },
  computed: {
    ...mapStores(useScheduleStore, useStatusWindowStore),
  },
  methods: {
    selectClass(){
      this.scheduleStore.selectedClass = 1; //
    },
    addMySchedule(){
      if(!this.showAddMySchedule){
        this.showAddMySchedule = true;
        return;
      }
      this.mySheduleText = this.mySchedule;
      this.mySchedule = '';
      this.showAddMySchedule = false;
      if(this.mySheduleText !== '') this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Расписание обновлено!');
    },
    closeMySchedule(){
      this.mySchedule = '';
      this.showAddMySchedule = false;
    }
  }
};
</script>