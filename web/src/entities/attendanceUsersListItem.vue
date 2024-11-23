<template>
  <div class="flex flex-row items-stretch flex-nowrap min-h-8 rounded-lg bg-white">
    <div class="flex flex-row justify-center items-center w-14 rounded-l-lg bg-color-bold">
      <p class="text-xl text-white">{{ index }}</p>
    </div>
    <div class="flex flex-row items-center px-2 flex-grow">
      <p class="text-xl">{{ data.name }}</p>
    </div>
    <div class="flex flex-row gap-x-1 px-2 py-1">
      <div class="w-6 flex flex-row justify-center items-center rounded cursor-pointer bg-color-bold">
        <img class="w-5 h-5" src="../assets/icons/icon-plus.svg"/>
      </div>
      <div class="w-6 flex flex-row justify-center items-center rounded cursor-pointer bg-color-bold">
        <p class="text-xl text-white">Н</p>
      </div>
      <div class="w-6 flex flex-row justify-center items-center rounded cursor-pointer bg-color-bold">
        <p class="text-xl text-white">У</p>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
export default {
  props: {
    index: {
      type: Number,
      required: true,
    },
    data: {
      type: Object,
      required: true,
    }
  }
};
</script>