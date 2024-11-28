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
      const arr:IUserGet[] = [];
      for(let student of this.universityStore.groupMembersList){
        if(student.group_id === this.groupCorrectPageStore.selectedGroupID){
          arr.push(student);
        }
      }
      return arr;
    },

    groupsSearchList(): ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.universityStore.groupsList.length === 0) return arr;
      for(let item of this.universityStore.groupsList){
        arr.push({id: item.id, search_field: item.name, data: item});
      }
      return arr;
    }
  },
  methods: {
    deleteStudent(studentDelete: IUserGet){
      for(let i = 0; i < this.universityStore.studentsList.length; i++){
        if(this.universityStore.studentsList[i].id === studentDelete.id){
          this.universityStore.studentsList.splice(i, 1);
        }
      }
    }
  },
};
</script>