<script setup>
import './../style.css';
import { reactive, defineProps, defineExpose, ref } from 'vue'
import PrngInput from './PrngInput.vue'
import { useForm, useField } from 'vee-validate'
import * as yup from 'yup'

function getSchema(type) {
  const base = {
    type: yup.string().required('Выберите тип основы'),
  }

  if (type === 'BASE_TYPE_CONSTANT') {
    base.value = yup.number().required('Введите значение')
  } else if (['BASE_TYPE_SINEWAVE', 'BASE_TYPE_BEZIER', 'BASE_TYPE_LINEAR'].includes(type)) {
    base.min_value = yup.number().required("Введите минимальное значение").max(yup.ref('max_value'), "Минимальное значение должно быть меньше максимального")
    base.max_value = yup.number().required("Введите максимальное значение")
    base.min_period = yup.string().required("Введите минимальный период изменения значения").matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Формат 123h45m6s')
    base.max_period = yup.string().required("Введите максимальный период изменения значения").matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Формат 123h45m6s')
  }

  return yup.object(base)
}

const prngRef = ref()

const { resetForm, errors, validate, values, setErrors } = useForm({
    initialValues: {
        type: 'BASE_TYPE_SINEWAVE',
    },
})

const getJson = async () => {
  const currentSchema = getSchema(type.value)

  try {
    let prng = null

    if (type.value !== 'BASE_TYPE_CONSTANT') {
        prng = await prngRef.value.getJson()
    }

    const result = await currentSchema.validate(values, {
      abortEarly: false,
    })

    if (prng == null && !type.value === 'BASE_TYPE_CONSTANT') {
        return null
    }

    const data = { 
      type: type.value,
      data: {}
    }

    if (type.value === 'BASE_TYPE_CONSTANT') {
      data.data.value = value.value
    } else {
      data.data.generator = prng
      data.data.min_value = min_value.value
      data.data.max_value = max_value.value
      data.data.min_period = min_period.value
      data.data.max_period = max_period.value
    }

    return data
  } catch (error) {
    if (error.inner) {
      const formattedErrors = {}
      for (const err of error.inner) {
        if (!formattedErrors[err.path]) {
          formattedErrors[err.path] = err.message
        }
      }
      setErrors(formattedErrors)
    }

    return null
  }
}

const clearData = async() => {
    if (!(type.value === 'BASE_TYPE_CONSTANT')) {
        console.log(type.value)
        await prngRef.value.clearData()
    }
    resetForm()
}

const { value: type } = useField('type')
const { value: min_value } = useField('min_value')
const { value: max_value } = useField('max_value')
const { value: min_period } = useField('min_period')
const { value: max_period } = useField('max_period')
const { value: value } = useField('value')
value
defineExpose({ 
    getJson,
    clearData
})

</script>

<template>
  <form @submit.prevent="onSubmit" class="base-panel">
    <label class="input">
      <div class="data">
        <div class="header-name">
          Основа
        </div>
          <select v-model="type" class="header-select">
          <option value="BASE_TYPE_SINEWAVE">Синусоида</option>
          <option value="BASE_TYPE_CONSTANT">Значение</option>
          <option value="BASE_TYPE_BEZIER">Безье</option>
          <option value="BASE_TYPE_LINEAR">Линейная</option>
        </select>
      </div>
      <div class="hint">
        {{ errors.type }}
      </div>
    </label>
    <label class="input" v-if="['BASE_TYPE_SINEWAVE', 'BASE_TYPE_BEZIER', 'BASE_TYPE_LINEAR'].includes(type)">
      <div class="data">
        <div class="name">
          Мин. знач.:
        </div>
        <input type="number" v-model="min_value" class="input-field"/>
      </div>
      <div class="hint">
        {{ errors.min_value }}
      </div>
    </label>
    <label class="input" v-if="['BASE_TYPE_SINEWAVE', 'BASE_TYPE_BEZIER', 'BASE_TYPE_LINEAR'].includes(type)">
      <div class="data">
        <div class="name">
          Макс. знач.:
        </div>
        <input type="number" v-model="max_value" class="input-field"/>
      </div>
      <div class="hint">
        {{ errors.max_value }}
      </div>
    </label>
    <label class="input" v-if="['BASE_TYPE_SINEWAVE', 'BASE_TYPE_BEZIER', 'BASE_TYPE_LINEAR'].includes(type)">
      <div class="data">
        <div class="name">
          Мин. период:
        </div>
        <input type="string" v-model="min_period" class="input-field"/>
      </div>
      <div class="hint">
        {{ errors.min_period }}
      </div>
    </label>
    <label class="input" v-if="['BASE_TYPE_SINEWAVE', 'BASE_TYPE_BEZIER', 'BASE_TYPE_LINEAR'].includes(type)">
      <div class="data">
        <div class="name">
          Макс. период:
        </div>
        <input type="string" v-model="max_period" class="input-field"/>
      </div>
      <div class="hint">
        {{ errors.max_period }}
      </div>
    </label>
    <label class="input" v-if="['BASE_TYPE_SINEWAVE', 'BASE_TYPE_BEZIER', 'BASE_TYPE_LINEAR'].includes(type)">
      <div class="data">
        <div class="name">
          Генератор:
        </div>
        <PrngInput  ref="prngRef"/>
      </div>
    </label>
    <label class="input" v-if="type === 'BASE_TYPE_CONSTANT'">
      <div class="data">
        <div class="name">
          Значение:
        </div>
        <input type="number" v-model="value" class="input-field"/>
      </div>
      <div class="hint">
        {{ errors.value }}
      </div>
    </label>
  </form>
</template>

<style scoped>
.base-panel{
    margin-top: 10px;
    display: flex;
    width: 100%;
    background-color: var(--color-dark_gray);
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
    background-color: var(--color-gray);
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
    background-color: var(--color-gray);
    font-family: 'Arial';
    font-weight: bold;
    font-size: 14px;
}
.hint{
    color: var(--color-red);
}
</style>