export const MODIFIER_TO_NAME = {
    "MODIFIER_TYPE_CONSTANT_OFFSET": "Сдвиг",
    "MODIFIER_TYPE_HYSTERESIS": "Гистерезис",
    "MODIFIER_TYPE_INERTIA": "Инертность",
    "MODIFIER_TYPE_NONLINEAR_DEPENDENCE": "Нелинейность",
    "MODIFIER_TYPE_PROGRESSING_OFFSET": "Прогрессирующий сдвиг",
    "MODIFIER_TYPE_QUANTIZATION": "Квантование",
    "MODIFIER_TYPE_RANDOM_ADD_DASH": "Случайный добавочный сдвиг",
    "MODIFIER_TYPE_RANDOM_FIXED_DASH": "Случайный фиксированный сдвиг",
    "MODIFIER_TYPE_WHITE_NOISE": "Белый шум",
    "MODIFIER_TYPE_DEPENDENCE": "Зависимость",
}

export const MODIFIER_DATA_TO_DATA_NAME = {
    "MODIFIER_TYPE_CONSTANT_OFFSET": {
        "offset": "Значение"
    },
    "MODIFIER_TYPE_HYSTERESIS": {
        "percentage": "Процент"
    },
    "MODIFIER_TYPE_INERTIA": {
        "value": "Макс. изменение за период",
        "period": "Период"
    },
    "MODIFIER_TYPE_NONLINEAR_DEPENDENCE": {
        "coefficient": "Коэффициент",
        "center": "Центровая точка"
    },
    "MODIFIER_TYPE_PROGRESSING_OFFSET": {
        "value": "Изменение за перод",
        "interval": "Период"
    },
    "MODIFIER_TYPE_QUANTIZATION": {
        "quant": "Квант"
    },
    "MODIFIER_TYPE_RANDOM_ADD_DASH": {
        "min_add_value": "Мин. добавочное значение",
        "max_add_value": "Макс. добавочное значение",
        "min_dash_duration": "Мин. длительность сдвига",
        "max_dash_duration": "Макс. длительность сдвига",
        "avg_period": "Период возникновения сдвига"
    },
    "MODIFIER_TYPE_RANDOM_FIXED_DASH": {
        "value": "Значение сдвига",
        "min_dash_duration": "Мин. длительность сдвига",
        "max_dash_duration": "Макс. длительность сдвига",
        "avg_period": "Период возникновения сдвига"
    },
    "MODIFIER_TYPE_WHITE_NOISE": {
        "max_value": "Максимальное значение"
    },
    "MODIFIER_TYPE_DEPENDENCE": {
        "simulator_name": "Зависим от", 
        "center": "Центровая точка",
        "coefficient": "Коэффициент"
    },
}
