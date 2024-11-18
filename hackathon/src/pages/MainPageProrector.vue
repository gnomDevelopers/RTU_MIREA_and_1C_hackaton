<template>
  <div class="w-full h-full scrollable flex flex-col items-center">
    <div class="w-full px-2 us:px-4 lg:w-10/12 lg:px-0">
      <p class="text-3xl my-5 text-center lg:text-left">Добро пожаловать, Проректор Проректорович</p>
      <div class="flex flex-col gap-y-3">

        <MainControl title="Управление вузом">
          <MainControlItem title="Управление деканами">

            <UserList 
              title="Введите ФИО декана" 
              placeholder="ФИО декана"
              :usersList="decanList" 
              :itemComponent="getListItemComponent"
            />

          </MainControlItem>
          <MainControlItem title="Управление преподавателями">
            
            <UserList 
              title="Введите ФИО преподавателя" 
              placeholder="ФИО преподавателя"
              :usersList="prepodList" 
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

        <MainControl title="Мониторинг посещаемости">
          <MainControlItem title="Просмотр отчетов о посещаемости вуза">
            скрытая часть
          </MainControlItem>
        </MainControl>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { ROLES_SET_PRORECTOR, ROLES_NAME } from '../helpers/constants';
import MainControl from '../shared/mainControl.vue';
import MainControlItem from '../shared/mainControlItem.vue';
import UserListItem from '../shared/userListItem.vue';
import UserList from '../entities/userList.vue';

import { API_SendFile } from '../api/api';

export default {
  components:{
    MainControl,
    MainControlItem,
    UserListItem,
    UserList,
  },
  data(){
    return{
      rolesList: ROLES_SET_PRORECTOR,
      decanList: [
        {id: 1, name: "Деканов1 Декан1 Деканович1", role: 2}, 
        {id: 2, name: "Деканов2 Декан2 Деканович2", role: 2}, 
        {id: 3, name: "Деканов3 Декан3 Деканович3", role: 2}
      ],
      prepodList: [
        {id: 1, name: "Преподов1 Препод1 Преподович1", role: 4}, 
        {id: 2, name: "Преподов2 Препод2 Преподович2", role: 4}, 
        {id: 3, name: "Преподов3 Препод3 Преподович3", role: 4}
      ],
    }
  },
  mounted(){
    for(let i = 0; i < 100; i++) this.decanList.push({id: i+4, name: `Деканов${i+4} Декан${i+4} Деканович${i+4}`, role: 2});
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