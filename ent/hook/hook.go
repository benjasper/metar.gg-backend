// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"metar.gg/ent"
)

// The AirportFunc type is an adapter to allow the use of ordinary
// function as Airport mutator.
type AirportFunc func(context.Context, *ent.AirportMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AirportFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AirportMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AirportMutation", m)
}

// The CountryFunc type is an adapter to allow the use of ordinary
// function as Country mutator.
type CountryFunc func(context.Context, *ent.CountryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CountryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CountryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CountryMutation", m)
}

// The ForecastFunc type is an adapter to allow the use of ordinary
// function as Forecast mutator.
type ForecastFunc func(context.Context, *ent.ForecastMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ForecastFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ForecastMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ForecastMutation", m)
}

// The FrequencyFunc type is an adapter to allow the use of ordinary
// function as Frequency mutator.
type FrequencyFunc func(context.Context, *ent.FrequencyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FrequencyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FrequencyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FrequencyMutation", m)
}

// The IcingConditionFunc type is an adapter to allow the use of ordinary
// function as IcingCondition mutator.
type IcingConditionFunc func(context.Context, *ent.IcingConditionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f IcingConditionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.IcingConditionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.IcingConditionMutation", m)
}

// The MetarFunc type is an adapter to allow the use of ordinary
// function as Metar mutator.
type MetarFunc func(context.Context, *ent.MetarMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MetarFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MetarMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MetarMutation", m)
}

// The RegionFunc type is an adapter to allow the use of ordinary
// function as Region mutator.
type RegionFunc func(context.Context, *ent.RegionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RegionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RegionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RegionMutation", m)
}

// The RunwayFunc type is an adapter to allow the use of ordinary
// function as Runway mutator.
type RunwayFunc func(context.Context, *ent.RunwayMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RunwayFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RunwayMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RunwayMutation", m)
}

// The SkyConditionFunc type is an adapter to allow the use of ordinary
// function as SkyCondition mutator.
type SkyConditionFunc func(context.Context, *ent.SkyConditionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SkyConditionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SkyConditionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SkyConditionMutation", m)
}

// The TafFunc type is an adapter to allow the use of ordinary
// function as Taf mutator.
type TafFunc func(context.Context, *ent.TafMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TafFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TafMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TafMutation", m)
}

// The TemperatureDataFunc type is an adapter to allow the use of ordinary
// function as TemperatureData mutator.
type TemperatureDataFunc func(context.Context, *ent.TemperatureDataMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TemperatureDataFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TemperatureDataMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TemperatureDataMutation", m)
}

// The TurbulenceConditionFunc type is an adapter to allow the use of ordinary
// function as TurbulenceCondition mutator.
type TurbulenceConditionFunc func(context.Context, *ent.TurbulenceConditionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TurbulenceConditionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TurbulenceConditionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TurbulenceConditionMutation", m)
}

// The WeatherStationFunc type is an adapter to allow the use of ordinary
// function as WeatherStation mutator.
type WeatherStationFunc func(context.Context, *ent.WeatherStationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WeatherStationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WeatherStationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WeatherStationMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
