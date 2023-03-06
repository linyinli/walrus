// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package clustercost

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/types"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldID, id))
}

// StartTime applies equality check predicate on the "startTime" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "endTime" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldEndTime, v))
}

// Minutes applies equality check predicate on the "minutes" field. It's identical to MinutesEQ.
func Minutes(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldMinutes, v))
}

// ConnectorID applies equality check predicate on the "connectorID" field. It's identical to ConnectorIDEQ.
func ConnectorID(v types.ID) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldConnectorID, v))
}

// ClusterName applies equality check predicate on the "clusterName" field. It's identical to ClusterNameEQ.
func ClusterName(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldClusterName, v))
}

// TotalCost applies equality check predicate on the "totalCost" field. It's identical to TotalCostEQ.
func TotalCost(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldTotalCost, v))
}

// Currency applies equality check predicate on the "currency" field. It's identical to CurrencyEQ.
func Currency(v int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldCurrency, v))
}

// CpuCost applies equality check predicate on the "cpuCost" field. It's identical to CpuCostEQ.
func CpuCost(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldCpuCost, v))
}

// GpuCost applies equality check predicate on the "gpuCost" field. It's identical to GpuCostEQ.
func GpuCost(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldGpuCost, v))
}

// RamCost applies equality check predicate on the "ramCost" field. It's identical to RamCostEQ.
func RamCost(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldRamCost, v))
}

// StorageCost applies equality check predicate on the "storageCost" field. It's identical to StorageCostEQ.
func StorageCost(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldStorageCost, v))
}

// AllocationCost applies equality check predicate on the "allocationCost" field. It's identical to AllocationCostEQ.
func AllocationCost(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldAllocationCost, v))
}

// IdleCost applies equality check predicate on the "idleCost" field. It's identical to IdleCostEQ.
func IdleCost(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldIdleCost, v))
}

// ManagementCost applies equality check predicate on the "managementCost" field. It's identical to ManagementCostEQ.
func ManagementCost(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldManagementCost, v))
}

// StartTimeEQ applies the EQ predicate on the "startTime" field.
func StartTimeEQ(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "startTime" field.
func StartTimeNEQ(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "startTime" field.
func StartTimeIn(vs ...time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "startTime" field.
func StartTimeNotIn(vs ...time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "startTime" field.
func StartTimeGT(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "startTime" field.
func StartTimeGTE(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "startTime" field.
func StartTimeLT(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "startTime" field.
func StartTimeLTE(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldStartTime, v))
}

// EndTimeEQ applies the EQ predicate on the "endTime" field.
func EndTimeEQ(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "endTime" field.
func EndTimeNEQ(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "endTime" field.
func EndTimeIn(vs ...time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "endTime" field.
func EndTimeNotIn(vs ...time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "endTime" field.
func EndTimeGT(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "endTime" field.
func EndTimeGTE(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "endTime" field.
func EndTimeLT(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "endTime" field.
func EndTimeLTE(v time.Time) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldEndTime, v))
}

// MinutesEQ applies the EQ predicate on the "minutes" field.
func MinutesEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldMinutes, v))
}

// MinutesNEQ applies the NEQ predicate on the "minutes" field.
func MinutesNEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldMinutes, v))
}

// MinutesIn applies the In predicate on the "minutes" field.
func MinutesIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldMinutes, vs...))
}

// MinutesNotIn applies the NotIn predicate on the "minutes" field.
func MinutesNotIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldMinutes, vs...))
}

// MinutesGT applies the GT predicate on the "minutes" field.
func MinutesGT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldMinutes, v))
}

// MinutesGTE applies the GTE predicate on the "minutes" field.
func MinutesGTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldMinutes, v))
}

// MinutesLT applies the LT predicate on the "minutes" field.
func MinutesLT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldMinutes, v))
}

// MinutesLTE applies the LTE predicate on the "minutes" field.
func MinutesLTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldMinutes, v))
}

// ConnectorIDEQ applies the EQ predicate on the "connectorID" field.
func ConnectorIDEQ(v types.ID) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldConnectorID, v))
}

// ConnectorIDNEQ applies the NEQ predicate on the "connectorID" field.
func ConnectorIDNEQ(v types.ID) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldConnectorID, v))
}

// ConnectorIDIn applies the In predicate on the "connectorID" field.
func ConnectorIDIn(vs ...types.ID) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldConnectorID, vs...))
}

// ConnectorIDNotIn applies the NotIn predicate on the "connectorID" field.
func ConnectorIDNotIn(vs ...types.ID) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldConnectorID, vs...))
}

// ConnectorIDGT applies the GT predicate on the "connectorID" field.
func ConnectorIDGT(v types.ID) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldConnectorID, v))
}

// ConnectorIDGTE applies the GTE predicate on the "connectorID" field.
func ConnectorIDGTE(v types.ID) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldConnectorID, v))
}

// ConnectorIDLT applies the LT predicate on the "connectorID" field.
func ConnectorIDLT(v types.ID) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldConnectorID, v))
}

// ConnectorIDLTE applies the LTE predicate on the "connectorID" field.
func ConnectorIDLTE(v types.ID) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldConnectorID, v))
}

// ConnectorIDContains applies the Contains predicate on the "connectorID" field.
func ConnectorIDContains(v types.ID) predicate.ClusterCost {
	vc := string(v)
	return predicate.ClusterCost(sql.FieldContains(FieldConnectorID, vc))
}

// ConnectorIDHasPrefix applies the HasPrefix predicate on the "connectorID" field.
func ConnectorIDHasPrefix(v types.ID) predicate.ClusterCost {
	vc := string(v)
	return predicate.ClusterCost(sql.FieldHasPrefix(FieldConnectorID, vc))
}

// ConnectorIDHasSuffix applies the HasSuffix predicate on the "connectorID" field.
func ConnectorIDHasSuffix(v types.ID) predicate.ClusterCost {
	vc := string(v)
	return predicate.ClusterCost(sql.FieldHasSuffix(FieldConnectorID, vc))
}

// ConnectorIDEqualFold applies the EqualFold predicate on the "connectorID" field.
func ConnectorIDEqualFold(v types.ID) predicate.ClusterCost {
	vc := string(v)
	return predicate.ClusterCost(sql.FieldEqualFold(FieldConnectorID, vc))
}

// ConnectorIDContainsFold applies the ContainsFold predicate on the "connectorID" field.
func ConnectorIDContainsFold(v types.ID) predicate.ClusterCost {
	vc := string(v)
	return predicate.ClusterCost(sql.FieldContainsFold(FieldConnectorID, vc))
}

// ClusterNameEQ applies the EQ predicate on the "clusterName" field.
func ClusterNameEQ(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldClusterName, v))
}

// ClusterNameNEQ applies the NEQ predicate on the "clusterName" field.
func ClusterNameNEQ(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldClusterName, v))
}

// ClusterNameIn applies the In predicate on the "clusterName" field.
func ClusterNameIn(vs ...string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldClusterName, vs...))
}

// ClusterNameNotIn applies the NotIn predicate on the "clusterName" field.
func ClusterNameNotIn(vs ...string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldClusterName, vs...))
}

// ClusterNameGT applies the GT predicate on the "clusterName" field.
func ClusterNameGT(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldClusterName, v))
}

// ClusterNameGTE applies the GTE predicate on the "clusterName" field.
func ClusterNameGTE(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldClusterName, v))
}

// ClusterNameLT applies the LT predicate on the "clusterName" field.
func ClusterNameLT(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldClusterName, v))
}

// ClusterNameLTE applies the LTE predicate on the "clusterName" field.
func ClusterNameLTE(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldClusterName, v))
}

// ClusterNameContains applies the Contains predicate on the "clusterName" field.
func ClusterNameContains(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldContains(FieldClusterName, v))
}

// ClusterNameHasPrefix applies the HasPrefix predicate on the "clusterName" field.
func ClusterNameHasPrefix(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldHasPrefix(FieldClusterName, v))
}

// ClusterNameHasSuffix applies the HasSuffix predicate on the "clusterName" field.
func ClusterNameHasSuffix(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldHasSuffix(FieldClusterName, v))
}

// ClusterNameEqualFold applies the EqualFold predicate on the "clusterName" field.
func ClusterNameEqualFold(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEqualFold(FieldClusterName, v))
}

// ClusterNameContainsFold applies the ContainsFold predicate on the "clusterName" field.
func ClusterNameContainsFold(v string) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldContainsFold(FieldClusterName, v))
}

// TotalCostEQ applies the EQ predicate on the "totalCost" field.
func TotalCostEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldTotalCost, v))
}

// TotalCostNEQ applies the NEQ predicate on the "totalCost" field.
func TotalCostNEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldTotalCost, v))
}

// TotalCostIn applies the In predicate on the "totalCost" field.
func TotalCostIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldTotalCost, vs...))
}

// TotalCostNotIn applies the NotIn predicate on the "totalCost" field.
func TotalCostNotIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldTotalCost, vs...))
}

// TotalCostGT applies the GT predicate on the "totalCost" field.
func TotalCostGT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldTotalCost, v))
}

// TotalCostGTE applies the GTE predicate on the "totalCost" field.
func TotalCostGTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldTotalCost, v))
}

// TotalCostLT applies the LT predicate on the "totalCost" field.
func TotalCostLT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldTotalCost, v))
}

// TotalCostLTE applies the LTE predicate on the "totalCost" field.
func TotalCostLTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldTotalCost, v))
}

// CurrencyEQ applies the EQ predicate on the "currency" field.
func CurrencyEQ(v int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldCurrency, v))
}

// CurrencyNEQ applies the NEQ predicate on the "currency" field.
func CurrencyNEQ(v int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldCurrency, v))
}

// CurrencyIn applies the In predicate on the "currency" field.
func CurrencyIn(vs ...int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldCurrency, vs...))
}

// CurrencyNotIn applies the NotIn predicate on the "currency" field.
func CurrencyNotIn(vs ...int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldCurrency, vs...))
}

// CurrencyGT applies the GT predicate on the "currency" field.
func CurrencyGT(v int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldCurrency, v))
}

// CurrencyGTE applies the GTE predicate on the "currency" field.
func CurrencyGTE(v int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldCurrency, v))
}

// CurrencyLT applies the LT predicate on the "currency" field.
func CurrencyLT(v int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldCurrency, v))
}

// CurrencyLTE applies the LTE predicate on the "currency" field.
func CurrencyLTE(v int) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldCurrency, v))
}

// CurrencyIsNil applies the IsNil predicate on the "currency" field.
func CurrencyIsNil() predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIsNull(FieldCurrency))
}

// CurrencyNotNil applies the NotNil predicate on the "currency" field.
func CurrencyNotNil() predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotNull(FieldCurrency))
}

// CpuCostEQ applies the EQ predicate on the "cpuCost" field.
func CpuCostEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldCpuCost, v))
}

// CpuCostNEQ applies the NEQ predicate on the "cpuCost" field.
func CpuCostNEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldCpuCost, v))
}

// CpuCostIn applies the In predicate on the "cpuCost" field.
func CpuCostIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldCpuCost, vs...))
}

// CpuCostNotIn applies the NotIn predicate on the "cpuCost" field.
func CpuCostNotIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldCpuCost, vs...))
}

// CpuCostGT applies the GT predicate on the "cpuCost" field.
func CpuCostGT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldCpuCost, v))
}

// CpuCostGTE applies the GTE predicate on the "cpuCost" field.
func CpuCostGTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldCpuCost, v))
}

// CpuCostLT applies the LT predicate on the "cpuCost" field.
func CpuCostLT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldCpuCost, v))
}

// CpuCostLTE applies the LTE predicate on the "cpuCost" field.
func CpuCostLTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldCpuCost, v))
}

// GpuCostEQ applies the EQ predicate on the "gpuCost" field.
func GpuCostEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldGpuCost, v))
}

// GpuCostNEQ applies the NEQ predicate on the "gpuCost" field.
func GpuCostNEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldGpuCost, v))
}

// GpuCostIn applies the In predicate on the "gpuCost" field.
func GpuCostIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldGpuCost, vs...))
}

// GpuCostNotIn applies the NotIn predicate on the "gpuCost" field.
func GpuCostNotIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldGpuCost, vs...))
}

// GpuCostGT applies the GT predicate on the "gpuCost" field.
func GpuCostGT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldGpuCost, v))
}

// GpuCostGTE applies the GTE predicate on the "gpuCost" field.
func GpuCostGTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldGpuCost, v))
}

// GpuCostLT applies the LT predicate on the "gpuCost" field.
func GpuCostLT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldGpuCost, v))
}

// GpuCostLTE applies the LTE predicate on the "gpuCost" field.
func GpuCostLTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldGpuCost, v))
}

// RamCostEQ applies the EQ predicate on the "ramCost" field.
func RamCostEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldRamCost, v))
}

// RamCostNEQ applies the NEQ predicate on the "ramCost" field.
func RamCostNEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldRamCost, v))
}

// RamCostIn applies the In predicate on the "ramCost" field.
func RamCostIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldRamCost, vs...))
}

// RamCostNotIn applies the NotIn predicate on the "ramCost" field.
func RamCostNotIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldRamCost, vs...))
}

// RamCostGT applies the GT predicate on the "ramCost" field.
func RamCostGT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldRamCost, v))
}

// RamCostGTE applies the GTE predicate on the "ramCost" field.
func RamCostGTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldRamCost, v))
}

// RamCostLT applies the LT predicate on the "ramCost" field.
func RamCostLT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldRamCost, v))
}

// RamCostLTE applies the LTE predicate on the "ramCost" field.
func RamCostLTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldRamCost, v))
}

// StorageCostEQ applies the EQ predicate on the "storageCost" field.
func StorageCostEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldStorageCost, v))
}

// StorageCostNEQ applies the NEQ predicate on the "storageCost" field.
func StorageCostNEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldStorageCost, v))
}

// StorageCostIn applies the In predicate on the "storageCost" field.
func StorageCostIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldStorageCost, vs...))
}

// StorageCostNotIn applies the NotIn predicate on the "storageCost" field.
func StorageCostNotIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldStorageCost, vs...))
}

// StorageCostGT applies the GT predicate on the "storageCost" field.
func StorageCostGT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldStorageCost, v))
}

// StorageCostGTE applies the GTE predicate on the "storageCost" field.
func StorageCostGTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldStorageCost, v))
}

// StorageCostLT applies the LT predicate on the "storageCost" field.
func StorageCostLT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldStorageCost, v))
}

// StorageCostLTE applies the LTE predicate on the "storageCost" field.
func StorageCostLTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldStorageCost, v))
}

// AllocationCostEQ applies the EQ predicate on the "allocationCost" field.
func AllocationCostEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldAllocationCost, v))
}

// AllocationCostNEQ applies the NEQ predicate on the "allocationCost" field.
func AllocationCostNEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldAllocationCost, v))
}

// AllocationCostIn applies the In predicate on the "allocationCost" field.
func AllocationCostIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldAllocationCost, vs...))
}

// AllocationCostNotIn applies the NotIn predicate on the "allocationCost" field.
func AllocationCostNotIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldAllocationCost, vs...))
}

// AllocationCostGT applies the GT predicate on the "allocationCost" field.
func AllocationCostGT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldAllocationCost, v))
}

// AllocationCostGTE applies the GTE predicate on the "allocationCost" field.
func AllocationCostGTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldAllocationCost, v))
}

// AllocationCostLT applies the LT predicate on the "allocationCost" field.
func AllocationCostLT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldAllocationCost, v))
}

// AllocationCostLTE applies the LTE predicate on the "allocationCost" field.
func AllocationCostLTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldAllocationCost, v))
}

// IdleCostEQ applies the EQ predicate on the "idleCost" field.
func IdleCostEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldIdleCost, v))
}

// IdleCostNEQ applies the NEQ predicate on the "idleCost" field.
func IdleCostNEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldIdleCost, v))
}

// IdleCostIn applies the In predicate on the "idleCost" field.
func IdleCostIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldIdleCost, vs...))
}

// IdleCostNotIn applies the NotIn predicate on the "idleCost" field.
func IdleCostNotIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldIdleCost, vs...))
}

// IdleCostGT applies the GT predicate on the "idleCost" field.
func IdleCostGT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldIdleCost, v))
}

// IdleCostGTE applies the GTE predicate on the "idleCost" field.
func IdleCostGTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldIdleCost, v))
}

// IdleCostLT applies the LT predicate on the "idleCost" field.
func IdleCostLT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldIdleCost, v))
}

// IdleCostLTE applies the LTE predicate on the "idleCost" field.
func IdleCostLTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldIdleCost, v))
}

// ManagementCostEQ applies the EQ predicate on the "managementCost" field.
func ManagementCostEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldEQ(FieldManagementCost, v))
}

// ManagementCostNEQ applies the NEQ predicate on the "managementCost" field.
func ManagementCostNEQ(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNEQ(FieldManagementCost, v))
}

// ManagementCostIn applies the In predicate on the "managementCost" field.
func ManagementCostIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldIn(FieldManagementCost, vs...))
}

// ManagementCostNotIn applies the NotIn predicate on the "managementCost" field.
func ManagementCostNotIn(vs ...float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldNotIn(FieldManagementCost, vs...))
}

// ManagementCostGT applies the GT predicate on the "managementCost" field.
func ManagementCostGT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGT(FieldManagementCost, v))
}

// ManagementCostGTE applies the GTE predicate on the "managementCost" field.
func ManagementCostGTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldGTE(FieldManagementCost, v))
}

// ManagementCostLT applies the LT predicate on the "managementCost" field.
func ManagementCostLT(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLT(FieldManagementCost, v))
}

// ManagementCostLTE applies the LTE predicate on the "managementCost" field.
func ManagementCostLTE(v float64) predicate.ClusterCost {
	return predicate.ClusterCost(sql.FieldLTE(FieldManagementCost, v))
}

// HasConnector applies the HasEdge predicate on the "connector" edge.
func HasConnector() predicate.ClusterCost {
	return predicate.ClusterCost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ConnectorTable, ConnectorColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Connector
		step.Edge.Schema = schemaConfig.ClusterCost
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConnectorWith applies the HasEdge predicate on the "connector" edge with a given conditions (other predicates).
func HasConnectorWith(preds ...predicate.Connector) predicate.ClusterCost {
	return predicate.ClusterCost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConnectorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ConnectorTable, ConnectorColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Connector
		step.Edge.Schema = schemaConfig.ClusterCost
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ClusterCost) predicate.ClusterCost {
	return predicate.ClusterCost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ClusterCost) predicate.ClusterCost {
	return predicate.ClusterCost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ClusterCost) predicate.ClusterCost {
	return predicate.ClusterCost(func(s *sql.Selector) {
		p(s.Not())
	})
}