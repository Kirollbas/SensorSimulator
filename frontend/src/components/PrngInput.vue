<script setup>
import './../style.css';
import { reactive, defineProps, defineExpose, ref } from 'vue'
import { useForm, useField } from 'vee-validate'
import * as yup from 'yup'

const schema = yup.object({
  type: yup.string()
    .required('Выберите тип'),
  seed: yup.number()
    .required()
    .integer('Только целое число')
    .min(-1, 'Введите положительное число, или -1 для случайного сида')
})

const { resetForm, errors, validate, values } = useForm({ 
    validationSchema: schema,
    initialValues: {
        type: 'PRNG_TYPE_XOSHIRO',
        seed: -1
    },
})

const { value: type } = useField('type')
const { value: seed } = useField('seed')

const getJson = async () => {
  const result = await validate()

  if (!result.valid) {
    return null
  }

  const data = JSON.parse(JSON.stringify(values))

  return data
}

const clearData = async() => {
    resetForm()
}

defineExpose({ 
    getJson,
    clearData
})


</script>

<template>
  <form @submit.prevent="onSubmit" class="prng-panel">
    <label class="input">
      <div class="data">
        <div class="header-name">
          Тип
        </div>
          <select v-model="type" class="header-select">
          <option value="PRNG_TYPE_PCG">PCG</option>
          <option value="PRNG_TYPE_XOSHIRO">Xoshiro</option>
        </select>
      </div>
      <div class="hint">
        {{ errors.type }}
      </div>
    </label>
    <label class="input">
      <div class="data">
        <div class="name">
          seed:
        </div>
        <input type="number" step="1" v-model="seed" class="input-field"/>
      </div>
      <div class="hint">
        {{ errors.seed }}
      </div>
    </label>
    </form>
</template>

<style scoped>
.prng-panel{
    margin-top: 10px;
    display: flex;
    width: 100%;
    background-color: var(--color-gray);
    flex-direction: column;
    padding: 10px;
    box-sizing: border-box;
    border-radius: 10px;
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
.header-name{
    display: flex;
    width: 70%;
    height: 100%;
    align-items: center;
    padding-right: 20px;
    padding-left: 20px;
    box-sizing: border-box;
    font-family: 'Arial';
    font-weight: bold;
    font-size: 20px;
}
.header-select{
    text-align: center;
    display: flex;
    width: 30%;
    border: none;
    outline: none;
    box-shadow: none;
    justify-content: flex-end;
    height: 30px;
    border-radius: 10px;
    background-color: var(--color-dark_gray);
    font-family: 'Arial';
    font-weight: bold;
    font-size: 14px;
}
.hint{
    color: var(--color-red);
}
</style>