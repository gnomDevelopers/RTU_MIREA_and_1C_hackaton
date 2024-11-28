<template>
  <div class="flex flex-col items-center scrollable px-4 lg:px-0">
    <div class="flex flex-col w-full lg:w-10/12 items-center gap-y-4 mb-4">
      
      <PageTitle title="Создание учебных групп" description="В этом разделе Вы можете создать новые или отредактировать уже существующие учебные группы. Для этого найдите в поиске нужную группу и выберите действие, которое хотите совершить.">
        <IconGroups class="w-20 us:w-36 h-20 us:h-36"/>
      </PageTitle>

      <div>
        <SearchList 
          title="" 
          placeholder="Введите название группы"
          :searchList="groupsSearchList" 
          :itemComponent="getListItemComponent"
        />
      </div>
      <div v-if="isGroupSelected" class="flex flex-col items-center gap-y-4 rounded-xl p-4 bg-color-light min-w-20 sm:min-w-[500px]">
        <div class="flex flex-col items-stretch w-full gap-y-1">
          <div v-for="(item, index) in getGroupMembers" class="flex flex-row flex-nowrap items-stretch min-h-10 rounded-lg bg-white ">
            <div class="flex flex-row justify-center items-center w-16 rounded-l-lg bg-color-bold">
              <p class="text-white text-xl">{{ index + 1 }}</p>
            </div>
            <div class="flex flex-row flex-grow justify-start items-center px-2">
              <p class="text-xl">{{ item.last_name }} {{ item.first_name }} {{ item.father_name }}</p>
            </div>
            <div class="flex flex-row justify-center items-center px-2">
              <svg @click="deleteStudent(item)" class="w-5 h-5 cursor-pointer" viewBox="0 0 19 19" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M6 9.5H13M18.25 9.5C18.25 14.3325 14.3325 18.25 9.5 18.25C4.66751 18.25 0.75 14.3325 0.75 9.5C0.75 4.66751 4.66751 0.75 9.5 0.75C14.3325 0.75 18.25 4.66751 18.25 9.5Z" stroke="#063C73" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </div>
          </div>
        </div>
        <div class="flex flex-row items-stretch gap-x-4 w-full">
          <div class="flex flex-row justify-center items-center">
            <svg @click="addStudent" class="w-7 h-7 cursor-pointer" viewBox="0 0 28 28" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M14 9V19M9 14H19M26.5 14C26.5 20.9036 20.9036 26.5 14 26.5C7.09644 26.5 1.5 20.9036 1.5 14C1.5 7.09644 7.09644 1.5 14 1.5C20.9036 1.5 26.5 7.09644 26.5 14Z" stroke="#063C73" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </div>
          <div class="flex flex-grow">
            <input 
              type="text"
              v-model="inputStudentFIO" 
              class="w-full px-2 py-1 min-w-20 max-w-none outline-none rounded-lg text-lg font-medium header-shadow" 
              placeholder="Введите ФИО студента">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useUniversityStore } from '@/stores/universityStore';
import { useGroupCorrectPageStore } from '@/stores/groupCorrectPageStore';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { type ISearchList, type IItemList, type IUserGet, StatusCodes } from '@/helpers/constants';

import GroupCorrectSearchListItem from '@/entities/listItems/groupCorrectSearchListItem.vue';
import SearchList from '@/entities/searchList.vue';
import IconGroups from '@/shared/iconGroups.vue';
import PageTitle from '@/shared/pageTitle.vue';

export default{
  components:{
    PageTitle,
    IconGroups,
    SearchList,
    GroupCorrectSearchListItem,
  },
  data(){
    return{
      groupsSearchList: [] as ISearchList[],

      inputStudentFIO: '',
    }
  },
  computed:{
    ...mapStores(useUniversityStore, useGroupCorrectPageStore, useStatusWindowStore),

    getListItemComponent(){
      return GroupCorrectSearchListItem;
    },

    isGroupSelected(){
      return this.groupCorrectPageStore.selectedGroupID !== null;
    },

    getGroupMembers(){
      return this.universityStore.groupMembersList;
    }
  },
  methods: {
    addStudent(){
      if(this.inputStudentFIO === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите фио студента!');
        return;
      }
      const input = this.inputStudentFIO.split(' ');
      if(input.length !== 3){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Некорректное фио студента!');
        return;
      }
      this.universityStore.groupMembersList.push({
        id: 12, 
        first_name: input[1], 
        last_name: input[0], 
        father_name: input[2], 
        role: 6, 
        faculty_id: 1, 
        department_id: 1, 
        educational_direction: 'Фуллстек разработка', 
        group_id: 1,
        university_id: 1,
      });
      this.universityStore.groupMembersList = this.universityStore.sortPeople(this.universityStore.groupMembersList);
      this.inputStudentFIO = '';
    },
    deleteStudent(studentDelete: IUserGet){
      for(let i = 0; i < this.universityStore.groupMembersList.length; i++){
        if(this.universityStore.groupMembersList[i].id === studentDelete.id){
          this.universityStore.groupMembersList.splice(i, 1);
        }
      }
    }
  },
  watch: {
    'universityStore.groupsList' : {
      handler(val: IUserGet[]){
        this.groupsSearchList = [];
        for(let item of val){
          this.groupsSearchList.push({id: item.id, search_field: `${item.last_name} ${item.first_name} ${item.father_name}`, data: item});
        }
      },
      immediate: true,
      deep: true,
    },
  }
};
</script>