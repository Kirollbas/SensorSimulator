<script setup>
import './../style.css';
import ModifierInput from './ModifierInput.vue'
import { reactive, defineProps, defineExpose, ref } from 'vue'
import { useForm, useField } from 'vee-validate'
import * as yup from 'yup'

const componentList = ref([])

const componentRefs = ref([])

function addComponent() {
  componentList.value.push({})
}

const getJson = async () => {
  const data = []

  let is_bad_data = false

  for (const child of componentRefs.value) {
    if (child && child.getJson) {
      const result = await child.getJson()
      if (result == null) {
        is_bad_data = true
        break
      }
      data.push(result)
    }
  }

  if (is_bad_data) {
    return null
  }

  return data
}

const clearData = async() => {
    componentRefs.value.forEach(async (child, i) => {
    if (child && child.clearData) {
      await child.clearData()
    }
    })

    componentList.value = []
    componentRefs.value = []
}

defineExpose({ 
    getJson,
    clearData
})
</script>

<template>
  <div class="modifiers-panel">
    <div class="modifier-name">
      Модификаторы
    </div>
    <ModifierInput
      v-for="(_, i) in componentList"
      :key="i"
      ref="componentRefs"
    />

    <button class="add-modifier-button" @click="addComponent">
      <i class="fas fa-plus add-icon"></i> 
    </button>
  </div>
</template>

<style scoped>
.modifiers-panel{
  margin-top: 10px;
  display: flex;
  width: 100%;
  flex-direction: column;
  padding-top: 10px;
  box-sizing: border-box;
  border-radius: 10px;
}
.modifier-name{
  width: 100%;
  align-content: center;
  padding: 10px;
  box-sizing: border-box;
  font-family: 'Arial';
  font-weight: bold;
  font-size: 20px;
}
.add-modifier-button{
  background-color: var(--color-dark_gray);
  border: none;
  outline: none;
  box-shadow: none;
  margin-top: 20px;
}
</style>