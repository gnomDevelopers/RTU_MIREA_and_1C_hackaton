<template>
  <div class="flex flex-col flex-shrink flex-grow basis-full mb:basis-[440px] gap-y-4 p-4 w-full mb:w-[440px] max-w-[500px] rounded-lg bg-color-light">
    <div class="flex flex-row gap-x-2 items-start">
      <img class="flex-grow w-1/2 aspect-video" src="../assets/images/1ass.webp"/>
      <p class="text-xl text-center font-semibold">{{ title }}</p>
    </div>
    <div>
      <p class="text-xl">Результат: {{ place }} место</p>
    </div>
    <div class="flex flex-row gap-x-2 items-center self-start rounded-xl py-1 px-2 cursor-pointer border-2 border-blue-800 border-solid">
      <p class="text-xl font-semibold">Загрузить сертификат</p>
      <svg class="w-8 h-11" viewBox="0 0 37 41" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M16.249 38.4996H7.24901C4.76372 38.4996 2.74901 36.4849 2.74902 33.9996L2.7492 6.99973C2.74921 4.51446 4.76393 2.49976 7.2492 2.49976H27.4997C29.985 2.49976 31.9997 4.51447 31.9997 6.99975V19.3748M34.2497 33.9199L29.648 38.4997M29.648 38.4997L25.2497 34.1269M29.648 38.4997V27.2498" stroke="#063C73" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
      </svg>
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