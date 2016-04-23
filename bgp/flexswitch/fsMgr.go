package FSMgr

import (
	"asicdServices"
	"bfdd"
	"errors"
	"fmt"
	"l3/bgp/rpc"
	"ribd"
	"utils/logging"
)

/*  Router manager will handle all the communication with ribd
 */
type FSRouteMgr struct {
	ribdClient *ribd.RIBDServicesClient
	plugin     string
	logger     *logging.Writer
}

/*  Interface manager will handle all the communication with asicd
 */
type FSIntfMgr struct {
	AsicdClient *asicdServices.ASICDServicesClient
	plugin      string
	logger      *logging.Writer
}

/*  @FUTURE: this will be using in future if FlexSwitch is planning to support
 *	     daemon which is handling policy statments
 */
type FSPolicyMgr struct {
	plugin string
	logger *logging.Writer
}

/*  BFD manager will handle all the communication with bfd daemon
 */
type FSBfdMgr struct {
	bfddClient *bfdd.BFDDServicesClient
	plugin     string
	logger     *logging.Writer
}

/*  Interface manager is responsible for handling asicd notifications and hence
 *  we are creating asicd client
 */
func NewFSIntfMgr(logger *logging.Writer, fileName string) (*FSIntfMgr, error) {
	var asicdClient *asicdServices.ASICDServicesClient = nil
	asicdClientChan := make(chan *asicdServices.ASICDServicesClient)

	logger.Info("Connecting to ASICd")
	go rpc.StartAsicdClient(logger, fileName, asicdClientChan)
	asicdClient = <-asicdClientChan
	if asicdClient == nil {
		logger.Err("Failed to connect to ASICd")
		return nil, errors.New("Failed to connect to ASICd")
	} else {
		logger.Info("Connected to ASICd")
	}
	mgr := &FSIntfMgr{
		plugin:      "ovsdb",
		AsicdClient: asicdClient,
		logger:      logger,
	}
	return mgr, nil
}

/*  Init policy manager with specific needs
 */
func NewFSPolicyMgr(logger *logging.Writer, fileName string) *FSPolicyMgr {
	mgr := &FSPolicyMgr{
		plugin: "ovsdb",
		logger: logger,
	}

	return mgr
}

/*  Init route manager with ribd client as its core
 */
func NewFSRouteMgr(logger *logging.Writer, fileName string) (*FSRouteMgr, error) {
	var ribdClient *ribd.RIBDServicesClient = nil
	ribdClientChan := make(chan *ribd.RIBDServicesClient)

	logger.Info("Connecting to RIBd")
	go rpc.StartRibdClient(logger, fileName, ribdClientChan)
	ribdClient = <-ribdClientChan
	if ribdClient == nil {
		logger.Err("Failed to connect to RIBd\n")
		return nil, errors.New("Failed to connect to RIBd")
	} else {
		logger.Info("Connected to RIBd")
	}

	mgr := &FSRouteMgr{
		plugin:     "ovsdb",
		ribdClient: ribdClient,
		logger:     logger,
	}

	return mgr, nil
}

/*  Init bfd manager with bfd client as its core
 */
func NewFSBfdMgr(logger *logging.Writer, fileName string) (*FSBfdMgr, error) {
	var bfddClient *bfdd.BFDDServicesClient = nil
	bfddClientChan := make(chan *bfdd.BFDDServicesClient)

	logger.Info("Connecting to BFDd")
	go rpc.StartBfddClient(logger, fileName, bfddClientChan)
	bfddClient = <-bfddClientChan
	if bfddClient == nil {
		logger.Err("Failed to connect to BFDd\n")
		return nil, errors.New("Failed to connect to BFDd")
	} else {
		logger.Info("Connected to BFDd")
	}
	mgr := &FSBfdMgr{
		plugin:     "ovsdb",
		logger:     logger,
		bfddClient: bfddClient,
	}

	return mgr, nil
}

func (mgr *FSRouteMgr) CreateRoute() {
	fmt.Println("Create Route called in", mgr.plugin)
}

func (mgr *FSRouteMgr) DeleteRoute() {

}

func (mgr *FSPolicyMgr) AddPolicy() {

}

func (mgr *FSPolicyMgr) RemovePolicy() {

}

func (mgr *FSIntfMgr) PortStateChange() {

}
