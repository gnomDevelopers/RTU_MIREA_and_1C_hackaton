<template>
  <div class="flex flex-col w-full items-center scrollable">
    <div class="flex flex-col gap-y-4 lg:flex-row xl:w-10/12 pt-6 px-2 us:px-0">
      <div class="flex flex-col w-full lg:w-1/2 items-center gap-y-5">
        <p class=" text-4xl font-medium text-center">Просмотр расписания</p>
        <div class="flex flex-row gap-x-2 p-2 rounded-lg bg-color-light w-full us:w-auto">
          <p class="text-lg cursor-pointer px-2 text-center">Моё расписание</p>
          <p class="text-lg cursor-pointer px-2 text-center bg-white rounded-md">Общее расписание</p>
        </div>
        <div class="flex flex-row flex-wrap justify-center gap-2 p-2 rounded-lg bg-color-light">
          <p class="text-lg cursor-pointer px-2 bg-white rounded-md">Группа</p>
          <p class="text-lg cursor-pointer px-2">Преподаватель</p>
          <p class="text-lg cursor-pointer px-2">Факультатив</p>
        </div>
        <div class="flex flex-row w-[314px] mb:w-auto rounded-lg header-shadow bg-white">
          <input 
            class="p-2 outline-none min-w-0 max-w-none flex-grow mb:w-96 bg-transparent" 
            type="text" 
            placeholder="Выберите группу для просмотра"
            v-model="searchFilter">
          <div 
            class="w-10 h-10 flex flex-col justify-center items-center cursor-pointer rounded-r-lg bg-slate-200 hover:bg-slate-300 active:bg-slate-400"
            @click="filterUserList">
            <img class="w-7 h-7" src="../assets/icons/icon-search.svg"/>
          </div>
        </div>
        <div class="flex flex-col items-center gap-y-2">
          <CalendarTable :month="11"/>
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
              :index="1" 
              :time="''" 
              :title="''" 
              :room="''" 
              :group="''"/>
            <ScheduleItem 
              :index="2" 
              :time="'10.40-12.10 ПР'" 
              :title="'Создание программного обеспечения'" 
              :room="'В-408 (В-78)'" 
              :group="'ЭФБО-01-23'"/>
            <ScheduleItem 
              :index="3" 
              :time="''" 
              :title="''" 
              :room="''" 
              :group="''"/>
            <ScheduleItem 
              :index="4" 
              :time="'10.40-12.10 ПР'" 
              :title="'Создание программного обеспечения'" 
              :room="'В-408 (В-78)'" 
              :group="'ЭФБО-01-23'"/>
            <ScheduleItem 
              :index="5" 
              :time="'10.40-12.10 ПР'" 
              :title="'Создание программного обеспечения'" 
              :room="'В-408 (В-78)'" 
              :group="'ЭФБО-01-23'"/>
            <ScheduleItem 
              :index="6" 
              :time="''" 
              :title="''" 
              :room="''" 
              :group="''"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
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
  methods:{
    filterUserList(){

    }
  }
};
</script>