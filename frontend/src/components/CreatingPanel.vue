<script setup>
import './../style.css';
import Simulator from './Simulator.vue'
import BaseInput from './BaseInput.vue'
import ModifiersInput from './ModifiersInput.vue'
import { reactive, defineProps, defineExpose, ref } from 'vue'
import { useForm, useField } from 'vee-validate'
import * as yup from 'yup'

const schema = yup.object({
  name: yup.string()
    .required('Введите название'),
  address: yup.number()
    .required("Введите адрес Modbus")
    .integer('Только целое число')
    .min(1, 'Адрес дожен находится в промежутке между 1 и 65535')
    .max(65535, 'Адрес дожен находится в промежутке между 1 и 65535'),
  duration: yup.string()
    .required("Введите длительность симуяции")
    .matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Длительность должна соответствовать шаблону 123h45m6s, неограниченная длительность при 0s')
})

const { resetForm, errors, validate, values } = useForm({ validationSchema: schema })

const { value: name } = useField('name')
const { value: address } = useField('address')
const { value: duration } = useField('duration')

const props = defineProps({
  isAddingSimulator: Boolean,
  isRunning: Boolean,
})

const baseRef = ref()
const modifiersRef = ref()

const getJson = async () => {
  const result = await validate()
  const base = await baseRef.value.getJson()
  const modifiers = await modifiersRef.value.getJson()

  if (!result.valid) {
    return null
  }

  if (base == null) {
    return null
  }

  if (modifiers == null) {
    return null
  }

  const data = JSON.parse(JSON.stringify(values))
  data.base = base
  data.modifiers = modifiers

  return data
}

const clearData = async() => {
    await baseRef.value.clearData()
    await modifiersRef.value.clearData()
    resetForm()
}

defineExpose({ 
    getJson,
    clearData
})

</script>

<template>
  <div :class="isRunning? 'inactive' : (isAddingSimulator ? 'panel' : 'inactive')">
    <form @submit.prevent="onSubmit" class="inputs">
      <label class="input">
        <div class="data">
            <div class="name">
                Название:
            </div>
            <input type="string" v-model="name" class="input-field"/>
        </div>
        <div class="hint">
            {{ errors.name }}
        </div>
      </label>
      <label class="input">
        <div class="data">
            <div class="name">
                Адрес:
            </div>
            <input type="number" step="1" v-model="address" class="input-field"/>
        </div>
        <div class="hint">
            {{ errors.address }}
        </div>
      </label>
      <label class="input">
        <div class="data">
            <div class="name">
                Длительность:
            </div>
            <input type="string" v-model="duration" class="input-field"/>
        </div>
        <div class="hint">
            {{ errors.duration }}
        </div>
      </label>
    </form>
    <BaseInput ref="baseRef"/>
    <ModifiersInput ref="modifiersRef"/>
  </div>
</template>

<style scoped>
.inactive {
    visibility: hidden;
    overflow: hidden;
}
.panel{
    background-color: var(--color-gray);
    width: 100%;
    height: 100%;
    border-radius: 20px;
    padding: 20px;
    box-sizing: border-box;
    overflow-y: auto;
    box-sizing: border-box;
}

.inputs{
    width: 100%;
    display: flex;
    flex-direction: column;
}
.input{
    width: 100%;
    display: flex;
    flex-direction: column;
    padding-top: 10px;
    padding-bottom: 10px;
    box-sizing: border-box;
}
.data{
    width: 100%;
    display: flex;
    flex-direction: row;
}
.name{
    display: flex;
    width: 30%;
    height: 100%;
    text-align: right;
    align-items: center;
    justify-content: flex-end;
    padding-right: 20px;
    box-sizing: border-box;
    font-family: 'Arial';
    font-weight: bold;
    font-size: 20px;
}
.input-field{
    text-align: center;
    display: flex;
    width: 70%;
    border: none;
    outline: none;
    box-shadow: none;
    height: 30px;
    border-radius: 10px;
    background-color: var(--color-dark_gray);
}
.hint{
    color: var(--color-red);
}
</style>
