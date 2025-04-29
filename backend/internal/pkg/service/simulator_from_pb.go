package service

import (
	"fmt"
	pb "sensor-simulator/gen/sensor_simulator/proto/simulator"
	"sensor-simulator/internal/pkg/domain/base"
	"sensor-simulator/internal/pkg/domain/generator"
	"sensor-simulator/internal/pkg/domain/modifier"
	"sensor-simulator/internal/pkg/domain/simulator"
	"time"
)

const frequency = time.Second / 100

var (
	startGen = generator.NewPcg(0)
)

type Generator interface {
	NextZeroToOne() float64
	Restart()
}

func (s *SimulatorService) simulatorFromPb(proto *pb.Simulator) (newSimulator *simulator.Simulator, newDependencies []string, err error) {
	base, err := s.baseFromPb(proto.Base)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create simulator base. Err: %w", err)
	}

	var startValue float64

	switch proto.Base.TypeData.(type) {
	case *pb.Base_Common:
		startValue = proto.Base.TypeData.(*pb.Base_Common).Common.GetMaxValue() -
			proto.Base.TypeData.(*pb.Base_Common).Common.GetMinValue()
	case *pb.Base_Constant:
		startValue = proto.Base.TypeData.(*pb.Base_Constant).Constant.GetValue()
	}

	modifiers := []simulator.Modifier{}
	dependencies := []string{}

	for _, protoModifier := range proto.Modifiers {
		modifier, newDependencies, err := s.modifierFromPb(
			startValue,
			base,
			protoModifier,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to create simulator modifier. Err: %w", err)
		}

		modifiers = append(modifiers, modifier)
		dependencies = append(dependencies, newDependencies...)
	}

	simulator, err := simulator.NewSimulator(
		proto.Name,
		uint16(proto.Address),
		uint64(time.Second/frequency),
		base,
		modifiers,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create simulator. Err: %w", err)
	}

	return simulator, dependencies, nil
}

func (s *SimulatorService) modifierFromPb(
	startValue float64,
	base simulator.Base,
	proto *pb.Modifier,
) (newModifier simulator.Modifier, newDependencies []string, err error) {
	switch proto.Type {
	case pb.ModifierType_MODIFIER_TYPE_CONSTANT_OFFSET:
		typeData, ok := proto.TypeData.(*pb.Modifier_ConstOffset)
		if !ok {
			return nil, nil, fmt.Errorf("unable to parse constant offset modifier data")
		}

		newModifier, err = modifier.NewConstantOffsetModifier(typeData.ConstOffset.GetOffset())
		return newModifier, nil, err

	case pb.ModifierType_MODIFIER_TYPE_HYSTERESIS:
		typeData, ok := proto.TypeData.(*pb.Modifier_Hysteresis)
		if !ok {
			return nil, nil, fmt.Errorf("unable to parse hysteresis modifier data")
		}

		hysteresis, err := modifier.NewHysteresisModifier(typeData.Hysteresis.GetPercentage())
		if err != nil {
			return nil, nil, err
		}

		base.AddStateSubscriber(hysteresis)

		return hysteresis, nil, nil

	case pb.ModifierType_MODIFIER_TYPE_INERTIA:
		typeData, ok := proto.TypeData.(*pb.Modifier_Inertia)
		if !ok {
			return nil, nil, fmt.Errorf("unable to parse hysteresis modifier data")
		}

		speed := typeData.Inertia.GetValue() * float64(typeData.Inertia.GetPeriod().AsDuration()/frequency)

		newModifier, err = modifier.NewInertiaModifier(speed, startValue)
		return newModifier, nil, err

	case pb.ModifierType_MODIFIER_TYPE_NONLINEAR_DEPENDENCE:
		typeData, ok := proto.TypeData.(*pb.Modifier_NonLinearDependance)
		if !ok {
			return nil, nil, fmt.Errorf("unable to parse nonlinear dependance modifier data")
		}

		newModifier, err = modifier.NewNonLinearModifier(
			typeData.NonLinearDependance.GetCoefficient(),
			typeData.NonLinearDependance.GetCenter(),
		)
		return newModifier, nil, err

	case pb.ModifierType_MODIFIER_TYPE_PROGRESSING_OFFSET:
		typeData, ok := proto.TypeData.(*pb.Modifier_ProgressingOffset)
		if !ok {
			return nil, nil, fmt.Errorf("unable to parse progressing offset modifier data")
		}

		interval := uint64(typeData.ProgressingOffset.GetInterval().AsDuration() / frequency)

		newModifier, err = modifier.NewProgressingOffsetModifier(
			typeData.ProgressingOffset.GetValue(),
			interval,
		)
		return newModifier, nil, err

	case pb.ModifierType_MODIFIER_TYPE_QUANTIZATION:
		typeData, ok := proto.TypeData.(*pb.Modifier_Quantization)
		if !ok {
			return nil, nil, fmt.Errorf("unable to parse quantization modifier data")
		}

		newModifier, err = modifier.NewQuantizationModifier(
			typeData.Quantization.GetQuant(),
		)
		return newModifier, nil, err

	case pb.ModifierType_MODIFIER_TYPE_RANDOM_ADD_DASH:
		typeData, ok := proto.TypeData.(*pb.Modifier_RandomAddDash)
		if !ok {
			return nil, nil, fmt.Errorf("unable to parse random add dash modifier data")
		}

		prng, err := s.genaratorFromPb(typeData.RandomAddDash.Generator)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to create generator for random add dash modifier. Err: %w", err)
		}

		minDashTicks := uint64(typeData.RandomAddDash.GetMinDashDuration().AsDuration() / frequency)
		maxDashTicks := uint64(typeData.RandomAddDash.GetMaxDashDuration().AsDuration() / frequency)
		avgDashTicks := uint64(typeData.RandomAddDash.GetAvgPeriod().AsDuration() / frequency)

		newModifier, err = modifier.NewRandomAddDashModifier(
			prng,
			maxDashTicks,
			minDashTicks,
			avgDashTicks,
			typeData.RandomAddDash.GetMinAddValue(),
			typeData.RandomAddDash.GetMaxAddValue(),
		)
		return newModifier, nil, err

	case pb.ModifierType_MODIFIER_TYPE_RANDOM_FIXED_DASH:
		typeData, ok := proto.TypeData.(*pb.Modifier_RandomFixedDash)
		if !ok {
			return nil, nil, fmt.Errorf("unable to parse random fixed dash modifier data")
		}

		prng, err := s.genaratorFromPb(typeData.RandomFixedDash.Generator)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to create generator for random fixed dash modifier. Err: %w", err)
		}

		minDashTicks := uint64(typeData.RandomFixedDash.GetMinDashDuration().AsDuration() / frequency)
		maxDashTicks := uint64(typeData.RandomFixedDash.GetMaxDashDuration().AsDuration() / frequency)
		avgDashTicks := uint64(typeData.RandomFixedDash.GetAvgPeriod().AsDuration() / frequency)

		newModifier, err = modifier.NewRandomFixedDashModifier(
			prng,
			typeData.RandomFixedDash.GetValue(),
			maxDashTicks,
			minDashTicks,
			avgDashTicks,
		)
		return newModifier, nil, err

	case pb.ModifierType_MODIFIER_TYPE_WHITE_NOISE:
		typeData, ok := proto.TypeData.(*pb.Modifier_WhiteNoise)
		if !ok {
			return nil, nil, fmt.Errorf("unable to parse white noise modifier data")
		}

		prng, err := s.genaratorFromPb(typeData.WhiteNoise.Generator)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to create generator for white noise modifier. Err: %w", err)
		}

		newModifier, err = modifier.NewWhiteNoiseModifier(
			prng,
			typeData.WhiteNoise.GetMaxValue(),
		)
		return newModifier, nil, err

	case pb.ModifierType_MODIFIER_TYPE_DEPENDENCE:
		typeData, ok := proto.TypeData.(*pb.Modifier_Dependence)
		if !ok {
			return nil, nil, fmt.Errorf("unable to parse dependence modifier data")
		}

		if _, ok := s.simulators[SimulatorName(typeData.Dependence.GetSimulatorName())]; !ok {
			return nil, nil, fmt.Errorf("there is no simulator with name %s to crate dependence modifier", typeData.Dependence.GetSimulatorName())
		}

		simulator := s.simulators[SimulatorName(typeData.Dependence.GetSimulatorName())]

		newModifier, err = modifier.NewDependenceModifier(
			simulator,
			typeData.Dependence.GetCoefficient(),
			typeData.Dependence.GetCenter(),
		)
		return newModifier, []string{typeData.Dependence.GetSimulatorName()}, err

	case pb.ModifierType_MODIFIER_TYPE_UNKNOWN:
		return nil, nil, fmt.Errorf("unknown modifier type")
	}

	return nil, nil, fmt.Errorf("unknown modifier type")
}

func (s *SimulatorService) baseFromPb(proto *pb.Base) (simulator.Base, error) {
	switch proto.Type {
	case pb.BaseType_BASE_TYPE_BEZIER:
		typeData, ok := proto.TypeData.(*pb.Base_Common)
		if !ok {
			return nil, fmt.Errorf("unable to parse bezier base data")
		}

		prng, err := s.genaratorFromPb(typeData.Common.GetGenerator())
		if err != nil {
			return nil, fmt.Errorf("unable to create generator for bezier base. Err: %w", err)
		}

		minTicks := uint64(typeData.Common.GetMinPeriod().AsDuration() / frequency)
		maxTicks := uint64(typeData.Common.GetMaxPeriod().AsDuration() / frequency)

		return base.NewBezierSimulator(
			prng,
			typeData.Common.GetMinValue(),
			typeData.Common.GetMaxValue(),
			minTicks,
			maxTicks,
		)
	case pb.BaseType_BASE_TYPE_CONSTANT:
		typeData, ok := proto.TypeData.(*pb.Base_Constant)
		if !ok {
			return nil, fmt.Errorf("unable to parse constant base data")
		}

		return base.NewConstantSimulator(
			typeData.Constant.GetValue(),
		), nil

	case pb.BaseType_BASE_TYPE_LINEAR:
		typeData, ok := proto.TypeData.(*pb.Base_Common)
		if !ok {
			return nil, fmt.Errorf("unable to parse linear base data")
		}

		prng, err := s.genaratorFromPb(typeData.Common.GetGenerator())
		if err != nil {
			return nil, fmt.Errorf("unable to create generator for linear base. Err: %w", err)
		}

		minTicks := uint64(typeData.Common.GetMinPeriod().AsDuration() / frequency)
		maxTicks := uint64(typeData.Common.GetMaxPeriod().AsDuration() / frequency)

		return base.NewLinearSimulator(
			prng,
			typeData.Common.GetMinValue(),
			typeData.Common.GetMaxValue(),
			minTicks,
			maxTicks,
		)

	case pb.BaseType_BASE_TYPE_SINEWAVE:
		typeData, ok := proto.TypeData.(*pb.Base_Common)
		if !ok {
			return nil, fmt.Errorf("unable to parse sinewave base data")
		}

		prng, err := s.genaratorFromPb(typeData.Common.GetGenerator())
		if err != nil {
			return nil, fmt.Errorf("unable to create generator for sinewave base. Err: %w", err)
		}

		minTicks := uint64(typeData.Common.GetMinPeriod().AsDuration() / frequency)
		maxTicks := uint64(typeData.Common.GetMaxPeriod().AsDuration() / frequency)

		return base.NewSineWaveSimulator(
			prng,
			typeData.Common.GetMinValue(),
			typeData.Common.GetMaxValue(),
			minTicks,
			maxTicks,
		)

	case pb.BaseType_BASE_TYPE_UNKNOWN:
		return nil, fmt.Errorf("unknown base type")
	}

	return nil, fmt.Errorf("unknown base type")
}

func (s *SimulatorService) genaratorFromPb(proto *pb.Prng) (Generator, error) {
	seed := proto.Seed

	if seed == -1 {
		seed = startGen.NextInt()
	}

	switch proto.Type {
	case pb.PrngType_PRNG_TYPE_PCG:
		return generator.NewPcg(seed), nil
	case pb.PrngType_PRNG_TYPE_XOSHIRO:
		return generator.NewXoshiro(seed), nil
	case pb.PrngType_PRNG_TYPE_UNKNOWN:
		return nil, fmt.Errorf("unknown generator type")
	}

	return nil, fmt.Errorf("unknown generator type")
}
