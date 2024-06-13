package main

import (
	"github.com/ministryofjustice/opg-data-lpa-store/internal/shared"
	"github.com/ministryofjustice/opg-data-lpa-store/internal/validate"
	"github.com/ministryofjustice/opg-data-lpa-store/lambda/update/parse"
)

type IdCheckComplete struct {
	Actor         idccActor
	IdentityCheck *shared.IdentityCheck
}

type idccActor string

var (
	donor               = idccActor("Donor")
	certificateProvider = idccActor("CertificateProvider")
)

func (idcc IdCheckComplete) Apply(lpa *shared.Lpa) []shared.FieldError {
	if idcc.Actor == donor {
		lpa.Donor.IdentityCheck = idcc.IdentityCheck
	} else {
		lpa.CertificateProvider.IdentityCheck = idcc.IdentityCheck
	}
	return nil
}

func validateConfirmIdentity(prefix string, actor idccActor, changes []shared.Change, lpa *shared.Lpa) (IdCheckComplete, []shared.FieldError) {
	var existing IdCheckComplete

	identityCheckParser := func(actor idccActor) func(p *parse.Parser) []shared.FieldError {
		return func(p *parse.Parser) []shared.FieldError {
			if existing.Actor != "" {
				return []shared.FieldError{{Source: "/", Detail: "id check for multiple actors is not allowed"}}
			}

			switch actor {
			case donor:
				existing.IdentityCheck = lpa.Donor.IdentityCheck
			case certificateProvider:
				existing.IdentityCheck = lpa.CertificateProvider.IdentityCheck
			}

			if existing.IdentityCheck == nil {
				existing.IdentityCheck = &shared.IdentityCheck{}
			}

			existing.Actor = actor

			return p.
				Field("/type", &existing.IdentityCheck.Type, parse.Validate(func() []shared.FieldError {
					return validate.IsValid("", existing.IdentityCheck.Type)
				}), parse.MustMatchExisting()).
				Field("/checkedAt", &existing.IdentityCheck.CheckedAt, parse.Validate(func() []shared.FieldError {
					return validate.Time("", existing.IdentityCheck.CheckedAt)
				}), parse.MustMatchExisting()).
				Field("/reference", &existing.IdentityCheck.Reference, parse.Validate(func() []shared.FieldError {
					return validate.Required("", existing.IdentityCheck.Reference)
				}), parse.MustMatchExisting()).
				Consumed()
		}
	}

	errors := parse.Changes(changes).
		Prefix(prefix, identityCheckParser(actor)).
		Errors()

	if existing.Actor == "" {
		return existing, append(errors, shared.FieldError{Source: "/", Detail: "id check for unknown actor type"})
	}

	return existing, errors
}

func validateDonorConfirmIdentity(changes []shared.Change, lpa *shared.Lpa) (IdCheckComplete, []shared.FieldError) {
	return validateConfirmIdentity("/donor/identityCheck", donor, changes, lpa)
}

func validateCertificateProviderConfirmIdentity(changes []shared.Change, lpa *shared.Lpa) (IdCheckComplete, []shared.FieldError) {
	return validateConfirmIdentity("/certificateProvider/identityCheck", certificateProvider, changes, lpa)
}
