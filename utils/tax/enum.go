package tax

import "database/sql/driver"

type Type string

const (
	TypeAeTrn     Type = "ae_trn"
	TypeAuAbn     Type = "au_abn"
	TypeAuArn     Type = "au_arn"
	TypeBgUic     Type = "bg_uic"
	TypeBrCnpj    Type = "br_cnpj"
	TypeBrCpf     Type = "br_cpf"
	TypeCaBn      Type = "ca_bn"
	TypeCaGst_hst Type = "ca_gst_hst"
	TypeCaPst_bc  Type = "ca_pst_bc"
	TypeCaPst_mb  Type = "ca_pst_mb"
	TypeCaPst_sk  Type = "ca_pst_sk"
	TypeCaQst     Type = "ca_qst"
	TypeChVat     Type = "ch_vat"
	TypeClTin     Type = "cl_tin"
	TypeEgTin     Type = "eg_tin"
	TypeEsCif     Type = "es_cif"
	TypeEuOssVat  Type = "eu_oss_vat"
	TypeEuVat     Type = "eu_vat"
	TypeGbVat     Type = "gb_vat"
	TypeGeVat     Type = "ge_vat"
	TypeHkBr      Type = "hk_br"
	TypeHuTin     Type = "hu_tin"
	TypeIdNpwp    Type = "id_npwp"
	TypeIlVat     Type = "il_vat"
	TypeInGst     Type = "in_gst"
	TypeIsVat     Type = "is_vat"
	TypeJpCn      Type = "jp_cn"
	TypeJpRn      Type = "jp_rn"
	TypeJpTrn     Type = "jp_trn"
	TypeKePin     Type = "ke_pin"
	TypeKrBrn     Type = "kr_brn"
	TypeLiUid     Type = "li_uid"
	TypeMxRfc     Type = "mx_rfc"
	TypeMyFrp     Type = "my_frp"
	TypeMyItn     Type = "my_itn"
	TypeMySst     Type = "my_sst"
	TypeNoVat     Type = "no_vat"
	TypeNzGst     Type = "nz_gst"
	TypePhTin     Type = "ph_tin"
	TypeRuInn     Type = "ru_inn"
	TypeRuKpp     Type = "ru_kpp"
	TypeSaVat     Type = "sa_vat"
	TypeSgGst     Type = "sg_gst"
	TypeSgUen     Type = "sg_uen"
	TypeSiTin     Type = "si_tin"
	TypeThVat     Type = "th_vat"
	TypeTrTin     Type = "tr_tin"
	TypeTwVat     Type = "tw_vat"
	TypeUaVat     Type = "ua_vat"
	TypeUsEin     Type = "us_ein"
	TypeZaVat     Type = "za_vat"
	TypeUnknown   Type = "unknown"
)

func (ct *Type) Scan(value interface{}) error {
	*ct = Type(value.(string))
	return nil
}

func (ct Type) Value() (driver.Value,
	error) {
	return string(ct),
		nil
}

type Status string

const (
	StatusPending     Status = "pending"
	StatusVerified    Status = "verified"
	StatusUnverified  Status = "unverified"
	StatusUnavailable Status = "unavailable"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

type TaxType string

const (
	TaxTypeVat      TaxType = "vat"
	TaxTypeSalesTax TaxType = "sales_tax"
)

func (ct *TaxType) Scan(value interface{}) error {
	*ct = TaxType(value.(string))
	return nil
}

func (ct TaxType) Value() (driver.Value, error) {
	return string(ct), nil
}
