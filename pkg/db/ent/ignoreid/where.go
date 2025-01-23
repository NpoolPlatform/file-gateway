// Code generated by ent, DO NOT EDIT.

package ignoreid

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/file-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// SampleCol applies equality check predicate on the "sample_col" field. It's identical to SampleColEQ.
func SampleCol(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSampleCol), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.IgnoreID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.IgnoreID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.IgnoreID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.IgnoreID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.IgnoreID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.IgnoreID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.IgnoreID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.IgnoreID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// SampleColEQ applies the EQ predicate on the "sample_col" field.
func SampleColEQ(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSampleCol), v))
	})
}

// SampleColNEQ applies the NEQ predicate on the "sample_col" field.
func SampleColNEQ(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSampleCol), v))
	})
}

// SampleColIn applies the In predicate on the "sample_col" field.
func SampleColIn(vs ...string) predicate.IgnoreID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSampleCol), v...))
	})
}

// SampleColNotIn applies the NotIn predicate on the "sample_col" field.
func SampleColNotIn(vs ...string) predicate.IgnoreID {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSampleCol), v...))
	})
}

// SampleColGT applies the GT predicate on the "sample_col" field.
func SampleColGT(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSampleCol), v))
	})
}

// SampleColGTE applies the GTE predicate on the "sample_col" field.
func SampleColGTE(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSampleCol), v))
	})
}

// SampleColLT applies the LT predicate on the "sample_col" field.
func SampleColLT(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSampleCol), v))
	})
}

// SampleColLTE applies the LTE predicate on the "sample_col" field.
func SampleColLTE(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSampleCol), v))
	})
}

// SampleColContains applies the Contains predicate on the "sample_col" field.
func SampleColContains(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSampleCol), v))
	})
}

// SampleColHasPrefix applies the HasPrefix predicate on the "sample_col" field.
func SampleColHasPrefix(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSampleCol), v))
	})
}

// SampleColHasSuffix applies the HasSuffix predicate on the "sample_col" field.
func SampleColHasSuffix(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSampleCol), v))
	})
}

// SampleColIsNil applies the IsNil predicate on the "sample_col" field.
func SampleColIsNil() predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSampleCol)))
	})
}

// SampleColNotNil applies the NotNil predicate on the "sample_col" field.
func SampleColNotNil() predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSampleCol)))
	})
}

// SampleColEqualFold applies the EqualFold predicate on the "sample_col" field.
func SampleColEqualFold(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSampleCol), v))
	})
}

// SampleColContainsFold applies the ContainsFold predicate on the "sample_col" field.
func SampleColContainsFold(v string) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSampleCol), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IgnoreID) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IgnoreID) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
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
func Not(p predicate.IgnoreID) predicate.IgnoreID {
	return predicate.IgnoreID(func(s *sql.Selector) {
		p(s.Not())
	})
}
