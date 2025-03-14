package role

func (r *Role) HavePermissionOrNot(svc, resource, act string) bool {
	for i := range r.Spec.Permissons {
		permission := r.Spec.Permissons[i]
		if permission.HavePermissionOrNot(svc, resource, act) {
			return true
		}
	}
	return false
}

func (p *Permission) HavePermissionOrNot(svc, resource, act string) bool {
	if p.Service != "*" && p.Service != svc {
		return false
	}

	if p.Resource != "*" && p.Resource != resource {
		return false
	}

	if !p.HaveActOrNot(act) {
		return false
	}
	return true
}

func (p *Permission) HaveActOrNot(act string) bool {
	for i := range p.Action {
		a := p.Action[i]
		if a == "*" || a == act {
			return true
		}
	}
	return false
}
