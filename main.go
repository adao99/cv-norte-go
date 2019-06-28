package main

import (
	"fmt"
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"time"
)

func main() {
	fmt.Println("__Program Starting__")
	fmt.Println("Opening SQL-Lite")
	conn, err := sqlite3.Open("database/database.sqlite")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	stmt, err := conn.Prepare(`SELECT 
			docheader.id AS docheader_id,
			docheader.deleted AS docheader_deleted,
			docheader.docconfig_id as docheader_doc_config_id, 
			docheader.entity_name AS docheader_entity_name, 
			docheader.doc_date AS docheader_doc_date,
			docheader.clock_date AS docheader_clock_date,
			docheader.doc_number AS docheader_doc_number,
			docheader.entity_id AS docheader_entity_id,
			docheader.entity_address_id AS docheader_entity_address_id,
			docheader.employee_id AS docheader_open_employee_id,
			docheader.currency_id AS docheader_currency_id,
			docheader.print_count AS docheader_print_count,
			docheader.docstatus_id AS docheader_docstatus_id,
			docheader.controlstatus_id AS docheader_controlstatus_id,
			docheader.local_id AS docheader_local_id,
			docheader.accumulator AS docheader_accumulator,
			docheader.taxscenario_id AS docheader_taxscenario_id,
			docheader.taxincluded AS docheader_taxincluded,
			docheader.cashbox_id AS docheader_cashbox_id,
			docmovdetail.id AS docmovdetail_id,
			docmovdetail.line_number AS docmovdetail_line_number,
			docmovdetail.parent_id AS docmovdetail_parent_id,
			docmovdetail.product_id AS docmovdetail_product_id,
			docmovdetail.discount_value AS docmovdetail_discount_value,
			docmovdetail.qnt AS docmovdetail_quantity,
			docmovdetail.unit_id AS docmovdetail_unit_id,
			docmovdetail.price AS docmovdetail_price,
			docmovdetail.precampaign_price AS docmovdetail_precampaign_price,
			docmovdetail.gross_total AS docmovdetail_gross_total,
			docmovdetail.warehouse_id AS docmovdetail_warehouse_id,
			docmovdetail.workstation_id AS docmovdetail_workstation_id,
			docmovdetail.productset_id AS docmovdetail_productset_id,
			docmovdetail.employee_id AS docmovdetail_employee_id,
			docmovdetail.detailstatus_id AS docmovdetail_detailstatus_id,
			docmovdetail.campaign_id AS docmovdetail_campaign_id,
			docmovdetail.total AS docmovdetail_total,
			product.id AS product_id,
			product.family_id AS product_family_id,
			product.model_id AS product_model_id,
			product.productdiscount_id AS product_productdiscount_id,
			product.productcommission_id AS product_productcommission_id,
			product.print_sort_order AS product_print_sort_order,
			product.printzone_id AS product_printzone_id,
			product.taxgroup_id AS product_taxgroup_id,
			product.base_unit_id AS product_base_unit_id,
			product.description AS product_description,
			product.code AS product_code,
			product.forpurchase AS product_for_purchase, 
			product.forsale AS product_forsale,
			product.generic_id AS product_generic_id,
			product.obs AS product_obs,
			product.stockconfig_id AS product_stockconfig_id,
			product.default_sale_unit_id AS product_default_sale_unit_id,
			product.default_purchase_unit_id AS product_default_purchase_unit_id,
			product.default_stock_unit_id AS product_default_stock_unit_id,
			product.label_unit_id AS product_label_unit_id,
			product.order_code AS product_order_code, 
			product.purchaseprice AS product_purchase_price, 
			product.saleprice AS product_saleprice, 
			product.product_type AS product_producttype,
			product.status AS product_status,
			product.deleted AS product_deleted,
			family.code AS family_code,
			family.description AS family_description,
			family.deleted AS family_deleted,
			local.code AS local_code, 
			local.description AS local_description,
			local.saleprice_id AS local_saleprice_id,
			local.paymethod_id AS local_paymethod_id,
			local.taxincluded AS local_taxincluded,
			employee.internal_name AS employee_name
		FROM docmovdetail
			INNER JOIN docheader ON docheader.id = docmovdetail.docheader_id
			INNER JOIN product ON docmovdetail.product_id = product.id
			INNER JOIN family ON family.id = product.family_id
			INNER JOIN local ON local.id = docheader.local_id
			INNER JOIN employee ON employee.entity_id = docmovdetail.employee_id
		WHERE docheader.deleted = 0 AND docheader.docconfig_id = 1005 OR docheader.docconfig_id = 1018`)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	for {
		hasRow, err := stmt.Step()
		if err != nil {
			fmt.Println(err)
		}
		if !hasRow {
			// The query is finished
			break
		}
		// Use Scan to access column data from a row
		var docheader_id string

		err = stmt.Scan(&docheader_id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("param_value:", docheader_id)
	}
	// It's always a good idea to set a busy timeout
	conn.BusyTimeout(5 * time.Second)
}
