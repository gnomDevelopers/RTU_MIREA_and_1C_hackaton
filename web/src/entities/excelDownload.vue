<template>
  <div 
    class=" btn text-white cursor-pointer"
    @click="downloadExcel"
    >
      Скачать данные
  </div>
</template>
<script lang="ts">
import xlsx from 'json-as-xlsx';
import type { PropType } from 'vue';

export default {
  props: {
    data: {
      type: Array as PropType<{
        last_name: string,
        first_name: string,
        father_name: string,
        email: string,
        password: string
      }[]>,
      required: true,
    }
  },
  methods: {
    downloadExcel(){
      const data = [
        {
          sheet: 'Sheet1',
          columns: [
            { label: 'Фамилия', value: 'last_name'},
            { label: 'Имя', value: 'first_name'},
            { label: 'Отчество', value: 'father_name'},
            { label: 'Email', value: 'email'},
            { label: 'Пароль', value: 'password'},
          ],
          content: this.data,
        },
      ];

      xlsx(data, {
      fileName: "Данные_пользователей",
      extraLength: 5,
      writeMode: "writeFile", 
      writeOptions: {}, 
      RTL: false, 
    })
    }
  }
};
</script>