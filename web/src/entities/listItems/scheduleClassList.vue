<template>
  <div class="flex flex-col items-center w-full mb:w-auto gap-y-4 bg-color-light rounded-xl p-4">
    <div class="flex flex-col gap-y-1 w-full items-stretch">
      <div class="text-center cursor-default text-xl bg-white rounded-lg py-1">{{ getCurrentDate }}, {{ getCurrentWeekDay }}</div>
      <div class="text-center cursor-default text-xl bg-white rounded-lg py-1">{{ getWeekNumber }} неделя</div>
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
    <div v-else class="flex flex-col justify-center items-center w-[500px] h-[350px]">
      <div class="flex flex-row gap-x-1 items-center">
        <img class="w-8 h-8" src="../../assets/icons/icon-fire.svg"/>
        <p class="text-xl">В этот день нет пар!</p>
        <img class="w-8 h-8" src="../../assets/icons/icon-fire.svg"/>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
import { useUniversityStore } from '@/stores/universityStore';
import ScheduleClassListItem from '@/shared/scheduleClassListItem.vue';
import { GET_DATE_FROM_STRING, GET_DAY_OF_WEEK, GET_WEEK_DIFFERENCE, GET_WEEK_NUMBER } from '@/helpers/constants';

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
      return GET_DAY_OF_WEEK(this.scheduleStore.selectedDate);
    },

    getWeekNumber(){
      return GET_WEEK_DIFFERENCE( GET_DATE_FROM_STRING(this.scheduleStore.startDate), GET_DATE_FROM_STRING(this.scheduleStore.selectedDate));
    }
  },
};
</script>