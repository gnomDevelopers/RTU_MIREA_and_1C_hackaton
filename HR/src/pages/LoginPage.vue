<template>
  <div class="flex h-screen">
    <div class="flex flex-col w-1/2 h-full items-start justify-start p-4">
      <p class="logo-hr">WorkVUZ+</p>
      <div class="flex flex-col w-full items-center justify-center">
        <div class="flex flex-col w-10/12 items-center m-10">
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
                HR-специалист
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
          <div v-if="userType === 'student'" class="w-full relative">
            <p class="auth-description">Войдите с помощью учётной записи VUZ+</p>
            <div v-if="currentStep === 1">
              <form @submit.prevent="nextStep">
                <div class="mb-4">

                  <loginInput type="text"  text="Логин:" @input-change="checkLogin"/>
                  <loginInput type="password" id="student-password" text="Ваш пароль:" @input-change="checkPassword"/>

                  <submitButton value="Войти" class="btn" @click="sendLogin"/>
                </div>
              </form>
            </div>

            <!-- Второй шаг -->
            <div v-else-if="currentStep === 2">
              <form @submit.prevent="nextStep">
                <div class="mb-4">
                  <label for="specialty" class="block text-gray-700 text-sm font-bold mb-2">Специальность:</label>
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
                  <label for="phone" class="block text-gray-700 text-sm font-bold mb-2">Номер телефона:</label>
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
                  <label for="experience" class="block text-gray-700 text-sm font-bold mb-2">Опыт работы:</label>
                  <input
                      type="text"
                      id="experience"
                      required
                      v-model="formData.experience"
                      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      placeholder="Введите ваш опыт работы"
                  />
                </div>
              </form>
            </div>

            <!-- Третий шаг -->
            <div v-else-if="currentStep === 3">
              <form @submit.prevent="completeForm">
                <div class="mb-4">
                  <label for="links" class="block text-gray-700 text-sm font-bold mb-2">Полезные ссылки:</label>
                  <input
                      type="text"
                      id="links"
                      v-model="formData.links"
                      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                      placeholder="Добавьте ссылки"
                  />
                </div>
                <div class="mb-4">
                  <label for="resume" class="block text-gray-700 text-sm font-bold mb-2">Резюме:</label>
                  <input
                      type="file"
                      id="resume"
                      @change="handleFileUpload"
                      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  />
                </div>
              </form>
            </div>

            <!-- Кнопка продолжения -->
            <div class=" left-0 w-full px-4">
              <button v-if="currentStep < 3" @click="nextStep" class="w-full btn h-10 rounded-md">
                Продолжить
              </button>
              <button v-else @click="completeForm" class="w-full btn h-10 rounded-md">
                Завершить заполнение
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="login-bg w-1/2"></div>
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
import { API_Login } from '@/api/api';
export default {

  components:{
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
        links: '',
        resume: null,
      },
    }
  },
  computed:{
    ...mapStores(useStatusWindowStore, useUserInfoStore),
  },
  methods: {
    sendLogin(){
      if(this.login.value !== '' && this.password.value !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);
        const data:IAPI_Login_Request = { email: this.login.value, password: this.password.value };
        API_Login(data)
            .then(response => {
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
    completeForm() {
      alert(`Форма завершена! Данные: ${JSON.stringify(this.formData)}`);
    },
    // handleFileUpload(event) {
    //   const file = event.target.files[0];
    //   this.formData.resume = file;
    // },
    togglePasswordVisibility() {
      this.showPassword = !this.showPassword;
    },
  },
};
</script>

