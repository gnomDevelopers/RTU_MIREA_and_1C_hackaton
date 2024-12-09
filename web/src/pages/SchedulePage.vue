<template>
  <div class="flex flex-col w-full items-center scrollable">
    <div class="flex flex-col gap-y-4 lg:flex-row w-full xl:w-10/12 pt-4 px-2 us:px-0">

      <div class="flex flex-col w-full lg:w-1/2 items-center gap-y-4">
        <p class="text-4xl font-medium text-center">Просмотр расписания</p>
        
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
          <SearchList 
            v-if="scheduleStore.scheduleType === 1"
            title="" 
            :placeholder="getScheduleTargetText"
            :searchList="getScheduleTargetList" 
            :itemComponent="getListItemComponent"
          />
        </Transition>
        
        <div class="flex flex-col items-center gap-y-2">
          <CalendarTable :month="10" @select-day="selectDay"/>
        </div>

      </div>
      <div class="flex flex-col gap-y-4 items-center lg:w-1/2">
        <ScheduleClassList v-if="scheduleStore.scheduleType === 0 || (scheduleStore.scheduleType === 1 && scheduleStore.selectedSheduleTarget !== null)" :canAddFaculties="scheduleStore.scheduleType === 0"/>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
import { useUniversityStore } from '@/stores/universityStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { SCHEDULE_TARGET_TEXT, type ISearchList, type IItemList, type IUserGet, type Day, type IGroup, GET_CORRECT_DATE } from '@/helpers/constants';

import CalendarTable from '@/entities/calendarTable.vue';
import SearchList from '@/entities/searchList.vue';
import ScheduleSearchListItem from '@/entities/listItems/scheduleSearchListItem.vue';
import ScheduleClassList from '@/entities/listItems/scheduleClassList.vue';

export default {
  components:{
    CalendarTable,
    SearchList,
    ScheduleSearchListItem,
    ScheduleClassList,
  },
  computed:{
    ...mapStores(useScheduleStore, useUniversityStore, useUserInfoStore),

    groupsSearchList():ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.scheduleStore.scheduleGroups.length === 0) return arr;
      for(let item of this.scheduleStore.scheduleGroups){
        arr.push({id: this.universityStore.tmpuserID++, search_field: `${item}`, data: {name: item}});
      }
      return arr;
    },

    teachersSearchList():ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.scheduleStore.scheduleTeachers.length === 0) return arr;
      for(let item of this.scheduleStore.scheduleTeachers){
        arr.push({id: this.universityStore.tmpuserID++, search_field: `${item}`, data: {name: item}});
      }
      return arr;
    },

    facultativesList():ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.scheduleStore.scheduleFacultatives.length === 0) return arr;
      for(let item of this.scheduleStore.scheduleFacultatives){
        arr.push({id: this.universityStore.tmpuserID++, search_field: `${item}`, data: {name: item}});
      }
      return arr;
    },

    getScheduleTargetText(){
      return SCHEDULE_TARGET_TEXT[this.scheduleStore.scheduleTarget];
    },
    getScheduleTargetList(){
      if (this.scheduleStore.scheduleTarget === 0) return this.groupsSearchList;
      else if(this.scheduleStore.scheduleTarget === 1) return this.teachersSearchList;
      else if(this.scheduleStore.scheduleTarget === 2) return this.facultativesList;
      return [] as ISearchList[];
    },
    getListItemComponent(){
      return ScheduleSearchListItem;
    }
  },
  async mounted(){
    //загружаем все группы с расписанием
    await this.scheduleStore.loadScheduleGroups();
    //загружаем всех преподов с расписанием
    await this.scheduleStore.loadScheduleTeachers();
    //загружаем все факультативы с расписанием
    await this.scheduleStore.loadScheduleFacultatives();

    //устанавливаем текущую дату
    const today = new Date();
    this.scheduleStore.selectedDate = GET_CORRECT_DATE(today.getDate(), today.getMonth() + 1, today.getFullYear());

    //если группа пользователя есть в списке групп с расписанием
    if(this.scheduleStore.scheduleGroups.includes(this.userInfoStore.group_name)){
      // загружаем расписание группы студента
      await this.scheduleStore.loadScheduleTableByGroupName(this.userInfoStore.group_name);
      // загружаем сегодняшнее расписание
      this.scheduleStore.selectScheduleDay(this.scheduleStore.selectedDate);
    }
    
  },
  methods:{
    selectScheduleType(type: number){
      if(type === 0) {
        this.scheduleStore.selectedSheduleTarget = null;
        this.scheduleStore.loadScheduleTableByGroupName(this.userInfoStore.group_name);
      }
      this.scheduleStore.scheduleType = type;
    },
    selectScheduleTarget(target: number){
      this.scheduleStore.selectedSheduleTarget = null;
      this.scheduleStore.scheduleTarget = target;
    },
    selectDay(day: Day){
      this.scheduleStore.selectScheduleDay(GET_CORRECT_DATE(day.day, day.month+1, day.year));
    }
  },
  watch: {
    'scheduleStore.selectedSheduleTarget': {
      handler(val){
        switch(this.scheduleStore.scheduleTarget){
          case 0: this.scheduleStore.loadScheduleTableByGroupName(val); break;
          case 1: this.scheduleStore.loadScheduleTableByTeacherName(val); break;
          case 2: this.scheduleStore.loadScheduleTableByClassName(val); break;
        }
        // загружаем текущее расписание
        this.scheduleStore.selectScheduleDay(this.scheduleStore.selectedDate);
      }
    }
  }
};
</script>