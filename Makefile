PRODUCT_SERVICE_ROOT := ./product-service
PORTAL_SERVICE_ROOT := ./portal-service

run_product_service:
	@cd $(PRODUCT_SERVICE_ROOT) && go run $$(find cmd/ -name main.go) > ./product-service.txt

run_portal_service:
	@cd $(PORTAL_SERVICE_ROOT) && go run $$(find cmd/ -name main.go) >> ./portal-service.txt

