<template>
  <div class="flex flex-col gap-x-5 items-center p-5 us:flex-row us:items-stretch border-2 border-color rounded-xl bg-white">
  <slot></slot>

  <div class="flex flex-col justify-between gap-y-2 items-center cursor-default us:items-start us:gap-y-0">
    <p class="text-3xl text-center us:text-left">{{ title }}</p>
    <p class="text-lg text-justify">{{ text }}</p>
    <div @click="$emit('clicked')" class="us:self-start py-2 px-4 rounded-lg text-lg cursor-pointer bg-color-bold text-color-light transition-colors main-table-item-btn">Перейти</div>
  </div>
</div>
</template>
<script lang="ts">

export default{
  emits: ['clicked'],
  props:{
    title: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
  }
};

</script>