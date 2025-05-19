<script setup>
import './../style.css';
import { defineProps } from 'vue'
import BaseView from './BaseView.vue'
import { apiUrl } from './../constants/address.js';
import ModifiersView from './ModifiersView.vue'

const props = defineProps({
  simulator: Object,
  fetchFunc: Function,
  isRunning: Boolean,
})

async function deleteSimulator(name) {
  try {
    const response = await fetch(`${apiUrl}/api/simulator/${name}`, {
        method: "DELETE",
    });
    
    props.fetchFunc();
  } catch (error) {
    props.fetchFunc();
    console.error('cannot delete simulator:', error);
  }
}

</script>

<template>
  <div class="simulator">
    <div class="header">
        <div class="name">{{ simulator.name }}</div>
        <button @click="deleteSimulator(simulator.name)" class="delete-button" v-if="!isRunning">
            <i class="fas fa-trash delete-icon"></i> 
        </button>
        <div class="delete-button" v-if="isRunning">
            <i :class="simulator.is_active ? 'fas fa-circle active-icon' : 'fas fa-circle delete-icon'"></i> 
        </div>
    </div>
    <div class="info">
        <div class="info-name">Название:</div>
        <div class="info-data">{{ simulator.name }}</div>
    </div>
    <div class="info">
        <div class="info-name">Адрес:</div>
        <div class="info-data">{{ simulator.address }}</div>
    </div>
    <div class="info">
        <div class="info-name">Длительность:</div>
        <div class="info-data">{{ simulator.duration }}</div>
    </div>
    <BaseView :base="simulator.base"/>
    <ModifiersView :modifiers="simulator.modifiers"/>
  </div>
</template>

<style scoped>
.simulator{
    background-color: var(--color-dark_gray);
    border-radius: 20px;
    box-sizing: border-box;
    padding: 10px;
}
.header{
    display: flex; 
    padding: 10px;
}
.name{
    display: flex; 
    width: 90%;
    font-size: 20px;     
    font-family: 'Arial';
    font-weight: bold;
}
.delete-button{
    width: 10%;
    display: flex;
    align-items: center;
    justify-content: center;
    background: none; 
    border: none; 
    padding: 0; 
    margin: 0; 
    cursor: pointer;
    outline: none;
}
.delete-icon{
    color: var(--color-red);   
}
.active-icon{
    color: var(--color-green);   
}
.info{
    display: flex;
}
.info-name{
    display: flex; 
    width: 50%;
}
.info-data{
    display: flex; 
    width: 50%;
}
</style>