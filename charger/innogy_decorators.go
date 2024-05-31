package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateInnogy(base *Innogy, meterEnergy func() (float64, error), phaseVoltages func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meterEnergy == nil && phaseVoltages == nil:
		return base

	case meterEnergy != nil && phaseVoltages == nil:
		return &struct {
			*Innogy
			api.MeterEnergy
		}{
			Innogy: base,
			MeterEnergy: &decorateInnogyMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meterEnergy == nil && phaseVoltages != nil:
		return &struct {
			*Innogy
			api.PhaseVoltages
		}{
			Innogy: base,
			PhaseVoltages: &decorateInnogyPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case meterEnergy != nil && phaseVoltages != nil:
		return &struct {
			*Innogy
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Innogy: base,
			MeterEnergy: &decorateInnogyMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateInnogyPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}
	}

	return nil
}

type decorateInnogyMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateInnogyMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateInnogyPhaseVoltagesImpl struct {
	phaseVoltages func() (float64, float64, float64, error)
}

func (impl *decorateInnogyPhaseVoltagesImpl) Voltages() (float64, float64, float64, error) {
	return impl.phaseVoltages()
}