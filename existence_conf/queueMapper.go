package existence_conf

import "convert-to-dpfm-product-master-from-sap-product-master/config"

type exconfQueueMapper map[string]string

func NewExconfQueueMapper(c *config.Conf) exconfQueueMapper {
	m := exconfQueueMapper{}
	ecQNames := c.RMQ.QueueToExConf()
	m["Product"] = ecQNames[0%len(ecQNames)]
	m["BusinessPartner"] = ecQNames[1%len(ecQNames)]
	return m
}

func (m exconfQueueMapper) getQueueNameByConfContent(content string) string {
	return m[content]
}
