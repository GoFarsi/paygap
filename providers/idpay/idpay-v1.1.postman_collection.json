{
	"info": {
		"_postman_id": "4473dda1-76e6-4ac3-af59-464e18598eca",
		"name": "IDPay - Web Service - v1.1",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Create",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "text"
					},
					{
						"key": "X-API-KEY",
						"value": "6a7f99eb-7c20-4412-a972-6dfb7cd253a4",
						"equals": true
					},
					{
						"key": "X-SANDBOX",
						"value": "0",
						"equals": true
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"order_id\": 101,\n  \"amount\": 10000,\n  \"name\": \"قاسم رادمان\",\n  \"phone\": \"09382198592\",\n  \"mail\": \"my@site.com\",\n  \"desc\": \"توضیحات پرداخت کننده\",\n  \"callback\": \"https://example.com/callback\"\n}"
				},
				"url": {
					"raw": "https://api.idpay.ir/v1.1/payment",
					"protocol": "https",
					"host": [
						"api",
						"idpay",
						"ir"
					],
					"path": [
						"v1.1",
						"payment"
					]
				}
			},
			"response": [
				{
					"name": "Success",
					"originalRequest": {
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							},
							{
								"key": "X-API-KEY",
								"value": "6a7f99eb-7c20-4412-a972-6dfb7cd253a4",
								"equals": true
							},
							{
								"key": "X-SANDBOX",
								"value": "0",
								"equals": true
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"order_id\": 101,\n  \"amount\": 10000,\n  \"name\": \"قاسم رادمان\",\n  \"phone\": \"09382198592\",\n  \"mail\": \"my@site.com\",\n  \"desc\": \"توضیحات پرداخت کننده\",\n  \"callback\": \"https://example.com/callback\"\n}"
						},
						"url": {
							"raw": "https://api.idpay.ir/v1.1/payment",
							"protocol": "https",
							"host": [
								"api",
								"idpay",
								"ir"
							],
							"path": [
								"v1.1",
								"payment"
							]
						}
					},
					"status": "Created",
					"code": 201,
					"_postman_previewlanguage": "json",
					"header": [
						{
							"key": "Content-Type",
							"value": "application/json"
						}
					],
					"cookie": [],
					"body": "{\n    \"id\": \"e22952579725883bbad9f8fa429134bf\",\n    \"link\": \"http://idpay.local/p/ws/e22952579725883bbad9f8fa429134bf\"\n}"
				},
				{
					"name": "Error Code 38",
					"originalRequest": {
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							},
							{
								"key": "X-API-KEY",
								"value": "6a7f99eb-7c20-4412-a972-6dfb7cd253a4",
								"equals": true
							},
							{
								"key": "X-SANDBOX",
								"value": "0",
								"equals": true
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"order_id\": 101,\n  \"amount\": 10000,\n  \"name\": \"قاسم رادمان\",\n  \"phone\": \"09382198592\",\n  \"mail\": \"my@site.com\",\n  \"desc\": \"توضیحات پرداخت کننده\",\n  \"callback\": \"https://new-domain.com/callback\"\n}"
						},
						"url": {
							"raw": "https://api.idpay.ir/v1.1/payment",
							"protocol": "https",
							"host": [
								"api",
								"idpay",
								"ir"
							],
							"path": [
								"v1.1",
								"payment"
							]
						}
					},
					"status": "Not Acceptable : Your request has been sent from the https://wrong-domain.com domain. The `callback` does not match the domain address on the web service.",
					"code": 406,
					"_postman_previewlanguage": "json",
					"header": [
						{
							"key": "Content-Type",
							"value": "application/json; charset=utf-8"
						}
					],
					"cookie": [],
					"body": "{\n    \"error_code\": 38,\n    \"error_message\": \"درخواست شما از آدرس https://new-domain.com ارسال شده است. دامنه آدرس بازگشت `callback` با آدرس ثبت شده در وب سرویس همخوانی ندارد.\"\n}"
				}
			]
		},
		{
			"name": "Verify",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "text"
					},
					{
						"key": "X-API-KEY",
						"value": "6a7f99eb-7c20-4412-a972-6dfb7cd253a4",
						"equals": true
					},
					{
						"key": "X-SANDBOX",
						"value": "0",
						"equals": true
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"id\": \"e22952579725883bbad9f8fa429134bf\",\n  \"order_id\": \"101\"\n}"
				},
				"url": {
					"raw": "https://api.idpay.ir/v1.1/payment/verify",
					"protocol": "https",
					"host": [
						"api",
						"idpay",
						"ir"
					],
					"path": [
						"v1.1",
						"payment",
						"verify"
					]
				}
			},
			"response": [
				{
					"name": "Success Status 100",
					"originalRequest": {
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							},
							{
								"key": "X-API-KEY",
								"value": "6a7f99eb-7c20-4412-a972-6dfb7cd253a4",
								"equals": true
							},
							{
								"key": "X-SANDBOX",
								"value": "0",
								"equals": true
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"id\": \"e22952579725883bbad9f8fa429134bf\",\n  \"order_id\": \"101\"\n}"
						},
						"url": {
							"raw": "https://api.idpay.ir/v1.1/payment/verify",
							"protocol": "https",
							"host": [
								"api",
								"idpay",
								"ir"
							],
							"path": [
								"v1.1",
								"payment",
								"verify"
							]
						}
					},
					"status": "OK",
					"code": 200,
					"_postman_previewlanguage": "json",
					"header": [
						{
							"key": "Content-Type",
							"value": "application/json"
						}
					],
					"cookie": [],
					"body": "{\n    \"status\": 100,\n    \"track_id\": \"3872979\",\n    \"id\": \"e22952579725883bbad9f8fa429134bf\",\n    \"order_id\": \"101\",\n    \"amount\": \"10000\",\n    \"date\": \"1546288200\",\n    \"payment\": {\n        \"track_id\": \"123456789001\",\n        \"amount\": \"10000\",\n        \"card_no\": \"123456******1234\",\n        \"hashed_card_no\": \"E59FA6241C94B8836E3D03120DF33E80FD988888BBA0A122240C2E7D23B48295\",\n        \"date\": \"1549185575\"\n    },\n    \"verify\": {\n        \"date\": \"1546288500\"\n    }\n}"
				},
				{
					"name": "Success Status 101",
					"originalRequest": {
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							},
							{
								"key": "X-API-KEY",
								"value": "6a7f99eb-7c20-4412-a972-6dfb7cd253a4",
								"equals": true
							},
							{
								"key": "X-SANDBOX",
								"value": "0",
								"equals": true
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"id\": \"e22952579725883bbad9f8fa429134bf\",\n  \"order_id\": \"101\"\n}"
						},
						"url": {
							"raw": "https://api.idpay.ir/v1.1/payment/verify",
							"protocol": "https",
							"host": [
								"api",
								"idpay",
								"ir"
							],
							"path": [
								"v1.1",
								"payment",
								"verify"
							]
						}
					},
					"status": "OK",
					"code": 200,
					"_postman_previewlanguage": "json",
					"header": [
						{
							"key": "Content-Type",
							"value": "application/json"
						}
					],
					"cookie": [],
					"body": "{\n    \"status\": 101,\n    \"track_id\": \"3872979\",\n    \"id\": \"e22952579725883bbad9f8fa429134bf\",\n    \"order_id\": \"101\",\n    \"amount\": \"10000\",\n    \"date\": \"1546288200\",\n    \"payment\": {\n        \"track_id\": \"123456789001\",\n        \"amount\": \"10000\",\n        \"card_no\": \"123456******1234\",\n        \"hashed_card_no\": \"E59FA6241C94B8836E3D03120DF33E80FD988888BBA0A122240C2E7D23B48295\",\n        \"date\": \"1549185575\"\n    },\n    \"verify\": {\n        \"date\": \"1546288500\"\n    }\n}"
				},
				{
					"name": "Error Code 53",
					"originalRequest": {
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							},
							{
								"key": "X-API-KEY",
								"value": "6a7f99eb-7c20-4412-a972-6dfb7cd253a4",
								"equals": true
							},
							{
								"key": "X-SANDBOX",
								"value": "0",
								"equals": true
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"id\": \"e22952579725883bbad9f8fa429134bf\",\n  \"order_id\": \"101\"\n}"
						},
						"url": {
							"raw": "https://api.idpay.ir/v1.1/payment/verify",
							"protocol": "https",
							"host": [
								"api",
								"idpay",
								"ir"
							],
							"path": [
								"v1.1",
								"payment",
								"verify"
							]
						}
					},
					"status": ": Payment verification is not possible.",
					"code": 405,
					"_postman_previewlanguage": "json",
					"header": [
						{
							"key": "Content-Type",
							"value": "application/json; charset=utf-8"
						}
					],
					"cookie": [],
					"body": "{\n    \"error_code\": 53,\n    \"error_message\": \"تایید پرداخت امکان پذیر نیست.\"\n}"
				},
				{
					"name": "Error Code 54",
					"originalRequest": {
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							},
							{
								"key": "X-API-KEY",
								"value": "6a7f99eb-7c20-4412-a972-6dfb7cd253a4",
								"equals": true
							},
							{
								"key": "X-SANDBOX",
								"value": "0",
								"equals": true
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"id\": \"e22952579725883bbad9f8fa429134bf\",\n  \"order_id\": \"101\"\n}"
						},
						"url": {
							"raw": "https://api.idpay.ir/v1.1/payment/verify",
							"protocol": "https",
							"host": [
								"api",
								"idpay",
								"ir"
							],
							"path": [
								"v1.1",
								"payment",
								"verify"
							]
						}
					},
					"status": ": Payment verification period has expired.",
					"code": 405,
					"_postman_previewlanguage": "json",
					"header": [
						{
							"key": "Content-Type",
							"value": "application/json; charset=utf-8"
						}
					],
					"cookie": [],
					"body": "{\n    \"error_code\": 54,\n    \"error_message\": \"مدت زمان تایید پرداخت سپری شده است.\"\n}"
				}
			]
		},
		{
			"name": "Inquiry",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "text"
					},
					{
						"key": "X-API-KEY",
						"value": "6a7f99eb-7c20-4412-a972-6dfb7cd253a4",
						"equals": true
					},
					{
						"key": "X-SANDBOX",
						"value": "0",
						"equals": true
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"id\": \"e22952579725883bbad9f8fa429134bf\",\n  \"order_id\": \"101\"\n}"
				},
				"url": {
					"raw": "https://api.idpay.ir/v1.1/payment/inquiry",
					"protocol": "https",
					"host": [
						"api",
						"idpay",
						"ir"
					],
					"path": [
						"v1.1",
						"payment",
						"inquiry"
					]
				}
			},
			"response": [
				{
					"name": "Success",
					"originalRequest": {
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							},
							{
								"key": "X-API-KEY",
								"value": "6a7f99eb-7c20-4412-a972-6dfb7cd253a4",
								"equals": true
							},
							{
								"key": "X-SANDBOX",
								"value": "0",
								"equals": true
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"id\": \"e22952579725883bbad9f8fa429134bf\",\n  \"order_id\": \"101\"\n}"
						},
						"url": {
							"raw": "https://api.idpay.ir/v1.1/payment/inquiry",
							"protocol": "https",
							"host": [
								"api",
								"idpay",
								"ir"
							],
							"path": [
								"v1.1",
								"payment",
								"inquiry"
							]
						}
					},
					"status": "OK",
					"code": 200,
					"_postman_previewlanguage": "json",
					"header": [
						{
							"key": "Content-Type",
							"value": "application/json"
						}
					],
					"cookie": [],
					"body": "{\n    \"status\": 100,\n    \"track_id\": \"3872979\",\n    \"id\": \"e22952579725883bbad9f8fa429134bf\",\n    \"order_id\": \"101\",\n    \"amount\": \"10000\",\n    \"wage\": {\n        \"by\": \"payee\",\n        \"type\": \"percent\",\n        \"amount\": \"2500\"\n    },\n    \"date\": \"1546288200\",\n    \"payer\": {\n        \"name\": \"قاسم رادمان\",\n        \"phone\": \"09382198592\",\n        \"mail\": \"my@site.com\",\n        \"desc\": \"توضیحات پرداخت کننده\"\n    },\n    \"payment\": {\n        \"track_id\": \"123456789001\",\n        \"amount\": \"10000\",\n        \"card_no\": \"123456******1234\",\n        \"hashed_card_no\": \"E59FA6241C94B8836E3D03120DF33E80FD988888BBA0A122240C2E7D23B48295\",\n        \"date\": \"1549185575\"\n    },\n    \"verify\": {\n        \"date\": \"1546288500\"\n    },\n    \"settlement\": {\n        \"track_id\": \"\",\n        \"amount\": \"7500\",\n        \"date\": \"1546398000\"\n    }\n}"
				}
			]
		}
	]
}