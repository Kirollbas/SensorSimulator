<script setup>
import './../style.css';
import { defineProps } from 'vue'
import { PRNG_TO_NAME } from './../constants/prng.js';
import { MODIFIER_TO_NAME, MODIFIER_DATA_TO_DATA_NAME } from './../constants/modifier.js';

const props = defineProps({
  modifiers: Array,
})

function getHint(modifier){
  let map = MODIFIER_DATA_TO_DATA_NAME[modifier.type]

  return Object.entries(modifier.data)
    .map(([field, value]) => field == "generator" ? parsePRNG(value) : `${map[field]}: ${value}`)
    .join('\n');
}

function parsePRNG(prng){
  return `ГПСЧ_тип: ${PRNG_TO_NAME[prng.type]}` + '\n' + `ГПСЧ_seed: ${prng.seed}`
}
</script>

<template>
  <div class="modifier-view">
    <div class="info-name">Модификаторы</div>
    <div class="info-data-container">
      <div class="info-data" v-for="(modifier, _) in modifiers">
        <div class="name">{{MODIFIER_TO_NAME[modifier.type]}}</div>
        <div class="icon">
          <i class="fas fa-circle-question delete-icon" :title="getHint(modifier)"></i>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.modifier-view{
  padding-top: 20px;
  align-items: center;  
}
.info-name{
  width: 100%;
}
.info-data-container{
  width: 100%;
  align-items: center;
}
.info-data{
  padding-top: 10px;
  padding-bottom: 10px;
  margin-top: 10px;
  display: flex;
  width: 100%;
  background-color: var(--color-gray);
  border-radius: 10px;
  align-items: center;
}
.name{
  width: 80%;
  word-wrap: break-word; /* перенос слов */
  white-space: normal;    /* разрешить перенос строк */
  overflow-wrap: break-word;
  display: block;
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