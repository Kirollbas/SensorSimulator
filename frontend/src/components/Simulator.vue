<script setup>
import './../style.css';
import { defineProps } from 'vue'
import BaseView from './BaseView.vue'
import ModifiersView from './ModifiersView.vue'

const props = defineProps({
  simulator: Object,
  fetchFunc: Function,
})

async function deleteSimulator(name) {
  try {
    const response = await fetch(`http://localhost:8080/api/simulator/${name}`, {
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
        <button @click="deleteSimulator(simulator.name)" class="delete-button">
            <i class="fas fa-trash delete-icon"></i> 
        </button>
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