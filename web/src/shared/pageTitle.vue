<template>
  <div class="flex flex-col us:flex-row gap-x-2 items-center us:items-stretch mt-5">
    <slot><!--logo--></slot> 
    <div class="flex flex-col gap-y-4">
      <p class="text-4xl font-medium text-center us:text-left">Создание учётных записей</p>
      <p class="text-xl text-justify">В этом разделе Вы можете создать учётную запись для администратора, заведующего кафедрой, преподавателя или студента. Для этого загрузите на своё устройство шаблон для заполнения данных о пользователе, а затем отправьте заполненный файл в специальное поле.</p>
    </div>
  </div>
</template>
<script lang="ts">
export default {
  props:{
    title:{
      type: String,
      required: true,
    },
    description:{
      type: String,
      required: true,
    },
  }
};
</script>