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

      <div v-if="getGroupGrades !== [] && tableType === 0 && isSelectedDiscipline" class="flex flex-row self-stretch items-start flex-wrap-0">
        <table class="flex-shrink-0 cursor-default table-decorate">
          <thead>
            <tr>
              <th class="w-10 h-9">№</th>
              <th @click="sortByName" class="text-nowrap h-9 cursor-pointer">ФИО</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in getGroupGrades" :key="item.id">
              <td class="font-semibold h-9">{{ index + 1 }}</td>
              <td class="max-w-96 h-9 overflow-hidden text-nowrap text-left">{{ item.last_name }} {{ item.first_name }} {{ item.father_name }}</td>
            </tr>
          </tbody>
        </table>

        <div class="overflow-x-scroll scrollable-table cursor-default">
          <table class="w-auto no-x-border table-decorate">
            <thead>
              <tr>
                <th v-for="item in getGroupGrades[0].grades" class="w-16 min-w-10 h-9">{{ item.date.slice(5) }}</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in getGroupGrades">
                <td v-for="score in item.grades" class="w-16 min-w-10 h-9">{{ (score.value === 0 ? '' : score) }}</td>
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
            <tr v-for="item in getGroupGrades">
              <td class="font-semibold h-9">{{ item.average_score.toFixed(2) }}</td>
              <td class="font-semibold h-9">{{ item.gpa.toFixed(2) }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <!--Мобильная версия таблицы-->
      <div v-if="getGroupGrades !== [] && tableType === 1 && isSelectedDiscipline" class="flex flex-row items-start flex-wrap-0 self-stretch overflow-x-scroll scrollable-table ">
        <table class="cursor-default table-decorate">
          <thead>
            <tr>
              <th class="w-10">№</th>
              <th @click="sortByName" class="max-w-96 overflow-hidden text-nowrap cursor-pointer">ФИО</th>
              <th v-for="item in getGroupGrades[0].grades">{{ item.date.slice(5) }}</th>
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
            <tr v-for="(item, index) in getGroupGrades">
              <td class="font-semibold">{{ index + 1 }}</td>
              <td>{{ item.last_name }} {{ item.first_name }} {{ item.father_name }}</td>
              <td v-for="score in item.grades">{{ (score.value !== 0 ? score.value : '') }}</td>
              <td class="font-semibold">{{ item.average_score.toFixed(2) }}</td>
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
import { type ISearchList, type IItemList, type IUserGet, type IReaorganizedGroupScore } from '@/helpers/constants';

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
    }
  },
  computed:{
    ...mapStores(useUserInfoStore, useUniversityStore, usePerformancePageStore, useScheduleStore),

    // возвращает список групп, доступных для поиска
    groupsSearchList():ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.scheduleStore.scheduleGroups.length === 0) return arr;
      for(let item of this.scheduleStore.scheduleGroups){
        arr.push({id: this.universityStore.tmpuserID++, search_field: `${item}`, data: {name: item}});
      }
      console.log('groupsList: ', arr);
      return arr;
    },

    getGroupGrades(){
      return this.performancePageStore.groupGrades;
    },

    // возвращает компонент для отрисовки группы
    getListItemComponent(){
      return PerformanceSearchListItem;
    },
    // возвращает выбранную группу
    getSelectedGroup(){
      return this.performancePageStore.selectedGroup;
    },
    // проверяет, выбрана ли группа
    isSelectedGroup(){
      return this.performancePageStore.selectedGroup !== null;
    },
    // проверяет, выбрана ли дисциплина
    isSelectedDiscipline(){
      return this.performancePageStore.selectedDiscipline !== null;
    },
  },
  async mounted() {
    //загружаем все группы с расписанием
    await this.scheduleStore.loadScheduleGroups();

    //устанавливаем тип таблицы расписания
    this.setTableType();
    window.addEventListener('resize', this.setTableType);

    //превыбираем группу пользователя если у него она есть
    if(this.userInfoStore.group_name !== '') {
      this.performancePageStore.selectedGroup = this.userInfoStore.group_name;
    }
  },
  methods:{
    // устанавливает размер группы
    setTableType(){
      if(window.innerWidth < 756) this.tableType = 1;
      else this.tableType = 0;
    },
    // сортирует по GPA
    sortByGPA(){
      // this.performancePageStore.groupGrades = this.universityStore.sortByGpa(this.performancePageStore.groupGrades);
    },
    // сортирует по ФИО студента
    sortByName(){
      // this.performancePageStore.groupGrades = this.universityStore.sortByName(this.performancePageStore.groupGrades);
    },
    // подгружает дисциплины выбранной группы
    loadSelectedGroupDiscipline(groupName: string){
      this.disciplineList = [];
      API_Disciplines_Group_Get(groupName)
      .then((response: any) => {
        for(let discipline of response.data){
          this.disciplineList.push(discipline.name);
        }
      })
      .catch(error => {
        // nothing
      });
    },

    
  },
  unmounted() {
    window.removeEventListener('resize', this.setTableType);
  },
  watch: {
    // при смнене дисциплины подгрузить успеваемость
    'performancePageStore.selectedDiscipline':{
      async handler(val){
        if(val !== null){
          console.log('load discipline');
          await this.performancePageStore.loadGroupGrades();
        }
      }
    },
    // при смене группы подгрузить ее дисциплины
    'performancePageStore.selectedGroup':{
      handler(val){
        if(val !== null){
          this.loadSelectedGroupDiscipline(val);
          return;
        }
        this.disciplineList = [];
      }
    },

  }
};
</script>