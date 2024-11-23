<template>
  <div class="w-full flex flex-col items-center scrollable bg-color-ultralight">
    <div class="flex flex-col h-full w-full md:w-10/12 gap-y-10 py-5 px-4 md:px-0">
      <div class="text-4xl text-center us:text-left">Здравствуйте, Генадий{{ userInfoStore.first_name }} Абрамович{{ userInfoStore.father_name }}!</div>
      <div class="flex flex-col gap-y-4 pb-4">

        <MainTableItem @clicked="$router.push({name: 'AccountsPage'})" title="Создание учетных записей" text="Создать учётную запись для администратора, заведующего кафедрой, преподавателя или студента">
          <IconAccounts/>
        </MainTableItem>

        <MainTableItem @clicked="$router.push({name: 'CreateSchedulePage'})" title="Создание расписания" text="Создать расписание на семестр">
          <IconShedule/>
        </MainTableItem>

        <MainTableItem @clicked="$router.push({name: 'SchedulePage'})" title="Расписание" text="Посмотреть расписание для своей или любой другой группы">
          <IconShedule/>
        </MainTableItem>

        <MainTableItem @clicked="$router.push({name: 'PerformancePage'})" title="Успеваемость" text="Посмотреть успеваемость студентов">
          <IconPerformance/>
        </MainTableItem>

        <MainTableItem @clicked="$router.push({name: 'GroupCorrectPage'})" title="Группы" text="Заполнение учебных групп">
          <IconGroups/>
        </MainTableItem>

        <MainTableItem @click="$router.push({name: 'AttendancePage'})" title="Посещаемость" text="Просмотр, добавление, редактирование посещаемости">
          <IconAttendance/>
        </MainTableItem>

        <MainTableItem title="WorkVUZ+" text="Удобный сервис для интеграции HR-специалистов и студентов, который позволяет искать и откликаться на вакансии">
          <IconAttendance/>
        </MainTableItem>

        <MainTableItem title="EventVUZ+" text="Удобный сервис для организации и просмотра доступных внутривузовских мероприятий">
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