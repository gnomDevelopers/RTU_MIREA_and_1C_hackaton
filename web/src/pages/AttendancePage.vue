<template>
  <div class="flex flex-col items-center scrollable px-4 xl:px-0">
    <div class="flex flex-col w-full xl:w-10/12 items-center gap-y-4 mb-4">

      <PageTitle title="Посещаемость" description="В этом разделе Вы можете посмотреть и отметить посещаемость студентов. Для этого найдите в поиске нужную группу, выберите дату и отметьте посещаемость.">
        <IconAttendance class="w-20 us:w-36 h-20 us:h-36"/>
      </PageTitle>

      <div class="flex flex-col lg:flex-row justify-between flex-nowrap w-full gap-4">
        <div class="flex flex-col gap-y-4 items-center">
          <SearchList 
            title="" 
            placeholder="Выберите группу" 
            :search-list="groupsList" 
            :item-component="getListItemComponent"
            class="h-80"
          />
          
          <CalendarTable :month="10" />

          <ScheduleClassList :canAddFaculties="false"/>

        </div>

        <div class="flex flex-col flex-shrink-0 gap-y-4 items-center">
          <div class="flex flex-col items-center gap-y-4 min-w-20 xl:min-w-[600px] p-4 rounded-xl bg-color-light">
            <div class="flex flex-col w-full items-stretch gap-y-1">
              <AttendanceUsersListItem 
                v-for="(item, index) in studentsList" 
                :key="item.id" 
                :index="index + 1"
                :data="item"
              />
            </div>
            <div class="py-1 px-6 rounded-xl cursor-pointer bg-color-bold">
              <p class="text-2xl text-white">Заверить</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
import { type ISearchList, type IItemList } from '@/helpers/constants';

import AttendanceSearchListItem from '@/entities/listItems/attendanceSearchListItem.vue';
import CalendarTable from '@/entities/calendarTable.vue';
import ScheduleClassList from '@/entities/listItems/scheduleClassList.vue';
import SearchList from '@/entities/searchList.vue';
import IconAttendance from '@/shared/iconAttendance.vue';
import PageTitle from '@/shared/pageTitle.vue';
import AttendanceUsersListItem from '@/entities/listItems/attendanceUsersListItem.vue';

export default {
  components:{
    PageTitle,
    IconAttendance,
    CalendarTable,
    SearchList,
    AttendanceSearchListItem,
    ScheduleClassList,
    AttendanceUsersListItem,
  },
  data(){
    return {
      groupsList: [] as ISearchList[],
      studentsList: [] as any[],
    }
  },
  computed:{
    ...mapStores(useScheduleStore),

    getListItemComponent(){
      return AttendanceSearchListItem;
    }
  },
  mounted() {
    this.scheduleStore.loadScheduleTable();
    
    for(let i = 1; i < 20; i++){
      const data: IItemList = {id: i, name: `ЭФБО-${(i < 10 ? '0' : '')}${i}-23`};
      this.groupsList.push({id: data.id, search_field: data.name, data: data});
    }
    for(let i = 1; i < 30; i++){
      const data:any = {id: i, name: `Иванов Иван Иванович`};
      this.studentsList.push(data);
    }
  },
};
</script>