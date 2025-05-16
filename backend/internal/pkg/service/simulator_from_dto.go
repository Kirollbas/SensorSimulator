package service

import (
	"fmt"
	"sensor-simulator/internal/configs"
	"sensor-simulator/internal/pkg/domain/base"
	"sensor-simulator/internal/pkg/domain/generator"
	"sensor-simulator/internal/pkg/domain/modifier"
	"sensor-simulator/internal/pkg/domain/simulator"
	"sensor-simulator/internal/pkg/dto"
	"sensor-simulator/internal/pkg/utils"
	"time"
)

type Generator interface {
	NextZeroToOne() float64
	Restart()
	ToDTO() dto.Prng
}

func (s *SimulatorService) simulatorFromDTO(dto dto.Simulator) (newSimulator *simulator.Simulator, newDependencies []string, err error) {
	base, startValue, err := s.baseFromDTO(dto.Base)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create simulator base. Err: %w", err)
	}

	modifiers := []simulator.Modifier{}
	dependencies := []string{}

	for _, modifierDTO := range dto.Modifiers {
		modifier, newDependencies, err := s.modifierFromDTO(
			startValue,
			base,
			modifierDTO,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to create simulator modifier. Err: %w", err)
		}

		modifiers = append(modifiers, modifier)
		dependencies = append(dependencies, newDependencies...)
	}

	frequency := configs.GetConfig().Simulator.Frequency

	simulator, err := simulator.NewSimulator(
		dto.Name,
		uint16(dto.Address),
		uint64(frequency),
		dto.Duration.Duration,
		base,
		modifiers,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create simulator. Err: %w", err)
	}

	return simulator, dependencies, nil
}

func (s *SimulatorService) modifierFromDTO(
	startValue float64,
	base simulator.Base,
	modifierDto dto.Modifier,
) (newModifier simulator.Modifier, newDependencies []string, err error) {
	tickPeriod := time.Second / time.Duration(configs.GetConfig().Simulator.Frequency)

	switch modifierDto.Type {
	case dto.ModifierTypeConstantOffset:
		typeData, err := utils.Reunmarshal[dto.ConstantOffsetModifier](modifierDto.Data)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse constant offset modifier data")
		}

		newModifier, err = modifier.NewConstantOffsetModifier(typeData.Offset)
		return newModifier, nil, err

	case dto.ModifierTypeHysteresis:
		typeData, err := utils.Reunmarshal[dto.HysteresisModifier](modifierDto.Data)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse hysteresis modifier data")
		}

		hysteresis, err := modifier.NewHysteresisModifier(uint64(typeData.Percentage))
		if err != nil {
			return nil, nil, err
		}

		base.AddStateSubscriber(hysteresis)

		return hysteresis, nil, nil

	case dto.ModifierTypeInertia:
		typeData, err := utils.Reunmarshal[dto.InertitaModifier](modifierDto.Data)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse inertia modifier data")
		}

		newModifier, err = modifier.NewInertiaModifier(
			typeData.Value,
			tickPeriod,
			typeData.Period.Duration,
			startValue,
		)
		return newModifier, nil, err

	case dto.ModifierTypeNonLinearDependence:
		typeData, err := utils.Reunmarshal[dto.NonLinearDependenceModifier](modifierDto.Data)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse nonlinear dependance modifier data")
		}

		newModifier, err = modifier.NewNonLinearModifier(
			typeData.Coefficient,
			typeData.Center,
		)
		return newModifier, nil, err

	case dto.ModifierTypeProgressingOffset:
		typeData, err := utils.Reunmarshal[dto.ProgressingOffsetModifier](modifierDto.Data)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse progressing offset modifier data")
		}

		newModifier, err = modifier.NewProgressingOffsetModifier(
			typeData.Value,
			typeData.Interval.Duration,
			tickPeriod,
		)
		return newModifier, nil, err

	case dto.ModifierTypeQuantization:
		typeData, err := utils.Reunmarshal[dto.QuantizationModifier](modifierDto.Data)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse quantization modifier data")
		}

		newModifier, err = modifier.NewQuantizationModifier(
			typeData.Quant,
		)
		return newModifier, nil, err

	case dto.ModifierTypeRandomAddDash:
		typeData, err := utils.Reunmarshal[dto.RandomAddDashModifier](modifierDto.Data)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse random add dash modifier data")
		}

		prng, err := s.genaratorFromDTO(typeData.Generator)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to create generator for random add dash modifier. Err: %w", err)
		}

		newModifier, err = modifier.NewRandomAddDashModifier(
			prng,
			typeData.MaxDashDuration.Duration,
			typeData.MinDashDuration.Duration,
			typeData.AvgPeriod.Duration,
			typeData.MinAddValue,
			typeData.MaxAddValue,
			tickPeriod,
		)
		return newModifier, nil, err

	case dto.ModifierTypeRandomFixedDash:
		typeData, err := utils.Reunmarshal[dto.RandomFixedDashModifier](modifierDto.Data)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse random fixed dash modifier data")
		}

		prng, err := s.genaratorFromDTO(typeData.Generator)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to create generator for random fixed dash modifier. Err: %w", err)
		}

		newModifier, err = modifier.NewRandomFixedDashModifier(
			prng,
			typeData.Value,
			typeData.MaxDashDuration.Duration,
			typeData.MinDashDuration.Duration,
			typeData.AvgPeriod.Duration,
			tickPeriod,
		)
		return newModifier, nil, err

	case dto.ModifierTypeWhiteNoise:
		typeData, err := utils.Reunmarshal[dto.WhiteNoiseModifier](modifierDto.Data)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse white noise modifier data")
		}

		prng, err := s.genaratorFromDTO(typeData.Generator)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to create generator for white noise modifier. Err: %w", err)
		}

		newModifier, err = modifier.NewWhiteNoiseModifier(
			prng,
			typeData.MaxValue,
		)
		return newModifier, nil, err

	case dto.ModifierTypeDependence:
		typeData, err := utils.Reunmarshal[dto.DependenceModifier](modifierDto.Data)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse dependence modifier data")
		}

		if _, ok := s.simulators[SimulatorName(typeData.SimulatorName)]; !ok {
			return nil, nil, fmt.Errorf("there is no simulator with name %s to crate dependence modifier", typeData.SimulatorName)
		}

		simulator := s.simulators[SimulatorName(typeData.SimulatorName)]

		newModifier, err = modifier.NewDependenceModifier(
			simulator,
			typeData.Coefficient,
			typeData.Center,
		)
		return newModifier, []string{typeData.SimulatorName}, err
	}

	return nil, nil, fmt.Errorf("unknown modifier type")
}

func (s *SimulatorService) baseFromDTO(baseDTO dto.Base) (res simulator.Base, startValue float64, err error) {
	tickPeriod := time.Second / time.Duration(configs.GetConfig().Simulator.Frequency)

	switch baseDTO.Type {
	case dto.BaseTypeBezier:
		typeData, err := utils.Reunmarshal[dto.CommonBase](baseDTO.Data)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to parse bezier base data")
		}

		prng, err := s.genaratorFromDTO(typeData.Generator)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to create generator for bezier base. Err: %w", err)
		}

		minValue := typeData.MinValue
		maxValue := typeData.MaxValue

		newBase, err := base.NewBezierSimulator(
			prng,
			minValue,
			maxValue,
			typeData.MinPeriod.Duration,
			typeData.MaxPeriod.Duration,
			tickPeriod,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to create bezier base. Err: %w", err)
		}

		return newBase, (maxValue - minValue) / 2, nil

	case dto.BaseTypeConstant:
		typeData, err := utils.Reunmarshal[dto.ConstantBase](baseDTO.Data)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to parse constant base data")
		}

		return base.NewConstantSimulator(
			typeData.Value,
		), typeData.Value, nil

	case dto.BaseTypeLinear:
		typeData, err := utils.Reunmarshal[dto.CommonBase](baseDTO.Data)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to parse linear base data")
		}

		prng, err := s.genaratorFromDTO(typeData.Generator)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to create generator for linear base. Err: %w", err)
		}

		minValue := typeData.MinValue
		maxValue := typeData.MaxValue

		newBase, err := base.NewLinearSimulator(
			prng,
			minValue,
			maxValue,
			typeData.MinPeriod.Duration,
			typeData.MaxPeriod.Duration,
			tickPeriod,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to create linear base. Err: %w", err)
		}

		return newBase, (maxValue - minValue) / 2, nil

	case dto.BaseTypeSinewave:
		typeData, err := utils.Reunmarshal[dto.CommonBase](baseDTO.Data)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to parse sinewave base data")
		}

		prng, err := s.genaratorFromDTO(typeData.Generator)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to create generator for sinewave base. Err: %w", err)
		}

		minValue := typeData.MinValue
		maxValue := typeData.MaxValue

		newBase, err := base.NewSineWaveSimulator(
			prng,
			minValue,
			maxValue,
			typeData.MinPeriod.Duration,
			typeData.MaxPeriod.Duration,
			tickPeriod,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to create sinewave base. Err: %w", err)
		}

		return newBase, (maxValue - minValue) / 2, nil
	}

	return nil, 0, fmt.Errorf("unknown base type")
}

func (s *SimulatorService) genaratorFromDTO(prngDTO dto.Prng) (Generator, error) {
	seed := prngDTO.Seed

	if seed == -1 {
		seed = configs.GetSeed()
	}

	switch prngDTO.Type {
	case dto.PRNGTypePCG:
		return generator.NewPcg(seed), nil
	case dto.PRNGTypeXoshiro:
		return generator.NewXoshiro(seed), nil
	}

	return nil, fmt.Errorf("unknown generator type")
}
