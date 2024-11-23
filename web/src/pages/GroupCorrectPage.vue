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
          :searchList="groupsList" 
          :itemComponent="getListItemComponent"
        />
      </div>
      <div class="flex flex-col items-center gap-y-4 rounded-xl p-4 bg-color-light min-w-20 sm:min-w-[500px]">
        <div class="flex flex-col items-stretch w-full gap-y-1">
          <div v-for="i in 30" class="flex flex-row flex-nowrap items-stretch min-h-10 rounded-lg bg-white ">
            <div class="flex flex-row justify-center items-center w-16 rounded-l-lg bg-color-bold">
              <p class="text-white text-xl">{{ i }}</p>
            </div>
            <div class="flex flex-row flex-grow justify-start items-center px-2">
              <p class="text-xl">Иванов Иван Иванович</p>
            </div>
            <div class="flex flex-row justify-center items-center px-2">
              <svg class="w-5 h-5 cursor-pointer" viewBox="0 0 19 19" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M6 9.5H13M18.25 9.5C18.25 14.3325 14.3325 18.25 9.5 18.25C4.66751 18.25 0.75 14.3325 0.75 9.5C0.75 4.66751 4.66751 0.75 9.5 0.75C14.3325 0.75 18.25 4.66751 18.25 9.5Z" stroke="#063C73" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </div>
          </div>
        </div>
        <div class="flex flex-row items-stretch gap-x-4 w-full">
          <div class="flex flex-row justify-center items-center">
            <svg class="w-7 h-7 cursor-pointer" viewBox="0 0 28 28" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M14 9V19M9 14H19M26.5 14C26.5 20.9036 20.9036 26.5 14 26.5C7.09644 26.5 1.5 20.9036 1.5 14C1.5 7.09644 7.09644 1.5 14 1.5C20.9036 1.5 26.5 7.09644 26.5 14Z" stroke="#063C73" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </div>
          <div class="flex flex-grow">
            <input class="w-full px-2 py-1 min-w-20 max-w-none outline-none rounded-lg text-lg font-medium header-shadow" type="text" placeholder="Введите ФИО студента">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import GroupCorrectSearchListItem from '@/entities/listItems/groupCorrectSearchListItem.vue';
import SearchList from '@/entities/searchList.vue';
import { type ISearchList, type IItemList } from '@/helpers/constants';
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
      groupsList: [] as ISearchList[],
    }
  },
  computed:{
    getListItemComponent(){
      return GroupCorrectSearchListItem;
    }
  },
  mounted(){
    for(let i = 1; i < 20; i++){
      const data: IItemList = {id: i, name: `ЭФБО-${(i < 10 ? '0' : '')}${i}-23`};
      this.groupsList.push({id: data.id, search_field: data.name, data: data});
    }
  },
};
</script>