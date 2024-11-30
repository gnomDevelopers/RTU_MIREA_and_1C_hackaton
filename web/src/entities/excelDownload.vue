<template>
  <div 
    v-if="data.length !== 0"
    class=" btn text-white cursor-pointer"
    @click="downloadExcel"
    >
      Скачать данные
  </div>
</template>
<script lang="js">
import { export_json_to_excel } from '../helpers/Export2Excel';

export default {
  props: {
    fileds: {
      type: Object,
      required: true,
    },
    data: {
      type: Array,
      required: true,
    }
  },
  methods: {
    downloadExcel(){
      export_json_to_excel({
        header: this.fields,
        data: this.data,
        sheetName: "Scheet 1",
        filename: "Данные",
        autoWidth: true,
        bookType: "xlsx",
      });
    }
  }
};
</script>