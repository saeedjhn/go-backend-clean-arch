package entity

type AdminPermission string

const (
	AdminAdminRegisterPermission                 = AdminPermission("admin-register")
	AdminKindBoxReqAcceptPermission              = AdminPermission("kindboxreq-accept")
	AdminKindBoxReqAddPermission                 = AdminPermission("kindboxreq-add")
	AdminKindBoxReqRejectPermission              = AdminPermission("kindboxreq-reject")
	AdminKindBoxReqGetAllPermission              = AdminPermission("kindboxreq-getall")
	AdminKindBoxReqDeliverPermission             = AdminPermission("kindboxreq-deliver")
	AdminKindBoxReqAssignSenderAgentPermission   = AdminPermission("kindboxreq-assign_sender_agent")
	AdminAdminGetAllAgentPermission              = AdminPermission("admin-getall_agent")
	AdminBenefactorGetAllPermission              = AdminPermission("benefactor-getall")
	AdminKindBoxReqGetAwaitingDeliveryPermission = AdminPermission("kindboxreq-get_awaiting_delivery")
	AdminKindBoxGetPermission                    = AdminPermission("kindbox-get")
	AdminKindBoxAssignReceiverAgentPermission    = AdminPermission("kindbox-assign_receiver_agent")
	AdminKindBoxGetAllPermission                 = AdminPermission("kindbox-getall")
	AdminKindBoxReqUpdatePermission              = AdminPermission("kindboxreq-update")
	AdminKindBoxReqGetPermission                 = AdminPermission("kindboxreq-get")
	AdminKindBoxGetAwaitingReturnPermission      = AdminPermission("kindbox-get_awaiting_return")
	AdminKindBoxReturnPermission                 = AdminPermission("kindbox-return")
	AdminKindBoxEnumeratePermission              = AdminPermission("kindbox-enumerate")
	AdminKindBoxUpdatePermission                 = AdminPermission("kindbox-update")
	AdminBenefactorGetPermission                 = AdminPermission("benefactor-get")
	AdminBenefactorUpdatePermission              = AdminPermission("benefactor-update")
	AdminBenefactorUpdateStatusPermission        = AdminPermission("benefactor-update-status")
	AdminReferTimeGetAllPermission               = AdminPermission("refertime-getall")
)
