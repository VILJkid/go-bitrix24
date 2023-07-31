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

	UserFieldTypeAdd       = "userfieldtype.add"
	UserFieldTypeDelete    = "userfieldtype.delete"
	CrmContactUserFieldAdd = "crm.contact.userfield.add"
	CrmLeadUserFieldAdd    = "crm.lead.userfield.add"
	CrmDealUserFieldAdd    = "crm.deal.userfield.add"
	CrmCompanyUserFieldAdd = "crm.company.userfield.add"

	CrmContactUserFieldList = "crm.contact.userfield.list"
	CrmLeadUserFieldList    = "crm.lead.userfield.list"
	CrmDealUserFieldList    = "crm.deal.userfield.list"
	CrmCompanyUserFieldList = "crm.company.userfield.list"

	CrmContactUserFieldDelete = "crm.contact.userfield.delete"
	CrmLeadUserFieldDelete    = "crm.lead.userfield.delete"
	CrmDealUserFieldDelete    = "crm.deal.userfield.delete"
	CrmCompanyUserFieldDelete = "crm.company.userfield.delete"

	UserFieldConfigList   = "userfieldconfig.list"
	UserFieldConfigDelete = "userfieldconfig.delete"

	BizprocRobotAdd  = "bizproc.robot.add"
	BizprocRobotDel  = "bizproc.robot.delete"
	BizProcEventSend = "bizproc.event.send"
)

const (
	OnAppUninstall = "ONAPPUNINSTALL"
	CrmContactTab  = "CRM_CONTACT_DETAIL_TAB"
	CrmLeadTab     = "CRM_LEAD_DETAIL_TAB"
	CrmDealTab     = "CRM_DEAL_DETAIL_TAB"
	CrmCompanyTab  = "CRM_COMPANY_DETAIL_TAB"
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
