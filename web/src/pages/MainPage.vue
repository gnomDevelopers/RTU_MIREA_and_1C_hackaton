<template>
  <div class="w-full flex flex-col items-center scrollable bg-color-ultralight">
    <div class="flex flex-col h-full w-full md:w-10/12 gap-y-10 py-5 px-4 md:px-0">
      <div class="text-4xl text-center us:text-left">Здравствуйте, Генадий{{ userInfoStore.first_name }} Абрамович{{ userInfoStore.father_name }}!</div>
      <div class="flex flex-col gap-y-4 pb-4">

        <MainTableItem title="Создание учетных записей" text="Создать учётную запись для администратора, заведующего кафедрой, преподавателя или студента">
          <IconAccounts/>
        </MainTableItem>

        <MainTableItem title="Расписание" text="Посмотреть расписание для своей или любой другой группы">
          <IconShedule/>
        </MainTableItem>

        <MainTableItem title="Успеваемость" text="Посмотреть успеваемость студентов">
          <IconPerformance/>
        </MainTableItem>

        <MainTableItem title="Группы" text="Заполнение учебных групп">
          <IconGroups/>
        </MainTableItem>

        <MainTableItem title="Посещаемость" text="Просмотр, добавление, редактирование посещаемости">
          <IconAttendance/>
        </MainTableItem>

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
import MainTableItem from '@/entities/mainTableItem.vue';

export default{
  components:{
    IconShedule,
    IconAccounts,
    IconPerformance,
    IconGroups,
    IconAttendance,
    MainTableItem,
  },
  computed: {
    ...mapStores(useUserInfoStore),
  }
};
</script>