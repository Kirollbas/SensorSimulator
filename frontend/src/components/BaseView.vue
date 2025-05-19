<script setup>
import './../style.css';
import { defineProps } from 'vue'
import { BASE_TO_NAME, BASE_DATA_TO_DATA_NAME } from './../constants/base.js';
import { PRNG_TO_NAME } from './../constants/prng.js';

const props = defineProps({
  base: Object,
})

function getHint(base){
  let map = BASE_DATA_TO_DATA_NAME[base.type]

  return Object.entries(base.data)
    .map(([field, value]) => field == "generator" ? parsePRNG(value) : `${map[field]}: ${value}`)
    .join('\n');
}

function parsePRNG(prng){
  return `ГПСЧ_тип: ${PRNG_TO_NAME[prng.type]}` + '\n' + `ГПСЧ_seed: ${prng.seed}`
}
</script>


<template>
  <div class="base-view">
    <div class="info-name">Основа:</div>
    <div class="info-data">
      <div class="name">{{BASE_TO_NAME[base.type]}}</div>
      <div class="icon">
        <i class="fas fa-circle-question delete-icon" :title="getHint(base)"></i>
      </div>
    </div>
  </div>
</template>

<style scoped>
.base-view{
  padding-top: 20px;
  display: flex;
  height: 30px;
  align-items: center;  
}
.info-name{
  display: flex; 
  width: 30%;
}
.info-data{
  display: flex; 
  width: 70%;
  background-color: var(--color-gray);
  height: 100%;
  border-radius: 10px;
  align-items: center;
  padding: 5px;
}
.name{
  width: 80%;   
  font-family: 'Arial';
  font-weight: bold;
  font-size: 14px;
}
.icon{
  width: 20%;
  opacity: 50%;
  cursor: pointer;
}
</style>