<template>
  <div class="flex flex-col w-full items-center scrollable">
    <div class="flex flex-col gap-y-4 lg:flex-row xl:w-10/12 pt-6 px-2 us:px-0">
      <div class="flex flex-col w-full lg:w-1/2 items-center gap-y-5">

        <p class=" text-4xl font-medium text-center">Просмотр расписания</p>
        <div class="flex flex-row gap-x-2 p-2 rounded-lg bg-color-light w-full us:w-auto">
          <p 
            @click="selectScheduleType(0)"
            class="text-lg cursor-pointer px-2 text-center rounded-md transition-colors"
            :class="{'bg-white': scheduleStore.scheduleType === 0}">
              Моё расписание
          </p>
          <p 
            @click="selectScheduleType(1)"
            class="text-lg cursor-pointer px-2 text-center rounded-md transition-colors"
            :class="{'bg-white': scheduleStore.scheduleType === 1}">
              Общее расписание
          </p>
        </div>

        <Transition name="slowRemove">
          <div v-if="scheduleStore.scheduleType === 1" class="flex flex-row flex-wrap justify-center gap-2 p-2 rounded-lg bg-color-light">
            <p 
              @click="selectScheduleTarget(0)"
              class="text-lg cursor-pointer px-2 rounded-md transition-colors"
              :class="{'bg-white': scheduleStore.scheduleTarget === 0}">
              Группа
            </p>
            <p 
              @click="selectScheduleTarget(1)"
              class="text-lg cursor-pointer px-2 rounded-md transition-colors"
              :class="{'bg-white': scheduleStore.scheduleTarget === 1}">
              Преподаватель
            </p>
            <p 
              @click="selectScheduleTarget(2)"
              class="text-lg cursor-pointer px-2 rounded-md transition-colors"
              :class="{'bg-white': scheduleStore.scheduleTarget === 2}">
              Факультатив
            </p>
          </div>
        </Transition>
        
        <Transition name="slowRemove">
          <div v-if="scheduleStore.scheduleType === 1" class="flex flex-row w-[314px] mb:w-auto rounded-lg header-shadow bg-white">
            <input 
              class="p-2 outline-none min-w-0 max-w-none flex-grow mb:w-96 bg-transparent" 
              type="text" 
              :placeholder="getScheduleTargetText"
              v-model="searchFilter">
            <div 
              class="w-10 h-10 flex flex-col justify-center items-center cursor-pointer rounded-r-lg btn"
              @click="filterUserList">
              <img class="w-7 h-7" src="../assets/icons/icon-search.svg"/>
            </div>
          </div>
        </Transition>
        
        <div class="flex flex-col items-center gap-y-2">
          <CalendarTable :month="10"/>
        </div>

      </div>
      <div class="flex flex-col gap-y-4 items-center lg:w-1/2">
        <div class="flex flex-col items-center gap-y-4 bg-color-light rounded-xl p-4">
          <div class="flex flex-col gap-y-1 w-full items-stretch">
            <div class="text-center cursor-default text-xl bg-white rounded-lg py-1">12.11.2024, Вторник</div>
            <div class="text-center cursor-default text-xl bg-white rounded-lg py-1">11 неделя</div>
          </div>
          <div class="flex flex-col gap-y-1 w-full items-stretch">
            <ScheduleItem 
              v-for="(item, ind) in scheduleStore.scheduleTableDay"
              :index="ind + 1" 
              :time="`${item.time} ${item.type}`" 
              :title="item.title" 
              :room="item.place" 
              :group="item.groups.join(', ')"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
import { SCHEDULE_TARGET_TEXT } from '@/helpers/constants';

import CalendarTable from '@/entities/calendarTable.vue';
import ScheduleItem from '@/shared/scheduleItem.vue';

export default {
  components:{
    ScheduleItem,
    CalendarTable,
  },
  data(){
    return{
      searchFilter: '' as string,
    }
  },
  computed:{
    ...mapStores(useScheduleStore),

    getScheduleTargetText(){
      return SCHEDULE_TARGET_TEXT[this.scheduleStore.scheduleTarget];
    }
  },
  mounted(){
    this.scheduleStore.loadScheduleTable();
  },
  methods:{
    filterUserList(){

    },
    selectScheduleType(type: number){
      this.scheduleStore.scheduleType = type;
    },
    selectScheduleTarget(target: number){
      this.scheduleStore.scheduleTarget = target;
    }
  }
};
</script>