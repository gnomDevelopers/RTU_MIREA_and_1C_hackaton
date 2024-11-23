<template>
    <div class="flex flex-col items-center scrollable px-4 lg:px-0">
    <div class="flex flex-col w-full lg:w-10/12 items-center gap-y-4 mb-4">

      <PageTitle title="Создание расписания" description="В этом разделе Вы можете создать расписание с помощью автогенерации, а также заполнить его вручную или импортом excel-файла по шаблону. Для этого загрузите на своё устройство шаблон для заполнения данных, а затем отправьте заполненный файл в специальное поле.">
        <IconShedule class="w-20 us:w-36 h-20 us:h-36"/>
      </PageTitle>

      <div class="flex flex-row gap-4 w-full justify-center items-start">
        <div class="flex flex-col w-1/2 gap-y-4 p-4 rounded-xl bg-color-light">
          <div class="p-2 rounded-lg bg-white">
            <p class="text-center text-xl">Автоматическая генерация расписания</p>
          </div>
          <div>
            <p class="text-justify text-lg">Для того, чтобы расписание сгенерировалось корректно, внимательно вносите в таблицу требуемые данные об образовательном процессе. </p>
          </div>  
          <div class="flex flex-row gap-x-4 justify-center">
            <div class="flex flex-row items-center gap-x-1 py-2 px-4 rounded-lg cursor-pointer header-shadow bg-white">
              <img class="w-8 h-11" src="../assets/icons/icon-download-file.svg"/>
              <p class="text-lg">.xlsx шаблон</p>
            </div>
            <div class="flex flex-row items-center gap-x-1 py-2 px-4 rounded-lg cursor-pointer header-shadow bg-white" @click="$refs.fileInput.click()">
              <img class="w-8 h-11" src="../assets/icons/icon-upload-file.svg"/>
              <p class="text-lg">Выберите файл</p>
              <input class=" hidden invisible w-0 h-0" type="file" ref="fileInput">
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>
<script lang="ts">
import IconShedule from '@/shared/iconShedule.vue';
import PageTitle from '@/shared/pageTitle.vue';


export default {
  components: {
    PageTitle,
    IconShedule,
  }
};
</script>