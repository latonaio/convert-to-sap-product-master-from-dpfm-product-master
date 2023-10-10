# convert-to-sap-product-master-from-dpfm-product-master

convert-to-sap-product-master-from-dpfm-product-master は、周辺業務システム　を データ連携基盤 と統合することを目的に、API で品目マスタデータを取得するマイクロサービスです。  
https://xxx.xxx.io/api/FUNCTION_CONVERT_TO_SAP_PRODUCT_MASTER_FROM_DPFM_RRODUCT_MASTER/function/

## 動作環境

convert-to-sap-product-master-from-dpfm-product-master の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
convert-to-sap-product-master-from-dpfm-product-master が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/FUNCTION_CONVERT_TO_SAP_PRODUCT_MASTER_FROM_DPFM_RRODUCT_MASTER/function/

## 本レポジトリ に 含まれる API名
convert-to-sap-product-master-from-dpfm-product-master には、次の API をコールするためのリソースが含まれています。  

* A_General（品目マスタ - 一般データ）
* A_ProductDesc（品目マスタ - 品目テキストデータ）
* A_BusinessPartner（品目マスタ - 取引先データ）
* A_ProductDescByBP（品目マスタ - ビジネスパートナ品目テキストデータ）
* A_BPPlant（品目マスタ - 取引先プラントデータ）
* A_StorageLocation（品目マスタ - 保管場所データ）
* A_StorageBin（品目マスタ - 棚番データ）
* A_MRPArea（品目マスタ - MRPエリアデータ）
* A_Production（品目マスタ - 製造データ）
* A_Quality（品目マスタ - 品質データ）
* A_Accounting（データ連携基盤 品目マスタ - 会計データ）
* A_Tax（品目マスタ - 税データ）
* A_Allergen（品目マスタ - アレルゲンデータ）
* A_NutritionalInfo（品目マスタ - 栄養成分データ）
* A_Calories（品目マスタ - 熱量データ）

## API への 値入力条件 の 初期値
convert-to-sap-product-master-from-dpfm-product-master において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"General" が指定されています。    
  
```
	"api_schema": "FunctionConvertToSAPProductMasterFromDPFMProductMaster",
	"accepter": ["General"],
	"order_id": null,
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "FunctionConvertToSAPProductMasterFromDPFMProductMaster",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は 品目マスタ の 一般データ が取得された結果の JSON の例です。  
以下の項目のうち、"Product" ～ "IsMarkedForDeletion" は、/DPFM_API_Output_Formatter/type.go 内 の Type General {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
XXX
```
