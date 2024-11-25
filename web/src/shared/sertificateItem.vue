<template>
  <div class="flex flex-col flex-shrink flex-grow basis-full mb:basis-[440px] gap-y-4 p-4 w-full mb:w-[440px] max-w-[500px] rounded-lg bg-color-light">
    <div class="flex flex-row gap-x-2 items-start h-1/2">
      <div class="flex-shrink-0 w-8/12 aspect-video">
        <slot></slot>
      </div>
      <p class="text-xl text-center font-semibold">{{ title }}</p>
    </div>
    <div>
      <p class="text-xl">Результат: {{ place }} место</p>
    </div>
    <div class="flex flex-row gap-x-2 items-center self-start rounded-xl py-1 px-2 cursor-pointer border-2 border-blue-800 border-solid">
      <p class="text-xl font-semibold">Загрузить сертификат</p>
      <img class="w-8 h-11" src="../assets/icons/icon-download-file.svg"/>
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
    place: {
      type: [Number, String],
      required: true,
    },
  }
};
</script>