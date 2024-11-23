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
          <CalendarTable :month="10"/>
        </div>

      </div>
      <div class="flex flex-col gap-y-4 items-center lg:w-1/2">
        <ScheduleClassList :canAddFaculties="true"/>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
import { SCHEDULE_TARGET_TEXT, type ISearchList, type IItemList } from '@/helpers/constants';

import CalendarTable from '@/entities/calendarTable.vue';
import SearchList from '@/entities/searchList.vue';
import ScheduleSearchListItem from '@/entities/scheduleSearchListItem.vue';
import ScheduleClassList from '@/entities/scheduleClassList.vue';

export default {
  components:{
    CalendarTable,
    SearchList,
    ScheduleSearchListItem,
    ScheduleClassList,
  },
  data(){
    return{
      searchFilter: '' as string,
      groupsList: [] as ISearchList[],
      teachersList: [] as ISearchList[],
      facultyList: [] as ISearchList[],
    }
  },
  computed:{
    ...mapStores(useScheduleStore),

    getScheduleTargetText(){
      return SCHEDULE_TARGET_TEXT[this.scheduleStore.scheduleTarget];
    },
    getScheduleTargetList(){
      if (this.scheduleStore.scheduleTarget === 0) return this.groupsList;
      else if(this.scheduleStore.scheduleTarget === 1) return this.teachersList;
      else if(this.scheduleStore.scheduleTarget === 2) return this.facultyList;
      return [] as ISearchList[];
    },
    getListItemComponent(){
      return ScheduleSearchListItem;
    }
  },
  mounted(){
    this.scheduleStore.loadScheduleTable();

    for(let i = 1; i < 20; i++){
      const data: IItemList = {id: i, name: `ЭФБО-${(i < 10 ? '0' : '')}${i}-23`};
      this.groupsList.push({id: data.id, search_field: data.name, data: data});
    }
    for(let i = 1; i < 20; i++){
      const data: IItemList = {id: i, name: `Преподов${i} Препод${i} Преподович${i}`};
      this.teachersList.push({id: data.id, search_field: data.name, data: data});
    }
    for(let i = 1; i < 20; i++){
      const data: IItemList = {id: i, name: `Мобильная разработка${i}`};
      this.facultyList.push({id: data.id, search_field: data.name, data: data});
    }
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