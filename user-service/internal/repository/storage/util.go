package storage

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/trunov/erply-assignement-task/user-service/internal/repository/erply"
)

func CustomScanCustomer(row *sqlx.Row, customer *erply.Customer) error {
	var externalIDsJSON string
	err := row.Scan(
		&customer.ID,
		&customer.CustomerID,
		&customer.FullName,
		&customer.CompanyName,
		&customer.CompanyTypeID,
		&customer.FirstName,
		&customer.LastName,
		&customer.PersonTitleID,
		&customer.EInvoiceEmail,
		&customer.EInvoiceReference,
		&customer.EmailEnabled,
		&customer.EInvoiceEnabled,
		&customer.DocuraEDIEnabled,
		&customer.MailEnabled,
		&customer.OperatorIdentifier,
		&customer.EDI,
		&customer.DoNotSell,
		&customer.PartialTaxExemption,
		&customer.GroupID,
		&customer.CountryID,
		&customer.PayerID,
		&customer.Phone,
		&customer.Mobile,
		&customer.Email,
		&customer.Fax,
		&customer.Code,
		&customer.Birthday,
		&customer.IntegrationCode,
		&customer.FlagStatus,
		&customer.ColorStatus,
		&customer.Credit,
		&customer.SalesBlocked,
		&customer.ReferenceNumber,
		&customer.CustomerCardNumber,
		&customer.FactoringContractNumber,
		&customer.GroupName,
		&customer.CustomerType,
		&customer.Address,
		&customer.Street,
		&customer.Address2,
		&customer.City,
		&customer.PostalCode,
		&customer.Country,
		&customer.State,
		&customer.AddressTypeID,
		&customer.AddressTypeName,
		&customer.IsPOSDefaultCustomer,
		&customer.EUCustomerType,
		&customer.EDIType,
		&customer.LastModifierUsername,
		&customer.LastModifierEmployeeID,
		&customer.TaxExempt,
		&customer.PaysViaFactoring,
		&customer.RewardPoints,
		&customer.TwitterID,
		&customer.FacebookName,
		&customer.CreditCardLastNumbers,
		&customer.GLN,
		&customer.DeliveryTypeID,
		&customer.Image,
		&customer.CustomerBalanceDisabled,
		&customer.RewardPointsDisabled,
		&customer.PosCouponsDisabled,
		&customer.EmailOptOut,
		&customer.SignUpStoreID,
		&customer.HomeStoreID,
		&customer.Gender,
		&customer.PeppolID,
		&externalIDsJSON,
	)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(externalIDsJSON), &customer.ExternalIDs); err != nil {
		return fmt.Errorf("failed to unmarshal externalIDs: %w", err)
	}

	return nil
}
