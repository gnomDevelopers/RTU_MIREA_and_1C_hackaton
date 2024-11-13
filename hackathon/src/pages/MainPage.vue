<template>
  <div class="w-full flex flex-col items-center scrollable bg-color-ultralight">
    <div class="flex flex-col h-full w-full md:w-10/12 gap-y-10 py-5 px-4 md:px-0">
      <div class="text-4xl">Здравствуйте, {{ userInfoStore.first_name }} {{ userInfoStore.father_name }}!</div>
      <div class="flex flex-col gap-y-4 pb-4">
        <div class="main-table-item">
          <IconAccounts/>

          <div class="main-table-item-text">
            <p class="main-table-item-title">Создание учетных записей</p>
            <p class="main-table-item-decription">Создать учётную запись для администратора, заведующего кафедрой, преподавателя или студента</p>
            <div class="main-table-item-btn">Перейти</div>
          </div>
        </div>
        
        <div class="main-table-item">
          <IconShedule/>

          <div class="main-table-item-text">
            <p class="main-table-item-title">Расписание</p>
            <p class="main-table-item-decription">Посмотреть расписание для своей или любой другой группы</p>
            <div class="main-table-item-btn">Перейти</div>
          </div>
        </div>

        <div class="main-table-item">
          <IconPerformance/>
          
          <div class="main-table-item-text">
            <p class="main-table-item-title">Успеваемость</p>
            <p class="main-table-item-decription">Посмотреть успеваемость студентов</p>
            <div class="main-table-item-btn">Перейти</div>
          </div>
        </div>

        <div class="main-table-item">
          <IconGroups/>

          <div class="main-table-item-text">
            <p class="main-table-item-title">Группы</p>
            <p class="main-table-item-decription">Заполнение учебных групп</p>
            <div class="main-table-item-btn">Перейти</div>
          </div>
        </div>

        <div class="main-table-item">
          <IconAttendance/>

          <div class="main-table-item-text">
            <p class="main-table-item-title">Посещаемость</p>
            <p class="main-table-item-decription">Просмотр, добавление, редактирование посещаемости</p>
            <div class="main-table-item-btn">Перейти</div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { mapStores } from 'pinia';
import { useUserInfoStore } from '@/stores/userInfoStore';

import IconShedule from '@/shared/iconShedule.vue';
import IconAccounts from '@/shared/iconAccounts.vue';
import IconPerformance from '@/shared/iconPerformance.vue';
import IconGroups from '@/shared/iconGroups.vue';
import IconAttendance from '@/shared/iconAttendance.vue';

export default{
  components:{
    IconShedule,
    IconAccounts,
    IconPerformance,
    IconGroups,
    IconAttendance
  },
  computed: {
    ...mapStores(useUserInfoStore),
  }
};
</script>