<template>
  <div class="flex flex-col">
    <div class="self-start">
      <p class="text-xl cursor-pointer select-none hover:underline" @click="() => showSlot = !showSlot">{{ title }}</p>
    </div>
    <div v-if="showSlot" class="flex flex-col us:ml-2 mt-2">
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