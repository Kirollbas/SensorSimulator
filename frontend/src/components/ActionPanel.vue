<script setup>
import './../style.css';
import CreatingPanel from './CreatingPanel.vue'
import { apiUrl } from './../constants/address.js';
import { ref, defineProps } from 'vue'

const isAddingSimulator = ref(false)
const formRef = ref()

const props = defineProps({
  changeRunState: Function,
  isRunning: Boolean,
})

const leftButton = ref("Добавить симулятор")
const rightButton = ref("Начать симуляцию")

function openCreation(){
    isAddingSimulator.value = true
}

async function closeCreation(){
    await formRef.value.clearData()
    isAddingSimulator.value = false
}

async function startSimulation() {
    const response = await fetch(`${apiUrl}/api/start`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    })

    if (!response.ok) {
        console.log(await response.json())
        return
    }

    props.changeRunState(true)
}

async function stopSimulation() {
    const response = await fetch(`${apiUrl}/api/stop`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    })

    if (!response.ok) {
        console.log(await response.json())
        return
    }
    props.changeRunState(false)
}

async function addSimulator() {
    const json = await formRef.value.getJson()

    if (json) {
        const response = await fetch(`${apiUrl}/api/simulator/add`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(json)
        })

        if (!response.ok) {
            console.log(await response.json())
        }
    } else {
        console.log("Форма невалидна")
    }
}

</script>

<template>
  <div class="panel">
    <div class="skip"></div>
    <div class="add-panel">
        <CreatingPanel :isAddingSimulator="isAddingSimulator" :isRunning="isRunning" ref="formRef"/>
    </div>
    <div class="buttons">
        <button class="first-button" @click="isAddingSimulator ? closeCreation(): openCreation()" v-if="!isRunning">
            {{ isAddingSimulator ? "Отмена" : "Добавить симулятор" }}
        </button>
        <div v-if="isRunning" class="skip"></div>
        <button :class="isRunning ? 'stop-button' : (isAddingSimulator ? 'adding-button' : 'start-button')" @click="isRunning ? stopSimulation() : (isAddingSimulator ? addSimulator() : startSimulation())">
            {{ isRunning ? "Остановить симуляцию" : (isAddingSimulator ? "Потдвердить" : "Начать симуляцию") }}
        </button>
    </div>
  </div>
</template>

<style scoped>
.panel{
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
}
.skip{
    display: flex;
    width: 100%;
    height: 8%;
}
.add-panel{
    display: flex;
    padding: 10px;
    width: 100%;
    height: 82%;
    box-sizing: border-box;
}
.buttons{
    display: flex;
    width: 100%;
    height: 10%;
    padding-bottom: 10px;
    padding-left: 10px;
    padding-right: 10px;
    box-sizing: border-box;
    
}
.first-button{
    border-radius: 10px;
    width: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    box-sizing: border-box;
    background-color: var(--color-gray); 
    border: none; 
    padding: 0; 
    margin: 0; 
    cursor: pointer;
    outline: none;
    margin-right: 10px;
}

.adding-button{
    border-radius: 10px;
    width: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    box-sizing: border-box;
    background-color: var(--color-gray); 
    border: none; 
    padding: 0; 
    margin: 0; 
    cursor: pointer;
    outline: none;
    margin-left: 10px;
}
.start-button {
    border-radius: 10px;
    width: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    box-sizing: border-box;
    background-color: var(--color-green); 
    border: none; 
    padding: 0; 
    margin: 0; 
    cursor: pointer;
    outline: none;
    margin-left: 10px;
}
.stop-button{
    border-radius: 10px;
    width: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    box-sizing: border-box;
    background-color: var(--color-red); 
    border: none; 
    padding: 0; 
    margin: 0; 
    cursor: pointer;
    outline: none;
    margin-left: 10px;
}
.hidden{
    visibility: hidden;
}
.skip{
    display: flex;
    width: 50%;
}
</style>
