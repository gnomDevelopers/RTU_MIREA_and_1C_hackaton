<template>
    <div class="flex flex-col items-center scrollable px-4 lg:px-0">
    <div class="flex flex-col w-full lg:w-10/12 items-center gap-y-4 mb-4">

      <PageTitle title="Создание расписания" description="В этом разделе Вы можете создать расписание с помощью автогенерации, а также заполнить его вручную или импортом excel-файла по шаблону. Для этого загрузите на своё устройство шаблон для заполнения данных, а затем отправьте заполненный файл в специальное поле.">
        <IconShedule class="w-20 us:w-36 h-20 us:h-36"/>
      </PageTitle>

      
    </div>
  </div>
</template>
<script lang="ts">
import IconShedule from '@/shared/iconShedule.vue';
import PageTitle from '@/shared/pageTitle.vue';


export default {
  components: {
    PageTitle,
    IconShedule,
  }
};
</script>