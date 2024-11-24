<template>
  <div class="w-full h-full scrollable flex flex-col items-center">
    <div class="flex flex-col gap-y-4 w-full px-2 us:px-4 lg:w-10/12 lg:px-0 mb-4">
      
      <PageTitle title="Создание учётных записей" description="В этом разделе Вы можете создать учётную запись для администратора, заведующего кафедрой, преподавателя или студента. Для этого загрузите на своё устройство шаблон для заполнения данных о пользователе, а затем отправьте заполненный файл в специальное поле.">
        <IconAccounts class="w-20 us:w-36 h-20 us:h-36"/>
      </PageTitle>

      <MainControl title="Инициализации">
        <MainControlItem title="Создание пользователей">

          <AccountsCreateUsers class="flex flex-col gap-y-4 items-stretch"/>

        </MainControlItem>
        <MainControlItem title="Создание аудиторий">

          <AccountsCreateAuditory class="flex flex-col gap-y-4 items-stretch"/>

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

          <SearchList 
            title="Введите название аудитории" 
            placeholder="Название аудитории"
            :searchList="auditoriesList" 
            :itemComponent="getAuditoryListItemComponent"
          />

        </MainControlItem>
      </MainControl>

    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUniversityStore } from '@/stores/universityStore';
import { 
  ROLES_SET_PRORECTOR, 
  AUDITORY_TYPE_LIST, 
  AUDITORY_PROFILE_LIST,
  StatusCodes, 
  type ISearchList, 
  type IUsersList, 
  type IAPI_Audience_Update, 
} from '../helpers/constants';


import MainControl from '../shared/mainControl.vue';
import MainControlItem from '../shared/mainControlItem.vue';
import UserListItem from '../entities/listItems/accountUserListItem.vue';
import SearchList from '../entities/searchList.vue';
import IconAccounts from '@/shared/iconAccounts.vue';
import PageTitle from '@/shared/pageTitle.vue';
import AccountsCreateUsers from '@/widgets/accountsCreateUsers.vue';
import AccountsCreateAuditory from '@/widgets/accountsCreateAuditory.vue';
import AccountAuditoryListItem from '@/entities/listItems/accountAuditoryListItem.vue';

export default {
  components:{
    PageTitle,
    IconAccounts,
    MainControl,
    MainControlItem,
    UserListItem,
    AccountAuditoryListItem,
    SearchList,
    AccountsCreateUsers,
    AccountsCreateAuditory,
  },
  data(){
    return{
      rolesList: ROLES_SET_PRORECTOR,
      decanList: [] as ISearchList[],
      prepodList: [] as ISearchList[],
      auditoriesList: [] as ISearchList[],

      auditoryID: 0,
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore, useUniversityStore),
    
    getListItemComponent(){
      return UserListItem;
    },
    getAuditoryListItemComponent(){
      return AccountAuditoryListItem;
    }
  },
  mounted(){
    //получение всей информации об университете и его составных
    this.universityStore.loadUniversityInfo();

    for(let i = 1; i < 110; i++) {
      const data: IUsersList = {id: i, name: `Деканов${i} Декан${i} Деканович${i}`, role: 2};
      this.decanList.push({id: data.id, search_field: data.name, data: data});
    }

    for(let i = 1; i < 110; i++) {
      const data: IUsersList = {id: i, name: `Преподов${i} Препод${i} Преподович${i}`, role: 4};
      this.prepodList.push({id: data.id, search_field: data.name, data: data});
    }

    for(let i = 1; i < 110; i++) {
      const data: IAPI_Audience_Update = {  
        id: i,
        campus_id: i % 3,
        capacity: 123,
        name: `A-${i}`,
        profile: AUDITORY_PROFILE_LIST[i % 8],
        type: AUDITORY_TYPE_LIST[i % 5],
      };
      this.auditoriesList.push({id: data.id, search_field: data.name, data: data});
    }
  },
  methods:{
    
  }
};
</script>