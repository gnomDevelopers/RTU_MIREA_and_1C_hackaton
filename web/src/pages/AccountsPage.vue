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
        <MainControlItem title="Создание кампусов">

          <AccountsCreateCampus class="flex flex-col gap-y-4 items-stretch"/>

        </MainControlItem>
      </MainControl>

      <MainControl title="Управление вузом">
        <MainControlItem title="Управление деканами">

          <SearchList 
            title="Введите ФИО декана" 
            placeholder="ФИО декана"
            :searchList="decansSearchList" 
            :itemComponent="getListItemComponent"
          />

        </MainControlItem>
        <MainControlItem title="Управление учебным отделом">
          
          <SearchList 
            title="Введите ФИО сотрудника учебного отдела" 
            placeholder="ФИО сотрудника учебного отдела"
            :searchList="educationDepartmentsSearchList" 
            :itemComponent="getListItemComponent"
          />

        </MainControlItem>
        <MainControlItem title="Управление заведущими кафедрами">
          
          <SearchList 
            title="Введите ФИО заведущего кафедрой" 
            placeholder="ФИО заведущего кафедрой"
            :searchList="zavCafsSearchList" 
            :itemComponent="getListItemComponent"
          />

        </MainControlItem>
        <MainControlItem title="Управление преподавателями">
          
          <SearchList 
            title="Введите ФИО преподавателя" 
            placeholder="ФИО преподавателя"
            :searchList="teachersSearchList" 
            :itemComponent="getListItemComponent"
          />

        </MainControlItem>
        <MainControlItem title="Управление студентами">
          
          <SearchList 
            title="Введите ФИО студента" 
            placeholder="ФИО студента"
            :searchList="studentsSearchList" 
            :itemComponent="getListItemComponent"
          />

        </MainControlItem>
      </MainControl>

      <MainControl title="Управление ресурсами">
        <MainControlItem title="Управление кампусами">

          <SearchList 
            title="Введите название кампуса" 
            placeholder="Название кампуса"
            :searchList="campusesSearchList"
            :itemComponent="getAuditoryListItemComponent"
          />

        </MainControlItem>
        <MainControlItem title="Управление аудиториями">

          <SearchList 
            title="Введите название аудитории" 
            placeholder="Название аудитории"
            :searchList="auditoriesSearchList"
            :itemComponent="getCampusListItemComponent"
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
  StatusCodes, 
  type ISearchList, 
  type IUserGet, 
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
import AccountsCreateCampus from '@/widgets/accountsCreateCampus.vue';
import AccountCampusListItem from '../entities/listItems/accountCampusLIstItem.vue';

export default {
  components:{
    PageTitle,
    IconAccounts,
    MainControl,
    MainControlItem,
    UserListItem,
    AccountAuditoryListItem,
    AccountCampusListItem,
    SearchList,
    AccountsCreateUsers,
    AccountsCreateAuditory,
    AccountsCreateCampus,
  },
  computed: {
    ...mapStores(useStatusWindowStore, useUniversityStore),
    
    getListItemComponent(){
      return UserListItem;
    },
    getAuditoryListItemComponent(){
      return AccountAuditoryListItem;
    },
    getCampusListItemComponent(){
      return AccountCampusListItem;
    },

    auditoriesSearchList(): ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.universityStore.auditoriesList.length === 0) return arr;
      for(let item of this.universityStore.auditoriesList){
        arr.push({id: item.id, search_field: item.name, data: item});
      }
      return arr;
    },
    campusesSearchList(): ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.universityStore.campusList.length === 0) return arr;
      for(let item of this.universityStore.campusList){
        arr.push({id: item.id, search_field: item.name, data: item});
      }
      return arr;
    },
    decansSearchList(): ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.universityStore.decansList.length === 0) return arr;
      for(let item of this.universityStore.decansList){
        arr.push({id: item.id, search_field: `${item.last_name} ${item.first_name} ${item.father_name}`, data: item});
      }
      return arr;
    },
    educationDepartmentsSearchList(): ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.universityStore.educationDepartmentsList.length === 0) return arr;
      for(let item of this.universityStore.educationDepartmentsList){
        arr.push({id: item.id, search_field: `${item.last_name} ${item.first_name} ${item.father_name}`, data: item});
      }
      return arr;
    },
    zavCafsSearchList(): ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.universityStore.zavCafsList.length === 0) return arr;
      for(let item of this.universityStore.zavCafsList){
        arr.push({id: item.id, search_field: `${item.last_name} ${item.first_name} ${item.father_name}`, data: item});
      }
      return arr;
    },
    teachersSearchList(): ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.universityStore.teachersList.length === 0) return arr;
      for(let item of this.universityStore.teachersList){
        arr.push({id: item.id, search_field: `${item.last_name} ${item.first_name} ${item.father_name}`, data: item});
      }
      return arr;
    },
    studentsSearchList(): ISearchList[]{
      const arr:ISearchList[] = [];
      if(this.universityStore.studentsList.length === 0) return arr;
      for(let item of this.universityStore.studentsList){
        arr.push({id: item.id, search_field: `${item.last_name} ${item.first_name} ${item.father_name}`, data: item});
      }
      return arr;
    },
  },
};
</script>