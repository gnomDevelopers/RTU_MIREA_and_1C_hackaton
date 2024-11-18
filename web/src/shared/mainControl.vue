<template>
  <div class="flex flex-col gap-y-3 p-2 us:p-3 rounded-lg bg-color-light">
    <p class="text-2xl text-gray-900">{{ title }}</p>
    <div class="flex flex-col gap-y-2 ml-0 us:ml-2">
      <slot></slot>
    </div>
  </div>
</template>
<script lang="ts">
export default{
  props:{
    title: {
      type: String,
      required: true,
    },
  }
};
</script>