<script setup>
import './style.css';
import Simulators from './components/Simulators.vue'
import ActionPanel from './components/ActionPanel.vue'
import { apiUrl } from './constants/address.js';
import { ref, onMounted, onBeforeUnmount } from 'vue'

const simulators = ref([])
let intervalId;

const isSimulatingRuns = ref(false)

function changeSimulationState(isRunning) {
  isSimulatingRuns.value = isRunning
}

onMounted(() => {
  fetchSimulatorsData();
  intervalId = setInterval(fetchSimulatorsData, 1000);
})

onBeforeUnmount(() => {
  clearInterval(intervalId);
});

async function fetchSimulatorsData() {
  try {
    const response = await fetch(`${apiUrl}/api/simulator`);
    simulators.value = (await response.json()).simulators;

    simulators.value.sort((a, b) => {
      return a.name.localeCompare(b.name)
    })
  } catch (error) {
    console.error('cannot fetch simulators data:', error);
  }
}
</script>

<template>
  <div class="main">
    <div class="simulators">
      <Simulators :simulators="simulators" :fetchFunc="fetchSimulatorsData" :isRunning="isSimulatingRuns"/>
    </div>

    <div class="action-panel">
      <ActionPanel :isRunning="isSimulatingRuns" :changeRunState="changeSimulationState"/>
    </div>
  </div>
</template>

<style scoped>

.main {
  width: 100%;
  display: flex; 
  height: 100vh;
}
.simulators{
  width: 60%; 
  background-color: var(--color-dark_gray); 
  display: flex; 
  box-sizing: border-box;
}
.action-panel{
  padding-right: 10px;
  width: 40%; 
  background-color: var(--color-dark_gray); 
  display: flex; 
}
</style>