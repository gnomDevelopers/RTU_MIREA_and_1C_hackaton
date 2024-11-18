<template>
  <div class="flex flex-col">
    <p class="text-xl cursor-pointer select-none hover:underline" @click="() => showSlot = !showSlot">{{ title }}</p>
    <div v-if="showSlot" class="flex flex-col us:ml-2">
      <slot></slot>
    </div>
  </div>
</template>
<script lang="ts">
export default{
  props: {
    title: {
      type: String,
      required: true,
    },
  },
  data(){
    return{
      showSlot: false,
    }
  }
};
</script>