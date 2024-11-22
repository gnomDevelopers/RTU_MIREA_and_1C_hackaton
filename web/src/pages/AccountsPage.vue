<template>
  <div class="w-full h-full scrollable flex flex-col items-center">
    <div class="flex flex-col gap-y-4 w-full px-2 us:px-4 lg:w-10/12 lg:px-0">
      
      <div class="flex flex-row gap-x-2 items-stretch mt-5">
        <IconAccounts class="w-20 us:w-36 h-20 us:h-36"/>
        <div class="flex flex-col gap-y-4">
          <p class="text-4xl font-medium">Создание учётных записей</p>
          <p class="text-xl">В этом разделе Вы можете создать учётную запись для администратора, заведующего кафедрой, преподавателя или студента. Для этого загрузите на своё устройство шаблон для заполнения данных о пользователе, а затем отправьте заполненный файл в специальное поле.</p>
        </div>
      </div>

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
          
          <label for="fileInput">
            Choose file
            <input 
              type="file" 
              multiple 
              name="fileInput" 
              accept="image/*, .png, .svg, .jpg, .webp" 
              @change="showFiles" 
            />
          </label>

        </MainControlItem>
      </MainControl>

    </div>
  </div>
</template>
<script lang="ts">
import { API_SendFile } from '../api/api';
import { ROLES_SET_PRORECTOR, ROLES_NAME, type ISearchList, type IUsersList } from '../helpers/constants';

import MainControl from '../shared/mainControl.vue';
import MainControlItem from '../shared/mainControlItem.vue';
import UserListItem from '../shared/userListItem.vue';
import SearchList from '../entities/searchList.vue';
import IconAccounts from '@/shared/iconAccounts.vue';

export default {
  components:{
    IconAccounts,
    MainControl,
    MainControlItem,
    UserListItem,
    SearchList,
  },
  data(){
    return{
      rolesList: ROLES_SET_PRORECTOR,
      decanList: [] as ISearchList[],
      prepodList: [] as ISearchList[],
    }
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
  computed: {
    getListItemComponent(){
      return UserListItem;
    }
  },
  methods:{
    async showFiles(event: any){

      const formData = new FormData();
      const files = event.target.files;

      for (const file of files) formData.append('files', file, file.name);
      
      API_SendFile(formData)
        .then(res => {
          console.log('success response');
        })
        .catch(err => {
          console.log('error response');
        });
    }
  }
};


</script>