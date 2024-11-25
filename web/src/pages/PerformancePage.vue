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
          <SubjectItem :data="{id: 1, name: 'Дискретная математика (часть 2/2) [I.24-25]'}"/>
          <SubjectItem :data="{id: 2, name: 'Иностранный язык (часть 3/4) [I.24-25]'}"/>
          <SubjectItem :data="{id: 3, name: 'Математический анализ ФД_(часть 3/3) [I.24-25]'}"/>
          <SubjectItem :data="{id: 4, name: 'Программирование корпоративных систем (часть 1/4) [I.24-25]'}"/>
          <SubjectItem :data="{id: 5, name: 'Создание программного обеспечения (часть 1/2) [I.24-25]'}"/>
          <SubjectItem :data="{id: 6, name: 'Технологии индустриального программирования (часть 3/3) [I.24-25]'}"/>
        </div>
      </div>

      <div v-if="tableType === 0 && isSelectedDiscipline" class="flex flex-row self-stretch items-start flex-wrap-0">
        <table class="flex-shrink-0 cursor-default table-decorate">
          <thead>
            <tr>
              <th class="w-10 h-9">№</th>
              <th class="text-nowrap h-9">ФИО</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in getGroupMembers" :key="item.id">
              <td class="font-semibold h-9">{{ index + 1 }}</td>
              <td class="max-w-96 h-9 overflow-hidden text-nowrap text-left">{{ item.surname }} {{ item.name }} {{ item.thirdname }}</td>
            </tr>
          </tbody>
        </table>

        <div class="overflow-x-scroll scrollable-table cursor-default">
          <table class="w-auto no-x-border table-decorate">
            <thead>
              <tr>
                <th v-for="i in getGroupMembersScores[0]?.scores" class="w-16 min-w-10 h-9">01.09</th>
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
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in getGroupMembersScores">
              <td class="font-semibold h-9">{{ item.avg.toFixed(2) }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="tableType === 1 && isSelectedDiscipline" class="flex flex-row items-start flex-wrap-0 self-stretch overflow-x-scroll scrollable-table ">
        <table class="cursor-default table-decorate">
          <thead>
            <tr>
              <th class="w-10">№</th>
              <th class="max-w-96 overflow-hidden text-nowrap">ФИО</th>
              <th v-for="i in getGroupMembersScores[0].scores">01.09</th>
              <th>Ср.балл</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in getGroupMembersScores">
              <td class="font-semibold">{{ index + 1 }}</td>
              <td>{{ item.user.surname }} {{ item.user.name }} {{ item.user.thirdname }}</td>
              <td v-for="score in item.scores">{{ (score !== 0 ? score : '') }}</td>
              <td class="font-semibold">{{ item.avg.toFixed(2) }}</td>
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
import { type ISearchList, type IItemList, type IUserGet, type IGroupScores } from '@/helpers/constants';

import SearchList from '@/entities/searchList.vue';
import IconPerformance from '@/shared/iconPerformance.vue';
import SubjectItem from '@/shared/subjectItem.vue';
import PerformanceSearchListItem from '@/entities/listItems/performanceSearchListItem.vue';
import PageTitle from '@/shared/pageTitle.vue';

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
      groupsSearchList: [] as ISearchList[],
    }
  },
  computed:{
    ...mapStores(useUserInfoStore, useUniversityStore, usePerformancePageStore),

    getListItemComponent(){
      return PerformanceSearchListItem;
    },

    isSelectedGroup(){
      return this.performancePageStore.selectedGroupID !== null;
    },
    isSelectedDiscipline(){
      return this.performancePageStore.selectedDiscipline !== null;
    },

    getSelectedGroup(){
      if(this.performancePageStore.selectedGroupID === null) return '';
      for(let group of this.universityStore.groupsList){
        if(group.id === this.performancePageStore.selectedGroupID) return group.name;
      }
      return '';
    },

    getGroupMembers(){
      return this.universityStore.groupMembersList;
    },

    getGroupMembersScores(){
      return this.universityStore.groupMembersScores;
    }
  },
  mounted() {
    this.setTableType();
    window.addEventListener('resize', this.setTableType);
  },
  methods:{
    setTableType(){
      if(window.innerWidth < 756) this.tableType = 1;
      else this.tableType = 0;
    },
  },
  unmounted() {
    window.removeEventListener('resize', this.setTableType);
  },
  watch: {
    'universityStore.groupsList' : {
      handler(val: IUserGet[]){
        this.groupsSearchList = [];
        for(let item of val){
          this.groupsSearchList.push({id: item.id, search_field: `${item.surname} ${item.name} ${item.thirdname}`, data: item});
        }
      },
      immediate: true,
      deep: true,
    },
  }
};
</script>