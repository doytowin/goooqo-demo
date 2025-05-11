package main

import "strings"

func (q RoleQuery) BuildConditions() ([]string, []any) {
	conditions := make([]string, 0, 4)
	args := make([]any, 0, 4)
	if q.IdIn != nil {
		phs := make([]string, 0, len(*q.IdIn))
		for _, arg := range *q.IdIn {
			args = append(args, arg)
			phs = append(phs, "?")
		}
		conditions = append(conditions, "id IN ("+strings.Join(phs, ", ")+")")
	}
	if q.RoleNameStart != nil && *q.RoleNameStart != "" {
		conditions = append(conditions, "role_name LIKE ?")
		args = append(args, *q.RoleNameStart+"%")
	}
	if q.Valid != nil {
		conditions = append(conditions, "valid = ?")
		args = append(args, *q.Valid)
	}
	return conditions, args
}
