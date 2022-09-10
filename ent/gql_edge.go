// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Airport) Station(ctx context.Context) (*Station, error) {
	result, err := a.Edges.StationOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryStation().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Airport) Frequencies(ctx context.Context) ([]*Frequency, error) {
	result, err := a.NamedFrequencies(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = a.QueryFrequencies().All(ctx)
	}
	return result, err
}

func (f *Frequency) Airport(ctx context.Context) (*Airport, error) {
	result, err := f.Edges.AirportOrErr()
	if IsNotLoaded(err) {
		result, err = f.QueryAirport().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (m *Metar) Station(ctx context.Context) (*Station, error) {
	result, err := m.Edges.StationOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryStation().Only(ctx)
	}
	return result, err
}

func (m *Metar) SkyConditions(ctx context.Context) ([]*SkyCondition, error) {
	result, err := m.NamedSkyConditions(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = m.QuerySkyConditions().All(ctx)
	}
	return result, err
}

func (r *Runway) Airport(ctx context.Context) (*Airport, error) {
	result, err := r.Edges.AirportOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryAirport().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (sc *SkyCondition) Metar(ctx context.Context) (*Metar, error) {
	result, err := sc.Edges.MetarOrErr()
	if IsNotLoaded(err) {
		result, err = sc.QueryMetar().Only(ctx)
	}
	return result, err
}

func (s *Station) Airport(ctx context.Context) (*Airport, error) {
	result, err := s.Edges.AirportOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryAirport().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Station) Metars(ctx context.Context) ([]*Metar, error) {
	result, err := s.NamedMetars(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = s.QueryMetars().All(ctx)
	}
	return result, err
}

func (s *Station) Tafs(ctx context.Context) ([]*Taf, error) {
	result, err := s.NamedTafs(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = s.QueryTafs().All(ctx)
	}
	return result, err
}

func (t *Taf) Station(ctx context.Context) (*Station, error) {
	result, err := t.Edges.StationOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryStation().Only(ctx)
	}
	return result, err
}

func (t *Taf) SkyConditions(ctx context.Context) ([]*SkyCondition, error) {
	result, err := t.NamedSkyConditions(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = t.QuerySkyConditions().All(ctx)
	}
	return result, err
}
