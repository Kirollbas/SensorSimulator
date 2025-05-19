<script setup>
import './../style.css';
import { reactive, defineProps, defineExpose, ref } from 'vue'
import { useForm, useField } from 'vee-validate'
import PrngInput from './PrngInput.vue'
import * as yup from 'yup'

const { resetForm, errors, validate, values, setErrors } = useForm({
    initialValues: {
        type: 'MODIFIER_TYPE_CONSTANT_OFFSET',
    },
})

const { value: type } = useField('type')

const { value: co_offset } = useField('co_offset')

const { value: h_percentage } = useField('h_percentage')

const { value: i_value } = useField('i_value')
const { value: i_period } = useField('i_period')

const { value: nl_coefficient } = useField('nl_coefficient')
const { value: nl_center } = useField('nl_center')

const { value: po_value } = useField('po_value')
const { value: po_interval } = useField('po_interval')

const { value: q_quant } = useField('q_quant')

const rad_prngRef = ref()
const { value: rad_min_add_value } = useField('rad_min_add_value')
const { value: rad_max_add_value } = useField('rad_max_add_value')
const { value: rad_min_dash_duration } = useField('rad_min_dash_duration')
const { value: rad_max_dash_duration } = useField('rad_max_dash_duration')
const { value: rad_avg_period } = useField('rad_avg_period')

const rfd_prngRef = ref()
const { value: rfd_value } = useField('rfd_value')
const { value: rfd_min_dash_duration } = useField('rfd_min_dash_duration')
const { value: rfd_max_dash_duration } = useField('rfd_max_dash_duration')
const { value: rfd_avg_period } = useField('rfd_avg_period')

const wn_prngRef = ref()
const { value: wn_max_value } = useField('wn_max_value')

const { value: d_simulator_name } = useField('d_simulator_name')
const { value: d_center } = useField('d_center')
const { value: d_coefficient } = useField('d_coefficient')

function getSchema(type) {
  const modifier = {
    type: yup.string().required('Выберите тип модификатора'),
  }

  if (type === 'MODIFIER_TYPE_CONSTANT_OFFSET') {
    modifier.co_offset = yup.number().required('Введите величину сдвига')
  } else if (type === 'MODIFIER_TYPE_HYSTERESIS') {
    modifier.h_percentage = yup.number().required('Введите процент гистерезиса').integer('Введите целое число')
  } else if (type === 'MODIFIER_TYPE_INERTIA') {
    modifier.i_value = yup.number().required('Введите максимальное изменение величины за период')
    modifier.i_period = yup.string().required("Введите период").matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Формат 123h45m6s')
  } else if (type === 'MODIFIER_TYPE_NONLINEAR_DEPENDENCE') {
    modifier.nl_coefficient = yup.number().required('Введите коэффициент нелинейности')
    modifier.nl_center = yup.number().required('Введите центровую точку нелинейности')
  } else if (type === 'MODIFIER_TYPE_PROGRESSING_OFFSET') {
    modifier.po_value = yup.number().required('Введите изменение величины за интервал')
    modifier.po_interval = yup.string().required("Введите интервал").matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Формат 123h45m6s')
  } else if (type === 'MODIFIER_TYPE_QUANTIZATION') {
    modifier.q_quant = yup.number().required('Введите квант')
  } else if (type === 'MODIFIER_TYPE_RANDOM_ADD_DASH') {
    modifier.rad_min_add_value = yup.number().required('Введите минимальную величину добавочного сдвига').max(yup.ref('rad_max_add_value'), "Минимальное значение должно быть меньше максимального")
    modifier.rad_max_add_value = yup.number().required('Введите максимальную величину добавочного сдвига')
    modifier.rad_min_dash_duration = yup.string().required("Введите минимальную длительность сдвига").matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Формат 123h45m6s')
    modifier.rad_max_dash_duration = yup.string().required("Введите максимальную длительность сдвига").matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Формат 123h45m6s')
    modifier.rad_avg_period = yup.string().required("Введите средний период возникновения сдвига").matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Формат 123h45m6s')
  } else if (type === 'MODIFIER_TYPE_RANDOM_FIXED_DASH') {
    modifier.rfd_value = yup.number().required('Введите значение сдвига')
    modifier.rfd_min_dash_duration = yup.string().required("Введите минимальную длительность сдвига").matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Формат 123h45m6s')
    modifier.rfd_max_dash_duration = yup.string().required("Введите максимальную длительность сдвига").matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Формат 123h45m6s')
    modifier.rfd_avg_period = yup.string().required("Введите средний период возникновения сдвига").matches(/^(\d+h)?(\d+m)?(\d+s)?$/, 'Формат 123h45m6s')
  } else if (type === 'MODIFIER_TYPE_WHITE_NOISE') {
    modifier.wn_max_value = yup.number().required('Введите максимальное значение шума')
  } else if (type === 'MODIFIER_TYPE_DEPENDENCE') {
    modifier.d_simulator_name = yup.string().required('Введите название симулятора для зависимости')
    modifier.d_center = yup.number().required('Введите центровую точку зависимости')
    modifier.d_coefficient = yup.number().required('Введите коэффициент зависимости')
  }

  return yup.object(modifier)
}

const getJson = async () => {
  const currentSchema = getSchema(type.value)

  try {
    const result = await currentSchema.validate(values, {
      abortEarly: false,
    })

    const currentType = type.value

    const modifier = {
      type: currentType,
      data: {}
    }

    if (currentType === 'MODIFIER_TYPE_CONSTANT_OFFSET') {
      modifier.data.offset = co_offset.value
      return modifier
    }

    if (currentType === 'MODIFIER_TYPE_HYSTERESIS') {
      modifier.data.percentage = h_percentage.value
      return modifier
    }

    if (currentType === 'MODIFIER_TYPE_INERTIA') {
      modifier.data.value = i_value.value
      modifier.data.period = i_period.value
      return modifier
    }

    if (currentType === 'MODIFIER_TYPE_NONLINEAR_DEPENDENCE') {
      modifier.data.coefficient = nl_coefficient.value
      modifier.data.center = nl_center.value
      return modifier
    }

    if (currentType === 'MODIFIER_TYPE_PROGRESSING_OFFSET') {
      modifier.data.value = po_value.value
      modifier.data.interval = po_interval.value
      return modifier
    }

    if (currentType === 'MODIFIER_TYPE_QUANTIZATION') {
      modifier.data.quant = q_quant.value
      return modifier
    }

    if (currentType === 'MODIFIER_TYPE_RANDOM_ADD_DASH') {
      const prngJson = await rad_prngRef.value?.getJson?.()

      if (prngJson == null) {
          return null
      }

      modifier.data.min_add_value = rad_min_add_value.value
      modifier.data.max_add_value = rad_max_add_value.value
      modifier.data.min_dash_duration = rad_min_dash_duration.value
      modifier.data.max_dash_duration = rad_max_dash_duration.value
      modifier.data.avg_period = rad_avg_period.value
      modifier.data.generator = prngJson
      return modifier
    }

    if (currentType === 'MODIFIER_TYPE_RANDOM_FIXED_DASH') {
      const prngJson = await rfd_prngRef.value?.getJson?.()

      if (prngJson == null) {
          return null
      }

      modifier.data.value = rfd_value.value
      modifier.data.min_dash_duration = rfd_min_dash_duration.value
      modifier.data.max_dash_duration = rfd_max_dash_duration.value
      modifier.data.avg_period = rfd_avg_period.value
      modifier.data.generator = prngJson
      return modifier
    }

    if (currentType === 'MODIFIER_TYPE_WHITE_NOISE') {
      const prngJson = await wn_prngRef.value?.getJson?.()
      
      if (prngJson == null) {
          return null
      }

      modifier.data.max_value = wn_max_value.value
      modifier.data.generator = prngJson
      return modifier
    }

    if (currentType === 'MODIFIER_TYPE_DEPENDENCE') {
      modifier.data.simulator_name = d_simulator_name.value
      modifier.data.center = d_center.value
      modifier.data.coefficient = d_coefficient.value
      return modifier
    }

    return modifier
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
    resetForm()
}

defineExpose({ 
    getJson,
    clearData
})
</script>

<template>
  <form @submit.prevent="onSubmit" class="modifier-panel">
    <label class="input">
      <div class="data">
        <div class="header-name">
          Тип
        </div>
          <select v-model="type" class="header-select">
          <option value="MODIFIER_TYPE_CONSTANT_OFFSET">Смещение</option>
          <option value="MODIFIER_TYPE_HYSTERESIS">Гистерезис</option>
          <option value="MODIFIER_TYPE_INERTIA">Инертность</option>
          <option value="MODIFIER_TYPE_NONLINEAR_DEPENDENCE">Нелинейность</option>
          <option value="MODIFIER_TYPE_PROGRESSING_OFFSET">Старение</option>
          <option value="MODIFIER_TYPE_QUANTIZATION">Квантование</option>
          <option value="MODIFIER_TYPE_RANDOM_ADD_DASH">Случайный добавочный сдвиг</option>
          <option value="MODIFIER_TYPE_RANDOM_FIXED_DASH">Случайный фиксированный сдвиг</option>
          <option value="MODIFIER_TYPE_WHITE_NOISE">Белый шум</option>
          <option value="MODIFIER_TYPE_DEPENDENCE">Зависимость</option>
        </select>
      </div>
      <div class="hint">
        {{ errors.type }}
      </div>
    </label>

    <template v-if="type === 'MODIFIER_TYPE_CONSTANT_OFFSET'">
      <label class="input">
        <div class="data">
          <div class="name">Смещение:</div>
          <input type="number" v-model="co_offset" class="input-field" />
        </div>
        <div class="hint">{{ errors.co_offset }}</div>
      </label>
    </template>

    <template v-if="type === 'MODIFIER_TYPE_HYSTERESIS'">
      <label class="input">
        <div class="data">
          <div class="name">Процент:</div>
          <input type="number" v-model="h_percentage" class="input-field" />
        </div>
        <div class="hint">{{ errors.h_percentage }}</div>
      </label>
    </template>

    <template v-if="type === 'MODIFIER_TYPE_INERTIA'">
      <label class="input">
        <div class="data">
          <div class="name">Изменение за период:</div>
          <input type="number" v-model="i_value" class="input-field" />
        </div>
        <div class="hint">{{ errors.i_value }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Период:</div>
          <input type="text" v-model="i_period" class="input-field" />
        </div>
        <div class="hint">{{ errors.i_period }}</div>
      </label>
    </template>

    <template v-if="type === 'MODIFIER_TYPE_NONLINEAR_DEPENDENCE'">
      <label class="input">
        <div class="data">
          <div class="name">Коэффициент:</div>
          <input type="number" v-model="nl_coefficient" class="input-field" />
        </div>
        <div class="hint">{{ errors.nl_coefficient }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Центр:</div>
          <input type="number" v-model="nl_center" class="input-field" />
        </div>
        <div class="hint">{{ errors.nl_center }}</div>
      </label>
    </template>

    <template v-if="type === 'MODIFIER_TYPE_PROGRESSING_OFFSET'">
      <label class="input">
        <div class="data">
          <div class="name">Изменение:</div>
          <input type="number" v-model="po_value" class="input-field" />
        </div>
        <div class="hint">{{ errors.po_value }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Интервал:</div>
          <input type="text" v-model="po_interval" class="input-field" />
        </div>
        <div class="hint">{{ errors.po_interval }}</div>
      </label>
    </template>

    <template v-if="type === 'MODIFIER_TYPE_QUANTIZATION'">
      <label class="input">
        <div class="data">
          <div class="name">Квант:</div>
          <input type="number" v-model="q_quant" class="input-field" />
        </div>
        <div class="hint">{{ errors.q_quant }}</div>
      </label>
    </template>

    <template v-if="type === 'MODIFIER_TYPE_RANDOM_ADD_DASH'">
      <label class="input">
        <div class="data">
          <div class="name">Мин. значение:</div>
          <input type="number" v-model="rad_min_add_value" class="input-field" />
        </div>
        <div class="hint">{{ errors.rad_min_add_value }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Макс. значение:</div>
          <input type="number" v-model="rad_max_add_value" class="input-field" />
        </div>
        <div class="hint">{{ errors.rad_max_add_value }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Мин. длительность:</div>
          <input type="text" v-model="rad_min_dash_duration" class="input-field" />
        </div>
        <div class="hint">{{ errors.rad_min_dash_duration }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Макс. длительность:</div>
          <input type="text" v-model="rad_max_dash_duration" class="input-field" />
        </div>
        <div class="hint">{{ errors.rad_max_dash_duration }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Средний период:</div>
          <input type="text" v-model="rad_avg_period" class="input-field" />
        </div>
        <div class="hint">{{ errors.rad_avg_period }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">
            Генератор:
          </div>
          <PrngInput  ref="rad_prngRef"/>
        </div>
      </label>
    </template>

    <template v-if="type === 'MODIFIER_TYPE_RANDOM_FIXED_DASH'">
      <label class="input">
        <div class="data">
          <div class="name">Значение:</div>
          <input type="number" v-model="rfd_value" class="input-field" />
        </div>
        <div class="hint">{{ errors.rfd_value }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Мин. длительность:</div>
          <input type="text" v-model="rfd_min_dash_duration" class="input-field" />
        </div>
        <div class="hint">{{ errors.rfd_min_dash_duration }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Макс. длительность:</div>
          <input type="text" v-model="rfd_max_dash_duration" class="input-field" />
        </div>
        <div class="hint">{{ errors.rfd_max_dash_duration }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Средний период:</div>
          <input type="text" v-model="rfd_avg_period" class="input-field" />
        </div>
        <div class="hint">{{ errors.rfd_avg_period }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">
            Генератор:
          </div>
          <PrngInput  ref="rfd_prngRef"/>
        </div>
      </label>
    </template>

    <template v-if="type === 'MODIFIER_TYPE_WHITE_NOISE'">
      <label class="input">
        <div class="data">
          <div class="name">Макс. значение:</div>
          <input type="number" v-model="wn_max_value" class="input-field" />
        </div>
        <div class="hint">{{ errors.wn_max_value }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">
            Генератор:
          </div>
          <PrngInput  ref="wn_prngRef"/>
        </div>
      </label>
    </template>

    <template v-if="type === 'MODIFIER_TYPE_DEPENDENCE'">
      <label class="input">
        <div class="data">
          <div class="name">Симулятор:</div>
          <input type="text" v-model="d_simulator_name" class="input-field" />
        </div>
        <div class="hint">{{ errors.d_simulator_name }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Центр:</div>
          <input type="number" v-model="d_center" class="input-field" />
        </div>
        <div class="hint">{{ errors.d_center }}</div>
      </label>
      <label class="input">
        <div class="data">
          <div class="name">Коэффициент:</div>
          <input type="number" v-model="d_coefficient" class="input-field" />
        </div>
        <div class="hint">{{ errors.d_coefficient }}</div>
      </label>
    </template>
  </form>
</template>

<style scoped>
.modifier-panel{
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
    width: 30%;
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
    width: 70%;
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