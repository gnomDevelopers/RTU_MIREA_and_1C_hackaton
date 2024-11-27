<template>
  <div class="flex flex-col w-full items-center scrollable">
    <div class="flex flex-col gap-y-4 lg:flex-row xl:w-10/12 pt-4 px-2 us:px-0">

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
            class="h-80"
          />
        </Transition>
        
        <div class="flex flex-col items-center gap-y-2">
          <CalendarTable :month="10" @select-day="selectDay"/>
        </div>

      </div>
      <div class="flex flex-col gap-y-4 items-center lg:w-1/2">
        <ScheduleClassList v-if="scheduleStore.scheduleType === 0 || (scheduleStore.scheduleType === 1 && scheduleStore.selectedSheduleGroup !== null)" :canAddFaculties="scheduleStore.scheduleType === 0"/>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
import { useUniversityStore } from '@/stores/universityStore';
import { SCHEDULE_TARGET_TEXT, type ISearchList, type IItemList, type IUserGet, type Day } from '@/helpers/constants';

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
  data(){
    return{
      groupsSearchList: [] as ISearchList[],
      teachersSearchList: [] as ISearchList[],
      facultativeList: [] as ISearchList[],
    }
  },
  computed:{
    ...mapStores(useScheduleStore, useUniversityStore),

    getScheduleTargetText(){
      return SCHEDULE_TARGET_TEXT[this.scheduleStore.scheduleTarget];
    },
    getScheduleTargetList(){
      if (this.scheduleStore.scheduleTarget === 0) return this.groupsSearchList;
      else if(this.scheduleStore.scheduleTarget === 1) return this.teachersSearchList;
      else if(this.scheduleStore.scheduleTarget === 2) return this.facultativeList;
      return [] as ISearchList[];
    },
    getListItemComponent(){
      return ScheduleSearchListItem;
    }
  },
  mounted(){
    this.scheduleStore.loadScheduleTableByGroupName('ЭФБО-01-23');
  },
  methods:{
    selectScheduleType(type: number){
      if(type === 0) this.scheduleStore.selectedSheduleGroup = null;
      this.scheduleStore.scheduleType = type;
    },
    selectScheduleTarget(target: number){
      this.scheduleStore.selectedSheduleGroup = null;
      this.scheduleStore.scheduleTarget = target;
    },
    selectDay(day: Day){
      this.universityStore.selectedDate = `${day.day}.${day.month+1}.${day.year}`;
      this.scheduleStore.scheduleTableDay = [];
      for(let item of this.scheduleStore.scheduleData){
        if(item.date === `${day.day}.${day.month+1}.${day.year}`){
          this.scheduleStore.scheduleTableDay = item.timeTable;
          break;
        }
      }
    }
  },
  watch: {
    'universityStore.groupsList' : {
      handler(val: IUserGet[]){
        this.groupsSearchList = [];
        for(let item of val){
          this.groupsSearchList.push({id: item.id, search_field: `${item.surname} ${item.name} ${item.thirdname}`, data: item});
        }
        console.log('universityStore.groupsList: ', this.universityStore.groupsList);
        console.log('groupsSearchList: ', this.groupsSearchList);
      },
      immediate: true,
      deep: true,
    },
    'universityStore.teachersList' : {
      handler(val: IUserGet[]){
        this.teachersSearchList = [];
        for(let item of val){
          this.teachersSearchList.push({id: item.id, search_field: `${item.surname} ${item.name} ${item.thirdname}`, data: item});
        }
        console.log('universityStore.teachersList: ', this.universityStore.teachersList);
        console.log('teachersSearchList: ', this.teachersSearchList);
      },
      immediate: true,
      deep: true,
    },
    'universityStore.facultativesList' : {
      handler(val: IUserGet[]){
        this.facultativeList = [];
        for(let item of val){
          this.facultativeList.push({id: item.id, search_field: item.name, data: item});
        }
        console.log('universityStore.facultativesList: ', this.universityStore.facultativesList);
        console.log('facultativesSearchList: ', this.facultativeList);
      },
      immediate: true,
      deep: true,
    },
  }
};
</script>