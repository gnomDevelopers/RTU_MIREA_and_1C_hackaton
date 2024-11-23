<template>
  <div class="flex flex-row items-stretch min-h-14 cursor-pointer border-2 border-solid rounded-xl overflow-hidden transition-colors border-transparent hover:border-blue-900">
    <div class="flex flex-col justify-center items-center w-10  bg-color-bold">
      <p class="text-xl text-color-light">{{ id }}</p>
    </div>
    <div class="flex flex-col justify-center px-2 bg-white ">
      <p class="text-lg">{{ name }}</p>
    </div>
  </div>
</template>
<script lang="ts">
export default{
  props: {
    id: {
      type: [Number, String],
      reuired: true,
    },
    name: {
      type: String,
      required: true,
    }
  }
};
</script>