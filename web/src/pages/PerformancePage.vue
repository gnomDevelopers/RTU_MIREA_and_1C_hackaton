<template>
  <div class="flex flex-col items-center scrollable px-4 lg:px-0">
    <div class="flex flex-col w-full lg:w-10/12 items-center gap-y-4 mb-4">

      <PageTitle title="Успеваемость" description="В этом разделе Вы можете посмотреть и изменить успеваемость студентов. Для этого найдите в поиске нужную группу и предмет, а затем посмотрите или отредактируйте успеваемость.">
        <IconPerformance class="w-20 us:w-36 h-20 us:h-36"/>
      </PageTitle>

      <div>
        <SearchList 
          title="" 
          placeholder="Введите название группы"
          :searchList="groupsSearchList" 
          :itemComponent="getListItemComponent"
        />
        <p v-if="isSelectedGroup" class="text-lg text-center font-medium cursor-default">
          Выбранная группа: 
          <span class="underline cursor-pointer">{{ getSelectedGroup }}</span>
        </p>
      </div>

      <div v-if="isSelectedGroup" class="flex flex-col p-4 rounded-lg gap-y-4 bg-color-light">
        <p class="text-center text-xl p-1 rounded-lg cursor-default bg-white">Выберите дисциплину</p>
        <div class="flex flex-col gap-y-1">
          <SubjectItem v-for="(item, index) of disciplineList" :data="{name: item}" :index="index + 1"/>
        </div>
      </div>

      <div v-if="tableType === 0 && isSelectedDiscipline" class="flex flex-row self-stretch items-start flex-wrap-0">
        <table class="flex-shrink-0 cursor-default table-decorate">
          <thead>
            <tr>
              <th class="w-10 h-9">№</th>
              <th @click="sortByName" class="text-nowrap h-9 cursor-pointer">ФИО</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in getGroupMembersScores" :key="item.user.id">
              <td class="font-semibold h-9">{{ index + 1 }}</td>
              <td class="max-w-96 h-9 overflow-hidden text-nowrap text-left">{{ item.user.last_name }} {{ item.user.first_name }} {{ item.user.father_name }}</td>
            </tr>
          </tbody>
        </table>

        <div class="overflow-x-scroll scrollable-table cursor-default">
          <table class="w-auto no-x-border table-decorate">
            <thead>
              <tr>
                <th v-for="(i, index) in getGroupMembersScores[0]?.scores" class="w-16 min-w-10 h-9">02.09</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in getGroupMembersScores">
                <td v-for="score in item.scores" class="w-16 min-w-10 h-9">{{ (score === 0 ? '' : score) }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <table class="flex-shrink-0 cursor-default table-decorate">
          <thead>
            <tr>
              <th class=" h-9">Ср.балл</th>
              <th class=" h-9">GPA</th>
              <th class=" bg-transparent border-none border-transparent cursor-pointer">
                <svg @click="sortByGPA" width="33" height="21" viewBox="0 0 33 21" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M18.6025 2.66484H31.0837M18.6025 9.68639H27.5176M18.6025 16.7079H23.9515M7.46689 1.75V19.25M7.46689 19.25L1.91699 13.9004M7.46689 19.25L13.2287 13.9004" stroke="#063C73" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in getGroupMembersScores">
              <td class="font-semibold h-9">{{ item.avg.toFixed(2) }}</td>
              <td class="font-semibold h-9">{{ item.gpa.toFixed(2) }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <!--Мобильная версия таблицы-->
      <div v-if="tableType === 1 && isSelectedDiscipline" class="flex flex-row items-start flex-wrap-0 self-stretch overflow-x-scroll scrollable-table ">
        <table class="cursor-default table-decorate">
          <thead>
            <tr>
              <th class="w-10">№</th>
              <th @click="sortByName" class="max-w-96 overflow-hidden text-nowrap cursor-pointer">ФИО</th>
              <th v-for="(i, index) in getGroupMembersScores[0].scores">02.09</th>
              <th>Ср.балл</th>
              <th>GPA</th>
              <th class=" bg-transparent border-none border-transparent cursor-pointer">
                <svg @click="sortByGPA" width="33" height="21" viewBox="0 0 33 21" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M18.6025 2.66484H31.0837M18.6025 9.68639H27.5176M18.6025 16.7079H23.9515M7.46689 1.75V19.25M7.46689 19.25L1.91699 13.9004M7.46689 19.25L13.2287 13.9004" stroke="#063C73" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in getGroupMembersScores">
              <td class="font-semibold">{{ index + 1 }}</td>
              <td>{{ item.user.last_name }} {{ item.user.first_name }} {{ item.user.father_name }}</td>
              <td v-for="score in item.scores">{{ (score !== 0 ? score : '') }}</td>
              <td class="font-semibold">{{ item.avg.toFixed(2) }}</td>
              <td class="font-semibold">{{ item.gpa.toFixed(2) }}</td>
            </tr>
          </tbody>
        </table>
      </div>

    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { useUniversityStore } from '@/stores/universityStore';
import { usePerformancePageStore } from '@/stores/performancePageStore';
import { useScheduleStore } from '@/stores/scheduleStore';
import { type ISearchList, type IItemList, type IUserGet, type IGroupScores } from '@/helpers/constants';

import SearchList from '@/entities/searchList.vue';
import IconPerformance from '@/shared/iconPerformance.vue';
import SubjectItem from '@/shared/subjectItem.vue';
import PerformanceSearchListItem from '@/entities/listItems/performanceSearchListItem.vue';
import PageTitle from '@/shared/pageTitle.vue';
import { API_Disciplines_Group_Get } from '@/api/api';

export default{
  components:{
    PageTitle,
    IconPerformance,
    SubjectItem,
    SearchList,
    PerformanceSearchListItem,
  },
  data(){
    return{
      tableType: 0 as number, // 0 - таблица для больших экранов (>=756px), 1 - таблица для маленьких экранов (<756px)
      searchFilter: '' as string,
      disciplineList: [] as string[],
      // API_Disciplines_Group_Get
    }
  },
  computed:{
    ...mapStores(useUserInfoStore, useUniversityStore, usePerformancePageStore, useScheduleStore),

    groupsSearchList():ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.scheduleStore.scheduleGroups.length === 0) return arr;
      for(let item of this.scheduleStore.scheduleGroups){
        arr.push({id: this.universityStore.tmpuserID++, search_field: `${item}`, data: {name: item}});
      }
      console.log('groupsList: ', arr);
      return arr;
    },

    getGroupMembers(){
      return this.universityStore.groupMembersList;
    },

    getGroupMembersScores(){
      return this.performancePageStore.groupMembersScores;
    },

    getListItemComponent(){
      return PerformanceSearchListItem;
    },
    getSelectedGroup(){
      return this.performancePageStore.selectedGroup;
    },
    isSelectedGroup(){
      return this.performancePageStore.selectedGroup !== null;
    },
    isSelectedDiscipline(){
      return this.performancePageStore.selectedDiscipline !== null;
    },
  },
  mounted() {
    this.setTableType();
    window.addEventListener('resize', this.setTableType);

    this.performancePageStore.selectedGroup = this.userInfoStore.group_name; // автовыбор

    this.disciplineList = [];
    API_Disciplines_Group_Get()
    .then((response: any) => {
      for(let discipline of response.data){
        this.disciplineList.push(discipline.name);
      }
    })
    .catch(error => {
      // nothing
    })
  },
  methods:{
    setTableType(){
      if(window.innerWidth < 756) this.tableType = 1;
      else this.tableType = 0;
    },
    sortByGPA(){
      this.performancePageStore.groupMembersScores = this.universityStore.sortByGpa(this.performancePageStore.groupMembersScores);
    },
    sortByName(){
      this.performancePageStore.groupMembersScores = this.universityStore.sortByName(this.performancePageStore.groupMembersScores);
    }
  },
  unmounted() {
    window.removeEventListener('resize', this.setTableType);
  },
};
</script>