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
            :search-list="groupsSearchList" 
            :item-component="getListItemComponent"
            class="h-48"
          />
          
          <ScheduleClassList v-if="isGroupSelected" :canAddFaculties="false"/>
          
          <CalendarTable v-if="isGroupSelected && isClassSelected" :month="10" @selectDay="handleSelectDay" />

        </div>

        <div class="flex flex-col flex-shrink-0 gap-y-4 items-center">
          <div v-if="isGroupSelected && isClassSelected" class="flex flex-col items-center gap-y-4 min-w-20 xl:min-w-[600px] p-4 rounded-xl bg-color-light">
            <div class="flex flex-col w-full items-stretch gap-y-1">
              <AttendanceUsersListItem 
                v-for="(item, index) in getAttendaceList" 
                :key="item.user.id" 
                :index="index + 1"
                :data="item.user"
              />
            </div>
            <div @click="saveAttendance" class="py-1 px-6 rounded-xl cursor-pointer bg-color-bold">
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
import { useUniversityStore } from '@/stores/universityStore';
import { useAttendacePageStore } from '@/stores/attendacePageStore';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { type ISearchList, type IItemList, type IUserGet, type IGroupAttendance, StatusCodes, type Day, GET_CORRECT_DATE } from '@/helpers/constants';

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
  computed:{
    ...mapStores(useScheduleStore, useUniversityStore, useAttendacePageStore, useStatusWindowStore, useUserInfoStore),

    groupsSearchList(): ISearchList[] {
      const arr:ISearchList[] = [];
      if(this.attendancePageStore.attendanceGroups.length === 0) return arr;
      for(let item of this.attendancePageStore.attendanceGroups){
        arr.push({id: this.universityStore.tmpuserID++, search_field: item, data: {name: item}});
      }
      return arr;
    },

    getListItemComponent(){
      return AttendanceSearchListItem;
    },
    isGroupSelected(){
      return this.attendancePageStore.selectedGroup !== null;
    },
    isClassSelected(){
      return this.scheduleStore.selectedClass !== null;
    },
    getAttendaceList(){
      return this.attendancePageStore.attendanceGroupMembers;
    },
  },
  mounted() {
    //загружаем доступные для посещения группы
    this.attendancePageStore.loadScheduleGroups();
    //если у пользователя есть группа, выбираем ее по дефолту
    if(this.userInfoStore.group_name) this.attendancePageStore.selectedGroup = this.userInfoStore.group_name;

  },
  methods: {
    saveAttendance(){
      // const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Сохраняем посещения...', -1);
      // setTimeout(() => {
      //   this.statusWindowStore.deteleStatusWindow(stID);
      //   this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Посещения сохранены!');
      // }, 300);
    },
    handleSelectDay(day: Day){
      this.attendancePageStore.selectedDate = GET_CORRECT_DATE(day.day, day.month, day.year);
    }
  },
};
</script>