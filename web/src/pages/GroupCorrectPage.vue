<template>
  <div class="flex flex-col items-center scrollable px-4 lg:px-0">
    <div class="flex flex-col w-full lg:w-10/12 items-center gap-y-4 mb-4">
      
      <div class="flex flex-row gap-x-2 items-stretch mt-5">
        <IconGroups class="w-20 us:w-36 h-20 us:h-36"/>
        <div class="flex flex-col gap-y-4">
          <p class="text-4xl font-medium">Создание учебных групп</p>
          <p class="text-xl">В этом разделе Вы можете создать новые или отредактировать уже существующие учебные группы. Для этого найдите в поиске нужную группу и выберите действие, которое хотите совершить.</p>
        </div>
      </div>

    </div>
  </div>
</template>
<script lang="ts">
import IconGroups from '@/shared/iconGroups.vue';

export default{
  components:{
    IconGroups,
  },
};
</script>