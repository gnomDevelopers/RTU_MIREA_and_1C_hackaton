<template>
  <div>
    <div>
      <div>
        <IconPerformance/>
        <div>
          <p>Успеваемость</p>
          <p>В этом разделе Вы можете посмотреть и изменить успеваемость студентов. Для этого найдите в поиске нужную группу и предмет, а затем посмотрите или отредактируйте успеваемость.</p>
        </div>
      </div>
      <div>
        <div class="flex flex-row w-[314px] mb:w-auto rounded-lg header-shadow bg-white">
          <input 
            class="p-2 outline-none min-w-0 max-w-none flex-grow mb:w-96 bg-transparent" 
            type="text" 
            :placeholder="'Введите название группы'"
            v-model="searchFilter">
          <div 
            class="w-10 h-10 flex flex-col justify-center items-center cursor-pointer rounded-r-lg btn"
            @click="filterUserList">
            <img class="w-7 h-7" src="../assets/icons/icon-search.svg"/>
          </div>
        </div>
        
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import IconPerformance from '@/shared/iconPerformance.vue';

export default{
  components:{
    IconPerformance
  },
  data(){
    return{
      searchFilter: '' as string,
    }
  },
  methods:{
    filterUserList(){

    }
  }
};
</script>