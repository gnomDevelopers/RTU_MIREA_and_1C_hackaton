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
        <MainControlItem title="Управление аудиториями">

          <SearchList 
            title="Введите название аудитории" 
            placeholder="Название аудитории"
            :searchList="auditoriesSearchList"
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
      // списки ресурсов
      auditoriesSearchList: [] as ISearchList[],

      //списки пользователей
      decansSearchList: [] as ISearchList[],
      educationDepartmentsSearchList: [] as ISearchList[],
      zavCafsSearchList: [] as ISearchList[],
      teachersSearchList: [] as ISearchList[],
      studentsSearchList: [] as ISearchList[],
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore, useUniversityStore),
    
    getListItemComponent(){
      return UserListItem;
    },
    getAuditoryListItemComponent(){
      return AccountAuditoryListItem;
    },
  },
  watch: {
    'universityStore.auditoriesList' : {
      handler(val: IAPI_Audience_Update[]){
        this.auditoriesSearchList = [];
        for(let item of val){
          this.auditoriesSearchList.push({id: item.id, search_field: item.name, data: item});
        }
      },
      immediate: true,
      deep: true,
    },

    'universityStore.decansList' : {
      handler(val: IUserGet[]){
        this.decansSearchList = [];
        for(let item of val){
          this.decansSearchList.push({id: item.id, search_field: `${item.surname} ${item.name} ${item.thirdname}`, data: item});
        }
      },
      immediate: true,
      deep: true,
    },

    'universityStore.educationDepartmentsList' : {
      handler(val: IUserGet[]){
        this.educationDepartmentsSearchList = [];
        for(let item of val){
          this.educationDepartmentsSearchList.push({id: item.id, search_field: `${item.surname} ${item.name} ${item.thirdname}`, data: item});
        }
      },
      immediate: true,
      deep: true,
    },

    'universityStore.zavCafsList' : {
      handler(val: IUserGet[]){
        this.zavCafsSearchList = [];
        for(let item of val){
          this.zavCafsSearchList.push({id: item.id, search_field: `${item.surname} ${item.name} ${item.thirdname}`, data: item});
        }
      },
      immediate: true,
      deep: true,
    },

    'universityStore.teachersList' : {
      handler(val: IUserGet[]){
        this.teachersSearchList = [];
        for(let item of val){
          this.teachersSearchList.push({id: item.id, search_field: `${item.surname} ${item.name} ${item.thirdname}`, data: item});
        }
      },
      immediate: true,
      deep: true,
    },

    'universityStore.studentsList' : {
      handler(val: IUserGet[]){
        this.studentsSearchList = [];
        for(let item of val){
          this.studentsSearchList.push({id: item.id, search_field: `${item.surname} ${item.name} ${item.thirdname}`, data: item});
        }
      },
      immediate: true,
      deep: true,
    },
  }
};
</script>