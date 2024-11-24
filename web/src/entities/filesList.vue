<template>
  <div v-if="fileList.length !== 0" class="flex flex-col gap-y-1 p-2 cursor-default">
    <div v-for="file in fileList" class="flex flex-row items-center py-1 px-2 rounded-lg bg-slate-200">
      <img class="w-8 h-11" alt="Файл" src="../assets/icons/icon-excel.svg"/>
      <div class="flex flex-col flex-grow min-w-40 px-2">
        <p class="text-lg">{{ file.name }}</p>
        <p class=" text-sm text-gray-500">Size: {{ compileFileSize(file.size) }}</p>
      </div>
      <div @click="deleteFile(file)" class="w-9 h-9 flex flex-row justify-center items-center rounded-lg cursor-pointer bg-red-400">
        <img class="w-6 h-8" src="../assets/icons/icon-delete.svg"/>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
export default {
  emits: ['deleteFile'],
  props: {
    fileList: {
      type: Array<File>,
      required: true,
    }
  },
  methods: {
    compileFileSize(bytes: number):string {
      if (bytes < 1024) {
        return `${bytes} B`;
      } else if (bytes < 1024 ** 2) {
        return `${(bytes / 1024).toFixed(1)} KB`;
      } else if (bytes < 1024 ** 3) {
        return `${(bytes / (1024 ** 2)).toFixed(1)} MB`;
      } else if (bytes < 1024 ** 4) {
        return `${(bytes / (1024 ** 3)).toFixed(1)} GB`;
      } else {
        return `${(bytes / (1024 ** 4)).toFixed(1)} TB`;
      }
    },
    deleteFile(file: File) {
      this.$emit('deleteFile', file);
      console.log('delete emit');
    }
  }
};
</script>