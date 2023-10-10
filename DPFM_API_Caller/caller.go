package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "convert-to-dpfm-product-master-from-sap-product-master/DPFM_API_Input_Reader"
	"convert-to-dpfm-product-master-from-sap-product-master/config"
	"convert-to-dpfm-product-master-from-sap-product-master/existence_conf"
	"convert-to-dpfm-product-master-from-sap-product-master/sub_func_complementer"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient

	confirmor    *existence_conf.ExistenceConf
	complementer *sub_func_complementer.SubFuncComplementer
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,

	confirmor *existence_conf.ExistenceConf,
	complementer *sub_func_complementer.SubFuncComplementer,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:          context.Background(),
		conf:         conf,
		rmq:          rmq,
		confirmor:    confirmor,
		complementer: complementer,
	}
}

func (c *DPFMAPICaller) AsyncOrderCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,

	log *logger.Logger,
	// msg rabbitmq.RabbitmqMessage,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)
	exconfAllExist := false

	subFuncFin := make(chan error)
	exconfFin := make(chan error)

	// 他PODへ問い合わせ
	wg.Add(1)
	go func() {
		defer wg.Done()
		var e []error
		exconfAllExist, e = c.confirmor.Conf(input, log)
		if len(e) != 0 {
			mtx.Lock()
			errs = append(errs, e...)
			mtx.Unlock()
			exconfFin <- xerrors.Errorf("exconf error")
			return
		}
		exconfFin <- nil
	}()

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, subFuncFin, log, errs, input)
		case "Item":
			// TODO: 実装
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}

	// 後処理
	ticker := time.NewTicker(10 * time.Second)
	select {
	case e := <-exconfFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		errs = append(errs, xerrors.Errorf("time out"))
		return errs
	}

	if !exconfAllExist {
		mtx.Lock()
		return nil
	}
	select {
	case e := <-subFuncFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		mtx.Lock()
		errs = append(errs, xerrors.Errorf("time out"))
		return errs
	}

	log.JsonParseOut(input)
	return nil
}

func (c *DPFMAPICaller) headerCreate(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs []error, sdc *dpfm_api_input_reader.SDC) {
	defer wg.Done()
	err := c.complementer.ComplementHeader(sdc, log)
	if err != nil {
		mtx.Lock()
		errs = append(errs, err)
		mtx.Unlock()
		errFin <- xerrors.Errorf("complement error")
	}
	errFin <- nil

	// data_platform_orders_header_dataの更新
	headerData := sdc.ConvertToHeader()
	err = c.rmq.Send(c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "OrdersHeader"})
	if err != nil {
		log.Error(err)
		return
	}

	// data_platform_orders_header_partner_dataの更新
	for i := range sdc.Orders.HeaderPartner {
		headerPartnerData := sdc.ConvertToHeaderPartner(i)
		err = c.rmq.Send(c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerPartnerData, "function": "OrdersHeaderPartner"})
		if err != nil {
			log.Error(err)
			return
		}
	}

	// data_platform_orders_header_partner_plant_dataの更新
	for i := range sdc.Orders.HeaderPartner {
		for j := range sdc.Orders.HeaderPartner[i].HeaderPartnerPlant {
			headerPartnerPlantData := sdc.ConvertToHeaderPartnerPlant(i, j)
			err = c.rmq.Send(c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerPartnerPlantData, "function": "OrdersHeaderPartnerPlant"})
			if err != nil {
				log.Error(err)
				return
			}
		}

		// // data_platform_orders_header_partner_contact_dataの更新
		// for i := range sdc.Orders.HeaderPartner {
		// 	for j := range sdc.Orders.HeaderPartner[i].HeaderPartnerContact {
		// 		headerPartnerContactData := sdc.ConvertToHeaderPartnerContact(i, j)
		// 		err = c.rmq.Send(c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerPartnerContactData, "function": "OrdersHeaderPartnerContact"})
		// 		if err != nil {
		// 			log.Error(err)
		// 			return
		// 		}
		// 	}
		// }
	}
}

func (c *DPFMAPICaller) itemCreate(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs []error, input *dpfm_api_input_reader.SDC) {
	return
}
