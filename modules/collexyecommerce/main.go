package collexyecommerce

import
(
	//"fmt"
	"collexy/core/lib"
	"log"
	"net/http"
	
	//coreglobals "collexy/core/globals"
	//"github.com/gorilla/mux"
	"collexy/modules/collexyecommerce/controllers"
)

func init(){
	log.Println("Registered module: Collexy Ecommerce")

	// setup routes
	rCollexyEcommerceSection := lib.Route{"collexyEcommerce", "/admin/collexy-ecommerce", "modules/collexyecommerce/public/views/collexy-ecommerce/index.html", true}

	rOrderSection := lib.Route{"collexyEcommerce.order", "/order", "modules/collexyecommerce/public/views/order/index.html", false}
	rTransactionSection := lib.Route{"collexyEcommerce.transaction", "/transaction", "modules/collexyecommerce/public/views/transaction/index.html", false}
	rPaymentGatewaySection := lib.Route{"collexyEcommerce.paymentGateway", "/payment-gateway", "modules/collexyecommerce/public/views/payment-gateway/index.html", false}
	rShippingZoneSection := lib.Route{"collexyEcommerce.shippingZone", "/shipping-zone", "modules/collexyecommerce/public/views/shipping-zone/index.html", false}
	rShippingMethodSection := lib.Route{"collexyEcommerce.shippingMethod", "/shipping-method", "modules/collexyecommerce/public/views/shipping-method/index.html", false}
	rAccessRuleSection := lib.Route{"collexyEcommerce.accessRule", "/access-rule", "modules/collexyecommerce/public/views/access-rule/index.html", false}
	rSubscriptionProfileSection := lib.Route{"collexyEcommerce.subscriptionProfile", "/subscription-profile", "modules/collexyecommerce/public/views/subscription-profile/index.html", false}
	rStatsSection := lib.Route{"collexyEcommerce.stats", "/stats", "modules/collexyecommerce/public/views/stats/index.html", false}

	rOrderTreeMethodEdit := lib.Route{"collexyEcommerce.order.edit", "/edit/:id", "modules/collexyecommerce/public/views/order/edit.html", false}
	rOrderTreeMethodNew := lib.Route{"collexyEcommerce.order.new", "/new?type_id&parent_id", "modules/collexyecommerce/public/views/order/new.html", false}

	rTransactionTreeMethodEdit := lib.Route{"collexyEcommerce.transaction.edit", "/edit/:id", "modules/collexyecommerce/public/views/transaction/edit.html", false}
	rTransactionTreeMethodNew := lib.Route{"collexyEcommerce.transaction.new", "/new?type_id&parent_id", "modules/collexyecommerce/public/views/transaction/new.html", false}

	rPaymentGatewayTreeMethodEdit := lib.Route{"collexyEcommerce.paymentGateway.edit", "/edit/:id", "modules/collexyecommerce/public/views/payment-gateway/edit.html", false}
	rPaymentGatewayTreeMethodNew := lib.Route{"collexyEcommerce.paymentGateway.new", "/new?type_id&parent_id", "modules/collexyecommerce/public/views/payment-gateway/new.html", false}

	rShippingZoneTreeMethodEdit := lib.Route{"collexyEcommerce.shippingZone.edit", "/edit/:id", "modules/collexyecommerce/public/views/shipping-zone/edit.html", false}
	rShippingZoneTreeMethodNew := lib.Route{"collexyEcommerce.shippingZone.new", "/new?type_id&parent_id", "modules/collexyecommerce/public/views/shipping-zone/new.html", false}

	rShippingMethodTreeMethodEdit := lib.Route{"collexyEcommerce.shippingMethod.edit", "/edit/:id", "modules/collexyecommerce/public/views/shipping-method/edit.html", false}
	rShippingMethodTreeMethodNew := lib.Route{"collexyEcommerce.shippingMethod.new", "/new?type_id&parent_id", "modules/collexyecommerce/public/views/shipping-method/new.html", false}

	rAccessRuleTreeMethodEdit := lib.Route{"collexyEcommerce.accessRule.edit", "/edit/:id", "modules/collexyecommerce/public/views/access-rule/edit.html", false}
	rAccessRuleTreeMethodNew := lib.Route{"collexyEcommerce.accessRule.new", "/new?type_id&parent_id", "modules/collexyecommerce/public/views/access-rule/new.html", false}

	rSubscriptionProfileTreeMethodEdit := lib.Route{"collexyEcommerce.subscriptionProfile.edit", "/edit/:id", "modules/collexyecommerce/public/views/subscription-profile/edit.html", false}
	rSubscriptionProfileTreeMethodNew := lib.Route{"collexyEcommerce.subscriptionProfile.new", "/new?type_id&parent_id", "modules/collexyecommerce/public/views/subscription-profile/new.html", false}


	// setup trees
	routesOrderTree := []lib.Route{rOrderTreeMethodEdit, rOrderTreeMethodNew}
	routesTransactionTree := []lib.Route{rTransactionTreeMethodEdit, rTransactionTreeMethodNew}
	routesPaymentGatewayTree := []lib.Route{rPaymentGatewayTreeMethodEdit, rPaymentGatewayTreeMethodNew}
	routesShippingZoneTree := []lib.Route{rShippingZoneTreeMethodEdit, rShippingZoneTreeMethodNew}
	routesShippingMethodTree := []lib.Route{rShippingMethodTreeMethodEdit, rShippingMethodTreeMethodNew}
	routesAccessRuleTree := []lib.Route{rAccessRuleTreeMethodEdit, rAccessRuleTreeMethodNew}
	routesSubscriptionProfileTree := []lib.Route{rSubscriptionProfileTreeMethodEdit, rSubscriptionProfileTreeMethodNew}

	tOrder := lib.Tree{"Orders", "orders", routesOrderTree}
	tTransaction := lib.Tree{"Transactions", "transactions", routesTransactionTree}
	tPaymentGateway := lib.Tree{"Payment Gateways", "payment-gateways", routesPaymentGatewayTree}
	tShippingZone := lib.Tree{"Shipping Zones", "shipping-zones", routesShippingZoneTree}
	tShippingMethod := lib.Tree{"Shipping Methods", "shipping-methods", routesShippingMethodTree}
	tAccessRule := lib.Tree{"Access Rules", "access-rules", routesAccessRuleTree}
	tSubscriptionProfile := lib.Tree{"Subscription Profile", "subscription-profiles", routesSubscriptionProfileTree}

	treesOrderSection := []*lib.Tree{&tOrder}
	treesTransactionSection := []*lib.Tree{&tTransaction}
	treesPaymentGatewaySection := []*lib.Tree{&tPaymentGateway}
	treesShippingZoneSection := []*lib.Tree{&tShippingZone}
	treesShippingMethodSection := []*lib.Tree{&tShippingMethod}
	treesAccessRuleSection := []*lib.Tree{&tAccessRule}
	treesSubscriptionProfileSection := []*lib.Tree{&tSubscriptionProfile}

	/*

	should be changed to collexy_ecommerce_section once the permission is created!
	temporarily use settings_section permission

	sCollexyEcommerce := lib.Section{"Collexy Ecommerce Section", "collexyEcommerceSection", "fa fa-gear fa-fw", &rCollexyEcommerceSection, nil, true, nil, nil, []string{"settings_section"}}

	*/

	sCollexyEcommerce := lib.Section{"Collexy Ecommerce Section", "collexyEcommerceSection", "fa fa-shopping-cart fa-fw", &rCollexyEcommerceSection, nil, true, nil, nil, []string{"settings_section"}}

	sOrder := lib.Section{"Order Section", "orderSection", "fa fa-shopping-cart fa-fw", &rOrderSection, treesOrderSection, false, nil, nil, []string{"content_type_section"}}
	sTransaction := lib.Section{"Transaction Section", "transactionSection", "fa fa-money fa-fw", &rTransactionSection, treesTransactionSection, false, nil, nil, []string{"content_type_section"}}
	sPaymentGateway := lib.Section{"Payment Gateway Section", "paymentGatewaySection", "fa fa-cc-visa fa-fw", &rPaymentGatewaySection, treesPaymentGatewaySection, false, nil, nil, []string{"content_type_section"}}
	sShippingZone := lib.Section{"Shipping Zone Section", "shippingZoneSection", "fa fa-globe fa-fw", &rShippingZoneSection, treesShippingZoneSection, false, nil, nil, []string{"content_type_section"}}
	sShippingMethod := lib.Section{"Shipping Method Section", "shippingMethodSection", "fa fa-truck fa-fw", &rShippingMethodSection, treesShippingMethodSection, false, nil, nil, []string{"content_type_section"}}
	sAccessRule := lib.Section{"Access Rule Section", "accessRuleSection", "fa fa-unlock fa-fw", &rAccessRuleSection, treesAccessRuleSection, false, nil, nil, []string{"content_type_section"}}
	sSubscriptionProfile := lib.Section{"Subscription Profile Section", "subscriptionProfileSection", "fa fa-users fa-fw", &rSubscriptionProfileSection, treesSubscriptionProfileSection, false, nil, nil, []string{"content_type_section"}}
	sStats := lib.Section{"Stats Section", "statsSection", "fa fa-line-chart fa-fw", &rStatsSection, nil, false, nil, nil, []string{"content_type_section"}}
	
	lol := []lib.Section{sOrder, sTransaction, sPaymentGateway, sShippingZone, sShippingMethod, sAccessRule, sSubscriptionProfile, sStats}
	sCollexyEcommerce.Children = lol

	// setup module
	sections := []lib.Section{sCollexyEcommerce}
	// params: name, alias, description, sections

	var serverRoutes []lib.ServerRoute



	

	// wildcardHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "if no other routes are matched then show this")
	// }

	shoppingCartController := controllers.ShoppingCartController{}
	srvRoute1 := lib.ServerRoute{`/{_dummy:.*\/add-to-cart}`, http.HandlerFunc(shoppingCartController.AddToCart), []string{"POST"}}
	srvRoute2 := lib.ServerRoute{`/test-product`, http.HandlerFunc(shoppingCartController.TestProduct), []string{"GET"}}
	srvRoute3 := lib.ServerRoute{`/test-cart`, http.HandlerFunc(shoppingCartController.TestCart), []string{"GET"}}
	srvRoute4 := lib.ServerRoute{`/test-checkout`, http.HandlerFunc(shoppingCartController.TestCheckout), []string{"GET", "POST"}}
	//srvRoute2 := lib.ServerRoute{"/{_dummy:.*}", http.HandlerFunc(wildcardHandlerFunc), []string{"GET"}}
	//srvRoute3 := lib.ServerRoute{"/", http.HandlerFunc(wildcardHandlerFunc), []string{"GET"}}
	serverRoutes = []lib.ServerRoute{srvRoute1,srvRoute2,srvRoute3,srvRoute4}

	moduleCollexyEcommerce := lib.Module{"Collexy Ecommerce Module", "collexyEcommerceModule", "Just a collexyEcommerce module", sections, serverRoutes, 500}

	// register module
	lib.RegisterModule(moduleCollexyEcommerce)



	// Setup FileServer for the collexyEcommerce module
	log.Println("Registered a handler for static files. [collexyecommerce::module]")

	http.Handle("/modules/collexyecommerce/public/", http.FileServer(http.Dir("./")))
}