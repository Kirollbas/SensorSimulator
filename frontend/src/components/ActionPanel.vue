<script setup>
import './../style.css';
import CreatingPanel from './CreatingPanel.vue'
import { ref, defineProps } from 'vue'

const isAddingSimulator = ref(false)

const leftButton = ref("Добавить симулятор")
const rightButton = ref("Начать симуляцию")

</script>

<template>
  <div class="panel">
    <div class="skip"></div>
    <div class="add-panel">
        <CreatingPanel :isAddingSimulator="isAddingSimulator"/>
    </div>
    <div class="buttons">
        <button @click="handleClick">
            Нажми меня
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
}
</style>
