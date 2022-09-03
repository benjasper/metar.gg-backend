// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"metar.gg/ent/airport"
	"metar.gg/ent/runway"
	"metar.gg/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	airportMixin := schema.Airport{}.Mixin()
	airportMixinFields0 := airportMixin[0].Fields()
	_ = airportMixinFields0
	airportMixinFields1 := airportMixin[1].Fields()
	_ = airportMixinFields1
	airportFields := schema.Airport{}.Fields()
	_ = airportFields
	// airportDescImportFlag is the schema descriptor for import_flag field.
	airportDescImportFlag := airportMixinFields0[2].Descriptor()
	// airport.DefaultImportFlag holds the default value on creation for the import_flag field.
	airport.DefaultImportFlag = airportDescImportFlag.Default.(bool)
	// airportDescCreateTime is the schema descriptor for create_time field.
	airportDescCreateTime := airportMixinFields1[0].Descriptor()
	// airport.DefaultCreateTime holds the default value on creation for the create_time field.
	airport.DefaultCreateTime = airportDescCreateTime.Default.(func() time.Time)
	// airportDescUpdateTime is the schema descriptor for update_time field.
	airportDescUpdateTime := airportMixinFields1[1].Descriptor()
	// airport.DefaultUpdateTime holds the default value on creation for the update_time field.
	airport.DefaultUpdateTime = airportDescUpdateTime.Default.(func() time.Time)
	// airport.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	airport.UpdateDefaultUpdateTime = airportDescUpdateTime.UpdateDefault.(func() time.Time)
	runwayMixin := schema.Runway{}.Mixin()
	runwayMixinFields0 := runwayMixin[0].Fields()
	_ = runwayMixinFields0
	runwayMixinFields1 := runwayMixin[1].Fields()
	_ = runwayMixinFields1
	runwayFields := schema.Runway{}.Fields()
	_ = runwayFields
	// runwayDescImportFlag is the schema descriptor for import_flag field.
	runwayDescImportFlag := runwayMixinFields0[2].Descriptor()
	// runway.DefaultImportFlag holds the default value on creation for the import_flag field.
	runway.DefaultImportFlag = runwayDescImportFlag.Default.(bool)
	// runwayDescCreateTime is the schema descriptor for create_time field.
	runwayDescCreateTime := runwayMixinFields1[0].Descriptor()
	// runway.DefaultCreateTime holds the default value on creation for the create_time field.
	runway.DefaultCreateTime = runwayDescCreateTime.Default.(func() time.Time)
	// runwayDescUpdateTime is the schema descriptor for update_time field.
	runwayDescUpdateTime := runwayMixinFields1[1].Descriptor()
	// runway.DefaultUpdateTime holds the default value on creation for the update_time field.
	runway.DefaultUpdateTime = runwayDescUpdateTime.Default.(func() time.Time)
	// runway.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	runway.UpdateDefaultUpdateTime = runwayDescUpdateTime.UpdateDefault.(func() time.Time)
}
