<script setup>
import './../style.css';
import Simulator from './Simulator.vue'
import { defineProps } from 'vue'

const props = defineProps({
  simulators: Array,
  fetchFunc: Function,
})

</script>

<template>
  <div class="container">
    <div class="simulators-header">
      Список симуляторов:
    </div>
    <div class="simulators-container">
      <div class="simulators-list-container">
        <div class="simulators-list">
          <Simulator v-for="(item, index) in simulators" :simulator="item" :fetchFunc="fetchFunc"/>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
}

.simulators-header{
  height: 8%;
  padding-left: 10px;
  text-align: left;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  font-size: 32px;         /* размер текста */
  font-family: 'Arial';    /* шрифт */
  font-weight: bold;  
}
.simulators-container{
  height: 92%;
  padding: 10px;
  box-sizing: border-box;
}
.simulators-list-container{
  height: 100%;
  width: 100%;
  background-color: var(--color-gray);
  border-radius: 20px;
  padding: 20px;
  overflow-y: auto;
  box-sizing: border-box;
  scrollbar-width: none; /* Firefox */
  -ms-overflow-style: none;  /* IE/Edge */
}

.simulators-list-container::-webkit-scrollbar {
  display: none; /* Chrome, Safari, Opera */
}

.simulators-list{
  box-sizing: border-box;
  display: grid;
  grid-template-columns: repeat(3, 30%);
  gap: 5%;
}
</style>
