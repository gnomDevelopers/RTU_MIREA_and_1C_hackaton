<template>
  <div class="scrollable flex h-screen uus:flex-col md:flex-row">
    <div class="flex flex-col justify-center md:w-1/2 uus:w-full h-auto md:m-10 uus:mt-5">
      <div class="flex flex-col justify-center">
        <p class="flex text-3xl m-8 justify-center">Личный кабинет</p>
        <div class="flex justify-center">
          <img class="w-48 h-48"  src="../assets/icons/icon-profile.svg">
        </div>
        <p class="flex justify-center text-2xl mt-4">Постников Даниил Сергеевич</p>
        <p class="flex justify-center text-lg mt-4 text-gray-400">postnikov_d_s@vuzplus.ru</p>
        <p class="flex justify-center text-xl mt-4">ID участника: 7</p>
      </div>
    </div>
    <div class="md:w-1/2 uus:w-full h-auto md:m-10 uus:mt-5">
      <div class="w-full h-auto border-4 rounded-xl border-ev-color">
        <div class="md:mt-4 md:ml-10 md:mr-10 md:mb-4 uus:mt-1 uus:ml-3 uus:mr-3 uus:mb-1">
          <p style="font-weight: bold" class="md:text-2xl uus:text-xl">Специальность</p>
          <div >
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">Backend-разработчик</p>
          </div>
          <p style="font-weight: bold" class="md:text-2xl uus:text-xl">Ключевые навыки</p>
          <div >
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">Backend-разработка на Go, проектирование REST API, PostgreSQL</p>
          </div>
          <p style="font-weight: bold; " class="md:text-2xl uus:text-xl">Опыт работы</p>
          <div>
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">В течение полугода проходил стажировку по специальности “Backend-разработка на Go” в крупной компании и занимался проектированием REST API</p>
          </div>
          <p style="font-weight: bold; " class="md:text-2xl uus:text-xl">Дополнительный опыт</p>
          <div>
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">Участвовал в хакатоне “РТУ МИРЭА и 1С” и создавал решение кейса ”Система управления образовательным процессом в вузе” в роли Backend-разработчика. Наша команда заняла 1 место.</p>
          </div>
          <p style="font-weight: bold; " class="md:text-2xl uus:text-xl">Контакты</p>
          <div>
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">+7(915)-921-00-00</p>
          </div>
          <p style="font-weight: bold; " class="md:text-2xl uus:text-xl">Полезные ссылки</p>
          <div>
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">https://github.com/gnomDeployers</p>
          </div>
        </div>
      </div>


    </div>
  </div>

</template>

<script lang="ts">
import { useUserFormStore } from '@/stores/userFormStore';

export default {
  setup() {
    const userFormStore = useUserFormStore();
    return {
      formData: userFormStore.formData,
    };
  },
  computed: {
    getDownloadLink() {
      if (this.formData.filesList.length > 0) {
        return URL.createObjectURL(this.formData.filesList[0]);
      }
      return '';
    },
  },
  methods: {
    hideProfile() {
      if (!this.formData.isClosed) {
        this.formData.isClosed = true;
      }
      else this.formData.isClosed = false;
    }
  }
};
</script>
