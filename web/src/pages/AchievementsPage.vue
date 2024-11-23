<template>
  <div class="flex flex-col items-center scrollable px-4 xl:px-0">
    <div class="flex flex-col w-full xl:w-10/12 items-center gap-y-4 mb-4">
      <p class="text-3xl mt-4 font-semibold">Мои достижения</p>
      <div class="flex flex-row gap-4 flex-wrap justify-start">

        <SertificateItem v-for="i in 10" title="Студенческий хакатон “РТУ МИРЭА и 1С”" :place="-1"/>


      </div>
    </div>
  </div>
</template>
<script lang="ts">
import SertificateItem from '@/shared/sertificateItem.vue';
export default {
  components:{
    SertificateItem,
  }
};
</script>