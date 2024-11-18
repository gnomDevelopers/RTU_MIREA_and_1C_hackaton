<template>
  <div class="flex flex-col gap-y-2 mt-2 max-h-[600px]">
    <p class=" text-lg">{{ title }}</p>
    <div class="flex flex-row mb:self-start rounded-lg bg-white header-shadow ">

      <input 
        class=" bg-transparent p-2 outline-none min-w-0 max-w-none w-full flex-grow mb:w-96" 
        type="text" 
        :placeholder
        v-model="searchFilter"
      >

      <div 
        class="w-10 h-10 flex flex-col justify-center items-center cursor-pointer rounded-r-lg bg-slate-200 hover:bg-slate-300 active:bg-slate-400"
        @click="filterUserList"
        >
        <img class="w-7 h-7" src="../assets/icons/icon-search.svg"/>
      </div>

    </div>
    <div class="flex flex-col p-2 pr-0 rounded-lg overflow-hidden bg-white">
      <div class="flex-grow flex flex-col gap-y-2 pr-2 scrollable">
        <component 
          v-for="item in filteredUsersList" 
          :key="item.id" 
          :is="itemComponent" 
          :id=item.id 
          :name=item.name 
          :role=item.role 
        /> 
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { type PropType } from 'vue';
import { type IUsersList } from '@/helpers/constants';

export default{
  props: {
    title: {
      type: String,
      required: true,
    },
    placeholder:{
      type: String,
      required: false,
      default: '',
    },
    usersList: {
      type: Array as PropType<IUsersList[]>,
      required: true,
    },
    itemComponent: {
      required: true,
    },
  },
  data(){
    return{
      searchFilter: '',
      filteredUsersList: this.usersList as IUsersList[],
    }
  },
  methods:{
    filterUsersListInput(){
      if(this.searchFilter === '' || this.usersList.length > 1000) return this.usersList;
      this.filterUserList();
    },
    filterUserList(){
      if(this.searchFilter === '') return this.usersList; 
      this.filteredUsersList = this.usersList.filter(item => item.name.toLowerCase().includes(this.searchFilter.toLowerCase()));
    }
  },
  watch:{
    searchFilter(val){
      if(val === '') {
        this.filteredUsersList = this.usersList;
        return;
      }
      if(this.usersList.length > 1000) return;
      else this.filterUserList();
    }
  }
};

</script>
