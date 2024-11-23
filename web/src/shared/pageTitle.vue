<template>
  <div class="flex flex-col us:flex-row gap-x-2 items-center us:items-stretch mt-5">
    <slot><!--logo--></slot> 
    <div class="flex flex-col gap-y-4">
      <p class="text-4xl font-medium text-center us:text-left">{{ title }}</p>
      <p class="text-xl text-justify">{{ description }}</p>
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