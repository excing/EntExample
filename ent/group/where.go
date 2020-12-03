// Code generated by entc, DO NOT EDIT.

package group

import (
	"ent_example/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIndex), v))
	})
}

// Min applies equality check predicate on the "min" field. It's identical to MinEQ.
func Min(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMin), v))
	})
}

// Max applies equality check predicate on the "max" field. It's identical to MaxEQ.
func Max(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMax), v))
	})
}

// Range applies equality check predicate on the "range" field. It's identical to RangeEQ.
func Range(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRange), v))
	})
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// Log applies equality check predicate on the "log" field. It's identical to LogEQ.
func Log(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLog), v))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNickname), v))
	})
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNickname), v...))
	})
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNickname), v...))
	})
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNickname), v))
	})
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNickname), v))
	})
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNickname), v))
	})
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNickname), v))
	})
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNickname), v))
	})
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNickname), v))
	})
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNickname), v))
	})
}

// NicknameIsNil applies the IsNil predicate on the "nickname" field.
func NicknameIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNickname)))
	})
}

// NicknameNotNil applies the NotNil predicate on the "nickname" field.
func NicknameNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNickname)))
	})
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNickname), v))
	})
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNickname), v))
	})
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCount), v))
	})
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCount), v...))
	})
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCount), v...))
	})
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCount), v))
	})
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCount), v))
	})
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCount), v))
	})
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCount), v))
	})
}

// CountIsNil applies the IsNil predicate on the "count" field.
func CountIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCount)))
	})
}

// CountNotNil applies the NotNil predicate on the "count" field.
func CountNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCount)))
	})
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// CodeIsNil applies the IsNil predicate on the "code" field.
func CodeIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCode)))
	})
}

// CodeNotNil applies the NotNil predicate on the "code" field.
func CodeNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCode)))
	})
}

// IndexEQ applies the EQ predicate on the "index" field.
func IndexEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIndex), v))
	})
}

// IndexNEQ applies the NEQ predicate on the "index" field.
func IndexNEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIndex), v))
	})
}

// IndexIn applies the In predicate on the "index" field.
func IndexIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIndex), v...))
	})
}

// IndexNotIn applies the NotIn predicate on the "index" field.
func IndexNotIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIndex), v...))
	})
}

// IndexGT applies the GT predicate on the "index" field.
func IndexGT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIndex), v))
	})
}

// IndexGTE applies the GTE predicate on the "index" field.
func IndexGTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIndex), v))
	})
}

// IndexLT applies the LT predicate on the "index" field.
func IndexLT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIndex), v))
	})
}

// IndexLTE applies the LTE predicate on the "index" field.
func IndexLTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIndex), v))
	})
}

// IndexIsNil applies the IsNil predicate on the "index" field.
func IndexIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIndex)))
	})
}

// IndexNotNil applies the NotNil predicate on the "index" field.
func IndexNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIndex)))
	})
}

// MinEQ applies the EQ predicate on the "min" field.
func MinEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMin), v))
	})
}

// MinNEQ applies the NEQ predicate on the "min" field.
func MinNEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMin), v))
	})
}

// MinIn applies the In predicate on the "min" field.
func MinIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMin), v...))
	})
}

// MinNotIn applies the NotIn predicate on the "min" field.
func MinNotIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMin), v...))
	})
}

// MinGT applies the GT predicate on the "min" field.
func MinGT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMin), v))
	})
}

// MinGTE applies the GTE predicate on the "min" field.
func MinGTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMin), v))
	})
}

// MinLT applies the LT predicate on the "min" field.
func MinLT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMin), v))
	})
}

// MinLTE applies the LTE predicate on the "min" field.
func MinLTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMin), v))
	})
}

// MinIsNil applies the IsNil predicate on the "min" field.
func MinIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMin)))
	})
}

// MinNotNil applies the NotNil predicate on the "min" field.
func MinNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMin)))
	})
}

// MaxEQ applies the EQ predicate on the "max" field.
func MaxEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMax), v))
	})
}

// MaxNEQ applies the NEQ predicate on the "max" field.
func MaxNEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMax), v))
	})
}

// MaxIn applies the In predicate on the "max" field.
func MaxIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMax), v...))
	})
}

// MaxNotIn applies the NotIn predicate on the "max" field.
func MaxNotIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMax), v...))
	})
}

// MaxGT applies the GT predicate on the "max" field.
func MaxGT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMax), v))
	})
}

// MaxGTE applies the GTE predicate on the "max" field.
func MaxGTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMax), v))
	})
}

// MaxLT applies the LT predicate on the "max" field.
func MaxLT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMax), v))
	})
}

// MaxLTE applies the LTE predicate on the "max" field.
func MaxLTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMax), v))
	})
}

// MaxIsNil applies the IsNil predicate on the "max" field.
func MaxIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMax)))
	})
}

// MaxNotNil applies the NotNil predicate on the "max" field.
func MaxNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMax)))
	})
}

// RangeEQ applies the EQ predicate on the "range" field.
func RangeEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRange), v))
	})
}

// RangeNEQ applies the NEQ predicate on the "range" field.
func RangeNEQ(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRange), v))
	})
}

// RangeIn applies the In predicate on the "range" field.
func RangeIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRange), v...))
	})
}

// RangeNotIn applies the NotIn predicate on the "range" field.
func RangeNotIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRange), v...))
	})
}

// RangeGT applies the GT predicate on the "range" field.
func RangeGT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRange), v))
	})
}

// RangeGTE applies the GTE predicate on the "range" field.
func RangeGTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRange), v))
	})
}

// RangeLT applies the LT predicate on the "range" field.
func RangeLT(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRange), v))
	})
}

// RangeLTE applies the LTE predicate on the "range" field.
func RangeLTE(v int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRange), v))
	})
}

// RangeIsNil applies the IsNil predicate on the "range" field.
func RangeIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRange)))
	})
}

// RangeNotNil applies the NotNil predicate on the "range" field.
func RangeNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRange)))
	})
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNote), v))
	})
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNote), v...))
	})
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNote), v...))
	})
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNote), v))
	})
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNote), v))
	})
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNote), v))
	})
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNote), v))
	})
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNote), v))
	})
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNote), v))
	})
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNote), v))
	})
}

// NoteIsNil applies the IsNil predicate on the "note" field.
func NoteIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNote)))
	})
}

// NoteNotNil applies the NotNil predicate on the "note" field.
func NoteNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNote)))
	})
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNote), v))
	})
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNote), v))
	})
}

// LogEQ applies the EQ predicate on the "log" field.
func LogEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLog), v))
	})
}

// LogNEQ applies the NEQ predicate on the "log" field.
func LogNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLog), v))
	})
}

// LogIn applies the In predicate on the "log" field.
func LogIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLog), v...))
	})
}

// LogNotIn applies the NotIn predicate on the "log" field.
func LogNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLog), v...))
	})
}

// LogGT applies the GT predicate on the "log" field.
func LogGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLog), v))
	})
}

// LogGTE applies the GTE predicate on the "log" field.
func LogGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLog), v))
	})
}

// LogLT applies the LT predicate on the "log" field.
func LogLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLog), v))
	})
}

// LogLTE applies the LTE predicate on the "log" field.
func LogLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLog), v))
	})
}

// LogContains applies the Contains predicate on the "log" field.
func LogContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLog), v))
	})
}

// LogHasPrefix applies the HasPrefix predicate on the "log" field.
func LogHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLog), v))
	})
}

// LogHasSuffix applies the HasSuffix predicate on the "log" field.
func LogHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLog), v))
	})
}

// LogIsNil applies the IsNil predicate on the "log" field.
func LogIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLog)))
	})
}

// LogNotNil applies the NotNil predicate on the "log" field.
func LogNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLog)))
	})
}

// LogEqualFold applies the EqualFold predicate on the "log" field.
func LogEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLog), v))
	})
}

// LogContainsFold applies the ContainsFold predicate on the "log" field.
func LogContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLog), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameIsNil applies the IsNil predicate on the "username" field.
func UsernameIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUsername)))
	})
}

// UsernameNotNil applies the NotNil predicate on the "username" field.
func UsernameNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUsername)))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
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
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		p(s.Not())
	})
}
