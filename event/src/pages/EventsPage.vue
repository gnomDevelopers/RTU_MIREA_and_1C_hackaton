<template>
  <div class="scrollable flex flex-col justify-center h-auto">
    <div class=" m-8 flex justify-center w-full h-auto text-3xl">
    Календарь мероприятий
    </div> 
    <!--прогресс бар по месяцам -->
    <div class="b-someclass">
      <div class="b-someclass-inner">
        <input type="radio" name="check" id="check-1" @change="showContainer = 'container1'">
        <label class="text-xl" for="check-1"><br><br>Июль</label>    
        
        <input type="radio" name="check" id="check-2" @change="showContainer = 'container2'">
        <label class="text-xl" for="check-2"><br><br>Август</label>
        
        <input type="radio" name="check" id="check-3" @change="showContainer = 'container3'">
        <label class="text-xl" for="check-3"><br><br>Сентябрь</label>
      
        <input type="radio" name="check" id="check-4" @change="showContainer = 'container4'">
        <label class="text-xl" for="check-4"><br><br>Октябрь</label>
        
        <input type="radio" name="check" id="check-5" @change="showContainer = 'container5'" checked> 
        <label class="text-xl" for="check-5"><br><br>Ноябрь</label>
        
        <input type="radio" name="check" id="check-6" @change="showContainer = 'container6'">
        <label class="text-xl" for="check-6"><br><br>Декабрь</label>
    
        <div class="b-someclass-line"></div>
      </div>     
    </div>
    <div v-if="showContainer==='container6'" class="flex flex-row m-8 h-auto" style="margin:20px">
      <div class="w-1/2" style="margin:30px">
        <div class="border-4 rounded-xl border-ev-color" style="padding:10px;height:550px">
          <div class="flex flex-row m-4">
            <img class="h-[100px]" style="margin-right:15px" src="../assets/images/image-nor.png">
            <div class="flex flex-col">
              <p class="font-bold md:text-2xl uus:text-xl">Норникель</p>
              <p class="md:text-xl uus:text-base">Хакатон</p>
              <p class="text-xl">Дата проведения: 6 декабря - 8 декабря</p>
            </div>
          </div>
          <div class="flex flex-col">
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Описание мероприятия</p>
            <p class="text-xl" style="margin-bottom: 20px">6-8 декабря проводится хакатон "Норникель: Интеллектуальные горизонты". Он проводится в онлайн-формате.
              Студенты там:<br>
              - получат практический опыт решения задач по информационной безопасности, продуктовой разработке и машинному обучению,<br>
              - разделят призовой фонд в 1.500.000 рублей<br>
              - получат обратную связь от экспертов<br><br>
            </p>
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Ссылка на мероприятие: <a style="color:#6a0793; text-decoration: underline"  href="https://clck.ru/3Er73P">https://clck.ru/3Er73P</a><br><br></p>
            <button v-if="!isRegistered" class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold btn w-full text-slate-100" @click="regToEvent">Зарегистрироваться на мероприятие</button>
            <button v-if="isRegistered" class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold  w-full " style="background-color:#D9D9D9">Вы зарегистрированы на мероприятие</button>

          </div>
        </div>
      </div>
      <div class="invisible w-1/2" style="margin:30px">
        <div class="border-4 rounded-xl border-ev-color" style="padding:10px; height:550px">
          <div class="flex flex-row m-4">
            <img class="max-h-[100px]" src="../assets/images/image-elma.png">
            <div class="flex flex-col">
              <p class="font-bold md:text-2xl uus:text-xl">ELMA Day 2024</p>
              <p class="md:text-xl uus:text-base">Конференция</p>
              <p class="text-xl">Дата проведения: 5 ноября</p>
            </div>
          </div>
          <div class="flex flex-col">
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Описание мероприятия</p>
            <p class="text-xl" style="margin-bottom: 20px">Конференция, посвящённая актуальным задачам
              бизнеса и единой Low-code платформе для их решения. Студенты узнают о передовых подходах
              к трансформации BPM-платформ в экосистемы, которые помогают компаниям более эффективно решать
              бизнес-задачи. Особое внимание будет уделено реализации экосистемного подхода на платформе ELMA365.<br><br>
            </p>
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Ссылка на мероприятие: <a style="color:#6a0793; text-decoration: underline"  href="https://elma365.com/ru/elmaday2024/">https://elma365.com/ru/elmaday2024/</a><br><br></p>
            <button class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold  w-full " style="background-color:#D9D9D9">Мероприятие уже завершилось</button>
          </div>
        </div>
      </div>
    </div>
    <div v-if="showContainer==='container5'" class="flex flex-row m-8 h-auto" style="margin:20px">
      <div class="w-1/2" style="margin:30px">
        <div class="border-4 rounded-xl border-ev-color" style="padding:10px;height:550px">
          <div class="flex flex-row m-4">
            <img class="h-[100px]" src="../assets/images/image-1c.png">
            <div class="flex flex-col">
              <p class="font-bold md:text-2xl uus:text-xl">РТУ МИРЭА и 1C</p>
              <p class="md:text-xl uus:text-base">Хакатон</p>
              <p class="text-xl">Дата проведения: 1 ноября - 30 ноября</p>
            </div>
          </div>
          <div class="flex flex-col">
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Описание мероприятия</p>
            <p class="text-xl" style="margin-bottom: 20px">Cтуденческий хакатон «Цифровые инструменты для бизнеса»,
              организован РТУ МИРЭА и компанией «1С». Студенческий хакатон проводится по двум трекам:<br>
              1. Трек «Разработка цифровых инструментов для бизнеса»<br>
              2. Трек «Аналитика для создания цифровых инструментов для бизнеса»<br>
              К участию приглашаются команды от 3 до 5 человек.
            </p>
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Ссылка на мероприятие: <a style="color:#6a0793; text-decoration: underline"  href="https://online-edu.mirea.ru/course/view.php?id=14488">https://online-edu.mirea.ru/course/view.php?id=14488</a></p>
            <button class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold btn w-full text-slate-100">Мероприятие ещё идёт</button>
          </div>
        </div>
      </div>
      <div class="w-1/2" style="margin:30px">
        <div class="border-4 rounded-xl border-ev-color" style="padding:10px; height:550px">
          <div class="flex flex-row m-4">
            <img class="max-h-[100px]" src="../assets/images/image-elma.png">
            <div class="flex flex-col">
              <p class="font-bold md:text-2xl uus:text-xl">ELMA Day 2024</p>
              <p class="md:text-xl uus:text-base">Конференция</p>
              <p class="text-xl">Дата проведения: 5 ноября</p>
            </div>
          </div>
          <div class="flex flex-col">
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Описание мероприятия</p>
            <p class="text-xl" style="margin-bottom: 20px">Конференция, посвящённая актуальным задачам
              бизнеса и единой Low-code платформе для их решения. Студенты узнают о передовых подходах
              к трансформации BPM-платформ в экосистемы, которые помогают компаниям более эффективно решать
              бизнес-задачи. Особое внимание будет уделено реализации экосистемного подхода на платформе ELMA365.<br><br>
            </p>
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Ссылка на мероприятие: <a style="color:#6a0793; text-decoration: underline"  href="https://elma365.com/ru/elmaday2024/">https://elma365.com/ru/elmaday2024/</a><br><br></p>
            <button class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold  w-full " style="background-color:#D9D9D9">Мероприятие уже завершилось</button>
          </div>
        </div>
      </div>
    </div>
    <div v-if="showContainer==='container4'" class="flex flex-row m-8 h-auto" style="margin:20px">
      <div class="w-1/2" style="margin:30px">
        <div class="border-4 rounded-xl border-ev-color" style="padding:10px;height:550px">
          <div class="flex flex-row m-4">
            <img class="h-[100px]" src="../assets/images/image-robo.png" style="margin-right:20px"/>
            <div class="flex flex-col">
              <p class="font-bold md:text-2xl uus:text-xl">Битва роботов</p>
              <p class="md:text-xl uus:text-base">Чемпионат</p>
              <p class="text-xl">Дата проведения: 26 октября</p>
            </div>
          </div>
          <div class="flex flex-col">
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Описание мероприятия</p>
            <p class="text-xl" style="margin-bottom: 20px">2-ой этап Битвы роботов пройдёт в г. Екатеринбурге.
              Чемпионат представляет собой масштабное событие с участием лучших изобретателей страны.
              Здесь лучшие инженерные умы могут продемонстрировать свои умения и побороться за призовой фонд в размере 6,6 миллионов рублей.<br><br>
            </p>
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Ссылка на мероприятие: <a style="color:#6a0793; text-decoration: underline"  href="https://www.gosuslugi.ru/bitvarobotov?utm_source=vk_group&utm_medium=opisanie&utm_term=sayt_gosuslug&utm_campaign=bitvarobotov">https://www.gosuslugi.ru/bitvarobotov?utm_source=vk_group</a><br><br></p>
            <button class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold  w-full " style="background-color:#D9D9D9">Мероприятие уже завершилось</button>

          </div>
        </div>
      </div>
      <div class="invisible w-1/2" style="margin:30px">
        <div class="border-4 rounded-xl border-ev-color" style="padding:10px; height:550px">
          <div class="flex flex-row m-4">
            <img class="max-h-[100px]" src="../assets/images/image-elma.png">
            <div class="flex flex-col">
              <p class="font-bold md:text-2xl uus:text-xl">ELMA Day 2024</p>
              <p class="md:text-xl uus:text-base">Конференция</p>
              <p class="text-xl">Дата проведения: 5 ноября</p>
            </div>
          </div>
          <div class="flex flex-col">
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Описание мероприятия</p>
            <p class="text-xl" style="margin-bottom: 20px">Конференция, посвящённая актуальным задачам
              бизнеса и единой Low-code платформе для их решения. Студенты узнают о передовых подходах
              к трансформации BPM-платформ в экосистемы, которые помогают компаниям более эффективно решать
              бизнес-задачи. Особое внимание будет уделено реализации экосистемного подхода на платформе ELMA365.<br><br>
            </p>
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Ссылка на мероприятие: <a style="color:#6a0793; text-decoration: underline"  href="https://elma365.com/ru/elmaday2024/">https://elma365.com/ru/elmaday2024/</a><br><br></p>
            <button class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold  w-full " style="background-color:#D9D9D9">Мероприятие уже завершилось</button>
          </div>
        </div>
      </div>
    </div>
    <div v-if="showContainer==='container3'" class="flex flex-row m-8 h-auto" style="margin:20px">
      <div class="w-1/2" style="margin:30px">
        <div class="border-4 rounded-xl border-ev-color" style="padding:10px;height:550px">
          <div class="flex flex-row m-4">
            <img class="h-[100px]" src="../assets/images/image-cif.png" style="margin-right:20px">
            <div class="flex flex-col">
              <p class="font-bold md:text-2xl uus:text-xl">Цифровой прорыв</p>
              <p class="md:text-xl uus:text-base">Хакатон</p>
              <p class="text-xl">Дата проведения: 6 - 8 сентября</p>
            </div>
          </div>
          <div class="flex flex-col">
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Описание мероприятия</p>
            <p class="text-xl" style="margin-bottom: 20px">С 6 по 8 сентября в Сибирском федеральном округе в гибридном формате
              пройдёт очередной окружной хакатон «Цифровой прорыв. Сезон: искусственный интеллект».
              На соревновании будет представлено шесть кейсов от партнёров хакатона в различных направлениях искусственного интеллекта.
              В ходе хакатона участники смогут решить задачи от государства и бизнеса по искусственному интеллекту и другим ИТ-тематикам и поборются за призовой фонд.
            </p>
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Ссылка на мероприятие: <a style="color:#6a0793; text-decoration: underline"  href="https://hacks-ai.ru/events/1077378">https://hacks-ai.ru/events/1077378</a><br><br></p>
            <button class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold  w-full " style="background-color:#D9D9D9">Мероприятие уже завершилось</button>
          </div>
        </div>
      </div>
      <div class="w-1/2" style="margin:30px">
        <div class="border-4 rounded-xl border-ev-color" style="padding:10px; height:550px">
          <div class="flex flex-row m-4">
            <img class="max-h-[100px]" src="../assets/images/image-kv.png" style="margin-right:20px">
            <div class="flex flex-col">
              <p class="font-bold md:text-2xl uus:text-xl">Семья - в сердце каждого</p>
              <p class="md:text-xl uus:text-base">Квартирник</p>
              <p class="text-xl">Дата проведения: 25 сентября</p>
            </div>
          </div>
          <div class="flex flex-col">
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Описание мероприятия</p>
            <p class="text-xl" style="margin-bottom: 20px">25 сентября в РТУ МИРЭА в кампусе на улице Стромынка, 20 состоится квартирник «Семья — в сердце каждого».

              На мероприятии студенты проведут время в вузе в уютной домашней атмосфере. Участники продемонстрируют свои таланты в танцах и вокале и зарядят гостей на позитив.

              Тёплая обстановка и по-настоящему семейный вечер подарит ребятам радость быть частью одной большой семьи — РТУ МИРЭА. <br>
            </p>
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Ссылка на мероприятие: <a style="color:#6a0793; text-decoration: underline"  href="https://www.mirea.ru/news/v-rtu-mirea-proshyel-kvartirnik-semya-v-serdtse-kazhdogo/">https://www.mirea.ru/news/v-rtu-mirea-proshyel-kvartirnik-semya-v-serdtse-kazhdogo//</a></p>
            <button class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold  w-full " style="background-color:#D9D9D9">Мероприятие уже завершилось</button>
          </div>
        </div>
      </div>
    </div>
    <div v-if="showContainer==='container2'" class="flex flex-row m-8 h-auto" style="margin:20px">
      <div class="w-1/2" style="margin:30px">
        <div class="border-4 rounded-xl border-ev-color" style="padding:10px;height:550px">
          <div class="flex flex-row m-4">
            <img class="h-[100px]" src="../assets/images/image-cur.png" style="margin-right: 20px">
            <div class="flex flex-col">
              <p class="font-bold md:text-2xl uus:text-xl">Школа кураторов</p>
              <p class="md:text-xl uus:text-base">Концерт</p>
              <p class="text-xl">Дата проведения: 12 августа</p>
            </div>
          </div>
          <div class="flex flex-col">
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Описание мероприятия</p>
            <p class="text-xl" style="margin-bottom: 20px">12 августа стартует проект «Школа кураторов»
              2024 — главное мероприятие программы адаптации студентов «Факел» Студенческого союза МИРЭА. С 12 по 22 августа будущих кураторов ждёт насыщенная программа,
              в рамках которой они получат все необходимые знания и навыки, чтобы помочь адаптироваться будущим первокурсникам нашего вуза.

            </p>
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Ссылка на мероприятие: <a style="color:#6a0793; text-decoration: underline"  href="https://www.mirea.ru/news/v-rtu-mirea-startoval-proekt-shkola-kuratorov-2024/">https://www.mirea.ru/news/v-rtu-mirea-startoval-proekt-shkola-kuratorov-2024/</a><br><br></p>
            <button class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold  w-full " style="background-color:#D9D9D9">Мероприятие уже завершилось</button>
          </div>
        </div>
      </div>
      <div class="invisible w-1/2" style="margin:30px">
        <div class="border-4 rounded-xl border-ev-color" style="padding:10px; height:550px">
          <div class="flex flex-row m-4">
            <img class="max-h-[100px]" src="../assets/images/image-elma.png">
            <div class="flex flex-col">
              <p class="font-bold md:text-2xl uus:text-xl">ELMA Day 2024</p>
              <p class="md:text-xl uus:text-base">Конференция</p>
              <p class="text-xl">Дата проведения: 5 ноября</p>
            </div>
          </div>
          <div class="flex flex-col">
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Описание мероприятия</p>
            <p class="text-xl" style="margin-bottom: 20px">Конференция, посвящённая актуальным задачам
              бизнеса и единой Low-code платформе для их решения. Студенты узнают о передовых подходах
              к трансформации BPM-платформ в экосистемы, которые помогают компаниям более эффективно решать
              бизнес-задачи. Особое внимание будет уделено реализации экосистемного подхода на платформе ELMA365.<br><br>
            </p>
            <p class="text-xl" style="margin-bottom: 20px; margin-top: 20px">Ссылка на мероприятие: <a style="color:#6a0793; text-decoration: underline"  href="https://elma365.com/ru/elmaday2024/">https://elma365.com/ru/elmaday2024/</a><br><br></p>
            <button class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-10 font-semibold  w-full " style="background-color:#D9D9D9">Мероприятие уже завершилось</button>
          </div>
        </div>
      </div>
    </div>
    <div v-if="showContainer==='container1'" class="flex flex-row justify-center m-8 h-auto" style="margin:20px">
    <p class="text-3xl" style="margin:50px">В этом месяце мероприятий не было</p>
    </div>
    <!-- Модальное окно -->
    <div v-if="isModalVisible" class="modal-overlay fixed top-0 left-0 w-full h-full flex items-center justify-center z-50">
      <!-- Затемнение фона -->
      <div class="absolute top-0 left-0 w-full h-full bg-black opacity-50" @click="closeModal"></div>

      <!-- Модальное окно -->
      <div v-if="!isRegistered" class="relative bg-white rounded-lg p-8 shadow-lg z-10">

        <!-- Содержимое модального окна -->
        <div class="flex flex-col items-center text-center modal-content">
          <p class="text-3xl mb-4">Регистрация на мероприятие</p>
          <div class="flex flex-row" style="margin:20px">
            <img src="../assets/images/image-nor.png" style="margin-right:20px">
            <div class="flex flex-col items-start">
              <p class="text-xl font-bold">Норникель</p>
              <p class="text-xl ">Хакатон</p>
              <p class="text-xl ">Дата проведения: 6 - 8 декабря</p>
            </div>

          </div>

          <p class="text-xl">Выберите команду для участия</p><br>
          <div class="flex justify-center w-full">
            <button v-if="!isChosen" @click="chooseTeam" class="cursor-pointer transition-colors py-2 px-5 text-xl rounded-md h-12 border-2 border-ev-color btn w-full" style="background-color: #F3F0FF"><img class="items-start" src="../assets/icons/icon-team.svg" style="padding-right: 20px">Здесь могла быть ваша реклама</button>
            <button v-if="isChosen" @click="cancelChoice"  class="cursor-pointer transition-colors py-2 px-5 text-xl rounded-md h-12 border-4 border-ev-color btn w-full" style="background-color: #F3F0FF"><img class="items-start" src="../assets/icons/icon-team.svg" style="padding-right: 20px" >Выбрана команда: Здесь могла быть ваша реклама</button>
          </div><br><br>
          <div class="flex justify-center w-3/5">
            <button v-if="!isChosen" class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-12 font-semibold btn w-full " style="background-color: #D9D9D9">Зарегистрироваться</button><br><br>
            <button v-if="isChosen" @click="register" class="cursor-pointer transition-colors py-2 px-5 text-lg rounded-md h-12 font-semibold btn w-full text-slate-100">Зарегистрироваться</button>
          </div>

        </div>
      </div>

      <div v-if="isRegistered" class="relative bg-white rounded-lg p-8 shadow-lg z-10">

        <!-- Содержимое модального окна -->
        <div class="flex flex-col items-center text-center modal-content">
          <p class="text-3xl mb-4">Вы успешно зарегистрированы на мероприятие!</p>
        </div>
      </div>
    </div>
  </div>

</template>

<script lang="ts">
import { useUserFormStore } from '@/stores/userFormStore';

export default {
  data() {
    return {
      showContainer: 'container5',
      isModalVisible: false,
      isChosen: false,
      isRegistered: false,
    };
  },
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
    regToEvent(){
      this.isModalVisible = true;
    },
    closeModal() {
      this.isModalVisible = false;
    },
    chooseTeam() {
      this.isChosen = true;
    },
    cancelChoice() {
      this.isChosen = false;
    },
    register() {
      this.isRegistered = true;
    }
  }
};
</script>
<style scoped>
* {
  padding: 0;
  margin: 0;
}
.b-someclass {
  text-align: center;
}
.b-someclass-inner {
  display: inline-block;
  vertical-align: top;
  position: relative;
}
.b-someclass-inner:before {
  content: '';
  position: absolute;
  top: 50%;
  left: 15px;
  width: calc(100% - 30px);
  height: 6px;
  background: #C9C2F4;
  margin-top: -3px;
  z-index: 2;
}
.b-someclass-inner input {
  display: none;
}
.b-someclass-inner input + label {
  display: inline-block;
  vertical-align: top;
  margin: 35px;
  width: 50px;
  height: 50px;
  background: #C9C2F4;
  border-radius: 50%;
  position: relative;
  z-index: 9;
}
.b-someclass-inner input + label:after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  height: 50%;
  width: 50%;
  margin: -25% 0 0 -25%;
  border-radius: 50%;
  background: #52026A;
  transform: scale(0);
  transition: .25s;
  transition-delay: .2s;
}
.b-someclass-inner input:checked + label:after {
  background: #52026A;
  transform: scale(1);
}
.b-someclass-line {
  position: absolute;
  top: 50%;
  margin-top: -3px;
  left: 15px;
  width: 0;
  height: 6px;
  background: #52026A;
  transition: .35s;
  z-index: 2;
  pointer-events: none;
}
.b-someclass-inner input:checked + label ~ .b-someclass-line {
  left: 50px;
}
.b-someclass-inner #check-2:checked + label ~ .b-someclass-line {
  width: calc(100% / 6);
}
.b-someclass-inner #check-3:checked + label ~ .b-someclass-line {
  width: calc(100% / 6 * 2);
}
.b-someclass-inner #check-4:checked + label ~ .b-someclass-line {
  width: calc(100% / 6 * 3);
}
.b-someclass-inner #check-5:checked + label ~ .b-someclass-line {
  width: calc(100% / 6 * 4);
}
.b-someclass-inner #check-6:checked + label ~ .b-someclass-line {
  width: calc(100% / 6 * 5);
}
.modal-overlay {

  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  z-index: 10;
}

.modal-content {
  width: 45%;
  height: 50%;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: white;
  border-radius: 8px;
  padding: 30px;
  margin: 30px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  z-index: 20;
}

button {
  transition: background-color 0.3s;
}
</style>
