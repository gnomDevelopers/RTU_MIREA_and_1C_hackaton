<template>
  <div class="flex h-screen">
    <div class="flex flex-col md:w-1/2 uus:w-full h-full items-start justify-start p-4">
      <p class="logo-hr">WorkVUZ+</p>
      <div class="flex flex-col w-full h-full items-center justify-center">
        <div class="flex flex-col w-10/12 h-full items-center m-10">
          <!-- Переключатель Студент/HR -->
          <div class="toggle-container">
            <div class="toggle">
              <button
                  :class="{ active: userType === 'student' }"
                  @click="userType = 'student'"
                  class="toggle-btn"
              >
                Студент
              </button>
              <button
                  :class="{ active: userType === 'hr' }"
                  @click="userType = 'hr'"
                  class="toggle-btn"
              >
                HR
              </button>
            </div>
          </div>

          <!-- Прогресс-бар -->
          <div v-if="userType === 'student'" class="progress-bar flex justify-center mb-4">
            <div
                v-for="step in 3"
                :key="step"
                :class="{ active: step === currentStep }"
                class="progress-step"
            ></div>
          </div>

          <!-- Форма для студента -->
          <div v-if="userType === 'student'" class="w-full h-full">
            <p class="auth-description">Войдите с помощью учётной записи VUZ+</p>
            <div v-if="currentStep === 1">
              <form @submit.prevent="nextStep">
                <div class="mb-4 h-full">

                  <loginInput type="text"  text="Логин:" @input-change="checkLogin"/>
                  <loginInput type="password" id="student-password" text="Ваш пароль:" @input-change="checkPassword"/>


                </div>
              </form>
              <submitButton value="Войти" class="btn" @click="sendLogin"/>
            </div>

            <!-- Второй шаг -->
            <div v-else-if="currentStep === 2">
              <form @submit.prevent="nextStep">
                <div class="mb-4">
                  <label for="specialty" class="block text-gray-700 text-lg font-bold mb-2">Специальность:</label>
                  <input
                      type="text"
                      id="specialty"
                      required
                      v-model="formData.specialty"
                      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      placeholder="Разработчик"
                  />
                </div>
                <div class="mb-4">
                  <label for="phone" class="block text-gray-700 text-lg font-bold mb-2">Номер телефона:</label>
                  <input
                      type="tel"
                      pattern="8-[0-9]{3}-[0-9]{3}-[0-9]{2}-[0-9]{2}"
                      required
                      id="phone"
                      v-model="formData.phone"
                      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      placeholder="8-999-999-99-99"
                  />
                </div>
                <div class="mb-4">
                  <label for="experience" class="block text-gray-700 text-lg font-bold mb-2">Опыт работы:</label>
                  <textarea
                      type="text"
                      id="experience"
                      required
                      v-model="formData.experience"
                      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      placeholder="Расскажите о Вашем опыте работы"
                  />
                </div>
              </form>
            </div>

            <!-- Третий шаг -->
            <div v-else-if="currentStep === 3">
              <form @submit.prevent="completeForm">
                <div class="mb-4">
                  <label for="additional-experience" class="block text-gray-700 text-lg font-bold mb-2">Дополнительный опыт:</label>
                  <textarea
                      type="text"
                      id="additional-experience"
                      required
                      v-model="formData.add_experience"
                      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      placeholder="Расскажите об опыте участия в различных мероприятиях или о своих личных проектах"
                  />
                </div>
                <div class="mb-4">
                  <label for="links" class="block text-gray-700 text-lg font-bold mb-2">Полезные ссылки:</label>
                  <textarea
                      type="text"
                      id="links"
                      v-model="formData.links"
                      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      placeholder="Добавьте ссылки, которые могут быть полезны"
                  />
                </div>

                <div class="mb-4">
                  <label for="resume" class="block text-gray-700 text-lg font-bold mb-2">Резюме:</label>
                  <div>
                    <div class="shadow flex flex-col items-center gap-1 py-2 px-4 rounded-lg cursor-pointer header-shadow border" >
                      <div
                          v-if="filesList.length === 0"
                          @click="$refs.fileInput.click()"
                          class="flex flex-row items-center gap-x-1">

                        <img class="w-8 h-11" src="../assets/icons/icon-upload-file.svg"/>
                        <p class="text-lg">Выберите .pdf файл</p>
                        <input class=" hidden invisible w-0 h-0" type="file" ref="fileInput" accept=".pdf" @change="handleFilesChange">
                      </div>

                      <FilesList :file-list="getFilesList" @delete-file="handleDeleteFile"/>
                    </div>
                  </div>
                </div>
                <!-- Чекбокс согласия на обработку персональных данных -->
                <div class="m-8 flex items-center justify-center">
                  <input
                      type="checkbox"
                      id="agreement"
                      v-model="formData.agreed"
                      class="mr-2"
                      style="accent-color: #8F0101; transform:scale(1.5);opacity:0.9; cursor:pointer;"
                  />
                  <label for="agreement" class="text-gray-700 text-lg font-bold">Согласен на обработку персональных данных</label>
                </div>
              </form>
            </div>


<!--            Кнопка продолжения (далее на первом шаге надо убрать, когда можно будет авторизацию проверять)-->
            <div class=" left-0 w-full px-4">
              <button v-if="currentStep < 3" @click="nextStep" class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-xl font-semibold btn w-full text-slate-100 ">
                Продолжить
              </button>
              <button v-if="currentStep == 3" @click="previousStep" class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-xl font-semibold btn w-full text-slate-100 mb-6">
                Назад
              </button>
              <button 
                v-if="currentStep == 3 && formData.agreed" 
                @click="completeForm" 
                class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-xl font-semibold btn w-full text-slate-100" 
                :class="{'btn-badd' : !formData.agreed}">
                
                  Завершить заполнение
              </button>
            </div>
          </div>
          <!--Форма для HR-->
          <div v-if="userType === 'hr'" class="w-full h-full">
            <div class="flex flex-row items-stretch justify-center gap-2">
              <p class="auth-description">Войдите с помощью аккаунта hh.ru</p>
              <img class="w-8 h-8" src="../assets/icons/icon-hhru.svg">
            </div>


            <div class="mb-4">
              <loginInput type="text" text="Логин:" @input-change="checkLogin"/>
              <loginInput type="password" text="Пароль:" @input-change="checkPassword"/>
            </div>

            <submitButton value="Войти" class="btn" @click="sendLoginHR"/>
            <button @click="nextPage" class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-xl font-semibold btn w-full text-slate-100 ">
              Продолжить
            </button>
          </div>
        </div>
      </div>
    </div>
    <div class="login-bg w-1/2 uus:hidden md:block"></div>
  </div>
</template>

<script lang="ts">

import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import loginInput from '../shared/loginInput.vue';
import submitButton from '../shared/submitButton.vue';
import { ValidUserLogin, ValidUserPassword } from '../helpers/validator';
import { type IValidAnswer, StatusCodes, type IAPI_Login_Request } from '../helpers/constants';
import {API_Login, API_LoginHR} from '@/api/api';
import FilesList from "@/entities/filesList.vue";
import { useUserFormStore } from '@/stores/userFormStore';
import { defineStore } from 'pinia';
import CandidatesSearchPage from "@/pages/CandidatesSearchPage.vue";

export default {

  components:{
    FilesList,
    loginInput,
    submitButton,
  },

  data(){
    return{
      userType: 'student', // Текущий выбранный тип пользователя
      currentStep: 1, // Шаг формы
      login: {value: '', error: ''} as IValidAnswer,
      password: {value: '', error: ''} as IValidAnswer,
      showPassword: false, // Для отображения пароля
      formData: {
        specialty: '',
        phone: '',
        experience: '',
        add_experience: '',
        links: '',
        filesList: [] as File[],
        agreed: false,
        isClosed: false,
      },
      filesList: [] as File[]
    }
  },
  computed:{
    ...mapStores(useStatusWindowStore, useUserInfoStore, useUserFormStore),

    CandidatesSearchPage() {
      return CandidatesSearchPage;
    },
    getFilesList(){
      return this.filesList;
    }
  },

  methods: {
    sendLogin(){
      if(this.login.value !== '' && this.password.value !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);
        
        const data:IAPI_Login_Request = { email: this.login.value, password: this.password.value };
        
        API_Login(data)
          .then(async (response:any) => {
            await this.userInfoStore.onAuthorized(response);

            this.statusWindowStore.deteleStatusWindow(stID);
            this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Авторизация успешна!');

            this.currentStep++;
          })
          .catch(error => {
            this.statusWindowStore.deteleStatusWindow(stID);

            if(error.status === 500 || error.status === 400) this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Неверный логин или пароль!');
            else this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при авторизации!');
          });
        return;
      }

      if(this.login.value === ''){
        if(this.login.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите логин!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.login.error);
      }
      if(this.password.value === ''){
        if(this.password.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите пароль!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.password.error);
      }
    },
    sendLoginHR(){
      if(this.login.value !== '' && this.password.value !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);
        const data:IAPI_Login_Request = { email: this.login.value, password: this.password.value };
        API_LoginHR(data)
          .then(async (response:any) => {
            await this.userInfoStore.onAuthorized(response);

            this.statusWindowStore.deteleStatusWindow(stID);
            this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Авторизация успешна!');

            this.$router.push({name:'CandidatesSearchPage'});
          })
          .catch(error => {
            this.statusWindowStore.deteleStatusWindow(stID);

            if(error.status === 500 || error.status === 400) this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Неверный логин или пароль!');
            else this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при авторизации!');
          });
        return;
      }

      if(this.login.value === ''){
        if(this.login.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите логин!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.login.error);
      }
      if(this.password.value === ''){
        if(this.password.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите пароль!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.password.error);
      }
    },
    checkLogin(value: string){
      this.login = ValidUserLogin(value);
      if(value === '') this.login.error = '';
    },
    checkPassword(value: string){
      this.password = ValidUserPassword(value);
      if(value === '') this.password.error = '';
    },

    nextStep() {
      if (this.currentStep < 3) {
        this.currentStep++;
      }
    },
    previousStep() {
      if (this.currentStep == 3) {
        this.currentStep--;
      }
    },
    handleFilesChange(event: any){
      this.filesList = Array.from(event.target.files!);
      const selectedFiles:File[] = Array.from(event.target.files!);
      const wrongFiles = [] as File[];
      for(let i = 0; i < this.filesList.length; i++){
        if(!this.filesList[i].name.endsWith('.pdf')) {
          if(wrongFiles.length === 0) this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Неверное расширение файла!');
          wrongFiles.push(this.filesList[i]);
        }
      }
      for (let file of selectedFiles) {
        if (!file.name.endsWith('.pdf')) {
          wrongFiles.push(file);
        } else {
          this.formData.filesList.push(file);
        }
      }
      if (wrongFiles.length > 0) {
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Неверное расширение файла!');
      }
      console.log('Form Data:', this.formData);
      for(let file of wrongFiles) this.handleDeleteFile(file);
    },
    handleDeleteFile(fileDelete: File){
      this.formData.filesList = this.formData.filesList.filter((file) => file !== fileDelete);
      this.filesList = this.filesList.filter((file) => file !== fileDelete);
    },
    completeForm() {
      if(!this.formData.agreed){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Необходимо принять согласие об обработке!');
        return;
      }

      this.userFormStore.setFormData(this.formData);

      //API...

      this.$router.push({name: 'ProfilePage'});
    },
    nextPage() {
      this.$router.push({name: 'CandidatesSearchPage'});
    }

  },
};
</script>

