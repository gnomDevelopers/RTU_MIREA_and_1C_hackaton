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
        class="w-10 h-10 flex flex-col justify-center items-center cursor-pointer rounded-r-lg btn"
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
          :data="item.data"
        /> 
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { type PropType } from 'vue';
import { type ISearchList } from '@/helpers/constants';

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
    searchList: {
      type: Array as PropType<ISearchList[]>,
      required: true,
    },
    itemComponent: {
      required: true,
    },
  },
  data(){
    return{
      searchFilter: '',
      filteredUsersList: this.searchList as ISearchList[],
      inputFilterMaxSize: 1000, // предел по кол-ву элементов в списке, после которого сортировка после ввода не будет работать
    }
  },
  methods:{
    filterUsersListInput(){
      if(this.searchFilter === '' || this.searchList.length > this.inputFilterMaxSize) return this.searchList;
      this.filterUserList();
    },
    filterUserList(){
      if(this.searchFilter === '') return this.searchList; 
      this.filteredUsersList = this.searchList.filter(item => item.search_field.toLowerCase().includes(this.searchFilter.toLowerCase()));
    }
  },
  watch:{
    searchFilter(val){
      if(val === '') {
        this.filteredUsersList = this.searchList;
        return;
      }
      if(this.searchList.length > this.inputFilterMaxSize) return;
      else this.filterUserList();
    },
    searchList(val){ // изменился входной массив данных списка
      this.searchFilter = '';
      this.filteredUsersList = this.searchList;
    }
  }
};

</script>
