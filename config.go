package b24

const (
	scheme     = "https"
	oAuthToken = "https://oauth.bitrix.info/oauth/token"
)

// Methods

const (
	Auth      = "auth"
	rest      = "/rest/"
	UserAdmin = "user.admin.json"

	PlacementBind   = "placement.bind"
	PlacementUnBind = "placement.unbind"

	CrmContactGet    = "crm.contact.get"
	CrmContactList   = "crm.contact.list"
	CrmContactAdd    = "crm.contact.add"
	CrmContactUpdate = "crm.contact.update"

	CrmLeadGet  = "crm.lead.get"
	CrmLeadList = "crm.lead.get"

	CrmDealGet  = "crm.deal.get"
	CrmDealList = "crm.deal.get"

	CrmCompanyGet  = "crm.company.get"
	CrmCompanyList = "crm.company.list"

	EventBind = "event.bind.json"

	UserFieldTypeAdd    = "userfieldtype.add"
	CrmContactUserField = "crm.contact.userfield.add"
	CrmLeadUserField    = "crm.lead.userfield.add"
	CrmDealUserField    = "crm.deal.userfield.add"
	CrmCompanyUserField = "crm.company.userfield.add"

	BizprocRobotAdd = "bizproc.robot.add"
	BizprocRobotDel = "bizproc.robot.delete"
)

const (
	OnAppUninstall = "ONAPPUNINSTALL"
)

// Places

const (
	CrmContactTab = "CRM_CONTACT_DETAIL_TAB"
)

type callMethodOptions struct {
	// Method is a request's method
	Method string
	// BaseURL is a url from constants above.
	BaseURL string
	// In is a struct, which will be marshalled to Request Body
	In interface{}
	// Out is a struct, which will be unmarshalled
	Out interface{}
	// Params is a URL Parameters
	Params *RequestParams
}
