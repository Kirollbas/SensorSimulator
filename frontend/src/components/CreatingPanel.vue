<script setup>
import './../style.css';
import Simulator from './Simulator.vue'
import { ref, defineProps } from 'vue'

const isAddingSimulator = ref(false)

const props = defineProps({
  isAddingSimulator: Boolean,
})

</script>

<template>
  <div class="panel">
  </div>
</template>

<style scoped>
.panel{
    background-color: var(--color-gray);
    width: 100%;
    height: 100%;
}
</style>
