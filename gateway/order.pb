
�
service/order.prototestDb"�
	order_req
user_id (RuserId/
order_receive_mes_id (RorderReceiveMesId2
details (2.testDb.order_detail_reqRdetails
payment (Rpayment!
payment_type (RpaymentType
postage (Rpostage"
id
id (Rid"�


receiver_phone (	R
receiver_province (	RreceiverProvince#

receiver_district (	RreceiverDistrict)
receiver_address (	RreceiverAddress"�
order_detail_res

product_id (R	productId,
current_unit_price (RcurrentUnitPrice)
product_quantity (RproductQuantity
total_price (R
totalPrice

created_at (	R	createdAt!
product_name (	RproductName"�
order_detail_req

product_id (R	productId,
current_unit_price (RcurrentUnitPrice)
product_quantity (RproductQuantity"�
whole_order
id (Rid
user_id (RuserId2
details (2.testDb.order_detail_resRdetails6
receive_mes (2.testDb.order_receiveR
receiveMes
payment (Rpayment!
payment_type (RpaymentType
postage (Rpostage"j
user_order_all
id (Rid
	user_name (	RuserName+
orders (2.testDb.whole_orderRorders2�

post_order_mes.testDb.order_req.testDb.whole_order,
	get_order
.testDb.id.testDb.whole_order4
get_user_order
.testDb.id.testDb.user_order_allBZ.bproto3