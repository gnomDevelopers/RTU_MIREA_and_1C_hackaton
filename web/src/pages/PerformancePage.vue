<template>
  <div class="flex flex-col items-center scrollable px-4 lg:px-0">
    <div class="flex flex-col w-full lg:w-10/12 items-center gap-y-4 mb-4">

      <PageTitle title="Успеваемость" description="В этом разделе Вы можете посмотреть и изменить успеваемость студентов. Для этого найдите в поиске нужную группу и предмет, а затем посмотрите или отредактируйте успеваемость.">
        <IconPerformance class="w-20 us:w-36 h-20 us:h-36"/>
      </PageTitle>

      <div>
        <SearchList 
          title="" 
          placeholder="Введите название группы"
          :searchList="groupsList" 
          :itemComponent="getListItemComponent"
        />
        <p class="text-lg text-center font-medium cursor-default">
          Выбранная группа: 
          <span class="underline cursor-pointer">ЭФБО-01-23</span>
        </p>
      </div>

      <div class="flex flex-col p-4 rounded-lg gap-y-4 bg-color-light">
        <p class="text-center text-xl p-1 rounded-lg cursor-default bg-white">Выберите дисциплину</p>
        <div class="flex flex-col gap-y-1">
          <SubjectItem v-for="i in 6" :id="i" :name="'Дискретная математика (часть 2/2) [I.24-25]'"/>
        </div>
      </div>

      <div v-show="tableType === 0" class="flex flex-row items-start flex-wrap-0 self-stretch">
        <table class="flex-shrink-0 cursor-default table-decorate">
          <thead>
            <tr>
              <th class="w-10">№</th>
              <th class="text-nowrap">ФИО</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="j in 30">
              <td class="font-semibold">{{ j }}</td>
              <td class="max-w-96 overflow-hidden text-nowrap">Иванов Иван Иванович</td>
            </tr>
          </tbody>
        </table>

        <div class="overflow-x-scroll scrollable-table cursor-default">
          <table class="w-auto no-x-border table-decorate">
            <thead>
              <tr>
                <th v-for="i in 50" class="w-16 min-w-10">01.09</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="j in 30">
                <td v-for="i in 50" class="w-16 min-w-10">5</td>
              </tr>
            </tbody>
          </table>
        </div>

        <table class="flex-shrink-0 cursor-default table-decorate">
          <thead>
            <tr>
              <th>GPA</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="j in 30">
              <td class="font-semibold">10.23</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-show="tableType === 1" class="flex flex-row items-start flex-wrap-0 self-stretch overflow-x-scroll scrollable-table ">
        <table class="cursor-default table-decorate">
          <thead>
            <tr>
              <th class="w-10">№</th>
              <th class="max-w-96 overflow-hidden text-nowrap">ФИО</th>
              <th v-for="i in 50">01.09</th>
              <th>GPA</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="j in 30">
              <td class="font-semibold">{{ j }}</td>
              <td>Иванов Иван Иванович</td>
              <td v-for="i in 50">5</td>
              <td class="font-semibold">10.23</td>
            </tr>
          </tbody>
        </table>
      </div>

    </div>
  </div>
</template>
<script lang="ts">
import SearchList from '@/entities/searchList.vue';
import IconPerformance from '@/shared/iconPerformance.vue';
import SubjectItem from '@/shared/subjectItem.vue';
import { type ISearchList, type IItemList } from '@/helpers/constants';
import PerformanceSearchListItem from '@/entities/listItems/performanceSearchListItem.vue';
import PageTitle from '@/shared/pageTitle.vue';

export default{
  components:{
    PageTitle,
    IconPerformance,
    SubjectItem,
    SearchList,
    PerformanceSearchListItem,
  },
  data(){
    return{
      tableType: 0 as number, // 0 - таблица для больших экранов (>=756px), 1 - таблица для маленьких экранов (<756px)
      searchFilter: '' as string,
      groupsList: [] as ISearchList[],
    }
  },
  computed:{
    getListItemComponent(){
      return PerformanceSearchListItem;
    }
  },
  mounted() {
    this.setTableType();
    window.addEventListener('resize', this.setTableType);

    for(let i = 1; i < 20; i++){
      const data: IItemList = {id: i, name: `ЭФБО-${(i < 10 ? '0' : '')}${i}-23`};
      this.groupsList.push({id: data.id, search_field: data.name, data: data});
    }
  },
  methods:{
    setTableType(){
      if(window.innerWidth < 756) this.tableType = 1;
      else this.tableType = 0;
    },
    filterUserList(){

    }
  },
  unmounted() {
    window.removeEventListener('resize', this.setTableType);
  },
};
</script>