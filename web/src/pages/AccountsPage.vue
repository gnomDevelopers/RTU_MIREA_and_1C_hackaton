<template>
  <div class="w-full h-full scrollable flex flex-col items-center">
    <div class="flex flex-col gap-y-4 w-full px-2 us:px-4 lg:w-10/12 lg:px-0">
      
      <PageTitle title="Создание учётных записей" description="В этом разделе Вы можете создать учётную запись для администратора, заведующего кафедрой, преподавателя или студента. Для этого загрузите на своё устройство шаблон для заполнения данных о пользователе, а затем отправьте заполненный файл в специальное поле.">
        <IconAccounts class="w-20 us:w-36 h-20 us:h-36"/>
      </PageTitle>

      <MainControl title="Создание">
        <MainControlItem title="Создание аккаунтов">

          <AccountsCreateUsers class="flex flex-col gap-y-4 items-stretch"/>
          
        </MainControlItem>
      </MainControl>

      <MainControl title="Управление вузом">
        <MainControlItem title="Управление деканами">

          <SearchList 
            title="Введите ФИО декана" 
            placeholder="ФИО декана"
            :searchList="decanList" 
            :itemComponent="getListItemComponent"
          />

        </MainControlItem>
        <MainControlItem title="Управление преподавателями">
          
          <SearchList 
            title="Введите ФИО преподавателя" 
            placeholder="ФИО преподавателя"
            :searchList="prepodList" 
            :itemComponent="getListItemComponent"
          />

        </MainControlItem>
      </MainControl>

      <MainControl title="Управление ресурсами">
        <MainControlItem title="Управление аудиториями">
          
          <AccountsCreateAuditory class="flex flex-col gap-y-4 items-stretch"/>

        </MainControlItem>
      </MainControl>

    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { 
  ROLES_SET_PRORECTOR, 
  StatusCodes, 
  type ISearchList, 
  type IUsersList, 
} from '../helpers/constants';


import MainControl from '../shared/mainControl.vue';
import MainControlItem from '../shared/mainControlItem.vue';
import UserListItem from '../shared/userListItem.vue';
import SearchList from '../entities/searchList.vue';
import IconAccounts from '@/shared/iconAccounts.vue';
import PageTitle from '@/shared/pageTitle.vue';
import AccountsCreateUsers from '@/widgets/accountsCreateUsers.vue';
import AccountsCreateAuditory from '@/widgets/accountsCreateAuditory.vue';

export default {
  components:{
    PageTitle,
    IconAccounts,
    MainControl,
    MainControlItem,
    UserListItem,
    SearchList,
    AccountsCreateUsers,
    AccountsCreateAuditory,
  },
  data(){
    return{
      rolesList: ROLES_SET_PRORECTOR,
      decanList: [] as ISearchList[],
      prepodList: [] as ISearchList[],
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore),
    
    getListItemComponent(){
      return UserListItem;
    },
  },
  mounted(){
    for(let i = 1; i < 110; i++) {
      const data: IUsersList = {id: i, name: `Деканов${i} Декан${i} Деканович${i}`, role: 2};
      this.decanList.push({id: data.id, search_field: data.name, data: data});
    }

    for(let i = 1; i < 110; i++) {
      const data: IUsersList = {id: i, name: `Преподов${i} Препод${i} Преподович${i}`, role: 4};
      this.prepodList.push({id: data.id, search_field: data.name, data: data});
    }
  },
  methods:{
    
  }
};
</script>