<template>
  <div class="scrollable flex h-screen uus:flex-col md:flex-row">
    <div class="md:w-1/2 uus:w-full h-auto md:m-10 uus:mt-5">
      <p style="margin-left: 30px; color: #8F0101;" class="md:text-3xl uus:text-xl">Анкета кандидата</p>
      <div class="w-full h-auto border-4 rounded-xl border-hr-color">
        <div class="md:mt-4 md:ml-10 md:mr-10 md:mb-4 uus:mt-1 uus:ml-3 uus:mr-3 uus:mb-1">
          <p style="font-weight: bold" class="md:text-2xl uus:text-xl">Специальность</p>
          <div>
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px" >{{ formData.specialty }}</p>
          </div>
          <p style="font-weight: bold" class="md:text-2xl uus:text-xl">Опыт работы</p>
          <div >
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">{{ formData.experience }}</p>
          </div>
          <p style="font-weight: bold" class="md:text-2xl uus:text-xl" >Дополнительный опыт</p>
          <div>
            <p class="break-words md:text-2xl uus:text-xl" style=" margin-bottom: 10px">{{ formData.add_experience }}</p>
          </div>
          <p style="font-weight: bold; " class="md:text-2xl uus:text-xl">Номер телефона</p>
          <div>
            <p class="break-words md:text-2xl uus:text-xl" style=" margin-bottom: 10px">{{ formData.phone }}</p>
          </div>
          <p style="font-weight: bold" class="md:text-2xl uus:text-xl">Полезные ссылки</p>
          <div>
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">{{ formData.links }}</p>
          </div>
          <p style="font-weight: bold" class="md:text-2xl uus:text-xl">Резюме</p>
          <div>
            <div
                @click="$refs.downloadLink.click()"
                class="flex flex-row md:justify-normal uus:justify-center gap-x-1 py-2 px-4 rounded-lg cursor-pointer header-shadow bg-white">


              <p class="text-lg shadow flex flex-row items-center gap-1 py-2 px-4 rounded-lg cursor-pointer header-shadow border"><img class="w-8 h-11" src="../assets/icons/icon-download-file.svg"/>Скачать резюме</p>
              <a class="hidden invisible w-0 h-0" ref="downloadLink" :href="getDownloadLink" download></a>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="md:w-1/2 uus:w-full h-auto md:m-10 uus:mt-5">
      <p style="margin-left: 30px; color: #8F0101;" class="md:text-3xl uus:text-xl">Учёба студента (данные с VUZ+)</p> <!--Надо брать данные с вуз+ хз-->
      <div class="w-full h-auto border-4 rounded-xl border-hr-color">
        <div class="md:mt-4 md:ml-10 md:mr-10 md:mb-4 uus:mt-1 uus:ml-3 uus:mr-3 uus:mb-1">
          <p style="font-weight: bold" class="md:text-2xl uus:text-xl">Образование</p>
          <div >
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">РТУ МИРЭА, Фуллстек-разработка (бакалавриат), 2 курс</p>
          </div>
          <p style="font-weight: bold" class="md:text-2xl uus:text-xl">Рейтинг студента (GPA - пятибалльная шкала)</p>
          <div >
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">4.6</p>
          </div>
          <p style="font-weight: bold" class="md:text-2xl uus:text-xl">Посещаемость</p>
          <div>
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">87%</p>
          </div>
          <p style="font-weight: bold; " class="md:text-2xl uus:text-xl">Успеваемость</p>
          <div>
            <p class="break-words md:text-2xl uus:text-xl" style="margin-bottom: 10px">4.7</p>
          </div>
        </div>
      </div>
      <div class="flex justify-center">
        <button v-if="!formData.isClosed" @click="hideProfile" class= "mt-16 w-1/2 cursor-pointer transition-colors py-2 px-5 text-lg rounded-xl font-semibold btn  text-slate-100 mb-8">
          Скрыть профиль
        </button>
        <button v-if="formData.isClosed" @click="hideProfile" class= "mt-16 w-1/2 cursor-pointer transition-colors py-2 px-5 text-lg rounded-xl font-semibold bg-gray-300 text-black mb-8">
          Открыть профиль
        </button>
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
