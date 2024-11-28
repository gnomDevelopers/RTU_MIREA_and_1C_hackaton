<template>
  <div class="flex flex-col items-center w-full mb:w-auto gap-y-4 bg-color-light rounded-xl p-4">
    <div class="flex flex-col gap-y-1 w-full items-stretch">
      <div class="text-center cursor-default text-xl bg-white rounded-lg py-1">{{ getCurrentDate }}, {{ getCurrentWeekDay }}</div>
      <div class="text-center cursor-default text-xl bg-white rounded-lg py-1">13 неделя</div>
    </div>
    <div v-if="getScheduleTable.length !== 0" class="flex flex-col gap-y-1 w-full items-stretch">
      
      <div v-for="(item, ind) in getScheduleTable">
        <ScheduleClassListItem 
          :index="ind + 1" 
          :data="item"
          :canAddFaculties
        />
      </div>
      
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
import { useUniversityStore } from '@/stores/universityStore';
import ScheduleClassListItem from '@/shared/scheduleClassListItem.vue';

export default {
  props:{
    canAddFaculties:{
      type: Boolean,
      required: false,
      default: false,
    }
  },
  components:{
    ScheduleClassListItem,
  },
  computed:{
    ...mapStores(useScheduleStore, useUniversityStore),

    getScheduleTable(){
      return this.scheduleStore.scheduleTableDay.length !== 0 ? this.scheduleStore.scheduleTableDay : [];
    },

    getCurrentDate(){
      return this.scheduleStore.selectedDate;
    },

    getCurrentWeekDay(){
      switch (this.scheduleStore.selectedDate){
        case '25.11.2024': return 'Понедельник';
        case '26.11.2024': return 'Вторник';
        case '27.11.2024': return 'Среда';
        case '28.11.2024': return 'Четверг';
        case '29.11.2024': return 'Пятница';
        case '30.11.2024': return 'Суббота';
        case '01.12.2024': return 'Воскресенье';
      }
      return 'Понедельник';
    }
  },
};
</script>